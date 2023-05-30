// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/policy/v2/security/client_tls_policy.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	v2 "github.com/solo-io/solo-apis/client-go/common.gloo.solo.io/v2"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ClientTLSPolicy explicitly controls the TLS/mTLS configuration for upstream connections
type ClientTLSPolicySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplyToDestinations []*v2.DestinationSelector `protobuf:"bytes,1,rep,name=apply_to_destinations,json=applyToDestinations,proto3" json:"apply_to_destinations,omitempty"`
	// Type of TLS used when establishing connections to the selected destinations.
	// Note that this mode will override any inherited TLS configuration when connecting to the destination,
	// for example it will override a mesh-wide mTLS configured via Istio PeerAuthentication.
	//
	// Types that are assignable to Mode:
	//
	//	*ClientTLSPolicySpec_Disable_
	//	*ClientTLSPolicySpec_Simple_
	//	*ClientTLSPolicySpec_Mutual_
	//	*ClientTLSPolicySpec_IstioMutual_
	Mode isClientTLSPolicySpec_Mode `protobuf_oneof:"mode"`
}

func (x *ClientTLSPolicySpec) Reset() {
	*x = ClientTLSPolicySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientTLSPolicySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientTLSPolicySpec) ProtoMessage() {}

func (x *ClientTLSPolicySpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientTLSPolicySpec.ProtoReflect.Descriptor instead.
func (*ClientTLSPolicySpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_rawDescGZIP(), []int{0}
}

func (x *ClientTLSPolicySpec) GetApplyToDestinations() []*v2.DestinationSelector {
	if x != nil {
		return x.ApplyToDestinations
	}
	return nil
}

func (m *ClientTLSPolicySpec) GetMode() isClientTLSPolicySpec_Mode {
	if m != nil {
		return m.Mode
	}
	return nil
}

func (x *ClientTLSPolicySpec) GetDisable() *ClientTLSPolicySpec_Disable {
	if x, ok := x.GetMode().(*ClientTLSPolicySpec_Disable_); ok {
		return x.Disable
	}
	return nil
}

func (x *ClientTLSPolicySpec) GetSimple() *ClientTLSPolicySpec_Simple {
	if x, ok := x.GetMode().(*ClientTLSPolicySpec_Simple_); ok {
		return x.Simple
	}
	return nil
}

func (x *ClientTLSPolicySpec) GetMutual() *ClientTLSPolicySpec_Mutual {
	if x, ok := x.GetMode().(*ClientTLSPolicySpec_Mutual_); ok {
		return x.Mutual
	}
	return nil
}

func (x *ClientTLSPolicySpec) GetIstioMutual() *ClientTLSPolicySpec_IstioMutual {
	if x, ok := x.GetMode().(*ClientTLSPolicySpec_IstioMutual_); ok {
		return x.IstioMutual
	}
	return nil
}

type isClientTLSPolicySpec_Mode interface {
	isClientTLSPolicySpec_Mode()
}

type ClientTLSPolicySpec_Disable_ struct {
	Disable *ClientTLSPolicySpec_Disable `protobuf:"bytes,2,opt,name=disable,proto3,oneof"`
}

type ClientTLSPolicySpec_Simple_ struct {
	Simple *ClientTLSPolicySpec_Simple `protobuf:"bytes,3,opt,name=simple,proto3,oneof"`
}

type ClientTLSPolicySpec_Mutual_ struct {
	Mutual *ClientTLSPolicySpec_Mutual `protobuf:"bytes,4,opt,name=mutual,proto3,oneof"`
}

type ClientTLSPolicySpec_IstioMutual_ struct {
	IstioMutual *ClientTLSPolicySpec_IstioMutual `protobuf:"bytes,5,opt,name=istioMutual,proto3,oneof"`
}

func (*ClientTLSPolicySpec_Disable_) isClientTLSPolicySpec_Mode() {}

func (*ClientTLSPolicySpec_Simple_) isClientTLSPolicySpec_Mode() {}

func (*ClientTLSPolicySpec_Mutual_) isClientTLSPolicySpec_Mode() {}

func (*ClientTLSPolicySpec_IstioMutual_) isClientTLSPolicySpec_Mode() {}

type TLSConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SNI string to present to the server during TLS handshake.
	// Recommended to be set, however if omitted, the first hostname associated with the destination
	// will be used
	Sni *wrappers.StringValue `protobuf:"bytes,1,opt,name=sni,proto3" json:"sni,omitempty"`
	// The name of the secret that holds the TLS certs for the client including the CA certificates.
	// Secret must exist in the same namespace with the proxy using the certificates.
	// The secret (of type generic) should contain the following keys and values:
	// key: <privateKey>, cert: <clientCert>, cacert: <CACertificate>.
	// Here CACertificate is used to verify the server certificate.
	// Secret of type tls for client certificates along with ca.crt key for CA certificates is also supported.
	// see: https://istio.io/latest/docs/reference/config/networking/destination-rule/#ClientTLSSettings
	CredentialName string `protobuf:"bytes,2,opt,name=credentialName,proto3" json:"credentialName,omitempty"`
}

