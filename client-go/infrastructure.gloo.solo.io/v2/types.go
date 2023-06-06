// Code generated by skv2. DO NOT EDIT.

// Definitions for the Kubernetes types
package v2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for CloudProvider
var CloudProviderGVK = schema.GroupVersionKind{
	Group:   "infrastructure.gloo.solo.io",
	Version: "v2",
	Kind:    "CloudProvider",
}

// CloudProvider is the Schema for the cloudProvider API
type CloudProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudProviderSpec   `json:"spec,omitempty"`
	Status CloudProviderStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (CloudProvider) GVK() schema.GroupVersionKind {
	return CloudProviderGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CloudProviderList contains a list of CloudProvider
type CloudProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudProvider `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +genclient:noStatus

// GroupVersionKind for CloudResources
var CloudResourcesGVK = schema.GroupVersionKind{
	Group:   "infrastructure.gloo.solo.io",
	Version: "v2",
	Kind:    "CloudResources",
}

// CloudResources is the Schema for the cloudResources API
type CloudResources struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec CloudResourcesSpec `json:"spec,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (CloudResources) GVK() schema.GroupVersionKind {
	return CloudResourcesGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CloudResourcesList contains a list of CloudResources
type CloudResourcesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudResources `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudProvider{}, &CloudProviderList{})
	SchemeBuilder.Register(&CloudResources{}, &CloudResourcesList{})
}
