// Code generated by skv2. DO NOT EDIT.

// Definitions for the Kubernetes types
package v1alpha1

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/runtime/schema"
    i030ea16e9a32c349e55238df46175e70 "github.com/solo-io/solo-apis/pkg/api/fed.ratelimit.solo.io/v1alpha1/types"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for FederatedRateLimitConfig
var FederatedRateLimitConfigGVK = schema.GroupVersionKind{
    Group: "fed.ratelimit.solo.io",
    Version: "v1alpha1",
    Kind: "FederatedRateLimitConfig",
}

// FederatedRateLimitConfig is the Schema for the federatedRateLimitConfig API
type FederatedRateLimitConfig struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec i030ea16e9a32c349e55238df46175e70.FederatedRateLimitConfigSpec `json:"spec,omitempty"`
    Status i030ea16e9a32c349e55238df46175e70.FederatedRateLimitConfigStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (FederatedRateLimitConfig)  GVK() schema.GroupVersionKind {
	return FederatedRateLimitConfigGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FederatedRateLimitConfigList contains a list of FederatedRateLimitConfig
type FederatedRateLimitConfigList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []FederatedRateLimitConfig `json:"items"`
}

func init() {
    SchemeBuilder.Register(&FederatedRateLimitConfig{}, &FederatedRateLimitConfigList{})
}
