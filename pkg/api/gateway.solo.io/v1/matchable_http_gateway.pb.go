// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo/gateway/v1/matchable_http_gateway.proto

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

type MatchableHttpGatewayStatus_State int32

const (
	// Pending status indicates the resource has not yet been validated
	MatchableHttpGatewayStatus_Pending MatchableHttpGatewayStatus_State = 0
	// Accepted indicates the resource has been validated
	MatchableHttpGatewayStatus_Accepted MatchableHttpGatewayStatus_State = 1
	// Rejected indicates an invalid configuration by the user
	// Rejected resources may be propagated to the xDS server depending on their severity
	MatchableHttpGatewayStatus_Rejected MatchableHttpGatewayStatus_State = 2
	// Warning indicates a partially invalid configuration by the user
	// Resources with Warnings may be partially accepted by a controller, depending on the implementation
	MatchableHttpGatewayStatus_Warning MatchableHttpGatewayStatus_State = 3
)

// Enum value maps for MatchableHttpGatewayStatus_State.
var (
	MatchableHttpGatewayStatus_State_name = map[int32]string{
		0: "Pending",
		1: "Accepted",
		2: "Rejected",
		3: "Warning",
	}
	MatchableHttpGatewayStatus_State_value = map[string]int32{
		"Pending":  0,
		"Accepted": 1,
		"Rejected": 2,
		"Warning":  3,
	}
)

func (x MatchableHttpGatewayStatus_State) Enum() *MatchableHttpGatewayStatus_State {
	p := new(MatchableHttpGatewayStatus_State)
	*p = x
	return p
}

func (x MatchableHttpGatewayStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MatchableHttpGatewayStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_enumTypes[0].Descriptor()
}

func (MatchableHttpGatewayStatus_State) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_enumTypes[0]
}

func (x MatchableHttpGatewayStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MatchableHttpGatewayStatus_State.Descriptor instead.
func (MatchableHttpGatewayStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_rawDescGZIP(), []int{1, 0}
}

//
//A MatchableHttpGateway describes a single FilterChain configured with:
//- The HttpConnectionManager NetworkFilter
//- A FilterChainMatch and TransportSocket that support TLS configuration and Source IP matching
//
//A Gateway CR may select one or more MatchableHttpGateways on a single listener.
//This enables separate teams to own Listener configuration (Gateway CR)
//and FilterChain configuration (MatchableHttpGateway CR)
type MatchableHttpGatewaySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Matcher creates a FilterChainMatch and TransportSocket for a FilterChain
	// For each MatchableHttpGateway on a Gateway CR, the matcher must be unique.
	// If there are any identical matchers, the Gateway will be rejected.
	// An empty matcher will produce an empty FilterChainMatch (https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/listener/v3/listener_components.proto#envoy-v3-api-msg-config-listener-v3-filterchainmatch)
	// effectively matching all incoming connections
	Matcher *MatchableHttpGatewaySpec_Matcher `protobuf:"bytes,3,opt,name=matcher,proto3" json:"matcher,omitempty"`
	// HttpGateway creates a FilterChain with an HttpConnectionManager
	HttpGateway *HttpGateway `protobuf:"bytes,4,opt,name=http_gateway,json=httpGateway,proto3" json:"http_gateway,omitempty"`
}

func (x *MatchableHttpGatewaySpec) Reset() {
	*x = MatchableHttpGatewaySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchableHttpGatewaySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchableHttpGatewaySpec) ProtoMessage() {}

func (x *MatchableHttpGatewaySpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchableHttpGatewaySpec.ProtoReflect.Descriptor instead.
func (*MatchableHttpGatewaySpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *MatchableHttpGatewaySpec) GetMatcher() *MatchableHttpGatewaySpec_Matcher {
	if x != nil {
		return x.Matcher
	}
	return nil
}

func (x *MatchableHttpGatewaySpec) GetHttpGateway() *HttpGateway {
	if x != nil {
		return x.HttpGateway
	}
	return nil
}

type MatchableHttpGatewayStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// State is the enum indicating the state of the resource
	State MatchableHttpGatewayStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=gateway.solo.io.MatchableHttpGatewayStatus_State" json:"state,omitempty"`
	// Reason is a description of the error for Rejected resources. If the resource is pending or accepted, this field will be empty
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	// Reference to the reporter who wrote this status
	ReportedBy string `protobuf:"bytes,3,opt,name=reported_by,json=reportedBy,proto3" json:"reported_by,omitempty"`
	// Reference to statuses (by resource-ref string: "Kind.Namespace.Name") of subresources of the parent resource
	SubresourceStatuses map[string]*MatchableHttpGatewayStatus `protobuf:"bytes,4,rep,name=subresource_statuses,json=subresourceStatuses,proto3" json:"subresource_statuses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Opaque details about status results
	Details *_struct.Struct `protobuf:"bytes,5,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *MatchableHttpGatewayStatus) Reset() {
	*x = MatchableHttpGatewayStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchableHttpGatewayStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchableHttpGatewayStatus) ProtoMessage() {}

