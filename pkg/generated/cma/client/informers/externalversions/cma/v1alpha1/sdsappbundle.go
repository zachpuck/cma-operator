/*
Copyright 2019 Samsung SDS Cloud Native Computing Team.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	cma_v1alpha1 "github.com/samsung-cnct/cma-operator/pkg/apis/cma/v1alpha1"
	versioned "github.com/samsung-cnct/cma-operator/pkg/generated/cma/client/clientset/versioned"
	internalinterfaces "github.com/samsung-cnct/cma-operator/pkg/generated/cma/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/samsung-cnct/cma-operator/pkg/generated/cma/client/listers/cma/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SDSAppBundleInformer provides access to a shared informer and lister for
// SDSAppBundles.
type SDSAppBundleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.SDSAppBundleLister
}

type sDSAppBundleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSDSAppBundleInformer constructs a new informer for SDSAppBundle type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSDSAppBundleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSDSAppBundleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSDSAppBundleInformer constructs a new informer for SDSAppBundle type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSDSAppBundleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CmaV1alpha1().SDSAppBundles(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CmaV1alpha1().SDSAppBundles(namespace).Watch(options)
			},
		},
		&cma_v1alpha1.SDSAppBundle{},
		resyncPeriod,
		indexers,
	)
}

func (f *sDSAppBundleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSDSAppBundleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sDSAppBundleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cma_v1alpha1.SDSAppBundle{}, f.defaultInformer)
}

func (f *sDSAppBundleInformer) Lister() v1alpha1.SDSAppBundleLister {
	return v1alpha1.NewSDSAppBundleLister(f.Informer().GetIndexer())
}
