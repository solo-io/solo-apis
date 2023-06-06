// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/networking/v2/external_service.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	v2 "github.com/solo-io/solo-apis/client-go/common.gloo.solo.io/v2"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExternalServiceSpec_Port_TlsConfigTlsMode int32

const (
	ExternalServiceSpec_Port_TlsConfig_SIMPLE       ExternalServiceSpec_Port_TlsConfigTlsMode = 0
	ExternalServiceSpec_Port_TlsConfig_DISABLE      ExternalServiceSpec_Port_TlsConfigTlsMode = 1
	ExternalServiceSpec_Port_TlsConfig_MUTUAL       ExternalServiceSpec_Port_TlsConfigTlsMode = 2
	ExternalServiceSpec_Port_TlsConfig_ISTIO_MUTUAL ExternalServiceSpec_Port_TlsConfigTlsMode = 3
)

// Enum value maps for ExternalServiceSpec_Port_TlsConfigTlsMode.
var (
	ExternalServiceSpec_Port_TlsConfigTlsMode_name = map[int32]string{
		0: "SIMPLE",
		1: "DISABLE",
		2: "MUTUAL",
		3: "ISTIO_MUTUAL",
	}
	ExternalServiceSpec_Port_TlsConfigTlsMode_value = map[string]int32{
		"SIMPLE":       0,
		"DISABLE":      1,
		"MUTUAL":       2,
		"ISTIO_MUTUAL": 3,
	}
)

func (x ExternalServiceSpec_Port_TlsConfigTlsMode) Enum() *ExternalServiceSpec_Port_TlsConfigTlsMode {
	p := new(ExternalServiceSpec_Port_TlsConfigTlsMode)
	*p = x
	return p
}

func (x ExternalServiceSpec_Port_TlsConfigTlsMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExternalServiceSpec_Port_TlsConfigTlsMode) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_enumTypes[0].Descriptor()
}

func (ExternalServiceSpec_Port_TlsConfigTlsMode) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_enumTypes[0]
}

func (x ExternalServiceSpec_Port_TlsConfigTlsMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExternalServiceSpec_Port_TlsConfigTlsMode.Descriptor instead.
func (ExternalServiceSpec_Port_TlsConfigTlsMode) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

// ExternalService defines a destination for routing which exist outside the mesh.
// This could for example be a web API or a set of virtual machines that are not running in Kubernetes.
// When an ExternalService is defined for a given workspace, it can be used as a
// Destination in Routes, as well as called directly via its specified hostname.
// Resolution of the IP addresses for external services can be done via DNS
// or provided statically using the ExternalEndpoint object.
type ExternalServiceSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (mutually exclusive with addresses): The List of hostnames which will resolve to this service.
	// These hosts must be unique among all ExternalServices and VirtualHosts within a workspace.
	// Both FQDN and wildcard prefix domains are supported.
	// TLS origination to ExternalServices is supported with use of the ClientsideTls property.
	Hosts []string `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	// (mutually exclusive with hosts): The List of ipv4 or ipv6 addresses which will be associated to this service. Can be CIDR prefixes.
	// These addresses must be unique among all ExternalServices within a workspace.
	Addresses []string `protobuf:"bytes,5,rep,name=addresses,proto3" json:"addresses,omitempty"`
	// The associated ports of the external service.
	Ports []*ExternalServiceSpec_Port `protobuf:"bytes,2,rep,name=ports,proto3" json:"ports,omitempty"`
	// Selecting ExternalEndpoints will provide the addresses used for routing traffic to on the ExternalService's hosts.
	// An empty selector will not select any endpoints.
	// If no endpoints are selected, requests will be routed to the provided hosts using dns resolution.
	Selector map[string]string `protobuf:"bytes,3,rep,name=selector,proto3" json:"selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A list of alternate names to verify the subject identity in the
	// certificate. If specified, the proxy will verify that the server
	// certificate's subject alt name matches one of the specified values.
	// Only applicable when using TLS to communicate with the ExternalService.
	SubjectAltNames []string `protobuf:"bytes,4,rep,name=subject_alt_names,json=subjectAltNames,proto3" json:"subject_alt_names,omitempty"`
}

func (x *ExternalServiceSpec) Reset() {
	*x = ExternalServiceSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalServiceSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalServiceSpec) ProtoMessage() {}

func (x *ExternalServiceSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalServiceSpec.ProtoReflect.Descriptor instead.
func (*ExternalServiceSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_rawDescGZIP(), []int{0}
}

func (x *ExternalServiceSpec) GetHosts() []string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *ExternalServiceSpec) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *ExternalServiceSpec) GetPorts() []*ExternalServiceSpec_Port {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *ExternalServiceSpec) GetSelector() map[string]string {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *ExternalServiceSpec) GetSubjectAltNames() []string {
	if x != nil {
		return x.SubjectAltNames
	}
	return nil
}

