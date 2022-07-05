// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/policy/v2/resilience/fault_injection_policy.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// FaultInjectionPolicy is used to inject faults (latency and errors) into requests sent through the mesh.
// Fault specification is part of a VirtualService rule. Faults include
// aborting the Http request from downstream service, and/or delaying
// proxying of requests. A fault rule MUST HAVE delay or abort or both.
// FaultInjectionPolicies are applied at the *Route* level.
type FaultInjectionPolicySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// select the routes where the policy will be applied
	// if left empty, will apply to all routes in the workspace.
	ApplyToRoutes []*v2.RouteSelector `protobuf:"bytes,1,rep,name=apply_to_routes,json=applyToRoutes,proto3" json:"apply_to_routes,omitempty"`
	// The details of the fault injection policy to apply to the selected routes.
	Config *FaultInjectionPolicySpec_Config `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *FaultInjectionPolicySpec) Reset() {
	*x = FaultInjectionPolicySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaultInjectionPolicySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaultInjectionPolicySpec) ProtoMessage() {}

func (x *FaultInjectionPolicySpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaultInjectionPolicySpec.ProtoReflect.Descriptor instead.
func (*FaultInjectionPolicySpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_rawDescGZIP(), []int{0}
}

func (x *FaultInjectionPolicySpec) GetApplyToRoutes() []*v2.RouteSelector {
	if x != nil {
		return x.ApplyToRoutes
	}
	return nil
}

func (x *FaultInjectionPolicySpec) GetConfig() *FaultInjectionPolicySpec_Config {
	if x != nil {
		return x.Config
	}
	return nil
}

// reflects the status of the FaultInjectionPolicy
type FaultInjectionPolicyStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Global *v2.GenericGlobalStatus `protobuf:"bytes,1,opt,name=global,proto3" json:"global,omitempty"`
	// The status of the resource in each workspace that it exists in.
	Workspaces map[string]*v2.WorkspaceStatus `protobuf:"bytes,2,rep,name=workspaces,proto3" json:"workspaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Routes selected by the policy
	SelectedRoutes []*v2.RouteReference `protobuf:"bytes,3,rep,name=selected_routes,json=selectedRoutes,proto3" json:"selected_routes,omitempty"`
}

func (x *FaultInjectionPolicyStatus) Reset() {
	*x = FaultInjectionPolicyStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaultInjectionPolicyStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaultInjectionPolicyStatus) ProtoMessage() {}

func (x *FaultInjectionPolicyStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaultInjectionPolicyStatus.ProtoReflect.Descriptor instead.
func (*FaultInjectionPolicyStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_rawDescGZIP(), []int{1}
}

func (x *FaultInjectionPolicyStatus) GetGlobal() *v2.GenericGlobalStatus {
	if x != nil {
		return x.Global
	}
	return nil
}

func (x *FaultInjectionPolicyStatus) GetWorkspaces() map[string]*v2.WorkspaceStatus {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

func (x *FaultInjectionPolicyStatus) GetSelectedRoutes() []*v2.RouteReference {
	if x != nil {
		return x.SelectedRoutes
	}
	return nil
}

// *Note:* Delay and abort faults are independent of one another, even if
// both are specified simultaneously.
type FaultInjectionPolicySpec_Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicate the amount of delay in seconds.
	// The optional _percentage_ field can be used to only delay a certain
	// percentage of requests. If left unspecified, all request will be delayed.
	Delay *FaultInjectionPolicySpec_Config_Delay `protobuf:"bytes,1,opt,name=delay,proto3" json:"delay,omitempty"`
	// Abort the request and return the specified error code back to traffic source.
	Abort *FaultInjectionPolicySpec_Config_Abort `protobuf:"bytes,2,opt,name=abort,proto3" json:"abort,omitempty"`
}

func (x *FaultInjectionPolicySpec_Config) Reset() {
	*x = FaultInjectionPolicySpec_Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaultInjectionPolicySpec_Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaultInjectionPolicySpec_Config) ProtoMessage() {}

func (x *FaultInjectionPolicySpec_Config) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaultInjectionPolicySpec_Config.ProtoReflect.Descriptor instead.
func (*FaultInjectionPolicySpec_Config) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_rawDescGZIP(), []int{0, 0}
}

func (x *FaultInjectionPolicySpec_Config) GetDelay() *FaultInjectionPolicySpec_Config_Delay {
	if x != nil {
		return x.Delay
	}
	return nil
}

