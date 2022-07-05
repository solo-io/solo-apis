// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/admin/v2/root_trust_policy.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	v2 "github.com/solo-io/solo-apis/pkg/api/common.gloo.solo.io/v2"
	tls "github.com/solo-io/solo-apis/pkg/api/security.policy.gloo.solo.io/v2/tls"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// RootTrustPolicy is used to designate the root of trust, including the trust domain and root certificates used by one or more service meshes.
// A shared RootTrustPolicy is currently required to support communication between workloads and destinations running in different meshes. In the future Gloo Mesh will support cross-mesh connectivity using a Limited Trust model (where participating meshes are permitted to use separate roots of trust).
type RootTrustPolicySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// select the meshes where the root of trust will be applied.
	// if left empty, will apply to all Meshes in the workspace.
	ApplyToMeshes []*v2.MeshSelector `protobuf:"bytes,1,rep,name=apply_to_meshes,json=applyToMeshes,proto3" json:"apply_to_meshes,omitempty"`
	// The details of the root of trust to apply to the selected meshes.
	Config *RootTrustPolicySpec_Config `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *RootTrustPolicySpec) Reset() {
	*x = RootTrustPolicySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RootTrustPolicySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RootTrustPolicySpec) ProtoMessage() {}

func (x *RootTrustPolicySpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RootTrustPolicySpec.ProtoReflect.Descriptor instead.
func (*RootTrustPolicySpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_rawDescGZIP(), []int{0}
}

func (x *RootTrustPolicySpec) GetApplyToMeshes() []*v2.MeshSelector {
	if x != nil {
		return x.ApplyToMeshes
	}
	return nil
}

func (x *RootTrustPolicySpec) GetConfig() *RootTrustPolicySpec_Config {
	if x != nil {
		return x.Config
	}
	return nil
}

// reflects the status of the RootTrustPolicy
type RootTrustPolicyStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The most recent generation observed in the the object's metadata.
	// If the `observedGeneration` does not match `metadata.generation`, Gloo Mesh has not processed the most
	// recent version of this object.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// Whether the resource has been accepted as valid and processed in the Gloo Mesh config translation.
	State v2.ApprovalState `protobuf:"varint,2,opt,name=state,proto3,enum=common.gloo.solo.io.ApprovalState" json:"state,omitempty"`
}

func (x *RootTrustPolicyStatus) Reset() {
	*x = RootTrustPolicyStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RootTrustPolicyStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RootTrustPolicyStatus) ProtoMessage() {}

func (x *RootTrustPolicyStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RootTrustPolicyStatus.ProtoReflect.Descriptor instead.
func (*RootTrustPolicyStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_rawDescGZIP(), []int{1}
}

func (x *RootTrustPolicyStatus) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *RootTrustPolicyStatus) GetState() v2.ApprovalState {
	if x != nil {
		return x.State
	}
	return v2.ApprovalState_PENDING
}

type RootTrustPolicySpec_Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// different options for choosing the CA used to provide the root cert
	//
	// Types that are assignable to CertificateAuthorityType:
	//	*RootTrustPolicySpec_Config_MgmtServerCa
	//	*RootTrustPolicySpec_Config_AgentCa
	CertificateAuthorityType isRootTrustPolicySpec_Config_CertificateAuthorityType `protobuf_oneof:"certificate_authority_type"`
	// Configuration options for generated intermediate certs.
	IntermediateCertOptions *tls.CommonCertOptions `protobuf:"bytes,3,opt,name=intermediate_cert_options,json=intermediateCertOptions,proto3" json:"intermediate_cert_options,omitempty"`
	// This setting specifies whether or not workload pods should be automatically restarted
	// upon completion of a successful certificate issuance.
	AutoRestartPods bool `protobuf:"varint,4,opt,name=auto_restart_pods,json=autoRestartPods,proto3" json:"auto_restart_pods,omitempty"`
}

func (x *RootTrustPolicySpec_Config) Reset() {
	*x = RootTrustPolicySpec_Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RootTrustPolicySpec_Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RootTrustPolicySpec_Config) ProtoMessage() {}

func (x *RootTrustPolicySpec_Config) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RootTrustPolicySpec_Config.ProtoReflect.Descriptor instead.
func (*RootTrustPolicySpec_Config) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_rawDescGZIP(), []int{0, 0}
}

func (m *RootTrustPolicySpec_Config) GetCertificateAuthorityType() isRootTrustPolicySpec_Config_CertificateAuthorityType {
	if m != nil {
		return m.CertificateAuthorityType
	}
	return nil
}

func (x *RootTrustPolicySpec_Config) GetMgmtServerCa() *RootTrustPolicySpec_Config_MgmtServerCertificateAuthority {
	if x, ok := x.GetCertificateAuthorityType().(*RootTrustPolicySpec_Config_MgmtServerCa); ok {
		return x.MgmtServerCa
	}
	return nil
}

func (x *RootTrustPolicySpec_Config) GetAgentCa() *tls.AgentCertificateAuthority {
	if x, ok := x.GetCertificateAuthorityType().(*RootTrustPolicySpec_Config_AgentCa); ok {
		return x.AgentCa
	}
	return nil
}

func (x *RootTrustPolicySpec_Config) GetIntermediateCertOptions() *tls.CommonCertOptions {
	if x != nil {
		return x.IntermediateCertOptions
	}
	return nil
}

func (x *RootTrustPolicySpec_Config) GetAutoRestartPods() bool {
	if x != nil {
		return x.AutoRestartPods
	}
	return false
}

type isRootTrustPolicySpec_Config_CertificateAuthorityType interface {
	isRootTrustPolicySpec_Config_CertificateAuthorityType()
}

type RootTrustPolicySpec_Config_MgmtServerCa struct {
	// Configure a Root Certificate Authority which will be shared by all Meshes associated with this RootTrustPolicy.
	// If this is not provided, a self-signed certificate will be generated by Gloo Mesh.
	MgmtServerCa *RootTrustPolicySpec_Config_MgmtServerCertificateAuthority `protobuf:"bytes,1,opt,name=mgmt_server_ca,json=mgmtServerCa,proto3,oneof"`
}

type RootTrustPolicySpec_Config_AgentCa struct {
	// Configures an Intermediate Certificate Authority which selected meshes will use to generate intermediate certificates.
	// The CA being used must be configured to generate the intermediate certificates.
	AgentCa *tls.AgentCertificateAuthority `protobuf:"bytes,2,opt,name=agent_ca,json=agentCa,proto3,oneof"`
}

func (*RootTrustPolicySpec_Config_MgmtServerCa) isRootTrustPolicySpec_Config_CertificateAuthorityType() {
}

func (*RootTrustPolicySpec_Config_AgentCa) isRootTrustPolicySpec_Config_CertificateAuthorityType() {}

// Specify parameters for configuring the root certificate authority for a RootTrustPolicy.
type RootTrustPolicySpec_Config_MgmtServerCertificateAuthority struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specify the source of the Root CA data which Gloo Mesh will use for the RootTrustPolicy.
	//
	// Types that are assignable to CaSource:
	//	*RootTrustPolicySpec_Config_MgmtServerCertificateAuthority_Generated
	//	*RootTrustPolicySpec_Config_MgmtServerCertificateAuthority_SecretRef
	CaSource isRootTrustPolicySpec_Config_MgmtServerCertificateAuthority_CaSource `protobuf_oneof:"ca_source"`
}

func (x *RootTrustPolicySpec_Config_MgmtServerCertificateAuthority) Reset() {
	*x = RootTrustPolicySpec_Config_MgmtServerCertificateAuthority{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RootTrustPolicySpec_Config_MgmtServerCertificateAuthority) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RootTrustPolicySpec_Config_MgmtServerCertificateAuthority) ProtoMessage() {}

func (x *RootTrustPolicySpec_Config_MgmtServerCertificateAuthority) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RootTrustPolicySpec_Config_MgmtServerCertificateAuthority.ProtoReflect.Descriptor instead.
func (*RootTrustPolicySpec_Config_MgmtServerCertificateAuthority) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (m *RootTrustPolicySpec_Config_MgmtServerCertificateAuthority) GetCaSource() isRootTrustPolicySpec_Config_MgmtServerCertificateAuthority_CaSource {
	if m != nil {
		return m.CaSource
	}
	return nil
}

func (x *RootTrustPolicySpec_Config_MgmtServerCertificateAuthority) GetGenerated() *tls.CommonCertOptions {
	if x, ok := x.GetCaSource().(*RootTrustPolicySpec_Config_MgmtServerCertificateAuthority_Generated); ok {
		return x.Generated
	}
	return nil
}

func (x *RootTrustPolicySpec_Config_MgmtServerCertificateAuthority) GetSecretRef() *v1.ObjectRef {
	if x, ok := x.GetCaSource().(*RootTrustPolicySpec_Config_MgmtServerCertificateAuthority_SecretRef); ok {
		return x.SecretRef
	}
	return nil
}

type isRootTrustPolicySpec_Config_MgmtServerCertificateAuthority_CaSource interface {
	isRootTrustPolicySpec_Config_MgmtServerCertificateAuthority_CaSource()
}

type RootTrustPolicySpec_Config_MgmtServerCertificateAuthority_Generated struct {
	// Generate a self-signed root certificate with the given options.
	Generated *tls.CommonCertOptions `protobuf:"bytes,1,opt,name=generated,proto3,oneof"`
}

type RootTrustPolicySpec_Config_MgmtServerCertificateAuthority_SecretRef struct {
	// Name of a Kubernetes Secret in the same namespace as the RootTrustPolicy containing the root certificate authority.
	// Provided certificates must conform to a specified format, [documented here]({{< versioned_link_path fromRoot="/setup/prod/certs/cert-architecture/#required-certificates" >}}).
	SecretRef *v1.ObjectRef `protobuf:"bytes,2,opt,name=secret_ref,json=secretRef,proto3,oneof"`
}

func (*RootTrustPolicySpec_Config_MgmtServerCertificateAuthority_Generated) isRootTrustPolicySpec_Config_MgmtServerCertificateAuthority_CaSource() {
}

func (*RootTrustPolicySpec_Config_MgmtServerCertificateAuthority_SecretRef) isRootTrustPolicySpec_Config_MgmtServerCertificateAuthority_CaSource() {
}

var File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_rawDesc = []byte{
	0x0a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x76, 0x32, 0x2f, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x5f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76,
	0x32, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x74, 0x6c, 0x73, 0x2f, 0x63,
	0x61, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x32, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83,
	0x06, 0x0a, 0x13, 0x52, 0x6f, 0x6f, 0x74, 0x54, 0x72, 0x75, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x49, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f,
	0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x4d, 0x65, 0x73, 0x68, 0x65,
	0x73, 0x12, 0x46, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x6f, 0x6f, 0x74, 0x54, 0x72, 0x75, 0x73, 0x74,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0xd8, 0x04, 0x0a, 0x06, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x75, 0x0a, 0x0e, 0x6d, 0x67, 0x6d, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4d, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x52, 0x6f, 0x6f, 0x74, 0x54, 0x72, 0x75, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d, 0x67, 0x6d,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x48, 0x00, 0x52, 0x0c, 0x6d,
	0x67, 0x6d, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x61, 0x12, 0x58, 0x0a, 0x08, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e,
	0x74, 0x6c, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x48, 0x00, 0x52, 0x07, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x43, 0x61, 0x12, 0x6f, 0x0a, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x17, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x43, 0x65, 0x72, 0x74, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x6f, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f,
	0x64, 0x73, 0x1a, 0xc1, 0x01, 0x0a, 0x1e, 0x4d, 0x67, 0x6d, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x53, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52,
	0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x12, 0x3d, 0x0a, 0x0a, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x48, 0x00, 0x52, 0x09,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x66, 0x42, 0x0b, 0x0a, 0x09, 0x63, 0x61, 0x5f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x1c, 0x0a, 0x1a, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x15, 0x52, 0x6f, 0x6f, 0x74, 0x54, 0x72, 0x75,
	0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f,
	0x0a, 0x13, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6f, 0x62, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x38, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x48, 0x5a, 0x3a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x32, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0,
	0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_goTypes = []interface{}{
	(*RootTrustPolicySpec)(nil),                                       // 0: admin.gloo.solo.io.RootTrustPolicySpec
	(*RootTrustPolicyStatus)(nil),                                     // 1: admin.gloo.solo.io.RootTrustPolicyStatus
	(*RootTrustPolicySpec_Config)(nil),                                // 2: admin.gloo.solo.io.RootTrustPolicySpec.Config
	(*RootTrustPolicySpec_Config_MgmtServerCertificateAuthority)(nil), // 3: admin.gloo.solo.io.RootTrustPolicySpec.Config.MgmtServerCertificateAuthority
	(*v2.MeshSelector)(nil),                                           // 4: common.gloo.solo.io.MeshSelector
	(v2.ApprovalState)(0),                                             // 5: common.gloo.solo.io.ApprovalState
	(*tls.AgentCertificateAuthority)(nil),                             // 6: tls.security.policy.gloo.solo.io.AgentCertificateAuthority
	(*tls.CommonCertOptions)(nil),                                     // 7: tls.security.policy.gloo.solo.io.CommonCertOptions
	(*v1.ObjectRef)(nil),                                              // 8: core.skv2.solo.io.ObjectRef
}
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_depIdxs = []int32{
	4, // 0: admin.gloo.solo.io.RootTrustPolicySpec.apply_to_meshes:type_name -> common.gloo.solo.io.MeshSelector
	2, // 1: admin.gloo.solo.io.RootTrustPolicySpec.config:type_name -> admin.gloo.solo.io.RootTrustPolicySpec.Config
	5, // 2: admin.gloo.solo.io.RootTrustPolicyStatus.state:type_name -> common.gloo.solo.io.ApprovalState
	3, // 3: admin.gloo.solo.io.RootTrustPolicySpec.Config.mgmt_server_ca:type_name -> admin.gloo.solo.io.RootTrustPolicySpec.Config.MgmtServerCertificateAuthority
	6, // 4: admin.gloo.solo.io.RootTrustPolicySpec.Config.agent_ca:type_name -> tls.security.policy.gloo.solo.io.AgentCertificateAuthority
	7, // 5: admin.gloo.solo.io.RootTrustPolicySpec.Config.intermediate_cert_options:type_name -> tls.security.policy.gloo.solo.io.CommonCertOptions
	7, // 6: admin.gloo.solo.io.RootTrustPolicySpec.Config.MgmtServerCertificateAuthority.generated:type_name -> tls.security.policy.gloo.solo.io.CommonCertOptions
	8, // 7: admin.gloo.solo.io.RootTrustPolicySpec.Config.MgmtServerCertificateAuthority.secret_ref:type_name -> core.skv2.solo.io.ObjectRef
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RootTrustPolicySpec); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RootTrustPolicyStatus); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RootTrustPolicySpec_Config); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RootTrustPolicySpec_Config_MgmtServerCertificateAuthority); i {
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
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*RootTrustPolicySpec_Config_MgmtServerCa)(nil),
		(*RootTrustPolicySpec_Config_AgentCa)(nil),
	}
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*RootTrustPolicySpec_Config_MgmtServerCertificateAuthority_Generated)(nil),
		(*RootTrustPolicySpec_Config_MgmtServerCertificateAuthority_SecretRef)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_root_trust_policy_proto_depIdxs = nil
}