// reflects the status of the ExternalService
type ExternalServiceStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Global *v2.GenericGlobalStatus `protobuf:"bytes,1,opt,name=global,proto3" json:"global,omitempty"`
	// The status of the resource in each workspace that it exists in.
	Workspaces map[string]*v2.WorkspaceStatus `protobuf:"bytes,2,rep,name=workspaces,proto3" json:"workspaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Map of policy GVK to policy references for all policies applied on this
	// resource.
	AppliedDestinationPolicies map[string]*v2.AppliedDestinationPortPolicies `protobuf:"bytes,3,rep,name=applied_destination_policies,json=appliedDestinationPolicies,proto3" json:"applied_destination_policies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The list of endpoints that this ExternalService is currently using.
	SelectedExternalEndpoints []*v2.ObjectReference `protobuf:"bytes,4,rep,name=selected_external_endpoints,json=selectedExternalEndpoints,proto3" json:"selected_external_endpoints,omitempty"`
	// Name of Workspace that owns ExternalService
	OwnerWorkspace *v2.OwnerWorkspace `protobuf:"bytes,5,opt,name=owner_workspace,json=ownerWorkspace,proto3" json:"owner_workspace,omitempty"`
}

func (x *ExternalServiceStatus) Reset() {
	*x = ExternalServiceStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalServiceStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalServiceStatus) ProtoMessage() {}

