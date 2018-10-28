package sdsappsuite

import (
	"fmt"
	"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"time"
	api "github.com/samsung-cnct/cma-operator/pkg/apis/cma/v1alpha1"
)

func (c *SDSAppSuiteController) processNextItem() bool {
	// Wait until there is a new item in the work queue
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	// Tell the queue that we are done with processing this key. This unblocks the key for other workers
	// This allows safe parallel processing because two SDSClusters with the same key are never processed in
	// parallel.
	defer c.queue.Done(key)

	// Invoke the method containing the business logic
	err := c.processItem(key.(string))
	// Handle the error if something went wrong during the execution of the business logic
	c.handleErr(err, key)
	return true
}

// processItem is the business logic of the controller
func (c *SDSAppSuiteController) processItem(key string) error {
	obj, exists, err :=  c.indexer.GetByKey(key)
	if err != nil {
		glog.Errorf("Fetching object with key -->%s<-- from store failed with -->%v<--", key, err)
		return err
	}

	if !exists {
		// Below we will warm up our cache with a SDSAppSuite, so that we will see a delete for one SDSAppSuite
		fmt.Printf("SDSAppSuite -->%s<-- does not exist anymore\n", key)
	} else {
		SDSAppSuite := obj.(*api.SDSAppSuite)
		clusterName := SDSAppSuite.GetName()
		fmt.Printf("SDSAppSuite -->%s<-- does exist (name=%s)!\n", key, clusterName)

		switch SDSAppSuite.Status.Phase {
		case api.AppSuitePhaseNone, api.AppSuitePhasePending:
			c.deployAppSuite(SDSAppSuite)
			break
		case api.AppSuitePhaseInstalling:
			c.waitForAppSuite(SDSAppSuite)
			break
		}
	}
	return nil
}

// handleErr checks if an error happend and makes sure we will retry later
func (c *SDSAppSuiteController) handleErr(err error, key interface{}) {
	if err == nil {
		// Forget about the #AddRateLimited history of the key on every successful synchronization.
		// This ensures that future processing of updates for this key is not delayed because of
		// an outdated error history.
		c.queue.Forget(key)
		return
	}

	// This controller retries 5 times if something goes wrong. After that, it stops trying.
	if c.queue.NumRequeues(key) < 5 {
		glog.Infof("Error syncing appSuite %v: %v", key, err)

		// Re-enqueue the key rate limited. Based on the rate limiter on the
		// queue and the re-enqueue history, the key will be processed later again.
		c.queue.AddRateLimited(key)
		return
	}

	c.queue.Forget(key)
	// Report to an external entity that, even after several retries, we could not successfully process this key
	runtime.HandleError(err)
	glog.Infof("Dropping appSuite %q out of the queue: %v", key, err)
}

func (c *SDSAppSuiteController) Run(threadiness int, stopCh chan struct{}) {
	defer runtime.HandleCrash()

	// Let the workers stop when we are done
	defer c.queue.ShutDown()
	glog.Info("Starting SDSAppSuite controller")

	go c.informer.Run(stopCh)

	// Wait for all involved caches to be synced, before processing items from the queue is started
	if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}

	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	<-stopCh
	glog.Info("Stopping SDSAppSuite controller")
}

func (c *SDSAppSuiteController) runWorker() {
	for c.processNextItem() {
	}
}
