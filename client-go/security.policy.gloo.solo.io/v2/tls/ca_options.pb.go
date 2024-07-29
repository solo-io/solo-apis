// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/policy/v2/security/tls/ca_options.proto

package tls

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// State of Certificate Rotation
// Possible states in which a CertificateRotation can exist.
type CertificateRotationState int32

const (
	// No Certificate rotation is currently happening
	CertificateRotationState_NOT_ROTATING CertificateRotationState = 0
	// Signing the certificate using the previously applied CA. This step is mostly used when `ADDING_NEW_ROOT`
	// fails, and the rotation has to be ROLLED_BACK
	CertificateRotationState_PREVIOUS_CA CertificateRotationState = 1
	// The CertificateRotation is underway, both roots are set, and the new root is being propagated
	CertificateRotationState_ADDING_NEW_ROOT CertificateRotationState = 2
	// The CertificateRotation is underway again.
	// The initial verification is over, the traffic continues to work with both roots present.
	// Now the old root is being removed, and the new root is being propagated alone to the data-plane clusters
	CertificateRotationState_PROPAGATING_NEW_INTERMEDIATE CertificateRotationState = 3
	// The CertificateRotation is underway again.
	// Removing the old-root from all data-plane clusters
	CertificateRotationState_DELETING_OLD_ROOT CertificateRotationState = 4
	// Verifying connectivity between workloads, the workflow will not progress until connectivity has been verified.
	// This can either be manual or in the future automated
	CertificateRotationState_VERIFYING CertificateRotationState = 5
	// The connectivity has been verified.
	CertificateRotationState_VERIFIED CertificateRotationState = 6
	// The connectivity has been deemed to not be functioning properly, rolling back to the last
	// known good state.
	CertificateRotationState_ROLLING_BACK CertificateRotationState = 7
	// The rotation has finished, the new root has been propagated to all data-plane clusters, and traffic has
	// been verified successfully.
	CertificateRotationState_FINISHED CertificateRotationState = 8
	// Processing the certificate rotation workflow failed.
	CertificateRotationState_FAILED CertificateRotationState = 9
)

// Enum value maps for CertificateRotationState.
var (
	CertificateRotationState_name = map[int32]string{
		0: "NOT_ROTATING",
		1: "PREVIOUS_CA",
		2: "ADDING_NEW_ROOT",
		3: "PROPAGATING_NEW_INTERMEDIATE",
		4: "DELETING_OLD_ROOT",
		5: "VERIFYING",
		6: "VERIFIED",
		7: "ROLLING_BACK",
		8: "FINISHED",
		9: "FAILED",
	}
	CertificateRotationState_value = map[string]int32{
		"NOT_ROTATING":                 0,
		"PREVIOUS_CA":                  1,
		"ADDING_NEW_ROOT":              2,
		"PROPAGATING_NEW_INTERMEDIATE": 3,
		"DELETING_OLD_ROOT":            4,
		"VERIFYING":                    5,
		"VERIFIED":                     6,
		"ROLLING_BACK":                 7,
		"FINISHED":                     8,
		"FAILED":                       9,
	}
)

func (x CertificateRotationState) Enum() *CertificateRotationState {
	p := new(CertificateRotationState)
	*p = x
	return p
}

func (x CertificateRotationState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CertificateRotationState) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_enumTypes[0].Descriptor()
}

func (CertificateRotationState) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_enumTypes[0]
}

func (x CertificateRotationState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CertificateRotationState.Descriptor instead.
func (CertificateRotationState) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_rawDescGZIP(), []int{0}
}

type CertificateRotationStrategy int32

const (
	// The default certificate rotation strategy. This strategy involves three steps which
	// ensure that traffic in the mesh will experience no downtime.
	// For an in depth explination of how this strategy works in Istio see the [following blog](https://blog.christianposta.com/diving-into-istio-1-6-certificate-rotation/)
	// The steps are as follows:
	//  1. ADDING_NEW_ROOT
	//     During this step the new root-cert will be appended to the old root-cert, and then distributed.
	//     The intermediate will continue to be signed by the original root.
	//  2. PROPAGATING_NEW_INTERMEDIATE
	//     During this step both root-certs will still be distributed. In addition the intermediate will now
	//     be signed by the new root key.
	//  3. DELETING_OLD_ROOT
	//     During this step the old root is no longer included, and the intermediate will continue to be signed
	//     by the new root key.
	CertificateRotationStrategy_MULTI_ROOT CertificateRotationStrategy = 0
	// Do not use any rotation strategy.
	// NOTE: This can lead to downtime while workloads transition
	// from one root of trust to another
	CertificateRotationStrategy_NONE CertificateRotationStrategy = 1
)