func (x *TLSConfig) Reset() {
	*x = TLSConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TLSConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLSConfig) ProtoMessage() {}

func (x *TLSConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLSConfig.ProtoReflect.Descriptor instead.
func (*TLSConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_rawDescGZIP(), []int{1}
}

func (x *TLSConfig) GetSni() *wrappers.StringValue {
	if x != nil {
		return x.Sni
	}
	return nil
}

func (x *TLSConfig) GetCredentialName() string {
	if x != nil {
		return x.CredentialName
	}
	return ""
}

type ClientTLSPolicyStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Common *v2.Status `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	// The number of destination ports selected by the policy.
	NumSelectedDestinationPorts uint32 `protobuf:"varint,2,opt,name=num_selected_destination_ports,json=numSelectedDestinationPorts,proto3" json:"num_selected_destination_ports,omitempty"`
}

func (x *ClientTLSPolicyStatus) Reset() {
	*x = ClientTLSPolicyStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientTLSPolicyStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientTLSPolicyStatus) ProtoMessage() {}

func (x *ClientTLSPolicyStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientTLSPolicyStatus.ProtoReflect.Descriptor instead.
func (*ClientTLSPolicyStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_rawDescGZIP(), []int{2}
}

func (x *ClientTLSPolicyStatus) GetCommon() *v2.Status {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *ClientTLSPolicyStatus) GetNumSelectedDestinationPorts() uint32 {
	if x != nil {
		return x.NumSelectedDestinationPorts
	}
	return 0
}

type ClientTLSPolicyReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status of the resource in each workspace that it exists in.
	Workspaces map[string]*v2.Report `protobuf:"bytes,1,rep,name=workspaces,proto3" json:"workspaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A list of destination ports selected by the policy.
	SelectedDestinationPorts []*v2.DestinationReference `protobuf:"bytes,2,rep,name=selected_destination_ports,json=selectedDestinationPorts,proto3" json:"selected_destination_ports,omitempty"`
}

func (x *ClientTLSPolicyReport) Reset() {
	*x = ClientTLSPolicyReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientTLSPolicyReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientTLSPolicyReport) ProtoMessage() {}

func (x *ClientTLSPolicyReport) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientTLSPolicyReport.ProtoReflect.Descriptor instead.
func (*ClientTLSPolicyReport) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_rawDescGZIP(), []int{3}
}

func (x *ClientTLSPolicyReport) GetWorkspaces() map[string]*v2.Report {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

func (x *ClientTLSPolicyReport) GetSelectedDestinationPorts() []*v2.DestinationReference {
	if x != nil {
		return x.SelectedDestinationPorts
	}
	return nil
}

// Explicitly do not establish a TLS connection to the destination
type ClientTLSPolicySpec_Disable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClientTLSPolicySpec_Disable) Reset() {
	*x = ClientTLSPolicySpec_Disable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientTLSPolicySpec_Disable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientTLSPolicySpec_Disable) ProtoMessage() {}

func (x *ClientTLSPolicySpec_Disable) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientTLSPolicySpec_Disable.ProtoReflect.Descriptor instead.
func (*ClientTLSPolicySpec_Disable) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_rawDescGZIP(), []int{0, 0}
}

// Initiate a mutual TLS connection using the Istio provided certificates. This is useful
// if a more broad policy/configuration has disabled Istio mTLS but you need it enabled
// for a specific destination.
type ClientTLSPolicySpec_IstioMutual struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClientTLSPolicySpec_IstioMutual) Reset() {
	*x = ClientTLSPolicySpec_IstioMutual{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientTLSPolicySpec_IstioMutual) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientTLSPolicySpec_IstioMutual) ProtoMessage() {}

