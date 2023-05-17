// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo/gateway/v1/matchable_tcp_gateway.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v3 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3"
	ssl "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/ssl"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MatchableTcpGatewayStatus_State int32

const (
	// Pending status indicates the resource has not yet been validated
	MatchableTcpGatewayStatus_Pending MatchableTcpGatewayStatus_State = 0
	// Accepted indicates the resource has been validated
	MatchableTcpGatewayStatus_Accepted MatchableTcpGatewayStatus_State = 1
	// Rejected indicates an invalid configuration by the user
	// Rejected resources may be propagated to the xDS server depending on their severity
	MatchableTcpGatewayStatus_Rejected MatchableTcpGatewayStatus_State = 2
	// Warning indicates a partially invalid configuration by the user
	// Resources with Warnings may be partially accepted by a controller, depending on the implementation
	MatchableTcpGatewayStatus_Warning MatchableTcpGatewayStatus_State = 3
)

// Enum value maps for MatchableTcpGatewayStatus_State.
var (
	MatchableTcpGatewayStatus_State_name = map[int32]string{
		0: "Pending",
		1: "Accepted",
		2: "Rejected",
		3: "Warning",
	}
	MatchableTcpGatewayStatus_State_value = map[string]int32{
		"Pending":  0,
		"Accepted": 1,
		"Rejected": 2,
		"Warning":  3,
	}
)

func (x MatchableTcpGatewayStatus_State) Enum() *MatchableTcpGatewayStatus_State {
	p := new(MatchableTcpGatewayStatus_State)
	*p = x
	return p
}

func (x MatchableTcpGatewayStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MatchableTcpGatewayStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_enumTypes[0].Descriptor()
}

func (MatchableTcpGatewayStatus_State) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_enumTypes[0]
}

func (x MatchableTcpGatewayStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MatchableTcpGatewayStatus_State.Descriptor instead.
func (MatchableTcpGatewayStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_rawDescGZIP(), []int{1, 0}
}

// A MatchableTcpGateway describes a single FilterChain configured with the TcpProxy network filter and a matcher.
//
// A Gateway CR may select one or more MatchableTcpGateways on a single listener.
// This enables separate teams to own Listener configuration (Gateway CR)
// and FilterChain configuration (MatchableTcpGateway CR).
type MatchableTcpGatewaySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Matcher creates a FilterChainMatch and TransportSocket for a FilterChain
	// For each MatchableTcpGateway on a Gateway CR, the matcher must be unique.
	// If there are any identical matchers, the Gateway will be rejected.
	// An empty matcher will produce an empty FilterChainMatch (https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/listener/v3/listener_components.proto#envoy-v3-api-msg-config-listener-v3-filterchainmatch)
	// effectively matching all incoming connections
	Matcher *MatchableTcpGatewaySpec_Matcher `protobuf:"bytes,3,opt,name=matcher,proto3" json:"matcher,omitempty"`
	// TcpGateway creates a FilterChain with a TcpProxy.
	TcpGateway *TcpGateway `protobuf:"bytes,4,opt,name=tcp_gateway,json=tcpGateway,proto3" json:"tcp_gateway,omitempty"`
}

func (x *MatchableTcpGatewaySpec) Reset() {
	*x = MatchableTcpGatewaySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchableTcpGatewaySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchableTcpGatewaySpec) ProtoMessage() {}

func (x *MatchableTcpGatewaySpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchableTcpGatewaySpec.ProtoReflect.Descriptor instead.
func (*MatchableTcpGatewaySpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *MatchableTcpGatewaySpec) GetMatcher() *MatchableTcpGatewaySpec_Matcher {
	if x != nil {
		return x.Matcher
	}
	return nil
}

func (x *MatchableTcpGatewaySpec) GetTcpGateway() *TcpGateway {
	if x != nil {
		return x.TcpGateway
	}
	return nil
}

type MatchableTcpGatewayStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// State is the enum indicating the state of the resource
	State MatchableTcpGatewayStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=gateway.solo.io.MatchableTcpGatewayStatus_State" json:"state,omitempty"`
	// Reason is a description of the error for Rejected resources. If the resource is pending or accepted, this field will be empty
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	// Reference to the reporter who wrote this status
	ReportedBy string `protobuf:"bytes,3,opt,name=reported_by,json=reportedBy,proto3" json:"reported_by,omitempty"`
	// Reference to statuses (by resource-ref string: "Kind.Namespace.Name") of subresources of the parent resource
	SubresourceStatuses map[string]*MatchableTcpGatewayStatus `protobuf:"bytes,4,rep,name=subresource_statuses,json=subresourceStatuses,proto3" json:"subresource_statuses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Opaque details about status results
	Details *_struct.Struct `protobuf:"bytes,5,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *MatchableTcpGatewayStatus) Reset() {
	*x = MatchableTcpGatewayStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchableTcpGatewayStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchableTcpGatewayStatus) ProtoMessage() {}

