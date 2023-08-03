// Code generated by skv2. DO NOT EDIT.

// Definitions for the Kubernetes types
package v1

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/runtime/schema")

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for Gateway
var GatewayGVK = schema.GroupVersionKind{
    Group: "gateway.solo.io",
    Version: "v1",
    Kind: "Gateway",
}

// Gateway is the Schema for the gateway API
type Gateway struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec GatewaySpec `json:"spec,omitempty"`
    Status GatewayStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (Gateway)  GVK() schema.GroupVersionKind {
	return GatewayGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GatewayList contains a list of Gateway
type GatewayList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []Gateway `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for MatchableHttpGateway
var MatchableHttpGatewayGVK = schema.GroupVersionKind{
    Group: "gateway.solo.io",
    Version: "v1",
    Kind: "MatchableHttpGateway",
}

// MatchableHttpGateway is the Schema for the matchableHttpGateway API
type MatchableHttpGateway struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec MatchableHttpGatewaySpec `json:"spec,omitempty"`
    Status MatchableHttpGatewayStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (MatchableHttpGateway)  GVK() schema.GroupVersionKind {
	return MatchableHttpGatewayGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MatchableHttpGatewayList contains a list of MatchableHttpGateway
type MatchableHttpGatewayList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []MatchableHttpGateway `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for RouteTable
var RouteTableGVK = schema.GroupVersionKind{
    Group: "gateway.solo.io",
    Version: "v1",
    Kind: "RouteTable",
}

// RouteTable is the Schema for the routeTable API
type RouteTable struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec RouteTableSpec `json:"spec,omitempty"`
    Status RouteTableStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (RouteTable)  GVK() schema.GroupVersionKind {
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

// GroupVersionKind for VirtualService
var VirtualServiceGVK = schema.GroupVersionKind{
    Group: "gateway.solo.io",
    Version: "v1",
    Kind: "VirtualService",
}

// VirtualService is the Schema for the virtualService API
type VirtualService struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec VirtualServiceSpec `json:"spec,omitempty"`
    Status VirtualServiceStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (VirtualService)  GVK() schema.GroupVersionKind {
	return VirtualServiceGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualServiceList contains a list of VirtualService
type VirtualServiceList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []VirtualService `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for VirtualHostOption
var VirtualHostOptionGVK = schema.GroupVersionKind{
    Group: "gateway.solo.io",
    Version: "v1",
    Kind: "VirtualHostOption",
}

// VirtualHostOption is the Schema for the virtualHostOption API
type VirtualHostOption struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec VirtualHostOptionSpec `json:"spec,omitempty"`
    Status VirtualHostOptionStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (VirtualHostOption)  GVK() schema.GroupVersionKind {
	return VirtualHostOptionGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualHostOptionList contains a list of VirtualHostOption
type VirtualHostOptionList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []VirtualHostOption `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for RouteOption
var RouteOptionGVK = schema.GroupVersionKind{
    Group: "gateway.solo.io",
    Version: "v1",
    Kind: "RouteOption",
}

// RouteOption is the Schema for the routeOption API
type RouteOption struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec RouteOptionSpec `json:"spec,omitempty"`
    Status RouteOptionStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (RouteOption)  GVK() schema.GroupVersionKind {
	return RouteOptionGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RouteOptionList contains a list of RouteOption
type RouteOptionList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []RouteOption `json:"items"`
}

func init() {
    SchemeBuilder.Register(&Gateway{}, &GatewayList{})
    SchemeBuilder.Register(&MatchableHttpGateway{}, &MatchableHttpGatewayList{})
    SchemeBuilder.Register(&RouteTable{}, &RouteTableList{})
    SchemeBuilder.Register(&VirtualService{}, &VirtualServiceList{})
    SchemeBuilder.Register(&VirtualHostOption{}, &VirtualHostOptionList{})
    SchemeBuilder.Register(&RouteOption{}, &RouteOptionList{})
}
