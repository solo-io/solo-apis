// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/policy/v2/trafficcontrol/mirror_policy.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	v2 "github.com/solo-io/solo-apis/client-go/common.gloo.solo.io/v2"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MirrorPolicy is used to mirror/shadow traffic to a another destination (traffic will be sent to its original destination in addition to the mirrored destinations).
// MirrorPolicies are applied at the *Route* level.
type MirrorPolicySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Routes to apply the policy to.
	// If empty, the policy applies to all workloads in the workspace.
	ApplyToRoutes []*v2.RouteSelector `protobuf:"bytes,1,rep,name=apply_to_routes,json=applyToRoutes,proto3" json:"apply_to_routes,omitempty"`
	// The details of the Mirror policy to apply to the selected routes
	Config *MirrorPolicySpec_Config `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *MirrorPolicySpec) Reset() {
	*x = MirrorPolicySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MirrorPolicySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MirrorPolicySpec) ProtoMessage() {}

func (x *MirrorPolicySpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MirrorPolicySpec.ProtoReflect.Descriptor instead.
func (*MirrorPolicySpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_rawDescGZIP(), []int{0}
}

func (x *MirrorPolicySpec) GetApplyToRoutes() []*v2.RouteSelector {
	if x != nil {
		return x.ApplyToRoutes
	}
	return nil
}

func (x *MirrorPolicySpec) GetConfig() *MirrorPolicySpec_Config {
	if x != nil {
		return x.Config
	}
	return nil
}

// The status of the policy after it is applied to your Gloo environment.
type MirrorPolicyStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The state and workspace conditions of the applied resource.
	Common *v2.Status `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	// The number of routes selected by this policy.
	NumSelectedRoutes uint32 `protobuf:"varint,2,opt,name=num_selected_routes,json=numSelectedRoutes,proto3" json:"num_selected_routes,omitempty"`
}

func (x *MirrorPolicyStatus) Reset() {
	*x = MirrorPolicyStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MirrorPolicyStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MirrorPolicyStatus) ProtoMessage() {}

