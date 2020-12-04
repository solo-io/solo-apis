// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo/gateway/v1/gateway.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
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

type GatewayStatus_State int32

const (
	// Pending status indicates the resource has not yet been validated
	GatewayStatus_Pending GatewayStatus_State = 0
	// Accepted indicates the resource has been validated
	GatewayStatus_Accepted GatewayStatus_State = 1
	// Rejected indicates an invalid configuration by the user
	// Rejected resources may be propagated to the xDS server depending on their severity
	GatewayStatus_Rejected GatewayStatus_State = 2
	// Warning indicates a partially invalid configuration by the user
	// Resources with Warnings may be partially accepted by a controller, depending on the implementation
	GatewayStatus_Warning GatewayStatus_State = 3
)

// Enum value maps for GatewayStatus_State.
var (
	GatewayStatus_State_name = map[int32]string{
		0: "Pending",
		1: "Accepted",
		2: "Rejected",
		3: "Warning",
	}
	GatewayStatus_State_value = map[string]int32{
		"Pending":  0,
		"Accepted": 1,
		"Rejected": 2,
		"Warning":  3,
	}
)

func (x GatewayStatus_State) Enum() *GatewayStatus_State {
	p := new(GatewayStatus_State)
	*p = x
	return p
}

func (x GatewayStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GatewayStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_enumTypes[0].Descriptor()
}

func (GatewayStatus_State) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_enumTypes[0]
}

func (x GatewayStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GatewayStatus_State.Descriptor instead.
func (GatewayStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_rawDescGZIP(), []int{3, 0}
}

//
//A Gateway describes a single Listener (bind address:port)
//and the routing configuration to upstreams that are reachable via a specific port on the Gateway Proxy itself.
type GatewaySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// if set to false, only use virtual services without ssl configured.
	// if set to true, only use virtual services with ssl configured.
	Ssl bool `protobuf:"varint,1,opt,name=ssl,proto3" json:"ssl,omitempty"`
	// the bind address the gateway should serve traffic on
	BindAddress string `protobuf:"bytes,3,opt,name=bind_address,json=bindAddress,proto3" json:"bind_address,omitempty"`
	// bind ports must not conflict across gateways for a single proxy
	BindPort uint32 `protobuf:"varint,4,opt,name=bind_port,json=bindPort,proto3" json:"bind_port,omitempty"`
	// top level optional configuration for all routes on the gateway
	Options *v1.ListenerOptions `protobuf:"bytes,5,opt,name=options,proto3" json:"options,omitempty"`
	// Enable ProxyProtocol support for this listener
	UseProxyProto *wrappers.BoolValue `protobuf:"bytes,8,opt,name=use_proxy_proto,json=useProxyProto,proto3" json:"use_proxy_proto,omitempty"`
	// The type of gateway being created
	// HttpGateway creates a listener with an http_connection_manager
	// TcpGateway creates a listener with a tcp proxy filter
	//
	// Types that are assignable to GatewayType:
	//	*GatewaySpec_HttpGateway
	//	*GatewaySpec_TcpGateway
	GatewayType isGatewaySpec_GatewayType `protobuf_oneof:"GatewayType"`
	//
	// Names of the [`Proxy`](https://gloo.solo.io/api/github.com/solo-io/solo-apis/api/gloo/gloo/v1/proxy.proto.sk/)
	// resources to generate from this gateway. If other gateways exist which point to the same proxy,
	// Gloo will join them together.
	//
	// Proxies have a one-to-many relationship with Envoy bootstrap configuration.
	// In order to connect to Gloo, the Envoy bootstrap configuration sets a `role` in
	// the [node metadata](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/core/base.proto#envoy-api-msg-core-node)
	// Envoy instances announce their `role` to Gloo, which maps to the  `{{ .Namespace }}~{{ .Name }}`
	// of the Proxy resource.
	//
	// The template for this value can be seen in the [Gloo Helm chart](https://github.com/solo-io/gloo/blob/master/install/helm/gloo/templates/9-gateway-proxy-configmap.yaml#L22)
	//
	// Note: this field also accepts fields written in camel-case. They will be converted
	// to kebab-case in the Proxy name. This allows use of the [Gateway Name Helm value](https://github.com/solo-io/gloo/blob/master/install/helm/gloo/values-gateway-template.yaml#L47)
	// for this field
	//
	// Defaults to `["gateway-proxy"]`
	ProxyNames []string `protobuf:"bytes,12,rep,name=proxy_names,json=proxyNames,proto3" json:"proxy_names,omitempty"`
}