// Enum value maps for CertificateRotationStrategy.
var (
	CertificateRotationStrategy_name = map[int32]string{
		0: "MULTI_ROOT",
		1: "NONE",
	}
	CertificateRotationStrategy_value = map[string]int32{
		"MULTI_ROOT": 0,
		"NONE":       1,
	}
)

func (x CertificateRotationStrategy) Enum() *CertificateRotationStrategy {
	p := new(CertificateRotationStrategy)
	*p = x
	return p
}

func (x CertificateRotationStrategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CertificateRotationStrategy) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_enumTypes[1].Descriptor()
}

func (CertificateRotationStrategy) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_enumTypes[1]
}

func (x CertificateRotationStrategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CertificateRotationStrategy.Descriptor instead.
func (CertificateRotationStrategy) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_rawDescGZIP(), []int{1}
}

// Configuration for generating a self-signed root certificate.
// Uses the X.509 format, RFC5280.
type CommonCertOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of days before root cert expires. Defaults to 365.
	TtlDays uint32 `protobuf:"varint,1,opt,name=ttl_days,json=ttlDays,proto3" json:"ttl_days,omitempty"`
	// Size in bytes of the root cert's private key. Defaults to 4096.
	RsaKeySizeBytes uint32 `protobuf:"varint,2,opt,name=rsa_key_size_bytes,json=rsaKeySizeBytes,proto3" json:"rsa_key_size_bytes,omitempty"`
	// Root cert organization name. Defaults to "gloo-mesh".
	OrgName string `protobuf:"bytes,3,opt,name=org_name,json=orgName,proto3" json:"org_name,omitempty"`
	// The ratio of the certificate lifetime to when Gloo starts the certificate rotation process.
	// The ratio must be between 0 and 1 (exclusive). For example, if a certificate is valid for
	// 1 day (or 24 hours), and you specify a ratio of 0.1, Gloo starts the certificate rotation
	// process 2.4 hours before it expires (24x0.1).
	SecretRotationGracePeriodRatio float32 `protobuf:"fixed32,4,opt,name=secret_rotation_grace_period_ratio,json=secretRotationGracePeriodRatio,proto3" json:"secret_rotation_grace_period_ratio,omitempty"`
}

func (x *CommonCertOptions) Reset() {
	*x = CommonCertOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonCertOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonCertOptions) ProtoMessage() {}

func (x *CommonCertOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonCertOptions.ProtoReflect.Descriptor instead.
func (*CommonCertOptions) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_rawDescGZIP(), []int{0}
}

func (x *CommonCertOptions) GetTtlDays() uint32 {
	if x != nil {
		return x.TtlDays
	}
	return 0
}

func (x *CommonCertOptions) GetRsaKeySizeBytes() uint32 {
	if x != nil {
		return x.RsaKeySizeBytes
	}
	return 0
}

func (x *CommonCertOptions) GetOrgName() string {
	if x != nil {
		return x.OrgName
	}
	return ""
}

func (x *CommonCertOptions) GetSecretRotationGracePeriodRatio() float32 {
	if x != nil {
		return x.SecretRotationGracePeriodRatio
	}
	return 0
}

// Specify parameters for configuring the root certificate authority for a VirtualMesh.
type AgentCertificateAuthority struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specify the source of the Root CA data which Gloo Mesh will use for the VirtualMesh.
	//
	// Types that are assignable to CaSource:
	//
	//	*AgentCertificateAuthority_Vault
	CaSource isAgentCertificateAuthority_CaSource `protobuf_oneof:"ca_source"`
}

func (x *AgentCertificateAuthority) Reset() {
	*x = AgentCertificateAuthority{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentCertificateAuthority) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentCertificateAuthority) ProtoMessage() {}

func (x *AgentCertificateAuthority) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentCertificateAuthority.ProtoReflect.Descriptor instead.
func (*AgentCertificateAuthority) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_rawDescGZIP(), []int{1}
}

func (m *AgentCertificateAuthority) GetCaSource() isAgentCertificateAuthority_CaSource {
	if m != nil {
		return m.CaSource
	}
	return nil
}

func (x *AgentCertificateAuthority) GetVault() *VaultCA {
	if x, ok := x.GetCaSource().(*AgentCertificateAuthority_Vault); ok {
		return x.Vault
	}
	return nil
}

