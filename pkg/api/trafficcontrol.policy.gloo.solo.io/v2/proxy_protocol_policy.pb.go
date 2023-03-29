// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/policy/v2/trafficcontrol/proxy_protocol_policy.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v2 "github.com/solo-io/solo-apis/pkg/api/common.gloo.solo.io/v2"
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

// ProxyProtocolPolicy is used to enable proxy protocol for gateway listener(s).
//
// For more details on ProxyProtocol, visit:
// https://www.haproxy.org/download/1.8/doc/proxy-protocol.txt.
//
// ```
// apiVersion: trafficcontrol.policy.gloo.solo.io/v2
// kind: ProxyProtocolPolicy
// metadata:
//
//	name: my-policy
//	namespace: bookinfo
//
// spec:
//
//	config:
//	  enabledProxyProtocol: true
//	applyToListeners:
//	- virtualGateway:
//	    name: virtual-gateway
//	    namespace: bookinfo
//	    cluster: my-cluster
//	  port:
//	    number: 443
//
// ```
type ProxyProtocolPolicySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required: The VirtualGateway(s) and Port(s) to which this policy applies.
	ApplyToListeners []*ProxyProtocolPolicySpec_ListenerSelector `protobuf:"bytes,1,rep,name=apply_to_listeners,json=applyToListeners,proto3" json:"apply_to_listeners,omitempty"`
	// Required: The configuration to apply to the listener.
	Config *ProxyProtocolPolicySpec_Config `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *ProxyProtocolPolicySpec) Reset() {
	*x = ProxyProtocolPolicySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyProtocolPolicySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyProtocolPolicySpec) ProtoMessage() {}

func (x *ProxyProtocolPolicySpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyProtocolPolicySpec.ProtoReflect.Descriptor instead.
func (*ProxyProtocolPolicySpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_rawDescGZIP(), []int{0}
}

func (x *ProxyProtocolPolicySpec) GetApplyToListeners() []*ProxyProtocolPolicySpec_ListenerSelector {
	if x != nil {
		return x.ApplyToListeners
	}
	return nil
}

func (x *ProxyProtocolPolicySpec) GetConfig() *ProxyProtocolPolicySpec_Config {
	if x != nil {
		return x.Config
	}
	return nil
}

// reflects the status of the ProxyProtocolPolicy
type ProxyProtocolPolicyStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DEPRECATED: Use global instead
	//
	// Deprecated: Do not use.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// DEPRECATED: Use workspaces instead
	//
	// Deprecated: Do not use.
	State v2.ApprovalState `protobuf:"varint,2,opt,name=state,proto3,enum=common.gloo.solo.io.ApprovalState" json:"state,omitempty"`
	// The most recent generation observed in the the object's metadata.
	// If the `observedGeneration` does not match `metadata.generation`, Gloo Mesh
	// has not processed the most recent version of this object.
	Global *v2.GenericGlobalStatus `protobuf:"bytes,3,opt,name=global,proto3" json:"global,omitempty"`
	// Whether the resource has been accepted as valid and processed in the Gloo
	// Mesh config translation.
	// The status of the resource in each workspace that it exists in.
	Workspaces map[string]*v2.WorkspaceStatus `protobuf:"bytes,4,rep,name=workspaces,proto3" json:"workspaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A list of virtual gateway listeners selected by the policy.
	SelectedListeners []*v2.ListenerPortReference `protobuf:"bytes,5,rep,name=selected_listeners,json=selectedListeners,proto3" json:"selected_listeners,omitempty"`
}

func (x *ProxyProtocolPolicyStatus) Reset() {
	*x = ProxyProtocolPolicyStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyProtocolPolicyStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyProtocolPolicyStatus) ProtoMessage() {}

func (x *ProxyProtocolPolicyStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyProtocolPolicyStatus.ProtoReflect.Descriptor instead.
func (*ProxyProtocolPolicyStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_rawDescGZIP(), []int{1}
}

// Deprecated: Do not use.
func (x *ProxyProtocolPolicyStatus) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

// Deprecated: Do not use.
func (x *ProxyProtocolPolicyStatus) GetState() v2.ApprovalState {
	if x != nil {
		return x.State
	}
	return v2.ApprovalState_PENDING
}

func (x *ProxyProtocolPolicyStatus) GetGlobal() *v2.GenericGlobalStatus {
	if x != nil {
		return x.Global
	}
	return nil
}

func (x *ProxyProtocolPolicyStatus) GetWorkspaces() map[string]*v2.WorkspaceStatus {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

func (x *ProxyProtocolPolicyStatus) GetSelectedListeners() []*v2.ListenerPortReference {
	if x != nil {
		return x.SelectedListeners
	}
	return nil
}

// $hide_from_docs
type ProxyProtocolPolicyNewStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Common *v2.Status `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	// The number of virtual gateway listeners selected by the policy.
	SelectedListeners uint64 `protobuf:"varint,2,opt,name=selected_listeners,json=selectedListeners,proto3" json:"selected_listeners,omitempty"`
}

func (x *ProxyProtocolPolicyNewStatus) Reset() {
	*x = ProxyProtocolPolicyNewStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyProtocolPolicyNewStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyProtocolPolicyNewStatus) ProtoMessage() {}

func (x *ProxyProtocolPolicyNewStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyProtocolPolicyNewStatus.ProtoReflect.Descriptor instead.
func (*ProxyProtocolPolicyNewStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_rawDescGZIP(), []int{2}
}

func (x *ProxyProtocolPolicyNewStatus) GetCommon() *v2.Status {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *ProxyProtocolPolicyNewStatus) GetSelectedListeners() uint64 {
	if x != nil {
		return x.SelectedListeners
	}
	return 0
}

// $hide_from_docs
type ProxyProtocolPolicyReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status of the resource in each workspace that it exists in.
	Workspaces map[string]*v2.Report `protobuf:"bytes,1,rep,name=workspaces,proto3" json:"workspaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A list of virtual gateway listeners selected by the policy.
	SelectedListeners []*v2.ListenerPortReference `protobuf:"bytes,2,rep,name=selected_listeners,json=selectedListeners,proto3" json:"selected_listeners,omitempty"`
}

func (x *ProxyProtocolPolicyReport) Reset() {
	*x = ProxyProtocolPolicyReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyProtocolPolicyReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyProtocolPolicyReport) ProtoMessage() {}

func (x *ProxyProtocolPolicyReport) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyProtocolPolicyReport.ProtoReflect.Descriptor instead.
func (*ProxyProtocolPolicyReport) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_rawDescGZIP(), []int{3}
}