func (x *ExternalServiceStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalServiceStatus.ProtoReflect.Descriptor instead.
func (*ExternalServiceStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_rawDescGZIP(), []int{1}
}

func (x *ExternalServiceStatus) GetGlobal() *v2.GenericGlobalStatus {
	if x != nil {
		return x.Global
	}
	return nil
}

func (x *ExternalServiceStatus) GetWorkspaces() map[string]*v2.WorkspaceStatus {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

func (x *ExternalServiceStatus) GetAppliedDestinationPolicies() map[string]*v2.AppliedDestinationPortPolicies {
	if x != nil {
		return x.AppliedDestinationPolicies
	}
	return nil
}

func (x *ExternalServiceStatus) GetSelectedExternalEndpoints() []*v2.ObjectReference {
	if x != nil {
		return x.SelectedExternalEndpoints
	}
	return nil
}

func (x *ExternalServiceStatus) GetOwnerWorkspace() *v2.OwnerWorkspace {
	if x != nil {
		return x.OwnerWorkspace
	}
	return nil
}

// Port establishes a new port that will be exposed on an ExternalService.
type ExternalServiceSpec_Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The port number. Must be a valid, non-negative integer port number.
	Number uint32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	// (optional): The port number or name used to match the corresponding port on the ExternalService's backing ExternalEndpoints.
	// All of the backing ExternalEndpoints for this ExternalService must contain
	// this port, matching by name or number.
	// If no backing ExternalEndpoints are provided, a name selector is invalid,
	// and requests will be routed to the provided hosts on the port number specified.
	// If unspecified, will default to the value of the port number field above.
	TargetPort *v2.PortSelector `protobuf:"bytes,2,opt,name=target_port,json=targetPort,proto3" json:"target_port,omitempty"`
	// A label for the port, eg "http".
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The protocol used in communication with this destination
	// MUST be one of the following: HTTP, HTTPS, GRPC, HTTP2, MONGO, TCP, TLS.
	Protocol string `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// The tls config for the given port.
	// If the protocol is HTTPS or TLS, the ExternalService will be configured to use TLS automatically.
	// If used in conjunction with targetPort this can be used for TLS Origination.
	// For example, port: 80, targetPort: 443, with ClientsideTls will cause port 80 traffic from workloads to be forwarded to 443 resulting in HTTPS traffic over the internet with TLS originating at the sidecar proxy.
	ClientsideTls *ExternalServiceSpec_Port_TlsConfig `protobuf:"bytes,5,opt,name=clientside_tls,json=clientsideTls,proto3" json:"clientside_tls,omitempty"`
}

func (x *ExternalServiceSpec_Port) Reset() {
	*x = ExternalServiceSpec_Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalServiceSpec_Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalServiceSpec_Port) ProtoMessage() {}

func (x *ExternalServiceSpec_Port) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalServiceSpec_Port.ProtoReflect.Descriptor instead.
func (*ExternalServiceSpec_Port) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ExternalServiceSpec_Port) GetNumber() uint32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *ExternalServiceSpec_Port) GetTargetPort() *v2.PortSelector {
	if x != nil {
		return x.TargetPort
	}
	return nil
}

func (x *ExternalServiceSpec_Port) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExternalServiceSpec_Port) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *ExternalServiceSpec_Port) GetClientsideTls() *ExternalServiceSpec_Port_TlsConfig {
	if x != nil {
		return x.ClientsideTls
	}
	return nil
}

type ExternalServiceSpec_Port_TlsConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SNI string to present to the server during TLS handshake.
	// If there is a single host in the hosts list, this will be used as the SNI string.
	Sni string `protobuf:"bytes,1,opt,name=sni,proto3" json:"sni,omitempty"`
	// tls mode passed to corresponding DestinationRules' tls traffic policy
	// MUST be one of the following: SIMPLE, MUTUAL, ISTIO_MUTUAL, DISABLED
	Mode ExternalServiceSpec_Port_TlsConfigTlsMode `protobuf:"varint,2,opt,name=mode,proto3,enum=networking.gloo.solo.io.ExternalServiceSpec_Port_TlsConfigTlsMode" json:"mode,omitempty"`
	// REQUIRED if mode is `MUTUAL`.
	// Should be empty if mode is `ISTIO_MUTUAL`.
	// The file path to the client-side TLS certificate to use, which should be mounted onto the proxy of the specific workload communicating with the external service.
	ClientCertificate string `protobuf:"bytes,3,opt,name=client_certificate,json=clientCertificate,proto3" json:"client_certificate,omitempty"`
	// REQUIRED if mode is `MUTUAL`.
	// Should be empty if mode is `ISTIO_MUTUAL`.
	// The file path to the client’s private key, which should be mounted onto the proxy of the specific workload communicating with the external service.
	PrivateKey string `protobuf:"bytes,4,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// Should be empty if mode is `ISTIO_MUTUAL`.
	// The file path to the file containing CA certificates used to verify server certificates, which should be mounted onto the proxy of the specific workload communicating with the external service.
	// If omitted, the proxy will not verify the server’s certificate.
	CaCertificates string `protobuf:"bytes,5,opt,name=ca_certificates,json=caCertificates,proto3" json:"ca_certificates,omitempty"`
}

func (x *ExternalServiceSpec_Port_TlsConfig) Reset() {
	*x = ExternalServiceSpec_Port_TlsConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalServiceSpec_Port_TlsConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalServiceSpec_Port_TlsConfig) ProtoMessage() {}

func (x *ExternalServiceSpec_Port_TlsConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalServiceSpec_Port_TlsConfig.ProtoReflect.Descriptor instead.
func (*ExternalServiceSpec_Port_TlsConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ExternalServiceSpec_Port_TlsConfig) GetSni() string {
	if x != nil {
		return x.Sni
	}
	return ""
}

func (x *ExternalServiceSpec_Port_TlsConfig) GetMode() ExternalServiceSpec_Port_TlsConfigTlsMode {
	if x != nil {
		return x.Mode
	}
	return ExternalServiceSpec_Port_TlsConfig_SIMPLE
}

func (x *ExternalServiceSpec_Port_TlsConfig) GetClientCertificate() string {
	if x != nil {
		return x.ClientCertificate
	}
	return ""
}

func (x *ExternalServiceSpec_Port_TlsConfig) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *ExternalServiceSpec_Port_TlsConfig) GetCaCertificates() string {
	if x != nil {
		return x.CaCertificates
	}
	return ""
}

var File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_rawDesc = []byte{
	0x0a, 0x5d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32,
	0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x53, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70,
	0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x82, 0x07, 0x0a, 0x13, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x73,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x47, 0x0a,
	0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52,
	0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x56, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x2a,
	0x0a, 0x11, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x6c, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x41, 0x6c, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x1a, 0xac, 0x04, 0x0a, 0x04, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x0b, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12,
	0x62, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x74, 0x6c,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x2e, 0x54, 0x6c, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x69, 0x64, 0x65,
	0x54, 0x6c, 0x73, 0x1a, 0xb3, 0x02, 0x0a, 0x09, 0x54, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6e, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x73, 0x6e, 0x69, 0x12, 0x58, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x44, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x2e,
	0x50, 0x6f, 0x72, 0x74, 0x2e, 0x54, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74,
	0x6c, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x2d, 0x0a,
	0x12, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x27, 0x0a,
	0x0f, 0x63, 0x61, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x22, 0x41, 0x0a, 0x08, 0x74, 0x6c, 0x73, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d,
	0x55, 0x54, 0x55, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x53, 0x54, 0x49, 0x4f,
	0x5f, 0x4d, 0x55, 0x54, 0x55, 0x41, 0x4c, 0x10, 0x03, 0x1a, 0x3b, 0x0a, 0x0d, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xea, 0x05, 0x0a, 0x15, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x40, 0x0a, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x47, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x12, 0x5e, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x12, 0x90, 0x01, 0x0a, 0x1c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x5f, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x64, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x1a, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x64, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x64, 0x0a, 0x1b, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x19, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x4c, 0x0a, 0x0f, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x0e, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x1a, 0x63, 0x0a, 0x0f, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3a,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x82,
	0x01, 0x0a, 0x1f, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x49, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x64, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0x58, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d,
	0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76,
	0x32, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_rawDescData = file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_goTypes = []interface{}{
	(ExternalServiceSpec_Port_TlsConfigTlsMode)(0), // 0: networking.gloo.solo.io.ExternalServiceSpec.Port.TlsConfig.tls_mode
	(*ExternalServiceSpec)(nil),                    // 1: networking.gloo.solo.io.ExternalServiceSpec
	(*ExternalServiceStatus)(nil),                  // 2: networking.gloo.solo.io.ExternalServiceStatus
	(*ExternalServiceSpec_Port)(nil),               // 3: networking.gloo.solo.io.ExternalServiceSpec.Port
	nil,                                            // 4: networking.gloo.solo.io.ExternalServiceSpec.SelectorEntry
	(*ExternalServiceSpec_Port_TlsConfig)(nil),     // 5: networking.gloo.solo.io.ExternalServiceSpec.Port.TlsConfig
	nil,                            // 6: networking.gloo.solo.io.ExternalServiceStatus.WorkspacesEntry
	nil,                            // 7: networking.gloo.solo.io.ExternalServiceStatus.AppliedDestinationPoliciesEntry
	(*v2.GenericGlobalStatus)(nil), // 8: common.gloo.solo.io.GenericGlobalStatus
	(*v2.ObjectReference)(nil),     // 9: common.gloo.solo.io.ObjectReference
	(*v2.OwnerWorkspace)(nil),      // 10: common.gloo.solo.io.OwnerWorkspace
	(*v2.PortSelector)(nil),        // 11: common.gloo.solo.io.PortSelector
	(*v2.WorkspaceStatus)(nil),     // 12: common.gloo.solo.io.WorkspaceStatus
	(*v2.AppliedDestinationPortPolicies)(nil), // 13: common.gloo.solo.io.AppliedDestinationPortPolicies
}
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_depIdxs = []int32{
	3,  // 0: networking.gloo.solo.io.ExternalServiceSpec.ports:type_name -> networking.gloo.solo.io.ExternalServiceSpec.Port
	4,  // 1: networking.gloo.solo.io.ExternalServiceSpec.selector:type_name -> networking.gloo.solo.io.ExternalServiceSpec.SelectorEntry
	8,  // 2: networking.gloo.solo.io.ExternalServiceStatus.global:type_name -> common.gloo.solo.io.GenericGlobalStatus
	6,  // 3: networking.gloo.solo.io.ExternalServiceStatus.workspaces:type_name -> networking.gloo.solo.io.ExternalServiceStatus.WorkspacesEntry
	7,  // 4: networking.gloo.solo.io.ExternalServiceStatus.applied_destination_policies:type_name -> networking.gloo.solo.io.ExternalServiceStatus.AppliedDestinationPoliciesEntry
	9,  // 5: networking.gloo.solo.io.ExternalServiceStatus.selected_external_endpoints:type_name -> common.gloo.solo.io.ObjectReference
	10, // 6: networking.gloo.solo.io.ExternalServiceStatus.owner_workspace:type_name -> common.gloo.solo.io.OwnerWorkspace
	11, // 7: networking.gloo.solo.io.ExternalServiceSpec.Port.target_port:type_name -> common.gloo.solo.io.PortSelector
	5,  // 8: networking.gloo.solo.io.ExternalServiceSpec.Port.clientside_tls:type_name -> networking.gloo.solo.io.ExternalServiceSpec.Port.TlsConfig
	0,  // 9: networking.gloo.solo.io.ExternalServiceSpec.Port.TlsConfig.mode:type_name -> networking.gloo.solo.io.ExternalServiceSpec.Port.TlsConfig.tls_mode
	12, // 10: networking.gloo.solo.io.ExternalServiceStatus.WorkspacesEntry.value:type_name -> common.gloo.solo.io.WorkspaceStatus
	13, // 11: networking.gloo.solo.io.ExternalServiceStatus.AppliedDestinationPoliciesEntry.value:type_name -> common.gloo.solo.io.AppliedDestinationPortPolicies
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_init()
}
func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_init() {
	if File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalServiceSpec); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalServiceStatus); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalServiceSpec_Port); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalServiceSpec_Port_TlsConfig); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto = out.File
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_networking_v2_external_service_proto_depIdxs = nil
}