func (x *GatewaySpec) Reset() {
	*x = GatewaySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewaySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewaySpec) ProtoMessage() {}

func (x *GatewaySpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewaySpec.ProtoReflect.Descriptor instead.
func (*GatewaySpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *GatewaySpec) GetSsl() bool {
	if x != nil {
		return x.Ssl
	}
	return false
}

func (x *GatewaySpec) GetBindAddress() string {
	if x != nil {
		return x.BindAddress
	}
	return ""
}

func (x *GatewaySpec) GetBindPort() uint32 {
	if x != nil {
		return x.BindPort
	}
	return 0
}

func (x *GatewaySpec) GetOptions() *v1.ListenerOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *GatewaySpec) GetUseProxyProto() *wrappers.BoolValue {
	if x != nil {
		return x.UseProxyProto
	}
	return nil
}

func (m *GatewaySpec) GetGatewayType() isGatewaySpec_GatewayType {
	if m != nil {
		return m.GatewayType
	}
	return nil
}

func (x *GatewaySpec) GetHttpGateway() *HttpGateway {
	if x, ok := x.GetGatewayType().(*GatewaySpec_HttpGateway); ok {
		return x.HttpGateway
	}
	return nil
}

func (x *GatewaySpec) GetTcpGateway() *TcpGateway {
	if x, ok := x.GetGatewayType().(*GatewaySpec_TcpGateway); ok {
		return x.TcpGateway
	}
	return nil
}

func (x *GatewaySpec) GetProxyNames() []string {
	if x != nil {
		return x.ProxyNames
	}
	return nil
}

type isGatewaySpec_GatewayType interface {
	isGatewaySpec_GatewayType()
}

type GatewaySpec_HttpGateway struct {
	HttpGateway *HttpGateway `protobuf:"bytes,9,opt,name=http_gateway,json=httpGateway,proto3,oneof"`
}

type GatewaySpec_TcpGateway struct {
	TcpGateway *TcpGateway `protobuf:"bytes,10,opt,name=tcp_gateway,json=tcpGateway,proto3,oneof"`
}

func (*GatewaySpec_HttpGateway) isGatewaySpec_GatewayType() {}

func (*GatewaySpec_TcpGateway) isGatewaySpec_GatewayType() {}

type HttpGateway struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Names & namespace refs of the virtual services which contain the actual routes for the gateway.
	// If the list is empty, all virtual services in all namespaces that Gloo watches will apply,
	// with accordance to `ssl` flag on `Gateway` above.
	// The default namespace matching behavior can be overridden via `virtual_service_namespaces` flag below.
	// Only one of `virtualServices` or `virtualServiceSelector` should be provided.
	VirtualServices []*core.ResourceRef `protobuf:"bytes,1,rep,name=virtual_services,json=virtualServices,proto3" json:"virtual_services,omitempty"`
	// Select virtual services by their label. If `virtual_service_namespaces` is provided below, this will apply only
	// to virtual services in the namespaces specified.
	// Only one of `virtualServices` or `virtualServiceSelector` should be provided.
	VirtualServiceSelector map[string]string `protobuf:"bytes,2,rep,name=virtual_service_selector,json=virtualServiceSelector,proto3" json:"virtual_service_selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Restrict the search by providing a list of valid search namespaces here.
	// Setting '*' will search all namespaces, equivalent to omitting this value.
	VirtualServiceNamespaces []string `protobuf:"bytes,3,rep,name=virtual_service_namespaces,json=virtualServiceNamespaces,proto3" json:"virtual_service_namespaces,omitempty"`
	// HTTP Gateway configuration
	Options *v1.HttpListenerOptions `protobuf:"bytes,8,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *HttpGateway) Reset() {
	*x = HttpGateway{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpGateway) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpGateway) ProtoMessage() {}

