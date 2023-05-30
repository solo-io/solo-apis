// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/internal/v2/issued_certificate.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	tls "github.com/solo-io/solo-apis/client-go/security.policy.gloo.solo.io/v2/tls"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Possible states in which an IssuedCertificate can exist.
type IssuedCertificateStatus_State int32

const (
	// The IssuedCertificate has yet to be picked up by the agent.
	IssuedCertificateStatus_PENDING IssuedCertificateStatus_State = 0
	// The agent has created a local private key
	// and a CertificateRequest for the IssuedCertificate.
	// In this state, the agent is waiting for the Issuer
	// to issue certificates for the CertificateRequest before proceeding.
	IssuedCertificateStatus_REQUESTED IssuedCertificateStatus_State = 1
	// The certificate has been issued. Any pods that require restarting will be restarted at this point.
	IssuedCertificateStatus_ISSUED IssuedCertificateStatus_State = 2
	// The reply from the Issuer has been processed and
	// the agent has placed the final certificate secret
	// in the target location specified by the IssuedCertificate.
	IssuedCertificateStatus_FINISHED IssuedCertificateStatus_State = 3
	// Processing the certificate workflow failed.
	IssuedCertificateStatus_FAILED IssuedCertificateStatus_State = 4
)

// Enum value maps for IssuedCertificateStatus_State.
var (
	IssuedCertificateStatus_State_name = map[int32]string{
		0: "PENDING",
		1: "REQUESTED",
		2: "ISSUED",
		3: "FINISHED",
		4: "FAILED",
	}
	IssuedCertificateStatus_State_value = map[string]int32{
		"PENDING":   0,
		"REQUESTED": 1,
		"ISSUED":    2,
		"FINISHED":  3,
		"FAILED":    4,
	}
)

func (x IssuedCertificateStatus_State) Enum() *IssuedCertificateStatus_State {
	p := new(IssuedCertificateStatus_State)
	*p = x
	return p
}

func (x IssuedCertificateStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IssuedCertificateStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_enumTypes[0].Descriptor()
}

func (IssuedCertificateStatus_State) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_enumTypes[0]
}

