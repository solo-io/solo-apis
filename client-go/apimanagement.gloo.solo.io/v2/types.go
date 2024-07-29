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

// GroupVersionKind for GraphQLStitchedSchema
var GraphQLStitchedSchemaGVK = schema.GroupVersionKind{
	Group:   "apimanagement.gloo.solo.io",
	Version: "v2",
	Kind:    "GraphQLStitchedSchema",
}

// GraphQLStitchedSchema is the Schema for the graphQLStitchedSchema API
type GraphQLStitchedSchema struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GraphQLStitchedSchemaSpec   `json:"spec,omitempty"`
	Status GraphQLStitchedSchemaStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (GraphQLStitchedSchema) GVK() schema.GroupVersionKind {
	return GraphQLStitchedSchemaGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GraphQLStitchedSchemaList contains a list of GraphQLStitchedSchema
type GraphQLStitchedSchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GraphQLStitchedSchema `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for GraphQLResolverMap
var GraphQLResolverMapGVK = schema.GroupVersionKind{
	Group:   "apimanagement.gloo.solo.io",
	Version: "v2",
	Kind:    "GraphQLResolverMap",
}

// GraphQLResolverMap is the Schema for the graphQLResolverMap API
type GraphQLResolverMap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GraphQLResolverMapSpec   `json:"spec,omitempty"`
	Status GraphQLResolverMapStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (GraphQLResolverMap) GVK() schema.GroupVersionKind {
	return GraphQLResolverMapGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GraphQLResolverMapList contains a list of GraphQLResolverMap
type GraphQLResolverMapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GraphQLResolverMap `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for GraphQLSchema
var GraphQLSchemaGVK = schema.GroupVersionKind{
	Group:   "apimanagement.gloo.solo.io",
	Version: "v2",
	Kind:    "GraphQLSchema",
}

// GraphQLSchema is the Schema for the graphQLSchema API
type GraphQLSchema struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GraphQLSchemaSpec   `json:"spec,omitempty"`
	Status GraphQLSchemaStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (GraphQLSchema) GVK() schema.GroupVersionKind {
	return GraphQLSchemaGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GraphQLSchemaList contains a list of GraphQLSchema
type GraphQLSchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GraphQLSchema `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for ApiDoc
var ApiDocGVK = schema.GroupVersionKind{
	Group:   "apimanagement.gloo.solo.io",
	Version: "v2",
	Kind:    "ApiDoc",
}

// ApiDoc is the Schema for the apiDoc API
type ApiDoc struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiDocSpec   `json:"spec,omitempty"`
	Status ApiDocStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (ApiDoc) GVK() schema.GroupVersionKind {
	return ApiDocGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ApiDocList contains a list of ApiDoc
type ApiDocList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiDoc `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for Portal
var PortalGVK = schema.GroupVersionKind{
	Group:   "apimanagement.gloo.solo.io",
	Version: "v2",
	Kind:    "Portal",
}

// Portal is the Schema for the portal API
type Portal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PortalSpec   `json:"spec,omitempty"`
	Status PortalStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (Portal) GVK() schema.GroupVersionKind {
	return PortalGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PortalList contains a list of Portal
type PortalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Portal `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for PortalGroup
var PortalGroupGVK = schema.GroupVersionKind{
	Group:   "apimanagement.gloo.solo.io",
	Version: "v2",
	Kind:    "PortalGroup",
}

// PortalGroup is the Schema for the portalGroup API
type PortalGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PortalGroupSpec   `json:"spec,omitempty"`
	Status PortalGroupStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (PortalGroup) GVK() schema.GroupVersionKind {
	return PortalGroupGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PortalGroupList contains a list of PortalGroup
type PortalGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PortalGroup `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for ApiSchemaDiscovery
var ApiSchemaDiscoveryGVK = schema.GroupVersionKind{
	Group:   "apimanagement.gloo.solo.io",
	Version: "v2",
	Kind:    "ApiSchemaDiscovery",
}

// ApiSchemaDiscovery is the Schema for the apiSchemaDiscovery API
type ApiSchemaDiscovery struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiSchemaDiscoverySpec   `json:"spec,omitempty"`
	Status ApiSchemaDiscoveryStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (ApiSchemaDiscovery) GVK() schema.GroupVersionKind {
	return ApiSchemaDiscoveryGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ApiSchemaDiscoveryList contains a list of ApiSchemaDiscovery
type ApiSchemaDiscoveryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiSchemaDiscovery `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GraphQLStitchedSchema{}, &GraphQLStitchedSchemaList{})
	SchemeBuilder.Register(&GraphQLResolverMap{}, &GraphQLResolverMapList{})
	SchemeBuilder.Register(&GraphQLSchema{}, &GraphQLSchemaList{})
	SchemeBuilder.Register(&ApiDoc{}, &ApiDocList{})
	SchemeBuilder.Register(&Portal{}, &PortalList{})
	SchemeBuilder.Register(&PortalGroup{}, &PortalGroupList{})
	SchemeBuilder.Register(&ApiSchemaDiscovery{}, &ApiSchemaDiscoveryList{})
}