func (x *MatchableHttpGatewayStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchableHttpGatewayStatus.ProtoReflect.Descriptor instead.
func (*MatchableHttpGatewayStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *MatchableHttpGatewayStatus) GetState() MatchableHttpGatewayStatus_State {
	if x != nil {
		return x.State
	}
	return MatchableHttpGatewayStatus_Pending
}

func (x *MatchableHttpGatewayStatus) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *MatchableHttpGatewayStatus) GetReportedBy() string {
	if x != nil {
		return x.ReportedBy
	}
	return ""
}

func (x *MatchableHttpGatewayStatus) GetSubresourceStatuses() map[string]*MatchableHttpGatewayStatus {
	if x != nil {
		return x.SubresourceStatuses
	}
	return nil
}

func (x *MatchableHttpGatewayStatus) GetDetails() *_struct.Struct {
	if x != nil {
		return x.Details
	}
	return nil
}

type MatchableHttpGatewayNamespacedStatuses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Statuses map[string]*MatchableHttpGatewayStatus `protobuf:"bytes,1,rep,name=statuses,proto3" json:"statuses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MatchableHttpGatewayNamespacedStatuses) Reset() {
	*x = MatchableHttpGatewayNamespacedStatuses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchableHttpGatewayNamespacedStatuses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchableHttpGatewayNamespacedStatuses) ProtoMessage() {}

func (x *MatchableHttpGatewayNamespacedStatuses) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchableHttpGatewayNamespacedStatuses.ProtoReflect.Descriptor instead.
func (*MatchableHttpGatewayNamespacedStatuses) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *MatchableHttpGatewayNamespacedStatuses) GetStatuses() map[string]*MatchableHttpGatewayStatus {
	if x != nil {
		return x.Statuses
	}
	return nil
}

type MatchableHttpGatewaySpec_Matcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CidrRange specifies an IP Address and a prefix length to construct the subnet mask for a CIDR range.
	// See https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/address.proto#envoy-v3-api-msg-config-core-v3-cidrrange
	SourcePrefixRanges []*v3.CidrRange `protobuf:"bytes,1,rep,name=source_prefix_ranges,json=sourcePrefixRanges,proto3" json:"source_prefix_ranges,omitempty"`
	// Ssl configuration applied to the FilterChain:
	//  - FilterChainMatch: https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/listener/v3/listener_components.proto#config-listener-v3-filterchainmatch)
	//  - TransportSocket: https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/base.proto#envoy-v3-api-msg-config-core-v3-transportsocket
	SslConfig *ssl.SslConfig `protobuf:"bytes,2,opt,name=ssl_config,json=sslConfig,proto3" json:"ssl_config,omitempty"`
}

func (x *MatchableHttpGatewaySpec_Matcher) Reset() {
	*x = MatchableHttpGatewaySpec_Matcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchableHttpGatewaySpec_Matcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchableHttpGatewaySpec_Matcher) ProtoMessage() {}

func (x *MatchableHttpGatewaySpec_Matcher) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchableHttpGatewaySpec_Matcher.ProtoReflect.Descriptor instead.
func (*MatchableHttpGatewaySpec_Matcher) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MatchableHttpGatewaySpec_Matcher) GetSourcePrefixRanges() []*v3.CidrRange {
	if x != nil {
		return x.SourcePrefixRanges
	}
	return nil
}

func (x *MatchableHttpGatewaySpec_Matcher) GetSslConfig() *ssl.SslConfig {
	if x != nil {
		return x.SslConfig
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x68, 0x74, 0x74,
	0x70, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x73, 0x6c, 0x2f, 0x73, 0x73, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x02,
	0x0a, 0x18, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x74, 0x74, 0x70, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4b, 0x0a, 0x07, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x74, 0x74, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x07,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x5f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x48, 0x74, 0x74, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x0b, 0x68, 0x74, 0x74,
	0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x1a, 0x9c, 0x01, 0x0a, 0x07, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x12, 0x59, 0x0a, 0x14, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x33, 0x2e, 0x43, 0x69, 0x64, 0x72, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x12, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12,
	0x36, 0x0a, 0x0a, 0x73, 0x73, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x53, 0x73, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x09, 0x73, 0x73,
	0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xfe, 0x03, 0x0a, 0x1a, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x74, 0x74, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x47, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x61, 0x62, 0x6c,
	0x65, 0x48, 0x74, 0x74, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x77, 0x0a, 0x14, 0x73, 0x75, 0x62, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x61, 0x62,
	0x6c, 0x65, 0x48, 0x74, 0x74, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x13, 0x73, 0x75,
	0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65,
	0x73, 0x12, 0x31, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x1a, 0x73, 0x0a, 0x18, 0x53, 0x75, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x41, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x74, 0x74,
	0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3d, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0c, 0x0a,
	0x08, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x57,
	0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x03, 0x22, 0xf5, 0x01, 0x0a, 0x26, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x74, 0x74, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x65, 0x73, 0x12, 0x61, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x61, 0x62, 0x6c,
	0x65, 0x48, 0x74, 0x74, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x1a, 0x68, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x41, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x61, 0x62, 0x6c, 0x65, 0x48, 0x74, 0x74, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x45, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0xc0, 0xf5, 0x04, 0x01, 0xb8,
	0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_goTypes = []interface{}{
	(MatchableHttpGatewayStatus_State)(0),          // 0: gateway.solo.io.MatchableHttpGatewayStatus.State
	(*MatchableHttpGatewaySpec)(nil),               // 1: gateway.solo.io.MatchableHttpGatewaySpec
	(*MatchableHttpGatewayStatus)(nil),             // 2: gateway.solo.io.MatchableHttpGatewayStatus
	(*MatchableHttpGatewayNamespacedStatuses)(nil), // 3: gateway.solo.io.MatchableHttpGatewayNamespacedStatuses
	(*MatchableHttpGatewaySpec_Matcher)(nil),       // 4: gateway.solo.io.MatchableHttpGatewaySpec.Matcher
	nil,                                            // 5: gateway.solo.io.MatchableHttpGatewayStatus.SubresourceStatusesEntry
	nil,                                            // 6: gateway.solo.io.MatchableHttpGatewayNamespacedStatuses.StatusesEntry
	(*HttpGateway)(nil),                            // 7: gateway.solo.io.HttpGateway
	(*_struct.Struct)(nil),                         // 8: google.protobuf.Struct
	(*v3.CidrRange)(nil),                           // 9: solo.io.envoy.config.core.v3.CidrRange
	(*ssl.SslConfig)(nil),                          // 10: gloo.solo.io.SslConfig
}
var file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_depIdxs = []int32{
	4,  // 0: gateway.solo.io.MatchableHttpGatewaySpec.matcher:type_name -> gateway.solo.io.MatchableHttpGatewaySpec.Matcher
	7,  // 1: gateway.solo.io.MatchableHttpGatewaySpec.http_gateway:type_name -> gateway.solo.io.HttpGateway
	0,  // 2: gateway.solo.io.MatchableHttpGatewayStatus.state:type_name -> gateway.solo.io.MatchableHttpGatewayStatus.State
	5,  // 3: gateway.solo.io.MatchableHttpGatewayStatus.subresource_statuses:type_name -> gateway.solo.io.MatchableHttpGatewayStatus.SubresourceStatusesEntry
	8,  // 4: gateway.solo.io.MatchableHttpGatewayStatus.details:type_name -> google.protobuf.Struct
	6,  // 5: gateway.solo.io.MatchableHttpGatewayNamespacedStatuses.statuses:type_name -> gateway.solo.io.MatchableHttpGatewayNamespacedStatuses.StatusesEntry
	9,  // 6: gateway.solo.io.MatchableHttpGatewaySpec.Matcher.source_prefix_ranges:type_name -> solo.io.envoy.config.core.v3.CidrRange
	10, // 7: gateway.solo.io.MatchableHttpGatewaySpec.Matcher.ssl_config:type_name -> gloo.solo.io.SslConfig
	2,  // 8: gateway.solo.io.MatchableHttpGatewayStatus.SubresourceStatusesEntry.value:type_name -> gateway.solo.io.MatchableHttpGatewayStatus
	2,  // 9: gateway.solo.io.MatchableHttpGatewayNamespacedStatuses.StatusesEntry.value:type_name -> gateway.solo.io.MatchableHttpGatewayStatus
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto != nil {
		return
	}
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_http_gateway_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchableHttpGatewaySpec); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchableHttpGatewayStatus); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchableHttpGatewayNamespacedStatuses); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchableHttpGatewaySpec_Matcher); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_matchable_http_gateway_proto_depIdxs = nil
}
