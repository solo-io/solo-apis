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

// GroupVersionKind for ExternalService
var ExternalServiceGVK = schema.GroupVersionKind{
	Group:   "networking.gloo.solo.io",
	Version: "v2",
	Kind:    "ExternalService",
}

// ExternalService is the Schema for the externalService API
type ExternalService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExternalServiceSpec   `json:"spec,omitempty"`
	Status ExternalServiceStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (ExternalService) GVK() schema.GroupVersionKind {
	return ExternalServiceGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExternalServiceList contains a list of ExternalService
type ExternalServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExternalService `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for ExternalEndpoint
var ExternalEndpointGVK = schema.GroupVersionKind{
	Group:   "networking.gloo.solo.io",
	Version: "v2",
	Kind:    "ExternalEndpoint",
}

// ExternalEndpoint is the Schema for the externalEndpoint API
type ExternalEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExternalEndpointSpec   `json:"spec,omitempty"`
	Status ExternalEndpointStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (ExternalEndpoint) GVK() schema.GroupVersionKind {
	return ExternalEndpointGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExternalEndpointList contains a list of ExternalEndpoint
type ExternalEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExternalEndpoint `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for RouteTable
var RouteTableGVK = schema.GroupVersionKind{
	Group:   "networking.gloo.solo.io",
	Version: "v2",
	Kind:    "RouteTable",
}

// RouteTable is the Schema for the routeTable API
type RouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RouteTableSpec   `json:"spec,omitempty"`
	Status RouteTableStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (RouteTable) GVK() schema.GroupVersionKind {
	return RouteTableGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RouteTableList contains a list of RouteTable
type RouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTable `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for VirtualDestination
var VirtualDestinationGVK = schema.GroupVersionKind{
	Group:   "networking.gloo.solo.io",
	Version: "v2",
	Kind:    "VirtualDestination",
}

// VirtualDestination is the Schema for the virtualDestination API
type VirtualDestination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualDestinationSpec   `json:"spec,omitempty"`
	Status VirtualDestinationStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (VirtualDestination) GVK() schema.GroupVersionKind {
	return VirtualDestinationGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualDestinationList contains a list of VirtualDestination
type VirtualDestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualDestination `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for VirtualGateway
var VirtualGatewayGVK = schema.GroupVersionKind{
	Group:   "networking.gloo.solo.io",
	Version: "v2",
	Kind:    "VirtualGateway",
}

// VirtualGateway is the Schema for the virtualGateway API
type VirtualGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualGatewaySpec   `json:"spec,omitempty"`
	Status VirtualGatewayStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (VirtualGateway) GVK() schema.GroupVersionKind {
	return VirtualGatewayGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualGatewayList contains a list of VirtualGateway
type VirtualGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualGateway `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ExternalService{}, &ExternalServiceList{})
	SchemeBuilder.Register(&ExternalEndpoint{}, &ExternalEndpointList{})
	SchemeBuilder.Register(&RouteTable{}, &RouteTableList{})
	SchemeBuilder.Register(&VirtualDestination{}, &VirtualDestinationList{})
	SchemeBuilder.Register(&VirtualGateway{}, &VirtualGatewayList{})
}