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

// GroupVersionKind for IssuedCertificate
var IssuedCertificateGVK = schema.GroupVersionKind{
	Group:   "internal.gloo.solo.io",
	Version: "v2",
	Kind:    "IssuedCertificate",
}

// IssuedCertificate is the Schema for the issuedCertificate API
type IssuedCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IssuedCertificateSpec   `json:"spec,omitempty"`
	Status IssuedCertificateStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (IssuedCertificate) GVK() schema.GroupVersionKind {
	return IssuedCertificateGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IssuedCertificateList contains a list of IssuedCertificate
type IssuedCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IssuedCertificate `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for CertificateRequest
var CertificateRequestGVK = schema.GroupVersionKind{
	Group:   "internal.gloo.solo.io",
	Version: "v2",
	Kind:    "CertificateRequest",
}

// CertificateRequest is the Schema for the certificateRequest API
type CertificateRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CertificateRequestSpec   `json:"spec,omitempty"`
	Status CertificateRequestStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (CertificateRequest) GVK() schema.GroupVersionKind {
	return CertificateRequestGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CertificateRequestList contains a list of CertificateRequest
type CertificateRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateRequest `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for PodBounceDirective
var PodBounceDirectiveGVK = schema.GroupVersionKind{
	Group:   "internal.gloo.solo.io",
	Version: "v2",
	Kind:    "PodBounceDirective",
}

// PodBounceDirective is the Schema for the podBounceDirective API
type PodBounceDirective struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PodBounceDirectiveSpec   `json:"spec,omitempty"`
	Status PodBounceDirectiveStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (PodBounceDirective) GVK() schema.GroupVersionKind {
	return PodBounceDirectiveGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodBounceDirectiveList contains a list of PodBounceDirective
type PodBounceDirectiveList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodBounceDirective `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for XdsConfig
var XdsConfigGVK = schema.GroupVersionKind{
	Group:   "internal.gloo.solo.io",
	Version: "v2",
	Kind:    "XdsConfig",
}

// XdsConfig is the Schema for the xdsConfig API
type XdsConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   XdsConfigSpec   `json:"spec,omitempty"`
	Status XdsConfigStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (XdsConfig) GVK() schema.GroupVersionKind {
	return XdsConfigGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// XdsConfigList contains a list of XdsConfig
type XdsConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []XdsConfig `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for DiscoveredGateway
var DiscoveredGatewayGVK = schema.GroupVersionKind{
	Group:   "internal.gloo.solo.io",
	Version: "v2",
	Kind:    "DiscoveredGateway",
}

// DiscoveredGateway is the Schema for the discoveredGateway API
type DiscoveredGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DiscoveredGatewaySpec   `json:"spec,omitempty"`
	Status DiscoveredGatewayStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (DiscoveredGateway) GVK() schema.GroupVersionKind {
	return DiscoveredGatewayGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DiscoveredGatewayList contains a list of DiscoveredGateway
type DiscoveredGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiscoveredGateway `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for Mesh
var MeshGVK = schema.GroupVersionKind{
	Group:   "internal.gloo.solo.io",
	Version: "v2",
	Kind:    "Mesh",
}

// Mesh is the Schema for the mesh API
type Mesh struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MeshSpec   `json:"spec,omitempty"`
	Status MeshStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (Mesh) GVK() schema.GroupVersionKind {
	return MeshGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MeshList contains a list of Mesh
type MeshList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Mesh `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for DiscoveredCNI
var DiscoveredCNIGVK = schema.GroupVersionKind{
	Group:   "internal.gloo.solo.io",
	Version: "v2",
	Kind:    "DiscoveredCNI",
}

// DiscoveredCNI is the Schema for the discoveredCNI API
type DiscoveredCNI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DiscoveredCNISpec   `json:"spec,omitempty"`
	Status DiscoveredCNIStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (DiscoveredCNI) GVK() schema.GroupVersionKind {
	return DiscoveredCNIGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DiscoveredCNIList contains a list of DiscoveredCNI
type DiscoveredCNIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiscoveredCNI `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for PortalConfig
var PortalConfigGVK = schema.GroupVersionKind{
	Group:   "internal.gloo.solo.io",
	Version: "v2",
	Kind:    "PortalConfig",
}

// PortalConfig is the Schema for the portalConfig API
type PortalConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PortalConfigSpec   `json:"spec,omitempty"`
	Status PortalConfigStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (PortalConfig) GVK() schema.GroupVersionKind {
	return PortalConfigGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PortalConfigList contains a list of PortalConfig
type PortalConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PortalConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IssuedCertificate{}, &IssuedCertificateList{})
	SchemeBuilder.Register(&CertificateRequest{}, &CertificateRequestList{})
	SchemeBuilder.Register(&PodBounceDirective{}, &PodBounceDirectiveList{})
	SchemeBuilder.Register(&XdsConfig{}, &XdsConfigList{})
	SchemeBuilder.Register(&DiscoveredGateway{}, &DiscoveredGatewayList{})
	SchemeBuilder.Register(&Mesh{}, &MeshList{})
	SchemeBuilder.Register(&DiscoveredCNI{}, &DiscoveredCNIList{})
	SchemeBuilder.Register(&PortalConfig{}, &PortalConfigList{})
}