type isAgentCertificateAuthority_CaSource interface {
	isAgentCertificateAuthority_CaSource()
}

type AgentCertificateAuthority_Vault struct {
	// Use vault as the intermediate CA source
	Vault *VaultCA `protobuf:"bytes,1,opt,name=vault,proto3,oneof"`
}

func (*AgentCertificateAuthority_Vault) isAgentCertificateAuthority_CaSource() {}

type CertificateRotationVerificationMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Method:
	//
	//	*CertificateRotationVerificationMethod_None
	//	*CertificateRotationVerificationMethod_Manual
	Method isCertificateRotationVerificationMethod_Method `protobuf_oneof:"method"`
}

func (x *CertificateRotationVerificationMethod) Reset() {
	*x = CertificateRotationVerificationMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateRotationVerificationMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateRotationVerificationMethod) ProtoMessage() {}

func (x *CertificateRotationVerificationMethod) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateRotationVerificationMethod.ProtoReflect.Descriptor instead.
func (*CertificateRotationVerificationMethod) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_rawDescGZIP(), []int{2}
}

func (m *CertificateRotationVerificationMethod) GetMethod() isCertificateRotationVerificationMethod_Method {
	if m != nil {
		return m.Method
	}
	return nil
}

func (x *CertificateRotationVerificationMethod) GetNone() *emptypb.Empty {
	if x, ok := x.GetMethod().(*CertificateRotationVerificationMethod_None); ok {
		return x.None
	}
	return nil
}

func (x *CertificateRotationVerificationMethod) GetManual() *emptypb.Empty {
	if x, ok := x.GetMethod().(*CertificateRotationVerificationMethod_Manual); ok {
		return x.Manual
	}
	return nil
}

type isCertificateRotationVerificationMethod_Method interface {
	isCertificateRotationVerificationMethod_Method()
}

type CertificateRotationVerificationMethod_None struct {
	// Verification not enabled. NOTE: This setting is only recommended for testing.
	// When enabled rotation will continue from step to step without any kind of verification.
	// For information about the value format, see the [Google protocol buffer documentation](https://protobuf.dev/reference/protobuf/google.protobuf/#empty).
	None *emptypb.Empty `protobuf:"bytes,1,opt,name=none,proto3,oneof"`
}

type CertificateRotationVerificationMethod_Manual struct {
	// Verification must be completed manually. This involves using our certificate verification
	// endpoint when the certificates are in a VERIFYING state
	// For information about the value format, see the [Google protocol buffer documentation](https://protobuf.dev/reference/protobuf/google.protobuf/#empty).
	Manual *emptypb.Empty `protobuf:"bytes,2,opt,name=manual,proto3,oneof"`
}

func (*CertificateRotationVerificationMethod_None) isCertificateRotationVerificationMethod_Method() {}

func (*CertificateRotationVerificationMethod_Manual) isCertificateRotationVerificationMethod_Method() {
}

// CertificateRotationCondition represents a timesptamped snapshot of the certificate
// rotation workflow. This is used to keep track of the steps which have been completed
// thus far.
type CertificateRotationCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The time at which this condition was recorded
	Timestamp string `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The current state of the cert rotation
	State CertificateRotationState `protobuf:"varint,2,opt,name=state,proto3,enum=tls.security.policy.gloo.solo.io.CertificateRotationState" json:"state,omitempty"`
	// A human readable message related to the current condition
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// Any errors which occurred during the current rotation stage
	Errors []string `protobuf:"bytes,4,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *CertificateRotationCondition) Reset() {
	*x = CertificateRotationCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateRotationCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateRotationCondition) ProtoMessage() {}

func (x *CertificateRotationCondition) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateRotationCondition.ProtoReflect.Descriptor instead.
func (*CertificateRotationCondition) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_rawDescGZIP(), []int{3}
}

func (x *CertificateRotationCondition) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *CertificateRotationCondition) GetState() CertificateRotationState {
	if x != nil {
		return x.State
	}
	return CertificateRotationState_NOT_ROTATING
}

