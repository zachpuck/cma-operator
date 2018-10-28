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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/samsung-cnct/cma-operator/pkg/generated/cma/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// SDSAppSuites returns a SDSAppSuiteInformer.
	SDSAppSuites() SDSAppSuiteInformer
	// SDSApplications returns a SDSApplicationInformer.
	SDSApplications() SDSApplicationInformer
	// SDSClusters returns a SDSClusterInformer.
	SDSClusters() SDSClusterInformer
	// SDSPackageManagers returns a SDSPackageManagerInformer.
	SDSPackageManagers() SDSPackageManagerInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// SDSAppSuites returns a SDSAppSuiteInformer.
func (v *version) SDSAppSuites() SDSAppSuiteInformer {
	return &sDSAppSuiteInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SDSApplications returns a SDSApplicationInformer.
func (v *version) SDSApplications() SDSApplicationInformer {
	return &sDSApplicationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SDSClusters returns a SDSClusterInformer.
func (v *version) SDSClusters() SDSClusterInformer {
	return &sDSClusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SDSPackageManagers returns a SDSPackageManagerInformer.
func (v *version) SDSPackageManagers() SDSPackageManagerInformer {
	return &sDSPackageManagerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