func (x *ProxyProtocolPolicyReport) GetWorkspaces() map[string]*v2.Report {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

func (x *ProxyProtocolPolicyReport) GetSelectedListeners() []*v2.ListenerPortReference {
	if x != nil {
		return x.SelectedListeners
	}
	return nil
}

type ProxyProtocolPolicySpec_ListenerSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VirtualGateway *v2.ObjectSelector `protobuf:"bytes,1,opt,name=virtual_gateway,json=virtualGateway,proto3" json:"virtual_gateway,omitempty"`
	Port           *v2.PortSelector   `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *ProxyProtocolPolicySpec_ListenerSelector) Reset() {
	*x = ProxyProtocolPolicySpec_ListenerSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyProtocolPolicySpec_ListenerSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyProtocolPolicySpec_ListenerSelector) ProtoMessage() {}

func (x *ProxyProtocolPolicySpec_ListenerSelector) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyProtocolPolicySpec_ListenerSelector.ProtoReflect.Descriptor instead.
func (*ProxyProtocolPolicySpec_ListenerSelector) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ProxyProtocolPolicySpec_ListenerSelector) GetVirtualGateway() *v2.ObjectSelector {
	if x != nil {
		return x.VirtualGateway
	}
	return nil
}

func (x *ProxyProtocolPolicySpec_ListenerSelector) GetPort() *v2.PortSelector {
	if x != nil {
		return x.Port
	}
	return nil
}

type ProxyProtocolPolicySpec_Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnableProxyProtocol bool `protobuf:"varint,1,opt,name=enable_proxy_protocol,json=enableProxyProtocol,proto3" json:"enable_proxy_protocol,omitempty"`
}

func (x *ProxyProtocolPolicySpec_Config) Reset() {
	*x = ProxyProtocolPolicySpec_Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyProtocolPolicySpec_Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyProtocolPolicySpec_Config) ProtoMessage() {}

func (x *ProxyProtocolPolicySpec_Config) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyProtocolPolicySpec_Config.ProtoReflect.Descriptor instead.
func (*ProxyProtocolPolicySpec_Config) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ProxyProtocolPolicySpec_Config) GetEnableProxyProtocol() bool {
	if x != nil {
		return x.EnableProxyProtocol
	}
	return false
}

var File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_rawDesc = []byte{
	0x0a, 0x6c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2f, 0x76, 0x32, 0x2f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22,
	0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65,
	0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x32, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x52, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6e, 0x65,
	0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x03, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x7a, 0x0a, 0x12, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x4c, 0x2e, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x12,
	0x5a, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x42, 0x2e, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x97, 0x01, 0x0a, 0x10,
	0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x4c, 0x0a, 0x0f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0e,
	0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x35,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x3c, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x32, 0x0a, 0x15, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x22, 0xff, 0x03, 0x0a, 0x19, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x33, 0x0a, 0x13, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02,
	0x18, 0x01, 0x52, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x02, 0x18, 0x01, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x69, 0x63, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x12, 0x6d, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4d, 0x2e, 0x74, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x59, 0x0a, 0x12, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x11, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73,
	0x1a, 0x63, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x82, 0x01, 0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4e, 0x65, 0x77,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x12, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x22, 0xc1, 0x02, 0x0a, 0x19, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x6d, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4d, 0x2e, 0x74,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x59, 0x0a, 0x12, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x11, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x73, 0x1a, 0x5a, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x58,
	0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x32, 0xc0, 0xf5, 0x04, 0x01,
	0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_goTypes = []interface{}{
	(*ProxyProtocolPolicySpec)(nil),                  // 0: trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicySpec
	(*ProxyProtocolPolicyStatus)(nil),                // 1: trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicyStatus
	(*ProxyProtocolPolicyNewStatus)(nil),             // 2: trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicyNewStatus
	(*ProxyProtocolPolicyReport)(nil),                // 3: trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicyReport
	(*ProxyProtocolPolicySpec_ListenerSelector)(nil), // 4: trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicySpec.ListenerSelector
	(*ProxyProtocolPolicySpec_Config)(nil),           // 5: trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicySpec.Config
	nil,                                              // 6: trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicyStatus.WorkspacesEntry
	nil,                                              // 7: trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicyReport.WorkspacesEntry
	(v2.ApprovalState)(0),                            // 8: common.gloo.solo.io.ApprovalState
	(*v2.GenericGlobalStatus)(nil),                   // 9: common.gloo.solo.io.GenericGlobalStatus
	(*v2.ListenerPortReference)(nil),                 // 10: common.gloo.solo.io.ListenerPortReference
	(*v2.Status)(nil),                                // 11: common.gloo.solo.io.Status
	(*v2.ObjectSelector)(nil),                        // 12: common.gloo.solo.io.ObjectSelector
	(*v2.PortSelector)(nil),                          // 13: common.gloo.solo.io.PortSelector
	(*v2.WorkspaceStatus)(nil),                       // 14: common.gloo.solo.io.WorkspaceStatus
	(*v2.Report)(nil),                                // 15: common.gloo.solo.io.Report
}
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_depIdxs = []int32{
	4,  // 0: trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicySpec.apply_to_listeners:type_name -> trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicySpec.ListenerSelector
	5,  // 1: trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicySpec.config:type_name -> trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicySpec.Config
	8,  // 2: trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicyStatus.state:type_name -> common.gloo.solo.io.ApprovalState
	9,  // 3: trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicyStatus.global:type_name -> common.gloo.solo.io.GenericGlobalStatus
	6,  // 4: trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicyStatus.workspaces:type_name -> trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicyStatus.WorkspacesEntry
	10, // 5: trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicyStatus.selected_listeners:type_name -> common.gloo.solo.io.ListenerPortReference
	11, // 6: trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicyNewStatus.common:type_name -> common.gloo.solo.io.Status
	7,  // 7: trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicyReport.workspaces:type_name -> trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicyReport.WorkspacesEntry
	10, // 8: trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicyReport.selected_listeners:type_name -> common.gloo.solo.io.ListenerPortReference
	12, // 9: trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicySpec.ListenerSelector.virtual_gateway:type_name -> common.gloo.solo.io.ObjectSelector
	13, // 10: trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicySpec.ListenerSelector.port:type_name -> common.gloo.solo.io.PortSelector
	14, // 11: trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicyStatus.WorkspacesEntry.value:type_name -> common.gloo.solo.io.WorkspaceStatus
	15, // 12: trafficcontrol.policy.gloo.solo.io.ProxyProtocolPolicyReport.WorkspacesEntry.value:type_name -> common.gloo.solo.io.Report
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyProtocolPolicySpec); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyProtocolPolicyStatus); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyProtocolPolicyNewStatus); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyProtocolPolicyReport); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyProtocolPolicySpec_ListenerSelector); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyProtocolPolicySpec_Config); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_proxy_protocol_policy_proto_depIdxs = nil
}