func (x *CertificateRotationCondition) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CertificateRotationCondition) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_rawDesc = []byte{
	0x0a, 0x60, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x74, 0x6c,
	0x73, 0x2f, 0x63, 0x61, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x20, 0x74, 0x6c, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65,
	0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x5e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x74, 0x6c, 0x73, 0x2f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x5f,
	0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x43, 0x65, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x74,
	0x74, 0x6c, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x74,
	0x74, 0x6c, 0x44, 0x61, 0x79, 0x73, 0x12, 0x2b, 0x0a, 0x12, 0x72, 0x73, 0x61, 0x5f, 0x6b, 0x65,
	0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0f, 0x72, 0x73, 0x61, 0x4b, 0x65, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4a,
	0x0a, 0x22, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x67, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x1e, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x61, 0x63, 0x65, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x22, 0x6b, 0x0a, 0x19, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x41, 0x0a, 0x05, 0x76, 0x61, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x43,
	0x41, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x63, 0x61,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x25, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x2c, 0x0a, 0x04, 0x6e, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x6f, 0x6e, 0x65, 0x12,
	0x30, 0x0a, 0x06, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x06, 0x6d, 0x61, 0x6e, 0x75, 0x61,
	0x6c, 0x42, 0x08, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0xc0, 0x01, 0x0a, 0x1c,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x50, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x74, 0x6c, 0x73, 0x2e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2a, 0xd4,
	0x01, 0x0a, 0x18, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x4e,
	0x4f, 0x54, 0x5f, 0x52, 0x4f, 0x54, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x50, 0x52, 0x45, 0x56, 0x49, 0x4f, 0x55, 0x53, 0x5f, 0x43, 0x41, 0x10, 0x01, 0x12, 0x13,
	0x0a, 0x0f, 0x41, 0x44, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x52, 0x4f, 0x4f,
	0x54, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c, 0x50, 0x52, 0x4f, 0x50, 0x41, 0x47, 0x41, 0x54, 0x49,
	0x4e, 0x47, 0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4d, 0x45, 0x44, 0x49,
	0x41, 0x54, 0x45, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e,
	0x47, 0x5f, 0x4f, 0x4c, 0x44, 0x5f, 0x52, 0x4f, 0x4f, 0x54, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09,
	0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x56,
	0x45, 0x52, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x4f, 0x4c,
	0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x46,
	0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49,
	0x4c, 0x45, 0x44, 0x10, 0x09, 0x2a, 0x37, 0x0a, 0x1b, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x5f, 0x52, 0x4f,
	0x4f, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x42, 0x5d,
	0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76,
	0x32, 0x2f, 0x74, 0x6c, 0x73, 0xc0, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_rawDescData = file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_goTypes = []interface{}{
	(CertificateRotationState)(0),                 // 0: tls.security.policy.gloo.solo.io.CertificateRotationState
	(CertificateRotationStrategy)(0),              // 1: tls.security.policy.gloo.solo.io.CertificateRotationStrategy
	(*CommonCertOptions)(nil),                     // 2: tls.security.policy.gloo.solo.io.CommonCertOptions
	(*AgentCertificateAuthority)(nil),             // 3: tls.security.policy.gloo.solo.io.AgentCertificateAuthority
	(*CertificateRotationVerificationMethod)(nil), // 4: tls.security.policy.gloo.solo.io.CertificateRotationVerificationMethod
	(*CertificateRotationCondition)(nil),          // 5: tls.security.policy.gloo.solo.io.CertificateRotationCondition
	(*VaultCA)(nil),                               // 6: tls.security.policy.gloo.solo.io.VaultCA
	(*emptypb.Empty)(nil),                         // 7: google.protobuf.Empty
}
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_depIdxs = []int32{
	6, // 0: tls.security.policy.gloo.solo.io.AgentCertificateAuthority.vault:type_name -> tls.security.policy.gloo.solo.io.VaultCA
	7, // 1: tls.security.policy.gloo.solo.io.CertificateRotationVerificationMethod.none:type_name -> google.protobuf.Empty
	7, // 2: tls.security.policy.gloo.solo.io.CertificateRotationVerificationMethod.manual:type_name -> google.protobuf.Empty
	0, // 3: tls.security.policy.gloo.solo.io.CertificateRotationCondition.state:type_name -> tls.security.policy.gloo.solo.io.CertificateRotationState
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_init()
}
func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_init() {
	if File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto != nil {
		return
	}
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_vault_ca_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonCertOptions); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentCertificateAuthority); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateRotationVerificationMethod); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateRotationCondition); i {
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
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*AgentCertificateAuthority_Vault)(nil),
	}
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*CertificateRotationVerificationMethod_None)(nil),
		(*CertificateRotationVerificationMethod_Manual)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto = out.File
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_tls_ca_options_proto_depIdxs = nil
}