func (x *HttpGateway) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpGateway.ProtoReflect.Descriptor instead.
func (*HttpGateway) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *HttpGateway) GetVirtualServices() []*core.ResourceRef {
	if x != nil {
		return x.VirtualServices
	}
	return nil
}

func (x *HttpGateway) GetVirtualServiceSelector() map[string]string {
	if x != nil {
		return x.VirtualServiceSelector
	}
	return nil
}

func (x *HttpGateway) GetVirtualServiceNamespaces() []string {
	if x != nil {
		return x.VirtualServiceNamespaces
	}
	return nil
}

func (x *HttpGateway) GetOptions() *v1.HttpListenerOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type TcpGateway struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TCP hosts that the gateway can route to
	TcpHosts []*v1.TcpHost `protobuf:"bytes,1,rep,name=tcp_hosts,json=tcpHosts,proto3" json:"tcp_hosts,omitempty"`
	// TCP Gateway configuration
	Options *v1.TcpListenerOptions `protobuf:"bytes,8,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *TcpGateway) Reset() {
	*x = TcpGateway{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TcpGateway) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TcpGateway) ProtoMessage() {}

func (x *TcpGateway) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TcpGateway.ProtoReflect.Descriptor instead.
func (*TcpGateway) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *TcpGateway) GetTcpHosts() []*v1.TcpHost {
	if x != nil {
		return x.TcpHosts
	}
	return nil
}

func (x *TcpGateway) GetOptions() *v1.TcpListenerOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type GatewayStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// State is the enum indicating the state of the resource
	State GatewayStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=gateway.solo.io.GatewayStatus_State" json:"state,omitempty"`
	// Reason is a description of the error for Rejected resources. If the resource is pending or accepted, this field will be empty
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	// Reference to the reporter who wrote this status
	ReportedBy string `protobuf:"bytes,3,opt,name=reported_by,json=reportedBy,proto3" json:"reported_by,omitempty"`
	// Reference to statuses (by resource-ref string: "Kind.Namespace.Name") of subresources of the parent resource
	SubresourceStatuses map[string]*GatewayStatus `protobuf:"bytes,4,rep,name=subresource_statuses,json=subresourceStatuses,proto3" json:"subresource_statuses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Opaque details about status results
	Details *_struct.Struct `protobuf:"bytes,5,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *GatewayStatus) Reset() {
	*x = GatewayStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayStatus) ProtoMessage() {}

func (x *GatewayStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayStatus.ProtoReflect.Descriptor instead.
func (*GatewayStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_rawDescGZIP(), []int{3}
}

func (x *GatewayStatus) GetState() GatewayStatus_State {
	if x != nil {
		return x.State
	}
	return GatewayStatus_Pending
}

func (x *GatewayStatus) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *GatewayStatus) GetReportedBy() string {
	if x != nil {
		return x.ReportedBy
	}
	return ""
}

func (x *GatewayStatus) GetSubresourceStatuses() map[string]*GatewayStatus {
	if x != nil {
		return x.SubresourceStatuses
	}
	return nil
}