func (x IssuedCertificateStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IssuedCertificateStatus_State.Descriptor instead.
func (IssuedCertificateStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_rawDescGZIP(), []int{2, 0}
}

// IssuedCertificates are used to issue SSL certificates
// to workload Kubernetes clusters from a central (out-of-cluster) Certificate Authority.
//
// When an IssuedCertificate is created, a certificate is issued to a workload cluster by
// a central Certificate Authority via the following workflow:
//
// 1. The Certificate Issuer creates the IssuedCertificate resource on the remote cluster
// 2. The Certificate Signature Requesting agent installed to the remote cluster generates
// a Certificate Signing Request (CSR) and writes it to the status of the IssuedCertificate
// 3. Finally, the Certificate Issuer generates a signed certificate for the CSR and writes
// it back as Kubernetes Secret in the remote cluster.
//
// Trust can therefore be established across the Gloo Mesh server cluster and agents on workload clusters
// without requiring private keys to ever leave workload clusters.
//
// The certificate requested here is for Gloo Mesh agents on workload clusters to securely establish communication
// with Gloo Mesh server. This is not related to certificates for services running in the mesh.
type IssuedCertificateSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of hostnames and IPs to generate a certificate for.
	// This can also be set to the identity running the workload,
	// e.g. a Kubernetes service account.
	//
	// Generally for an Istio CA this will take the form `spiffe://cluster.local/ns/istio-system/sa/citadel`.
	//
	// "cluster.local" may be replaced by the root of trust domain for the mesh.
	Hosts []string `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	// The secret containing the SSL certificate to be generated for this IssuedCertificate (located in the Gloo Mesh agent's cluster).
	// If nil, the sidecar agent stores the signing certificate in memory. (Enterprise only)
	IssuedCertificateSecret *v1.ObjectRef `protobuf:"bytes,2,opt,name=issued_certificate_secret,json=issuedCertificateSecret,proto3" json:"issued_certificate_secret,omitempty"`
	// Set of options to configure the intermediate certificate being generated
	CertOptions *tls.CommonCertOptions `protobuf:"bytes,3,opt,name=cert_options,json=certOptions,proto3" json:"cert_options,omitempty"`
	// The location of the certificate authority to sign this certificate
	//
	// Types that are assignable to CertificateAuthority:
	//
	//	*IssuedCertificateSpec_MgmtServerCa
	//	*IssuedCertificateSpec_AgentCa
	CertificateAuthority isIssuedCertificateSpec_CertificateAuthority `protobuf_oneof:"certificate_authority"`
	// Reference to the mesh on which this cert is being issued for.
	MeshRef *v1.ObjectRef `protobuf:"bytes,6,opt,name=mesh_ref,json=meshRef,proto3" json:"mesh_ref,omitempty"`
	// Signals whether or not the workload pods should be restarted
	// to pick up the new cert.
	AutoRestartPods bool `protobuf:"varint,7,opt,name=auto_restart_pods,json=autoRestartPods,proto3" json:"auto_restart_pods,omitempty"`
	// A list of certificate authorities that should also be trusted by workloads
	PassiveCertificateAuthorities []*MgmtServerCertificateAuthority `protobuf:"bytes,10,rep,name=passive_certificate_authorities,json=passiveCertificateAuthorities,proto3" json:"passive_certificate_authorities,omitempty"`
}

func (x *IssuedCertificateSpec) Reset() {
	*x = IssuedCertificateSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssuedCertificateSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssuedCertificateSpec) ProtoMessage() {}

func (x *IssuedCertificateSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssuedCertificateSpec.ProtoReflect.Descriptor instead.
func (*IssuedCertificateSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_rawDescGZIP(), []int{0}
}

func (x *IssuedCertificateSpec) GetHosts() []string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *IssuedCertificateSpec) GetIssuedCertificateSecret() *v1.ObjectRef {
	if x != nil {
		return x.IssuedCertificateSecret
	}
	return nil
}

func (x *IssuedCertificateSpec) GetCertOptions() *tls.CommonCertOptions {
	if x != nil {
		return x.CertOptions
	}
	return nil
}

func (m *IssuedCertificateSpec) GetCertificateAuthority() isIssuedCertificateSpec_CertificateAuthority {
	if m != nil {
		return m.CertificateAuthority
	}
	return nil
}

func (x *IssuedCertificateSpec) GetMgmtServerCa() *MgmtServerCertificateAuthority {
	if x, ok := x.GetCertificateAuthority().(*IssuedCertificateSpec_MgmtServerCa); ok {
		return x.MgmtServerCa
	}
	return nil
}

func (x *IssuedCertificateSpec) GetAgentCa() *tls.AgentCertificateAuthority {
	if x, ok := x.GetCertificateAuthority().(*IssuedCertificateSpec_AgentCa); ok {
		return x.AgentCa
	}
	return nil
}

func (x *IssuedCertificateSpec) GetMeshRef() *v1.ObjectRef {
	if x != nil {
		return x.MeshRef
	}
	return nil
}

func (x *IssuedCertificateSpec) GetAutoRestartPods() bool {
	if x != nil {
		return x.AutoRestartPods
	}
	return false
}

func (x *IssuedCertificateSpec) GetPassiveCertificateAuthorities() []*MgmtServerCertificateAuthority {
	if x != nil {
		return x.PassiveCertificateAuthorities
	}
	return nil
}

type isIssuedCertificateSpec_CertificateAuthority interface {
	isIssuedCertificateSpec_CertificateAuthority()
}

type IssuedCertificateSpec_MgmtServerCa struct {
	// Gloo Mesh CA options
	MgmtServerCa *MgmtServerCertificateAuthority `protobuf:"bytes,4,opt,name=mgmt_server_ca,json=mgmtServerCa,proto3,oneof"`
}

type IssuedCertificateSpec_AgentCa struct {
	// Agent CA options
	AgentCa *tls.AgentCertificateAuthority `protobuf:"bytes,5,opt,name=agent_ca,json=agentCa,proto3,oneof"`
}

func (*IssuedCertificateSpec_MgmtServerCa) isIssuedCertificateSpec_CertificateAuthority() {}

func (*IssuedCertificateSpec_AgentCa) isIssuedCertificateSpec_CertificateAuthority() {}

// Set of options which represent the certificate authorities the management cluster can use
// to sign the intermediate certs.
type MgmtServerCertificateAuthority struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Certificate authority which gloo-mesh management will use to sign the intermediate cert
	//
	// Types that are assignable to CertificateAuthority:
	//
	//	*MgmtServerCertificateAuthority_SigningCertificateSecret
	CertificateAuthority isMgmtServerCertificateAuthority_CertificateAuthority `protobuf_oneof:"certificate_authority"`
}

func (x *MgmtServerCertificateAuthority) Reset() {
	*x = MgmtServerCertificateAuthority{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MgmtServerCertificateAuthority) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MgmtServerCertificateAuthority) ProtoMessage() {}

func (x *MgmtServerCertificateAuthority) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MgmtServerCertificateAuthority.ProtoReflect.Descriptor instead.
func (*MgmtServerCertificateAuthority) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_rawDescGZIP(), []int{1}
}

func (m *MgmtServerCertificateAuthority) GetCertificateAuthority() isMgmtServerCertificateAuthority_CertificateAuthority {
	if m != nil {
		return m.CertificateAuthority
	}
	return nil
}

func (x *MgmtServerCertificateAuthority) GetSigningCertificateSecret() *v1.ObjectRef {
	if x, ok := x.GetCertificateAuthority().(*MgmtServerCertificateAuthority_SigningCertificateSecret); ok {
		return x.SigningCertificateSecret
	}
	return nil
}

type isMgmtServerCertificateAuthority_CertificateAuthority interface {
	isMgmtServerCertificateAuthority_CertificateAuthority()
}

type MgmtServerCertificateAuthority_SigningCertificateSecret struct {
	SigningCertificateSecret *v1.ObjectRef `protobuf:"bytes,1,opt,name=signing_certificate_secret,json=signingCertificateSecret,proto3,oneof"`
}

func (*MgmtServerCertificateAuthority_SigningCertificateSecret) isMgmtServerCertificateAuthority_CertificateAuthority() {
}

// The IssuedCertificate status is written by the CertificateRequesting agent.
type IssuedCertificateStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The most recent generation observed in the the IssuedCertificate metadata.
	// If the `observedGeneration` does not match `metadata.generation`, the Gloo Mesh agent has not processed the most
	// recent version of this IssuedCertificate.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// Any error observed which prevented the CertificateRequest from being processed.
	// If the error is empty, the request has been processed successfully.
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// The current state of the IssuedCertificate workflow, reported by the agent.
	State IssuedCertificateStatus_State `protobuf:"varint,3,opt,name=state,proto3,enum=internal.gloo.solo.io.IssuedCertificateStatus_State" json:"state,omitempty"`
}

func (x *IssuedCertificateStatus) Reset() {
	*x = IssuedCertificateStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssuedCertificateStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssuedCertificateStatus) ProtoMessage() {}

func (x *IssuedCertificateStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssuedCertificateStatus.ProtoReflect.Descriptor instead.
func (*IssuedCertificateStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_rawDescGZIP(), []int{2}
}

func (x *IssuedCertificateStatus) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *IssuedCertificateStatus) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *IssuedCertificateStatus) GetState() IssuedCertificateStatus_State {
	if x != nil {
		return x.State
	}
	return IssuedCertificateStatus_PENDING
}

var File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_rawDesc = []byte{
	0x0a, 0x5d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x5f, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x15, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x60, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72,
	0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x74, 0x6c, 0x73, 0x2f, 0x63, 0x61, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f,
	0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x05, 0x0a,
	0x15, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x58, 0x0a, 0x19,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x52, 0x17, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x64, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x56, 0x0a, 0x0c, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x74,
	0x6c, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5d,
	0x0a, 0x0e, 0x6d, 0x67, 0x6d, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4d,
	0x67, 0x6d, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x48, 0x00, 0x52,
	0x0c, 0x6d, 0x67, 0x6d, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x61, 0x12, 0x58, 0x0a,
	0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3b, 0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x48, 0x00, 0x52, 0x07,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x12, 0x37, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x68, 0x5f,
	0x72, 0x65, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x68, 0x52, 0x65, 0x66,
	0x12, 0x2a, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x70, 0x6f, 0x64, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x61, 0x75, 0x74,
	0x6f, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f, 0x64, 0x73, 0x12, 0x7d, 0x0a, 0x1f,
	0x70, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4d, 0x67,
	0x6d, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x52, 0x1d, 0x70, 0x61,
	0x73, 0x73, 0x69, 0x76, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x17, 0x0a, 0x15, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x22, 0x97, 0x01, 0x0a, 0x1e, 0x4d, 0x67, 0x6d, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x5c, 0x0a, 0x1a, 0x73, 0x69, 0x67, 0x6e, 0x69,
	0x6e, 0x67, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x48, 0x00, 0x52, 0x18, 0x73, 0x69, 0x67,
	0x6e, 0x69, 0x6e, 0x67, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x42, 0x17, 0x0a, 0x15, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0xf7,
	0x01, 0x0a, 0x17, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x4a, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x34, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x49, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x53, 0x53, 0x55, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0c,
	0x0a, 0x08, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x42, 0x56, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72,
	0x69, 0x73, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2f, 0x76, 0x32, 0xc0, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_rawDescData = file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_goTypes = []interface{}{
	(IssuedCertificateStatus_State)(0),     // 0: internal.gloo.solo.io.IssuedCertificateStatus.State
	(*IssuedCertificateSpec)(nil),          // 1: internal.gloo.solo.io.IssuedCertificateSpec
	(*MgmtServerCertificateAuthority)(nil), // 2: internal.gloo.solo.io.MgmtServerCertificateAuthority
	(*IssuedCertificateStatus)(nil),        // 3: internal.gloo.solo.io.IssuedCertificateStatus
	(*v1.ObjectRef)(nil),                   // 4: core.skv2.solo.io.ObjectRef
	(*tls.CommonCertOptions)(nil),          // 5: tls.security.policy.gloo.solo.io.CommonCertOptions
	(*tls.AgentCertificateAuthority)(nil),  // 6: tls.security.policy.gloo.solo.io.AgentCertificateAuthority
}
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_depIdxs = []int32{
	4, // 0: internal.gloo.solo.io.IssuedCertificateSpec.issued_certificate_secret:type_name -> core.skv2.solo.io.ObjectRef
	5, // 1: internal.gloo.solo.io.IssuedCertificateSpec.cert_options:type_name -> tls.security.policy.gloo.solo.io.CommonCertOptions
	2, // 2: internal.gloo.solo.io.IssuedCertificateSpec.mgmt_server_ca:type_name -> internal.gloo.solo.io.MgmtServerCertificateAuthority
	6, // 3: internal.gloo.solo.io.IssuedCertificateSpec.agent_ca:type_name -> tls.security.policy.gloo.solo.io.AgentCertificateAuthority
	4, // 4: internal.gloo.solo.io.IssuedCertificateSpec.mesh_ref:type_name -> core.skv2.solo.io.ObjectRef
	2, // 5: internal.gloo.solo.io.IssuedCertificateSpec.passive_certificate_authorities:type_name -> internal.gloo.solo.io.MgmtServerCertificateAuthority
	4, // 6: internal.gloo.solo.io.MgmtServerCertificateAuthority.signing_certificate_secret:type_name -> core.skv2.solo.io.ObjectRef
	0, // 7: internal.gloo.solo.io.IssuedCertificateStatus.state:type_name -> internal.gloo.solo.io.IssuedCertificateStatus.State
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_init()
}
func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_init() {
	if File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssuedCertificateSpec); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MgmtServerCertificateAuthority); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssuedCertificateStatus); i {
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
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*IssuedCertificateSpec_MgmtServerCa)(nil),
		(*IssuedCertificateSpec_AgentCa)(nil),
	}
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*MgmtServerCertificateAuthority_SigningCertificateSecret)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto = out.File
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_issued_certificate_proto_depIdxs = nil
}
