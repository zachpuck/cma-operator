package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	AppSuitePhaseNone        = ""
	AppSuitePhasePending     = "Pending"
	AppSuitePhaseInstalling  = "Installing"
	AppSuitePhaseUpgrading   = "Upgrading"
	AppSuitePhaseImplemented = "Implemented"
	AppSuitePhaseFailed      = "Failed"
)

// SDSDAppSuiteList is a list of sds app suites.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type SDSAppSuiteList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SDSAppSuite `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type SDSAppSuite struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SDSAppSuiteSpec   `json:"spec"`
	Status            SDSAppSuiteStatus `json:"status"`
}

// SDSAppSuiteSpec represents a SDSAppSuite spec
// +k8s:openapi-gen=true
type SDSAppSuiteSpec struct {
	// What is the app suite name
	Name string `json:"name"`
}

// SDSAppSuiteStatus has the status of the system
// +k8s:openapi-gen=true
type SDSAppSuiteStatus struct {
	Phase      ApplicationPhase `json:"phase"`
	Ready      bool             `json:"ready"`
	Conditions []Condition      `json:"conditions"`
}