func (x *FaultInjectionPolicySpec_Config) GetAbort() *FaultInjectionPolicySpec_Config_Abort {
	if x != nil {
		return x.Abort
	}
	return nil
}

// Abort Http request attempts and return error codes back to downstream
// service, giving the impression that the upstream service is faulty.
type FaultInjectionPolicySpec_Config_Abort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. HTTP status code to use to abort the request.
	HttpStatus int32 `protobuf:"varint,1,opt,name=http_status,json=httpStatus,proto3" json:"http_status,omitempty"`
	// Percentage of requests to be aborted. Values range between 0 and 100. If omitted all requests will be aborted.
	Percentage *wrappers.DoubleValue `protobuf:"bytes,2,opt,name=percentage,proto3" json:"percentage,omitempty"`
}

func (x *FaultInjectionPolicySpec_Config_Abort) Reset() {
	*x = FaultInjectionPolicySpec_Config_Abort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaultInjectionPolicySpec_Config_Abort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaultInjectionPolicySpec_Config_Abort) ProtoMessage() {}

func (x *FaultInjectionPolicySpec_Config_Abort) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaultInjectionPolicySpec_Config_Abort.ProtoReflect.Descriptor instead.
func (*FaultInjectionPolicySpec_Config_Abort) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *FaultInjectionPolicySpec_Config_Abort) GetHttpStatus() int32 {
	if x != nil {
		return x.HttpStatus
	}
	return 0
}

func (x *FaultInjectionPolicySpec_Config_Abort) GetPercentage() *wrappers.DoubleValue {
	if x != nil {
		return x.Percentage
	}
	return nil
}

// Delay requests before forwarding, emulating various failures such as
// network issues, overloaded upstream service, etc.
type FaultInjectionPolicySpec_Config_Delay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Add a delay of a fixed duration before sending the request. Format: `1h`/`1m`/`1s`/`1ms`. MUST be >=1ms.
	FixedDelay *duration.Duration `protobuf:"bytes,1,opt,name=fixed_delay,json=fixedDelay,proto3" json:"fixed_delay,omitempty"`
	// Percentage of requests on which the delay will be injected. Values range between 0 and 100. If omitted all requests will be delayed.
	Percentage *wrappers.DoubleValue `protobuf:"bytes,2,opt,name=percentage,proto3" json:"percentage,omitempty"`
}

func (x *FaultInjectionPolicySpec_Config_Delay) Reset() {
	*x = FaultInjectionPolicySpec_Config_Delay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaultInjectionPolicySpec_Config_Delay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaultInjectionPolicySpec_Config_Delay) ProtoMessage() {}

func (x *FaultInjectionPolicySpec_Config_Delay) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaultInjectionPolicySpec_Config_Delay.ProtoReflect.Descriptor instead.
func (*FaultInjectionPolicySpec_Config_Delay) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *FaultInjectionPolicySpec_Config_Delay) GetFixedDelay() *duration.Duration {
	if x != nil {
		return x.FixedDelay
	}
	return nil
}

