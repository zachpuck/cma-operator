/*
Copyright 2018 Samsung SDS Cloud Native Computing Team.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/samsung-cnct/cma-operator/pkg/apis/cma/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SDSAppSuiteLister helps list SDSAppSuites.
type SDSAppSuiteLister interface {
	// List lists all SDSAppSuites in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SDSAppSuite, err error)
	// SDSAppSuites returns an object that can list and get SDSAppSuites.
	SDSAppSuites(namespace string) SDSAppSuiteNamespaceLister
	SDSAppSuiteListerExpansion
}

// sDSAppSuiteLister implements the SDSAppSuiteLister interface.
type sDSAppSuiteLister struct {
	indexer cache.Indexer
}

// NewSDSAppSuiteLister returns a new SDSAppSuiteLister.
func NewSDSAppSuiteLister(indexer cache.Indexer) SDSAppSuiteLister {
	return &sDSAppSuiteLister{indexer: indexer}
}

// List lists all SDSAppSuites in the indexer.
func (s *sDSAppSuiteLister) List(selector labels.Selector) (ret []*v1alpha1.SDSAppSuite, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SDSAppSuite))
	})
	return ret, err
}

// SDSAppSuites returns an object that can list and get SDSAppSuites.
func (s *sDSAppSuiteLister) SDSAppSuites(namespace string) SDSAppSuiteNamespaceLister {
	return sDSAppSuiteNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SDSAppSuiteNamespaceLister helps list and get SDSAppSuites.
type SDSAppSuiteNamespaceLister interface {
	// List lists all SDSAppSuites in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SDSAppSuite, err error)
	// Get retrieves the SDSAppSuite from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SDSAppSuite, error)
	SDSAppSuiteNamespaceListerExpansion
}

// sDSAppSuiteNamespaceLister implements the SDSAppSuiteNamespaceLister
// interface.
type sDSAppSuiteNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SDSAppSuites in the indexer for a given namespace.
func (s sDSAppSuiteNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SDSAppSuite, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SDSAppSuite))
	})
	return ret, err
}

// Get retrieves the SDSAppSuite from the indexer for a given namespace and name.
func (s sDSAppSuiteNamespaceLister) Get(name string) (*v1alpha1.SDSAppSuite, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sdsappsuite"), name)
	}
	return obj.(*v1alpha1.SDSAppSuite), nil
}