func (x *ClientTLSPolicySpec_IstioMutual) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientTLSPolicySpec_IstioMutual.ProtoReflect.Descriptor instead.
func (*ClientTLSPolicySpec_IstioMutual) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_rawDescGZIP(), []int{0, 1}
}

// Initiate a mutual TLS connection and present client certificates via the provided credential/secret.
// This is separate from any TLS/mTLS provided by Istio.
type ClientTLSPolicySpec_Mutual struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *TLSConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *ClientTLSPolicySpec_Mutual) Reset() {
	*x = ClientTLSPolicySpec_Mutual{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientTLSPolicySpec_Mutual) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientTLSPolicySpec_Mutual) ProtoMessage() {}

func (x *ClientTLSPolicySpec_Mutual) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientTLSPolicySpec_Mutual.ProtoReflect.Descriptor instead.
func (*ClientTLSPolicySpec_Mutual) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_rawDescGZIP(), []int{0, 2}
}

func (x *ClientTLSPolicySpec_Mutual) GetConfig() *TLSConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

// Initiate a basic TLS connection, and possibly verify the server certificate if provided a CA via credential/secret.
// This is separate from any TLS/mTLS provided by Istio.
type ClientTLSPolicySpec_Simple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *TLSConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *ClientTLSPolicySpec_Simple) Reset() {
	*x = ClientTLSPolicySpec_Simple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientTLSPolicySpec_Simple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientTLSPolicySpec_Simple) ProtoMessage() {}

func (x *ClientTLSPolicySpec_Simple) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientTLSPolicySpec_Simple.ProtoReflect.Descriptor instead.
func (*ClientTLSPolicySpec_Simple) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_rawDescGZIP(), []int{0, 3}
}