func (x *GatewayStatus) GetDetails() *_struct.Struct {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x03, 0x0a,
	0x0b, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x73, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x73, 0x73, 0x6c, 0x12, 0x21,
	0x0a, 0x0c, 0x62, 0x69, 0x6e, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x69, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x37,
	0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x42, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x75, 0x73,
	0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x41, 0x0a, 0x0c, 0x68,
	0x74, 0x74, 0x70, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x48,
	0x00, 0x52, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x3e,
	0x0a, 0x0b, 0x74, 0x63, 0x70, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x54, 0x63, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x48, 0x00, 0x52, 0x0a, 0x74, 0x63, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x0c, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x42,
	0x0d, 0x0a, 0x0b, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0x8d,
	0x03, 0x0a, 0x0b, 0x48, 0x74, 0x74, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x44,
	0x0a, 0x10, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x66, 0x52, 0x0f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x12, 0x72, 0x0a, 0x18, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x16, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x3c, 0x0a, 0x1a, 0x76, 0x69, 0x72, 0x74,
	0x75, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x18, 0x76, 0x69,
	0x72, 0x74, 0x75, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x1a, 0x49, 0x0a, 0x1b, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x7c,
	0x0a, 0x0a, 0x54, 0x63, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x32, 0x0a, 0x09,
	0x74, 0x63, 0x70, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x54,
	0x63, 0x70, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x08, 0x74, 0x63, 0x70, 0x48, 0x6f, 0x73, 0x74, 0x73,
	0x12, 0x3a, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x54, 0x63, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xca, 0x03, 0x0a,
	0x0d, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3a,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x6a, 0x0a, 0x14, 0x73, 0x75, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x37, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x53, 0x75, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x13, 0x73, 0x75, 0x62, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x12,
	0x31, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x1a, 0x66, 0x0a, 0x18, 0x53, 0x75, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x34, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3d, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x00,
	0x12, 0x0c, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07,
	0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x03, 0x42, 0x41, 0x5a, 0x37, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2f, 0x76, 0x31, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_goTypes = []interface{}{
	(GatewayStatus_State)(0),       // 0: gateway.solo.io.GatewayStatus.State
	(*GatewaySpec)(nil),            // 1: gateway.solo.io.GatewaySpec
	(*HttpGateway)(nil),            // 2: gateway.solo.io.HttpGateway
	(*TcpGateway)(nil),             // 3: gateway.solo.io.TcpGateway
	(*GatewayStatus)(nil),          // 4: gateway.solo.io.GatewayStatus
	nil,                            // 5: gateway.solo.io.HttpGateway.VirtualServiceSelectorEntry
	nil,                            // 6: gateway.solo.io.GatewayStatus.SubresourceStatusesEntry
	(*v1.ListenerOptions)(nil),     // 7: gloo.solo.io.ListenerOptions
	(*wrappers.BoolValue)(nil),     // 8: google.protobuf.BoolValue
	(*core.ResourceRef)(nil),       // 9: core.solo.io.ResourceRef
	(*v1.HttpListenerOptions)(nil), // 10: gloo.solo.io.HttpListenerOptions
	(*v1.TcpHost)(nil),             // 11: gloo.solo.io.TcpHost
	(*v1.TcpListenerOptions)(nil),  // 12: gloo.solo.io.TcpListenerOptions
	(*_struct.Struct)(nil),         // 13: google.protobuf.Struct
}
var file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_depIdxs = []int32{
	7,  // 0: gateway.solo.io.GatewaySpec.options:type_name -> gloo.solo.io.ListenerOptions
	8,  // 1: gateway.solo.io.GatewaySpec.use_proxy_proto:type_name -> google.protobuf.BoolValue
	2,  // 2: gateway.solo.io.GatewaySpec.http_gateway:type_name -> gateway.solo.io.HttpGateway
	3,  // 3: gateway.solo.io.GatewaySpec.tcp_gateway:type_name -> gateway.solo.io.TcpGateway
	9,  // 4: gateway.solo.io.HttpGateway.virtual_services:type_name -> core.solo.io.ResourceRef
	5,  // 5: gateway.solo.io.HttpGateway.virtual_service_selector:type_name -> gateway.solo.io.HttpGateway.VirtualServiceSelectorEntry
	10, // 6: gateway.solo.io.HttpGateway.options:type_name -> gloo.solo.io.HttpListenerOptions
	11, // 7: gateway.solo.io.TcpGateway.tcp_hosts:type_name -> gloo.solo.io.TcpHost
	12, // 8: gateway.solo.io.TcpGateway.options:type_name -> gloo.solo.io.TcpListenerOptions
	0,  // 9: gateway.solo.io.GatewayStatus.state:type_name -> gateway.solo.io.GatewayStatus.State
	6,  // 10: gateway.solo.io.GatewayStatus.subresource_statuses:type_name -> gateway.solo.io.GatewayStatus.SubresourceStatusesEntry
	13, // 11: gateway.solo.io.GatewayStatus.details:type_name -> google.protobuf.Struct
	4,  // 12: gateway.solo.io.GatewayStatus.SubresourceStatusesEntry.value:type_name -> gateway.solo.io.GatewayStatus
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_init() }
func file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewaySpec); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpGateway); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TcpGateway); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayStatus); i {
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
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*GatewaySpec_HttpGateway)(nil),
		(*GatewaySpec_TcpGateway)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_gateway_proto_depIdxs = nil
}
