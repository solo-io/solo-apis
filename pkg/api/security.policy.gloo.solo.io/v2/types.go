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

// GroupVersionKind for AccessPolicy
var AccessPolicyGVK = schema.GroupVersionKind{
	Group:   "security.policy.gloo.solo.io",
	Version: "v2",
	Kind:    "AccessPolicy",
}

// AccessPolicy is the Schema for the accessPolicy API
type AccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AccessPolicySpec   `json:"spec,omitempty"`
	Status AccessPolicyStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (AccessPolicy) GVK() schema.GroupVersionKind {
	return AccessPolicyGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AccessPolicyList contains a list of AccessPolicy
type AccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessPolicy `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for CORSPolicy
var CORSPolicyGVK = schema.GroupVersionKind{
	Group:   "security.policy.gloo.solo.io",
	Version: "v2",
	Kind:    "CORSPolicy",
}

// CORSPolicy is the Schema for the cORSPolicy API
type CORSPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CORSPolicySpec   `json:"spec,omitempty"`
	Status CORSPolicyStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (CORSPolicy) GVK() schema.GroupVersionKind {
	return CORSPolicyGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CORSPolicyList contains a list of CORSPolicy
type CORSPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CORSPolicy `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for CSRFPolicy
var CSRFPolicyGVK = schema.GroupVersionKind{
	Group:   "security.policy.gloo.solo.io",
	Version: "v2",
	Kind:    "CSRFPolicy",
}

// CSRFPolicy is the Schema for the cSRFPolicy API
type CSRFPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CSRFPolicySpec   `json:"spec,omitempty"`
	Status CSRFPolicyStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (CSRFPolicy) GVK() schema.GroupVersionKind {
	return CSRFPolicyGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CSRFPolicyList contains a list of CSRFPolicy
type CSRFPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CSRFPolicy `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for ExtAuthPolicy
var ExtAuthPolicyGVK = schema.GroupVersionKind{
	Group:   "security.policy.gloo.solo.io",
	Version: "v2",
	Kind:    "ExtAuthPolicy",
}

// ExtAuthPolicy is the Schema for the extAuthPolicy API
type ExtAuthPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExtAuthPolicySpec   `json:"spec,omitempty"`
	Status ExtAuthPolicyStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (ExtAuthPolicy) GVK() schema.GroupVersionKind {
	return ExtAuthPolicyGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExtAuthPolicyList contains a list of ExtAuthPolicy
type ExtAuthPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExtAuthPolicy `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for WAFPolicy
var WAFPolicyGVK = schema.GroupVersionKind{
	Group:   "security.policy.gloo.solo.io",
	Version: "v2",
	Kind:    "WAFPolicy",
}

// WAFPolicy is the Schema for the wAFPolicy API
type WAFPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WAFPolicySpec   `json:"spec,omitempty"`
	Status WAFPolicyStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (WAFPolicy) GVK() schema.GroupVersionKind {
	return WAFPolicyGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WAFPolicyList contains a list of WAFPolicy
type WAFPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WAFPolicy `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for JWTPolicy
var JWTPolicyGVK = schema.GroupVersionKind{
	Group:   "security.policy.gloo.solo.io",
	Version: "v2",
	Kind:    "JWTPolicy",
}

// JWTPolicy is the Schema for the jWTPolicy API
type JWTPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   JWTPolicySpec   `json:"spec,omitempty"`
	Status JWTPolicyStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (JWTPolicy) GVK() schema.GroupVersionKind {
	return JWTPolicyGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// JWTPolicyList contains a list of JWTPolicy
type JWTPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JWTPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AccessPolicy{}, &AccessPolicyList{})
	SchemeBuilder.Register(&CORSPolicy{}, &CORSPolicyList{})
	SchemeBuilder.Register(&CSRFPolicy{}, &CSRFPolicyList{})
	SchemeBuilder.Register(&ExtAuthPolicy{}, &ExtAuthPolicyList{})
	SchemeBuilder.Register(&WAFPolicy{}, &WAFPolicyList{})
	SchemeBuilder.Register(&JWTPolicy{}, &JWTPolicyList{})
}