func (x *ClientTLSPolicySpec_Simple) GetConfig() *TLSConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_rawDesc = []byte{
	0x0a, 0x63, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x6c, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x1a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73,
	0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x53, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x32, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x05, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x54, 0x4c, 0x53, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x5c,
	0x0a, 0x15, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x13, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x6f,
	0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x55, 0x0a, 0x07,
	0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x54, 0x4c, 0x53, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63,
	0x2e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x07, 0x64, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x52, 0x0a, 0x06, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x4c, 0x53, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x48, 0x00, 0x52,
	0x06, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x52, 0x0a, 0x06, 0x6d, 0x75, 0x74, 0x75, 0x61,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x4c, 0x53,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x4d, 0x75, 0x74, 0x75, 0x61,
	0x6c, 0x48, 0x00, 0x52, 0x06, 0x6d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x12, 0x61, 0x0a, 0x0b, 0x69,
	0x73, 0x74, 0x69, 0x6f, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3d, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x4c, 0x53, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53,
	0x70, 0x65, 0x63, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x48,
	0x00, 0x52, 0x0b, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x1a, 0x09,
	0x0a, 0x07, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x1a, 0x0d, 0x0a, 0x0b, 0x49, 0x73, 0x74,
	0x69, 0x6f, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x1a, 0x49, 0x0a, 0x06, 0x4d, 0x75, 0x74, 0x75,
	0x61, 0x6c, 0x12, 0x3f, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x54, 0x4c, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x1a, 0x49, 0x0a, 0x06, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x3f, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x54, 0x4c, 0x53,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x06,
	0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x63, 0x0a, 0x09, 0x54, 0x4c, 0x53, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x2e, 0x0a, 0x03, 0x73, 0x6e, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03,
	0x73, 0x6e, 0x69, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x15,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x4c, 0x53, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x1e, 0x6e, 0x75,
	0x6d, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x1b, 0x6e, 0x75, 0x6d, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x22,
	0xc1, 0x02, 0x0a, 0x15, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x4c, 0x53, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x63, 0x0a, 0x0a, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x54, 0x4c, 0x53, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x67,
	0x0a, 0x1a, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x18, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x1a, 0x5a, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0x5d, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d,
	0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x32, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5,
	0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_rawDescData = file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_goTypes = []interface{}{
	(*ClientTLSPolicySpec)(nil),             // 0: security.policy.gloo.solo.io.ClientTLSPolicySpec
	(*TLSConfig)(nil),                       // 1: security.policy.gloo.solo.io.TLSConfig
	(*ClientTLSPolicyStatus)(nil),           // 2: security.policy.gloo.solo.io.ClientTLSPolicyStatus
	(*ClientTLSPolicyReport)(nil),           // 3: security.policy.gloo.solo.io.ClientTLSPolicyReport
	(*ClientTLSPolicySpec_Disable)(nil),     // 4: security.policy.gloo.solo.io.ClientTLSPolicySpec.Disable
	(*ClientTLSPolicySpec_IstioMutual)(nil), // 5: security.policy.gloo.solo.io.ClientTLSPolicySpec.IstioMutual
	(*ClientTLSPolicySpec_Mutual)(nil),      // 6: security.policy.gloo.solo.io.ClientTLSPolicySpec.Mutual
	(*ClientTLSPolicySpec_Simple)(nil),      // 7: security.policy.gloo.solo.io.ClientTLSPolicySpec.Simple
	nil,                                     // 8: security.policy.gloo.solo.io.ClientTLSPolicyReport.WorkspacesEntry
	(*v2.DestinationSelector)(nil),          // 9: common.gloo.solo.io.DestinationSelector
	(*wrappers.StringValue)(nil),            // 10: google.protobuf.StringValue
	(*v2.Status)(nil),                       // 11: common.gloo.solo.io.Status
	(*v2.DestinationReference)(nil),         // 12: common.gloo.solo.io.DestinationReference
	(*v2.Report)(nil),                       // 13: common.gloo.solo.io.Report
}
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_depIdxs = []int32{
	9,  // 0: security.policy.gloo.solo.io.ClientTLSPolicySpec.apply_to_destinations:type_name -> common.gloo.solo.io.DestinationSelector
	4,  // 1: security.policy.gloo.solo.io.ClientTLSPolicySpec.disable:type_name -> security.policy.gloo.solo.io.ClientTLSPolicySpec.Disable
	7,  // 2: security.policy.gloo.solo.io.ClientTLSPolicySpec.simple:type_name -> security.policy.gloo.solo.io.ClientTLSPolicySpec.Simple
	6,  // 3: security.policy.gloo.solo.io.ClientTLSPolicySpec.mutual:type_name -> security.policy.gloo.solo.io.ClientTLSPolicySpec.Mutual
	5,  // 4: security.policy.gloo.solo.io.ClientTLSPolicySpec.istioMutual:type_name -> security.policy.gloo.solo.io.ClientTLSPolicySpec.IstioMutual
	10, // 5: security.policy.gloo.solo.io.TLSConfig.sni:type_name -> google.protobuf.StringValue
	11, // 6: security.policy.gloo.solo.io.ClientTLSPolicyStatus.common:type_name -> common.gloo.solo.io.Status
	8,  // 7: security.policy.gloo.solo.io.ClientTLSPolicyReport.workspaces:type_name -> security.policy.gloo.solo.io.ClientTLSPolicyReport.WorkspacesEntry
	12, // 8: security.policy.gloo.solo.io.ClientTLSPolicyReport.selected_destination_ports:type_name -> common.gloo.solo.io.DestinationReference
	1,  // 9: security.policy.gloo.solo.io.ClientTLSPolicySpec.Mutual.config:type_name -> security.policy.gloo.solo.io.TLSConfig
	1,  // 10: security.policy.gloo.solo.io.ClientTLSPolicySpec.Simple.config:type_name -> security.policy.gloo.solo.io.TLSConfig
	13, // 11: security.policy.gloo.solo.io.ClientTLSPolicyReport.WorkspacesEntry.value:type_name -> common.gloo.solo.io.Report
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_init()
}
func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_init() {
	if File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientTLSPolicySpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TLSConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientTLSPolicyStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientTLSPolicyReport); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientTLSPolicySpec_Disable); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientTLSPolicySpec_IstioMutual); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientTLSPolicySpec_Mutual); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientTLSPolicySpec_Simple); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ClientTLSPolicySpec_Disable_)(nil),
		(*ClientTLSPolicySpec_Simple_)(nil),
		(*ClientTLSPolicySpec_Mutual_)(nil),
		(*ClientTLSPolicySpec_IstioMutual_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto = out.File
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_client_tls_policy_proto_depIdxs = nil
}
