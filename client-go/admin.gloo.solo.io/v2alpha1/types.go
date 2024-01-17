// Code generated by skv2. DO NOT EDIT.

// Definitions for the Kubernetes types
package v2alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for WaypointLifecycleManager
var WaypointLifecycleManagerGVK = schema.GroupVersionKind{
	Group:   "admin.gloo.solo.io",
	Version: "v2alpha1",
	Kind:    "WaypointLifecycleManager",
}

// WaypointLifecycleManager is the Schema for the waypointLifecycleManager API
type WaypointLifecycleManager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WaypointLifecycleManagerSpec   `json:"spec,omitempty"`
	Status WaypointLifecycleManagerStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (WaypointLifecycleManager) GVK() schema.GroupVersionKind {
	return WaypointLifecycleManagerGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WaypointLifecycleManagerList contains a list of WaypointLifecycleManager
type WaypointLifecycleManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WaypointLifecycleManager `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for InsightsConfig
var InsightsConfigGVK = schema.GroupVersionKind{
	Group:   "admin.gloo.solo.io",
	Version: "v2alpha1",
	Kind:    "InsightsConfig",
}

// InsightsConfig is the Schema for the insightsConfig API
type InsightsConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InsightsConfigSpec   `json:"spec,omitempty"`
	Status InsightsConfigStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (InsightsConfig) GVK() schema.GroupVersionKind {
	return InsightsConfigGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// InsightsConfigList contains a list of InsightsConfig
type InsightsConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InsightsConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WaypointLifecycleManager{}, &WaypointLifecycleManagerList{})
	SchemeBuilder.Register(&InsightsConfig{}, &InsightsConfigList{})
}