func (x *FaultInjectionPolicySpec_Config_Delay) GetPercentage() *wrappers.DoubleValue {
	if x != nil {
		return x.Percentage
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_rawDesc = []byte{
	0x0a, 0x69, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x5f, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x72, 0x65, 0x73,
	0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x51, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x32, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf0, 0x04, 0x0a, 0x18, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x6a, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x4a, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0d, 0x61, 0x70,
	0x70, 0x6c, 0x79, 0x54, 0x6f, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x57, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x72, 0x65,
	0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x61, 0x75,
	0x6c, 0x74, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x1a, 0xae, 0x03, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x5b, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45,
	0x2e, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x46, 0x61, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x44, 0x65, 0x6c, 0x61, 0x79, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x5b, 0x0a, 0x05,
	0x61, 0x62, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x72, 0x65,
	0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x61, 0x75,
	0x6c, 0x74, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x62, 0x6f,
	0x72, 0x74, 0x52, 0x05, 0x61, 0x62, 0x6f, 0x72, 0x74, 0x1a, 0x66, 0x0a, 0x05, 0x41, 0x62, 0x6f,
	0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x3c, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x1a, 0x81, 0x01, 0x0a, 0x05, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x3a, 0x0a, 0x0b, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x3c, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x22, 0xfd, 0x02, 0x0a, 0x1a, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x49,
	0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x40, 0x0a, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x69, 0x63, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x12, 0x6a, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x72, 0x65, 0x73,
	0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x61, 0x75, 0x6c,
	0x74, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x12, 0x4c, 0x0a, 0x0f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x0e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x1a, 0x63, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x54, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x32, 0xc0,
	0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_goTypes = []interface{}{
	(*FaultInjectionPolicySpec)(nil),              // 0: resilience.policy.gloo.solo.io.FaultInjectionPolicySpec
	(*FaultInjectionPolicyStatus)(nil),            // 1: resilience.policy.gloo.solo.io.FaultInjectionPolicyStatus
	(*FaultInjectionPolicySpec_Config)(nil),       // 2: resilience.policy.gloo.solo.io.FaultInjectionPolicySpec.Config
	(*FaultInjectionPolicySpec_Config_Abort)(nil), // 3: resilience.policy.gloo.solo.io.FaultInjectionPolicySpec.Config.Abort
	(*FaultInjectionPolicySpec_Config_Delay)(nil), // 4: resilience.policy.gloo.solo.io.FaultInjectionPolicySpec.Config.Delay
	nil,                            // 5: resilience.policy.gloo.solo.io.FaultInjectionPolicyStatus.WorkspacesEntry
	(*v2.RouteSelector)(nil),       // 6: common.gloo.solo.io.RouteSelector
	(*v2.GenericGlobalStatus)(nil), // 7: common.gloo.solo.io.GenericGlobalStatus
	(*v2.RouteReference)(nil),      // 8: common.gloo.solo.io.RouteReference
	(*wrappers.DoubleValue)(nil),   // 9: google.protobuf.DoubleValue
	(*duration.Duration)(nil),      // 10: google.protobuf.Duration
	(*v2.WorkspaceStatus)(nil),     // 11: common.gloo.solo.io.WorkspaceStatus
}
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_depIdxs = []int32{
	6,  // 0: resilience.policy.gloo.solo.io.FaultInjectionPolicySpec.apply_to_routes:type_name -> common.gloo.solo.io.RouteSelector
	2,  // 1: resilience.policy.gloo.solo.io.FaultInjectionPolicySpec.config:type_name -> resilience.policy.gloo.solo.io.FaultInjectionPolicySpec.Config
	7,  // 2: resilience.policy.gloo.solo.io.FaultInjectionPolicyStatus.global:type_name -> common.gloo.solo.io.GenericGlobalStatus
	5,  // 3: resilience.policy.gloo.solo.io.FaultInjectionPolicyStatus.workspaces:type_name -> resilience.policy.gloo.solo.io.FaultInjectionPolicyStatus.WorkspacesEntry
	8,  // 4: resilience.policy.gloo.solo.io.FaultInjectionPolicyStatus.selected_routes:type_name -> common.gloo.solo.io.RouteReference
	4,  // 5: resilience.policy.gloo.solo.io.FaultInjectionPolicySpec.Config.delay:type_name -> resilience.policy.gloo.solo.io.FaultInjectionPolicySpec.Config.Delay
	3,  // 6: resilience.policy.gloo.solo.io.FaultInjectionPolicySpec.Config.abort:type_name -> resilience.policy.gloo.solo.io.FaultInjectionPolicySpec.Config.Abort
	9,  // 7: resilience.policy.gloo.solo.io.FaultInjectionPolicySpec.Config.Abort.percentage:type_name -> google.protobuf.DoubleValue
	10, // 8: resilience.policy.gloo.solo.io.FaultInjectionPolicySpec.Config.Delay.fixed_delay:type_name -> google.protobuf.Duration
	9,  // 9: resilience.policy.gloo.solo.io.FaultInjectionPolicySpec.Config.Delay.percentage:type_name -> google.protobuf.DoubleValue
	11, // 10: resilience.policy.gloo.solo.io.FaultInjectionPolicyStatus.WorkspacesEntry.value:type_name -> common.gloo.solo.io.WorkspaceStatus
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaultInjectionPolicySpec); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaultInjectionPolicyStatus); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaultInjectionPolicySpec_Config); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaultInjectionPolicySpec_Config_Abort); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaultInjectionPolicySpec_Config_Delay); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_fault_injection_policy_proto_depIdxs = nil
}
