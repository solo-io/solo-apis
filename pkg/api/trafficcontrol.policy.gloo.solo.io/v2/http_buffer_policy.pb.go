// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/policy/v2/trafficcontrol/http_buffer_policy.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// HTTPBufferPolicy is used to enable http buffers for routes
//
// For more details on HTTPBuffer, visit:
// https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/buffer_filter
//
// ```
// apiVersion: trafficcontrol.policy.gloo.solo.io/v2
// kind: HTTPBufferPolicy
// metadata:
//
//	name: buffer
//	namespace: bookinfo
//
// spec:
//
//	config:
//	  maxRequestBytes: 16384
//	applyToRoutes:
//	- route:
//	    labels:
//	      route: productpage
//
// ```
type HTTPBufferPolicySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplyToRoutes []*v2.RouteSelector          `protobuf:"bytes,1,rep,name=apply_to_routes,json=applyToRoutes,proto3" json:"apply_to_routes,omitempty"`
	Config        *HTTPBufferPolicySpec_Config `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *HTTPBufferPolicySpec) Reset() {
	*x = HTTPBufferPolicySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPBufferPolicySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPBufferPolicySpec) ProtoMessage() {}

func (x *HTTPBufferPolicySpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPBufferPolicySpec.ProtoReflect.Descriptor instead.
func (*HTTPBufferPolicySpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_rawDescGZIP(), []int{0}
}

func (x *HTTPBufferPolicySpec) GetApplyToRoutes() []*v2.RouteSelector {
	if x != nil {
		return x.ApplyToRoutes
	}
	return nil
}

func (x *HTTPBufferPolicySpec) GetConfig() *HTTPBufferPolicySpec_Config {
	if x != nil {
		return x.Config
	}
	return nil
}

// reflects the status of the HTTPBufferPolicy
type HTTPBufferPolicyStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Global *v2.GenericGlobalStatus `protobuf:"bytes,1,opt,name=global,proto3" json:"global,omitempty"`
	// The status of the resource in each workspace that it exists in
	Workspaces map[string]*v2.WorkspaceStatus `protobuf:"bytes,2,rep,name=workspaces,proto3" json:"workspaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HTTPBufferPolicyStatus) Reset() {
	*x = HTTPBufferPolicyStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPBufferPolicyStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPBufferPolicyStatus) ProtoMessage() {}

func (x *HTTPBufferPolicyStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPBufferPolicyStatus.ProtoReflect.Descriptor instead.
func (*HTTPBufferPolicyStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_rawDescGZIP(), []int{1}
}

func (x *HTTPBufferPolicyStatus) GetGlobal() *v2.GenericGlobalStatus {
	if x != nil {
		return x.Global
	}
	return nil
}

func (x *HTTPBufferPolicyStatus) GetWorkspaces() map[string]*v2.WorkspaceStatus {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

type HTTPBufferPolicySpec_Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum request size that the filter will buffer before
	// the connection manager will stop buffering and return a 413
	// response
	MaxRequestBytes *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=max_request_bytes,json=maxRequestBytes,proto3" json:"max_request_bytes,omitempty"`
}

func (x *HTTPBufferPolicySpec_Config) Reset() {
	*x = HTTPBufferPolicySpec_Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPBufferPolicySpec_Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPBufferPolicySpec_Config) ProtoMessage() {}

func (x *HTTPBufferPolicySpec_Config) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPBufferPolicySpec_Config.ProtoReflect.Descriptor instead.
func (*HTTPBufferPolicySpec_Config) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_rawDescGZIP(), []int{0, 0}
}

func (x *HTTPBufferPolicySpec_Config) GetMaxRequestBytes() *wrappers.UInt32Value {
	if x != nil {
		return x.MaxRequestBytes
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_rawDesc = []byte{
	0x0a, 0x69, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2f, 0x76, 0x32, 0x2f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x74, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a,
	0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65,
	0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x02, 0x0a, 0x14, 0x48, 0x54, 0x54, 0x50, 0x42,
	0x75, 0x66, 0x66, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x4a, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0d, 0x61, 0x70,
	0x70, 0x6c, 0x79, 0x54, 0x6f, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x57, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x74, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x48, 0x54, 0x54, 0x50, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x1a, 0x52, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x48,
	0x0a, 0x11, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0xab, 0x02, 0x0a, 0x16, 0x48, 0x54, 0x54,
	0x50, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x40, 0x0a, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69,
	0x63, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x12, 0x6a, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x74, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x48,
	0x54, 0x54, 0x50, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x73, 0x1a, 0x63, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x58, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2f, 0x76, 0x32, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_goTypes = []interface{}{
	(*HTTPBufferPolicySpec)(nil),        // 0: trafficcontrol.policy.gloo.solo.io.HTTPBufferPolicySpec
	(*HTTPBufferPolicyStatus)(nil),      // 1: trafficcontrol.policy.gloo.solo.io.HTTPBufferPolicyStatus
	(*HTTPBufferPolicySpec_Config)(nil), // 2: trafficcontrol.policy.gloo.solo.io.HTTPBufferPolicySpec.Config
	nil,                                 // 3: trafficcontrol.policy.gloo.solo.io.HTTPBufferPolicyStatus.WorkspacesEntry
	(*v2.RouteSelector)(nil),            // 4: common.gloo.solo.io.RouteSelector
	(*v2.GenericGlobalStatus)(nil),      // 5: common.gloo.solo.io.GenericGlobalStatus
	(*wrappers.UInt32Value)(nil),        // 6: google.protobuf.UInt32Value
	(*v2.WorkspaceStatus)(nil),          // 7: common.gloo.solo.io.WorkspaceStatus
}
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_depIdxs = []int32{
	4, // 0: trafficcontrol.policy.gloo.solo.io.HTTPBufferPolicySpec.apply_to_routes:type_name -> common.gloo.solo.io.RouteSelector
	2, // 1: trafficcontrol.policy.gloo.solo.io.HTTPBufferPolicySpec.config:type_name -> trafficcontrol.policy.gloo.solo.io.HTTPBufferPolicySpec.Config
	5, // 2: trafficcontrol.policy.gloo.solo.io.HTTPBufferPolicyStatus.global:type_name -> common.gloo.solo.io.GenericGlobalStatus
	3, // 3: trafficcontrol.policy.gloo.solo.io.HTTPBufferPolicyStatus.workspaces:type_name -> trafficcontrol.policy.gloo.solo.io.HTTPBufferPolicyStatus.WorkspacesEntry
	6, // 4: trafficcontrol.policy.gloo.solo.io.HTTPBufferPolicySpec.Config.max_request_bytes:type_name -> google.protobuf.UInt32Value
	7, // 5: trafficcontrol.policy.gloo.solo.io.HTTPBufferPolicyStatus.WorkspacesEntry.value:type_name -> common.gloo.solo.io.WorkspaceStatus
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPBufferPolicySpec); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPBufferPolicyStatus); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPBufferPolicySpec_Config); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_trafficcontrol_http_buffer_policy_proto_depIdxs = nil
}