func (x *MatchableTcpGatewayStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchableTcpGatewayStatus.ProtoReflect.Descriptor instead.
func (*MatchableTcpGatewayStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *MatchableTcpGatewayStatus) GetState() MatchableTcpGatewayStatus_State {
	if x != nil {
		return x.State
	}
	return MatchableTcpGatewayStatus_Pending
}

func (x *MatchableTcpGatewayStatus) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *MatchableTcpGatewayStatus) GetReportedBy() string {
	if x != nil {
		return x.ReportedBy
	}
	return ""
}

func (x *MatchableTcpGatewayStatus) GetSubresourceStatuses() map[string]*MatchableTcpGatewayStatus {
	if x != nil {
		return x.SubresourceStatuses
	}
	return nil
}

func (x *MatchableTcpGatewayStatus) GetDetails() *_struct.Struct {
	if x != nil {
		return x.Details
	}
	return nil
}

type MatchableTcpGatewayNamespacedStatuses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Statuses map[string]*MatchableTcpGatewayStatus `protobuf:"bytes,1,rep,name=statuses,proto3" json:"statuses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MatchableTcpGatewayNamespacedStatuses) Reset() {
	*x = MatchableTcpGatewayNamespacedStatuses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchableTcpGatewayNamespacedStatuses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchableTcpGatewayNamespacedStatuses) ProtoMessage() {}

func (x *MatchableTcpGatewayNamespacedStatuses) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchableTcpGatewayNamespacedStatuses.ProtoReflect.Descriptor instead.
func (*MatchableTcpGatewayNamespacedStatuses) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *MatchableTcpGatewayNamespacedStatuses) GetStatuses() map[string]*MatchableTcpGatewayStatus {
	if x != nil {
		return x.Statuses
	}
	return nil
}

type MatchableTcpGatewaySpec_Matcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CidrRange specifies an IP Address and a prefix length to construct the subnet mask for a CIDR range.
	// See https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/address.proto#envoy-v3-api-msg-config-core-v3-cidrrange
	SourcePrefixRanges []*v3.CidrRange `protobuf:"bytes,1,rep,name=source_prefix_ranges,json=sourcePrefixRanges,proto3" json:"source_prefix_ranges,omitempty"`
	// Ssl configuration applied to the FilterChain, if using passthrough should not include secrets :
	//   - FilterChainMatch: https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/listener/v3/listener_components.proto#config-listener-v3-filterchainmatch)
	//   - TransportSocket: https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/base.proto#envoy-v3-api-msg-config-core-v3-transportsocket
	SslConfig *ssl.SslConfig `protobuf:"bytes,2,opt,name=ssl_config,json=sslConfig,proto3" json:"ssl_config,omitempty"`
	// Enterprise-only: Passthrough cipher suites is an allow-list of OpenSSL cipher suite names for which TLS passthrough will be enabled.
	// If a client does not support any ciphers that are natively supported by Envoy, but does support one of the ciphers in the passthrough list,
	// then traffic will be routed via TCP Proxy to a destination specified by the TcpGateway, where TLS can then be terminated.
	PassthroughCipherSuites []string `protobuf:"bytes,3,rep,name=passthrough_cipher_suites,json=passthroughCipherSuites,proto3" json:"passthrough_cipher_suites,omitempty"`
}

func (x *MatchableTcpGatewaySpec_Matcher) Reset() {
	*x = MatchableTcpGatewaySpec_Matcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchableTcpGatewaySpec_Matcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchableTcpGatewaySpec_Matcher) ProtoMessage() {}

func (x *MatchableTcpGatewaySpec_Matcher) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchableTcpGatewaySpec_Matcher.ProtoReflect.Descriptor instead.
func (*MatchableTcpGatewaySpec_Matcher) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MatchableTcpGatewaySpec_Matcher) GetSourcePrefixRanges() []*v3.CidrRange {
	if x != nil {
		return x.SourcePrefixRanges
	}
	return nil
}

func (x *MatchableTcpGatewaySpec_Matcher) GetSslConfig() *ssl.SslConfig {
	if x != nil {
		return x.SslConfig
	}
	return nil
}

func (x *MatchableTcpGatewaySpec_Matcher) GetPassthroughCipherSuites() []string {
	if x != nil {
		return x.PassthroughCipherSuites
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x63, 0x70,
	0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65,
	0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x73, 0x6c, 0x2f, 0x73, 0x73, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x02, 0x0a, 0x17, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x63, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x4a, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65,
	0x54, 0x63, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12,
	0x3c, 0x0a, 0x0b, 0x74, 0x63, 0x70, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x54, 0x63, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x52, 0x0a, 0x74, 0x63, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x1a, 0xd8, 0x01,
	0x0a, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x59, 0x0a, 0x14, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x69, 0x64, 0x72, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x12, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x0a, 0x73, 0x73, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x73, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x09, 0x73, 0x73, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3a, 0x0a, 0x19,
	0x70, 0x61, 0x73, 0x73, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x5f, 0x63, 0x69, 0x70, 0x68,
	0x65, 0x72, 0x5f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x17, 0x70, 0x61, 0x73, 0x73, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x43, 0x69, 0x70, 0x68,
	0x65, 0x72, 0x53, 0x75, 0x69, 0x74, 0x65, 0x73, 0x22, 0xfa, 0x03, 0x0a, 0x19, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x63, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x46, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x61, 0x62, 0x6c,
	0x65, 0x54, 0x63, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x76, 0x0a, 0x14, 0x73, 0x75, 0x62, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x61, 0x62, 0x6c,
	0x65, 0x54, 0x63, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x53, 0x75, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x13, 0x73, 0x75, 0x62, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x12,
	0x31, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x1a, 0x72, 0x0a, 0x18, 0x53, 0x75, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x40, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x63, 0x70, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3d, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x61, 0x72, 0x6e,
	0x69, 0x6e, 0x67, 0x10, 0x03, 0x22, 0xf2, 0x01, 0x0a, 0x25, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x61,
	0x62, 0x6c, 0x65, 0x54, 0x63, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x12,
	0x60, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x44, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x63, 0x70,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65,
	0x73, 0x1a, 0x67, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x40, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x54,
	0x63, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x45, 0x5a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2f, 0x76, 0x31, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04,
	0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_goTypes = []interface{}{
	(MatchableTcpGatewayStatus_State)(0),          // 0: gateway.solo.io.MatchableTcpGatewayStatus.State
	(*MatchableTcpGatewaySpec)(nil),               // 1: gateway.solo.io.MatchableTcpGatewaySpec
	(*MatchableTcpGatewayStatus)(nil),             // 2: gateway.solo.io.MatchableTcpGatewayStatus
	(*MatchableTcpGatewayNamespacedStatuses)(nil), // 3: gateway.solo.io.MatchableTcpGatewayNamespacedStatuses
	(*MatchableTcpGatewaySpec_Matcher)(nil),       // 4: gateway.solo.io.MatchableTcpGatewaySpec.Matcher
	nil,                                           // 5: gateway.solo.io.MatchableTcpGatewayStatus.SubresourceStatusesEntry
	nil,                                           // 6: gateway.solo.io.MatchableTcpGatewayNamespacedStatuses.StatusesEntry
	(*TcpGateway)(nil),                            // 7: gateway.solo.io.TcpGateway
	(*_struct.Struct)(nil),                        // 8: google.protobuf.Struct
	(*v3.CidrRange)(nil),                          // 9: solo.io.envoy.config.core.v3.CidrRange
	(*ssl.SslConfig)(nil),                         // 10: gloo.solo.io.SslConfig
}
var file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_depIdxs = []int32{
	4,  // 0: gateway.solo.io.MatchableTcpGatewaySpec.matcher:type_name -> gateway.solo.io.MatchableTcpGatewaySpec.Matcher
	7,  // 1: gateway.solo.io.MatchableTcpGatewaySpec.tcp_gateway:type_name -> gateway.solo.io.TcpGateway
	0,  // 2: gateway.solo.io.MatchableTcpGatewayStatus.state:type_name -> gateway.solo.io.MatchableTcpGatewayStatus.State
	5,  // 3: gateway.solo.io.MatchableTcpGatewayStatus.subresource_statuses:type_name -> gateway.solo.io.MatchableTcpGatewayStatus.SubresourceStatusesEntry
	8,  // 4: gateway.solo.io.MatchableTcpGatewayStatus.details:type_name -> google.protobuf.Struct
	6,  // 5: gateway.solo.io.MatchableTcpGatewayNamespacedStatuses.statuses:type_name -> gateway.solo.io.MatchableTcpGatewayNamespacedStatuses.StatusesEntry
	9,  // 6: gateway.solo.io.MatchableTcpGatewaySpec.Matcher.source_prefix_ranges:type_name -> solo.io.envoy.config.core.v3.CidrRange
	10, // 7: gateway.solo.io.MatchableTcpGatewaySpec.Matcher.ssl_config:type_name -> gloo.solo.io.SslConfig
	2,  // 8: gateway.solo.io.MatchableTcpGatewayStatus.SubresourceStatusesEntry.value:type_name -> gateway.solo.io.MatchableTcpGatewayStatus
	2,  // 9: gateway.solo.io.MatchableTcpGatewayNamespacedStatuses.StatusesEntry.value:type_name -> gateway.solo.io.MatchableTcpGatewayStatus
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_init() }
func file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto != nil {
		return
	}
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchableTcpGatewaySpec); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchableTcpGatewayStatus); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchableTcpGatewayNamespacedStatuses); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchableTcpGatewaySpec_Matcher); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_tcp_gateway_proto_depIdxs = nil
}
