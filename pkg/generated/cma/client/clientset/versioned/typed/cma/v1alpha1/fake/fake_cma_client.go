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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/samsung-cnct/cma-operator/pkg/generated/cma/client/clientset/versioned/typed/cma/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCmaV1alpha1 struct {
	*testing.Fake
}

func (c *FakeCmaV1alpha1) SDSAppSuites(namespace string) v1alpha1.SDSAppSuiteInterface {
	return &FakeSDSAppSuites{c, namespace}
}

func (c *FakeCmaV1alpha1) SDSApplications(namespace string) v1alpha1.SDSApplicationInterface {
	return &FakeSDSApplications{c, namespace}
}

func (c *FakeCmaV1alpha1) SDSClusters(namespace string) v1alpha1.SDSClusterInterface {
	return &FakeSDSClusters{c, namespace}
}

func (c *FakeCmaV1alpha1) SDSPackageManagers(namespace string) v1alpha1.SDSPackageManagerInterface {
	return &FakeSDSPackageManagers{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCmaV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
