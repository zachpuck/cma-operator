package sdsappsuite

import (
	"github.com/spf13/viper"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"time"

	api "github.com/samsung-cnct/cma-operator/pkg/apis/cma/v1alpha1"
	"github.com/samsung-cnct/cma-operator/pkg/generated/cma/client/clientset/versioned"
	"github.com/samsung-cnct/cma-operator/pkg/util/cmagrpc"
	"github.com/samsung-cnct/cma-operator/pkg/util/k8sutil"
)

const (
	KubernetesNamespaceViperVariableName = "kubernetes-namespace"
)

type SDSAppSuiteController struct {
	indexer  cache.Indexer
	queue    workqueue.RateLimitingInterface
	informer cache.Controller

	client        *versioned.Clientset
	cmaGRPCClient cmagrpc.ClientInterface
}

func NewSDSAppSuiteController(config *rest.Config) (output *SDSAppSuiteController, err error) {
	cmaGRPCClient, err := cmagrpc.CreateNewDefaultClient()
	if err != nil {
		return nil, err
	}
	if config == nil {
		config = k8sutil.DefaultConfig
	}
	client := versioned.NewForConfigOrDie(config)

	// create sdsappsuite list/watcher
	sdsAppSuiteListWatcher := cache.NewListWatchFromClient(
		client.CmaV1alpha1().RESTClient(),
		api.SDSAppSuiteResourcePlural,
		viper.GetString(KubernetesNamespaceViperVariableName),
		fields.Everything())

	// create the workqueue
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	// Bind the workqueue to a cache with the help of an informer. This way we make sure that
	// whenever the cache is updated, the SDSCluster key is added to the workqueue.
	// Note that when we finally process the item from the workqueue, we might see a newer version
	// of the SDSAppSuite than the version which was responsible for triggering the update.
	indexer, informer := cache.NewIndexerInformer(sdsAppSuiteListWatcher, &api.SDSAppSuite{}, 30*time.Second, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
		UpdateFunc: func(old interface{}, new interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(new)
			if err == nil {
				queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			// IndexerInformer uses a delta queue, therefore for deletes we have to use this
			// key function.
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
	}, cache.Indexers{})

	output = &SDSAppSuiteController{
		informer: informer,
		indexer: indexer,
		queue: queue,
		client: client,
		cmaGRPCClient: cmaGRPCClient,
	}
	output.SetLogger()
	return
}

