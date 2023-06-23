// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/admin/v2alpha1/waypoint_lifecycle_manager.proto

package v2alpha1

import (
	reflect "reflect"
	sync "sync"

	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/solo-io/cue/encoding/protobuf/cue"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	v2 "github.com/solo-io/solo-apis/client-go/common.gloo.solo.io/v2"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Manages the deployment of Ambient L7 policy proxy (waypoint proxy) the selected Kubernetes service accounts.
// A RouteTable in ambient mode creates a server-side proxy to handle L7 routing, and to apply policies.
// Outlier detection and failover policies in an ambient mesh create a client-side L7 proxy.
type WaypointLifecycleManagerSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Kubernetes service account to that you want to apply the override settings to. Selects only the service accounts in the matching clusters.
	// If empty, the settings are applied to all service accounts in the workspace.
	ApplyToServiceAccount []*v2.ObjectSelector `protobuf:"bytes,1,rep,name=apply_to_service_account,json=applyToServiceAccount,proto3" json:"apply_to_service_account,omitempty"`
	// The deployment specification that you want to use to override the default L7 policy proxy deployment values.
	// The values that you specify are merged with the default deployment values.
	// For information about the value format, see the [Google protocol buffer documentation](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/struct).
	DeploymentSpec *_struct.Struct `protobuf:"bytes,2,opt,name=deployment_spec,json=deploymentSpec,proto3" json:"deployment_spec,omitempty"`
}

func (x *WaypointLifecycleManagerSpec) Reset() {
	*x = WaypointLifecycleManagerSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WaypointLifecycleManagerSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaypointLifecycleManagerSpec) ProtoMessage() {}

func (x *WaypointLifecycleManagerSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaypointLifecycleManagerSpec.ProtoReflect.Descriptor instead.
func (*WaypointLifecycleManagerSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_rawDescGZIP(), []int{0}
}

func (x *WaypointLifecycleManagerSpec) GetApplyToServiceAccount() []*v2.ObjectSelector {
	if x != nil {
		return x.ApplyToServiceAccount
	}
	return nil
}

func (x *WaypointLifecycleManagerSpec) GetDeploymentSpec() *_struct.Struct {
	if x != nil {
		return x.DeploymentSpec
	}
	return nil
}

// This field represents the status of the WaypointLifecycleManagerSpec.
type WaypointLifecycleManagerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Common *v2.Status `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
}

func (x *WaypointLifecycleManagerStatus) Reset() {
	*x = WaypointLifecycleManagerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WaypointLifecycleManagerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaypointLifecycleManagerStatus) ProtoMessage() {}

func (x *WaypointLifecycleManagerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaypointLifecycleManagerStatus.ProtoReflect.Descriptor instead.
func (*WaypointLifecycleManagerStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_rawDescGZIP(), []int{1}
}

func (x *WaypointLifecycleManagerStatus) GetCommon() *v2.Status {
	if x != nil {
		return x.Common
	}
	return nil
}

type WaypointLifecycleManagerReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspaces map[string]*v2.Report `protobuf:"bytes,1,rep,name=workspaces,proto3" json:"workspaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *WaypointLifecycleManagerReport) Reset() {
	*x = WaypointLifecycleManagerReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WaypointLifecycleManagerReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaypointLifecycleManagerReport) ProtoMessage() {}

func (x *WaypointLifecycleManagerReport) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaypointLifecycleManagerReport.ProtoReflect.Descriptor instead.
func (*WaypointLifecycleManagerReport) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_rawDescGZIP(), []int{2}
}

func (x *WaypointLifecycleManagerReport) GetWorkspaces() map[string]*v2.Report {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

var File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_rawDesc = []byte{
	0x0a, 0x68, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x1f,
	0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x63, 0x75, 0x65, 0x2f, 0x63, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73,
	0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x1c, 0x57, 0x61, 0x79, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x5c, 0x0a, 0x18, 0x61, 0x70, 0x70, 0x6c, 0x79,
	0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x15,
	0x61, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x05, 0xea, 0x42, 0x02, 0x10, 0x01, 0x52, 0x0e,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x22, 0x55,
	0x0a, 0x1e, 0x57, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x33, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0xe0, 0x01, 0x0a, 0x1e, 0x57, 0x61, 0x79, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x62, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x57, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x1a, 0x5a, 0x0a, 0x0f,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x59, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72,
	0x69, 0x73, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0,
	0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_rawDescData = file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_goTypes = []interface{}{
	(*WaypointLifecycleManagerSpec)(nil),   // 0: admin.gloo.solo.io.WaypointLifecycleManagerSpec
	(*WaypointLifecycleManagerStatus)(nil), // 1: admin.gloo.solo.io.WaypointLifecycleManagerStatus
	(*WaypointLifecycleManagerReport)(nil), // 2: admin.gloo.solo.io.WaypointLifecycleManagerReport
	nil,                                    // 3: admin.gloo.solo.io.WaypointLifecycleManagerReport.WorkspacesEntry
	(*v2.ObjectSelector)(nil),              // 4: common.gloo.solo.io.ObjectSelector
	(*_struct.Struct)(nil),                 // 5: google.protobuf.Struct
	(*v2.Status)(nil),                      // 6: common.gloo.solo.io.Status
	(*v2.Report)(nil),                      // 7: common.gloo.solo.io.Report
}
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_depIdxs = []int32{
	4, // 0: admin.gloo.solo.io.WaypointLifecycleManagerSpec.apply_to_service_account:type_name -> common.gloo.solo.io.ObjectSelector
	5, // 1: admin.gloo.solo.io.WaypointLifecycleManagerSpec.deployment_spec:type_name -> google.protobuf.Struct
	6, // 2: admin.gloo.solo.io.WaypointLifecycleManagerStatus.common:type_name -> common.gloo.solo.io.Status
	3, // 3: admin.gloo.solo.io.WaypointLifecycleManagerReport.workspaces:type_name -> admin.gloo.solo.io.WaypointLifecycleManagerReport.WorkspacesEntry
	7, // 4: admin.gloo.solo.io.WaypointLifecycleManagerReport.WorkspacesEntry.value:type_name -> common.gloo.solo.io.Report
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_init()
}
func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_init() {
	if File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WaypointLifecycleManagerSpec); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WaypointLifecycleManagerStatus); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WaypointLifecycleManagerReport); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto = out.File
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2alpha1_waypoint_lifecycle_manager_proto_depIdxs = nil
}
