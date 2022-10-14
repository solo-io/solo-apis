// Code generated by skv2. DO NOT EDIT.

// Definitions for the Kubernetes types
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for Settings
var SettingsGVK = schema.GroupVersionKind{
	Group:   "gloo.solo.io",
	Version: "v1",
	Kind:    "Settings",
}

// Settings is the Schema for the settings API
type Settings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SettingsSpec   `json:"spec,omitempty"`
	Status SettingsStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (Settings) GVK() schema.GroupVersionKind {
	return SettingsGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SettingsList contains a list of Settings
type SettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Settings `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for Upstream
var UpstreamGVK = schema.GroupVersionKind{
	Group:   "gloo.solo.io",
	Version: "v1",
	Kind:    "Upstream",
}

// Upstream is the Schema for the upstream API
type Upstream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UpstreamSpec   `json:"spec,omitempty"`
	Status UpstreamStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (Upstream) GVK() schema.GroupVersionKind {
	return UpstreamGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UpstreamList contains a list of Upstream
type UpstreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Upstream `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for UpstreamGroup
var UpstreamGroupGVK = schema.GroupVersionKind{
	Group:   "gloo.solo.io",
	Version: "v1",
	Kind:    "UpstreamGroup",
}

// UpstreamGroup is the Schema for the upstreamGroup API
type UpstreamGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UpstreamGroupSpec   `json:"spec,omitempty"`
	Status UpstreamGroupStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (UpstreamGroup) GVK() schema.GroupVersionKind {
	return UpstreamGroupGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UpstreamGroupList contains a list of UpstreamGroup
type UpstreamGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UpstreamGroup `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for Proxy
var ProxyGVK = schema.GroupVersionKind{
	Group:   "gloo.solo.io",
	Version: "v1",
	Kind:    "Proxy",
}

// Proxy is the Schema for the proxy API
type Proxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProxySpec   `json:"spec,omitempty"`
	Status ProxyStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (Proxy) GVK() schema.GroupVersionKind {
	return ProxyGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProxyList contains a list of Proxy
type ProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Proxy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Settings{}, &SettingsList{})
	SchemeBuilder.Register(&Upstream{}, &UpstreamList{})
	SchemeBuilder.Register(&UpstreamGroup{}, &UpstreamGroupList{})
	SchemeBuilder.Register(&Proxy{}, &ProxyList{})
}