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

// GroupVersionKind for GraphQLPersistedQueryCachePolicy
var GraphQLPersistedQueryCachePolicyGVK = schema.GroupVersionKind{
	Group:   "resilience.policy.gloo.solo.io",
	Version: "v2",
	Kind:    "GraphQLPersistedQueryCachePolicy",
}

// GraphQLPersistedQueryCachePolicy is the Schema for the graphQLPersistedQueryCachePolicy API
type GraphQLPersistedQueryCachePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GraphQLPersistedQueryCachePolicySpec   `json:"spec,omitempty"`
	Status GraphQLPersistedQueryCachePolicyStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (GraphQLPersistedQueryCachePolicy) GVK() schema.GroupVersionKind {
	return GraphQLPersistedQueryCachePolicyGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GraphQLPersistedQueryCachePolicyList contains a list of GraphQLPersistedQueryCachePolicy
type GraphQLPersistedQueryCachePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GraphQLPersistedQueryCachePolicy `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for FailoverPolicy
var FailoverPolicyGVK = schema.GroupVersionKind{
	Group:   "resilience.policy.gloo.solo.io",
	Version: "v2",
	Kind:    "FailoverPolicy",
}

// FailoverPolicy is the Schema for the failoverPolicy API
type FailoverPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FailoverPolicySpec   `json:"spec,omitempty"`
	Status FailoverPolicyStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (FailoverPolicy) GVK() schema.GroupVersionKind {
	return FailoverPolicyGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FailoverPolicyList contains a list of FailoverPolicy
type FailoverPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FailoverPolicy `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for OutlierDetectionPolicy
var OutlierDetectionPolicyGVK = schema.GroupVersionKind{
	Group:   "resilience.policy.gloo.solo.io",
	Version: "v2",
	Kind:    "OutlierDetectionPolicy",
}

// OutlierDetectionPolicy is the Schema for the outlierDetectionPolicy API
type OutlierDetectionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OutlierDetectionPolicySpec   `json:"spec,omitempty"`
	Status OutlierDetectionPolicyStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (OutlierDetectionPolicy) GVK() schema.GroupVersionKind {
	return OutlierDetectionPolicyGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OutlierDetectionPolicyList contains a list of OutlierDetectionPolicy
type OutlierDetectionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OutlierDetectionPolicy `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for FaultInjectionPolicy
var FaultInjectionPolicyGVK = schema.GroupVersionKind{
	Group:   "resilience.policy.gloo.solo.io",
	Version: "v2",
	Kind:    "FaultInjectionPolicy",
}

// FaultInjectionPolicy is the Schema for the faultInjectionPolicy API
type FaultInjectionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FaultInjectionPolicySpec   `json:"spec,omitempty"`
	Status FaultInjectionPolicyStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (FaultInjectionPolicy) GVK() schema.GroupVersionKind {
	return FaultInjectionPolicyGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FaultInjectionPolicyList contains a list of FaultInjectionPolicy
type FaultInjectionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FaultInjectionPolicy `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for RetryTimeoutPolicy
var RetryTimeoutPolicyGVK = schema.GroupVersionKind{
	Group:   "resilience.policy.gloo.solo.io",
	Version: "v2",
	Kind:    "RetryTimeoutPolicy",
}

// RetryTimeoutPolicy is the Schema for the retryTimeoutPolicy API
type RetryTimeoutPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RetryTimeoutPolicySpec   `json:"spec,omitempty"`
	Status RetryTimeoutPolicyStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (RetryTimeoutPolicy) GVK() schema.GroupVersionKind {
	return RetryTimeoutPolicyGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RetryTimeoutPolicyList contains a list of RetryTimeoutPolicy
type RetryTimeoutPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RetryTimeoutPolicy `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for ConnectionPolicy
var ConnectionPolicyGVK = schema.GroupVersionKind{
	Group:   "resilience.policy.gloo.solo.io",
	Version: "v2",
	Kind:    "ConnectionPolicy",
}

// ConnectionPolicy is the Schema for the connectionPolicy API
type ConnectionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConnectionPolicySpec   `json:"spec,omitempty"`
	Status ConnectionPolicyStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (ConnectionPolicy) GVK() schema.GroupVersionKind {
	return ConnectionPolicyGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConnectionPolicyList contains a list of ConnectionPolicy
type ConnectionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConnectionPolicy `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for TrimProxyConfigPolicy
var TrimProxyConfigPolicyGVK = schema.GroupVersionKind{
	Group:   "resilience.policy.gloo.solo.io",
	Version: "v2",
	Kind:    "TrimProxyConfigPolicy",
}

// TrimProxyConfigPolicy is the Schema for the trimProxyConfigPolicy API
type TrimProxyConfigPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TrimProxyConfigPolicySpec   `json:"spec,omitempty"`
	Status TrimProxyConfigPolicyStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (TrimProxyConfigPolicy) GVK() schema.GroupVersionKind {
	return TrimProxyConfigPolicyGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TrimProxyConfigPolicyList contains a list of TrimProxyConfigPolicy
type TrimProxyConfigPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrimProxyConfigPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GraphQLPersistedQueryCachePolicy{}, &GraphQLPersistedQueryCachePolicyList{})
	SchemeBuilder.Register(&FailoverPolicy{}, &FailoverPolicyList{})
	SchemeBuilder.Register(&OutlierDetectionPolicy{}, &OutlierDetectionPolicyList{})
	SchemeBuilder.Register(&FaultInjectionPolicy{}, &FaultInjectionPolicyList{})
	SchemeBuilder.Register(&RetryTimeoutPolicy{}, &RetryTimeoutPolicyList{})
	SchemeBuilder.Register(&ConnectionPolicy{}, &ConnectionPolicyList{})
	SchemeBuilder.Register(&TrimProxyConfigPolicy{}, &TrimProxyConfigPolicyList{})
}