func (x *MirrorPolicyStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MirrorPolicyStatus.ProtoReflect.Descriptor instead.
func (*MirrorPolicyStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_rawDescGZIP(), []int{1}
}

func (x *MirrorPolicyStatus) GetCommon() *v2.Status {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *MirrorPolicyStatus) GetNumSelectedRoutes() uint32 {
	if x != nil {
		return x.NumSelectedRoutes
	}
	return 0
}

type MirrorPolicyReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status of the resource in each workspace that it exists in.
	Workspaces map[string]*v2.Report `protobuf:"bytes,1,rep,name=workspaces,proto3" json:"workspaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A list of references to all routes selected by the policy.
	SelectedRoutes []*v2.RouteReference `protobuf:"bytes,2,rep,name=selected_routes,json=selectedRoutes,proto3" json:"selected_routes,omitempty"`
}

func (x *MirrorPolicyReport) Reset() {
	*x = MirrorPolicyReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MirrorPolicyReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MirrorPolicyReport) ProtoMessage() {}

func (x *MirrorPolicyReport) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MirrorPolicyReport.ProtoReflect.Descriptor instead.
func (*MirrorPolicyReport) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_rawDescGZIP(), []int{2}
}

func (x *MirrorPolicyReport) GetWorkspaces() map[string]*v2.Report {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

func (x *MirrorPolicyReport) GetSelectedRoutes() []*v2.RouteReference {
	if x != nil {
		return x.SelectedRoutes
	}
	return nil
}

type MirrorPolicySpec_Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The destination to which to send mirrored traffic
	Destination *v2.DestinationReference `protobuf:"bytes,1,opt,name=destination,proto3" json:"destination,omitempty"`
	// Percentage of traffic to mirror. If omitted all traffic will be mirrored. Values must be between 0 and 100.
	// For information about the value format, see the [Google protocol buffer documentation](https://protobuf.dev/reference/protobuf/google.protobuf/#double-value).
	Percentage *wrapperspb.DoubleValue `protobuf:"bytes,2,opt,name=percentage,proto3" json:"percentage,omitempty"`
}

func (x *MirrorPolicySpec_Config) Reset() {
	*x = MirrorPolicySpec_Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MirrorPolicySpec_Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MirrorPolicySpec_Config) ProtoMessage() {}

func (x *MirrorPolicySpec_Config) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MirrorPolicySpec_Config.ProtoReflect.Descriptor instead.
func (*MirrorPolicySpec_Config) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MirrorPolicySpec_Config) GetDestination() *v2.DestinationReference {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *MirrorPolicySpec_Config) GetPercentage() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Percentage
	}
	return nil
}

var File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_rawDesc = []byte{
	0x0a, 0x68, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x74, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12,
	0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68,
	0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x55, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72,
	0x69, 0x73, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x32, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x02, 0x0a, 0x10, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4a, 0x0a, 0x0f, 0x61, 0x70,
	0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x6f,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x53, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4d, 0x69, 0x72, 0x72,
	0x6f, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x93, 0x01, 0x0a, 0x06,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4b, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x22, 0x79, 0x0a, 0x12, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x13,
	0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x6e, 0x75, 0x6d, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0xa6, 0x02, 0x0a,
	0x12, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x66, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4d, 0x69, 0x72,
	0x72, 0x6f, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x4c, 0x0a, 0x0f, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0e, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x5a, 0x0a, 0x0f, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x66, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65,
	0x2f, 0x76, 0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76,
	0x32, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_rawDescData = file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_goTypes = []interface{}{
	(*MirrorPolicySpec)(nil),        // 0: trafficcontrol.policy.gloo.solo.io.MirrorPolicySpec
	(*MirrorPolicyStatus)(nil),      // 1: trafficcontrol.policy.gloo.solo.io.MirrorPolicyStatus
	(*MirrorPolicyReport)(nil),      // 2: trafficcontrol.policy.gloo.solo.io.MirrorPolicyReport
	(*MirrorPolicySpec_Config)(nil), // 3: trafficcontrol.policy.gloo.solo.io.MirrorPolicySpec.Config
	nil,                             // 4: trafficcontrol.policy.gloo.solo.io.MirrorPolicyReport.WorkspacesEntry
	(*v2.RouteSelector)(nil),        // 5: common.gloo.solo.io.RouteSelector
	(*v2.Status)(nil),               // 6: common.gloo.solo.io.Status
	(*v2.RouteReference)(nil),       // 7: common.gloo.solo.io.RouteReference
	(*v2.DestinationReference)(nil), // 8: common.gloo.solo.io.DestinationReference
	(*wrapperspb.DoubleValue)(nil),  // 9: google.protobuf.DoubleValue
	(*v2.Report)(nil),               // 10: common.gloo.solo.io.Report
}
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_depIdxs = []int32{
	5,  // 0: trafficcontrol.policy.gloo.solo.io.MirrorPolicySpec.apply_to_routes:type_name -> common.gloo.solo.io.RouteSelector
	3,  // 1: trafficcontrol.policy.gloo.solo.io.MirrorPolicySpec.config:type_name -> trafficcontrol.policy.gloo.solo.io.MirrorPolicySpec.Config
	6,  // 2: trafficcontrol.policy.gloo.solo.io.MirrorPolicyStatus.common:type_name -> common.gloo.solo.io.Status
	4,  // 3: trafficcontrol.policy.gloo.solo.io.MirrorPolicyReport.workspaces:type_name -> trafficcontrol.policy.gloo.solo.io.MirrorPolicyReport.WorkspacesEntry
	7,  // 4: trafficcontrol.policy.gloo.solo.io.MirrorPolicyReport.selected_routes:type_name -> common.gloo.solo.io.RouteReference
	8,  // 5: trafficcontrol.policy.gloo.solo.io.MirrorPolicySpec.Config.destination:type_name -> common.gloo.solo.io.DestinationReference
	9,  // 6: trafficcontrol.policy.gloo.solo.io.MirrorPolicySpec.Config.percentage:type_name -> google.protobuf.DoubleValue
	10, // 7: trafficcontrol.policy.gloo.solo.io.MirrorPolicyReport.WorkspacesEntry.value:type_name -> common.gloo.solo.io.Report
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_init()
}
func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_init() {
	if File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MirrorPolicySpec); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MirrorPolicyStatus); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MirrorPolicyReport); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MirrorPolicySpec_Config); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto = out.File
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_trafficcontrol_mirror_policy_proto_depIdxs = nil
}
