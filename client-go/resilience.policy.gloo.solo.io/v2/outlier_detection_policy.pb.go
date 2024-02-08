// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/policy/v2/resilience/outlier_detection_policy.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// OutlierDetectionPolicy is used to configure [outlier detection](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/outlier) on the selected destinations.
// Specifying this field requires an empty `source_selector` because it must apply to all traffic.
// OutlierDetectionPolicies are applied at the *Destination* level.
//
// For VirtualDestinations, traffic will not be sent to deployments that are unavailable by default.
// An OutlierDetectionPolicy will add configuration to also eject a deployment that is returning too many 5xx HTTP status codes.
type OutlierDetectionPolicySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Destinations to apply the policy to.
	// If empty, the policy applies to all destinations in the workspace.
	ApplyToDestinations []*v2.DestinationSelector `protobuf:"bytes,1,rep,name=apply_to_destinations,json=applyToDestinations,proto3" json:"apply_to_destinations,omitempty"`
	// The details of the OutlierDetectionPolicy to apply to the selected destinations.
	Config *OutlierDetectionPolicySpec_Config `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *OutlierDetectionPolicySpec) Reset() {
	*x = OutlierDetectionPolicySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutlierDetectionPolicySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutlierDetectionPolicySpec) ProtoMessage() {}

func (x *OutlierDetectionPolicySpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutlierDetectionPolicySpec.ProtoReflect.Descriptor instead.
func (*OutlierDetectionPolicySpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_rawDescGZIP(), []int{0}
}

func (x *OutlierDetectionPolicySpec) GetApplyToDestinations() []*v2.DestinationSelector {
	if x != nil {
		return x.ApplyToDestinations
	}
	return nil
}

func (x *OutlierDetectionPolicySpec) GetConfig() *OutlierDetectionPolicySpec_Config {
	if x != nil {
		return x.Config
	}
	return nil
}

// The status of the policy after it is applied to your Gloo environment.
type OutlierDetectionPolicyStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The state and workspace conditions of the applied resource.
	Common *v2.Status `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	// The number of destination ports selected by the policy.
	NumSelectedDestinationPorts uint32 `protobuf:"varint,2,opt,name=num_selected_destination_ports,json=numSelectedDestinationPorts,proto3" json:"num_selected_destination_ports,omitempty"`
}

func (x *OutlierDetectionPolicyStatus) Reset() {
	*x = OutlierDetectionPolicyStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutlierDetectionPolicyStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutlierDetectionPolicyStatus) ProtoMessage() {}

func (x *OutlierDetectionPolicyStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutlierDetectionPolicyStatus.ProtoReflect.Descriptor instead.
func (*OutlierDetectionPolicyStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_rawDescGZIP(), []int{1}
}

func (x *OutlierDetectionPolicyStatus) GetCommon() *v2.Status {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *OutlierDetectionPolicyStatus) GetNumSelectedDestinationPorts() uint32 {
	if x != nil {
		return x.NumSelectedDestinationPorts
	}
	return 0
}

type OutlierDetectionPolicyReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of workspaces in which the policy can apply to workloads.
	Workspaces map[string]*v2.Report `protobuf:"bytes,1,rep,name=workspaces,proto3" json:"workspaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A list of destination ports selected by the policy.
	SelectedDestinationPorts []*v2.DestinationReference `protobuf:"bytes,2,rep,name=selected_destination_ports,json=selectedDestinationPorts,proto3" json:"selected_destination_ports,omitempty"`
}

func (x *OutlierDetectionPolicyReport) Reset() {
	*x = OutlierDetectionPolicyReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutlierDetectionPolicyReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutlierDetectionPolicyReport) ProtoMessage() {}

func (x *OutlierDetectionPolicyReport) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutlierDetectionPolicyReport.ProtoReflect.Descriptor instead.
func (*OutlierDetectionPolicyReport) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_rawDescGZIP(), []int{2}
}

func (x *OutlierDetectionPolicyReport) GetWorkspaces() map[string]*v2.Report {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

func (x *OutlierDetectionPolicyReport) GetSelectedDestinationPorts() []*v2.DestinationReference {
	if x != nil {
		return x.SelectedDestinationPorts
	}
	return nil
}

type OutlierDetectionPolicySpec_Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of 5xx errors before a destination is removed from the healthy connection pool. The default, if this field is not set, is 5.
	ConsecutiveErrors *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=consecutive_errors,json=consecutiveErrors,proto3" json:"consecutive_errors,omitempty"`
	// The amount of time between analyzing destinations for ejection.
	// Set this value as an integer plus a unit of time, in the format `1h`, `1m`, `1s`, or `1ms`. The value must be at least `1ms`, and defaults to `10s`.
	// For information about the value format, see the [Google protocol buffer documentation](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/duration).
	Interval *duration.Duration `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	// The minimum time duration for ejection, or the time when a destination is considered unhealthy and not used for load balancing.
	// Set this value as an integer plus a unit of time, in the format `1h`, `1m`, `1s`, or `1ms`. The value must be at least `1ms`, and defaults to `30s`.
	// For information about the value format, see the [Google protocol buffer documentation](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/duration).
	BaseEjectionTime *duration.Duration `protobuf:"bytes,3,opt,name=base_ejection_time,json=baseEjectionTime,proto3" json:"base_ejection_time,omitempty"`
	// The maximum percentage of destinations that can be removed from the healthy connection pool at a time.
	// For example, if you have 10 total destinations that the policy selects, and set this value to 50 percent, 5 destinations can be removed at once.
	// At least 1 destination can always be removed, regardless of the value you set. You can set this value between `0` and `100`, with a default of `100`.
	MaxEjectionPercent uint32 `protobuf:"varint,4,opt,name=max_ejection_percent,json=maxEjectionPercent,proto3" json:"max_ejection_percent,omitempty"`
	// Number of gateway errors before a host is ejected from the connection pool.
	// When the upstream host is accessed over HTTP, a 502, 503, or 504 return
	// code qualifies as a gateway error. When the upstream host is accessed over
	// an opaque TCP connection, connect timeouts and connection error/failure
	// events qualify as a gateway error.
	// This feature is disabled by default or when set to the value 0.
	//
	// Note that consecutive_gateway_errors and consecutive_errors can be
	// used separately or together. Because the errors counted by
	// consecutive_gateway_errors are also included in consecutive_errors,
	// if the value of consecutive_gateway_errors is greater than or equal to
	// the value of consecutive_errors, consecutive_gateway_errors will have
	// no effect.
	ConsecutiveGatewayErrors *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=consecutive_gateway_errors,json=consecutiveGatewayErrors,proto3" json:"consecutive_gateway_errors,omitempty"`
}

func (x *OutlierDetectionPolicySpec_Config) Reset() {
	*x = OutlierDetectionPolicySpec_Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutlierDetectionPolicySpec_Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutlierDetectionPolicySpec_Config) ProtoMessage() {}

func (x *OutlierDetectionPolicySpec_Config) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutlierDetectionPolicySpec_Config.ProtoReflect.Descriptor instead.
func (*OutlierDetectionPolicySpec_Config) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_rawDescGZIP(), []int{0, 0}
}

func (x *OutlierDetectionPolicySpec_Config) GetConsecutiveErrors() *wrappers.UInt32Value {
	if x != nil {
		return x.ConsecutiveErrors
	}
	return nil
}

func (x *OutlierDetectionPolicySpec_Config) GetInterval() *duration.Duration {
	if x != nil {
		return x.Interval
	}
	return nil
}

func (x *OutlierDetectionPolicySpec_Config) GetBaseEjectionTime() *duration.Duration {
	if x != nil {
		return x.BaseEjectionTime
	}
	return nil
}

func (x *OutlierDetectionPolicySpec_Config) GetMaxEjectionPercent() uint32 {
	if x != nil {
		return x.MaxEjectionPercent
	}
	return 0
}

func (x *OutlierDetectionPolicySpec_Config) GetConsecutiveGatewayErrors() *wrappers.UInt32Value {
	if x != nil {
		return x.ConsecutiveGatewayErrors
	}
	return nil
}

var File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_rawDesc = []byte{
	0x0a, 0x6f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x2f, 0x6f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1e, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d,
	0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x76,
	0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x55, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65,
	0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x32,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x04, 0x0a, 0x1a, 0x4f, 0x75, 0x74,
	0x6c, 0x69, 0x65, 0x72, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x5c, 0x0a, 0x15, 0x61, 0x70, 0x70, 0x6c, 0x79,
	0x5f, 0x74, 0x6f, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x13, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x59, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65,
	0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x1a, 0xe3, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4b, 0x0a, 0x12, 0x63,
	0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x76, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12,
	0x47, 0x0a, 0x12, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x62, 0x61, 0x73, 0x65, 0x45, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x5f,
	0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x45, 0x6a, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x5a, 0x0a, 0x1a, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x18, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x98, 0x01, 0x0a, 0x1c, 0x4f, 0x75, 0x74, 0x6c, 0x69,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x1e,
	0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x1b, 0x6e, 0x75, 0x6d, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74,
	0x73, 0x22, 0xd1, 0x02, 0x0a, 0x1c, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x6c, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4c, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x44,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x12, 0x67, 0x0a, 0x1a, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x18, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x1a, 0x5a, 0x0a, 0x0f, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x62, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65,
	0x2f, 0x76, 0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x69,
	0x6c, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x32, 0xc0, 0xf5, 0x04,
	0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_rawDescData = file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_goTypes = []interface{}{
	(*OutlierDetectionPolicySpec)(nil),        // 0: resilience.policy.gloo.solo.io.OutlierDetectionPolicySpec
	(*OutlierDetectionPolicyStatus)(nil),      // 1: resilience.policy.gloo.solo.io.OutlierDetectionPolicyStatus
	(*OutlierDetectionPolicyReport)(nil),      // 2: resilience.policy.gloo.solo.io.OutlierDetectionPolicyReport
	(*OutlierDetectionPolicySpec_Config)(nil), // 3: resilience.policy.gloo.solo.io.OutlierDetectionPolicySpec.Config
	nil,                             // 4: resilience.policy.gloo.solo.io.OutlierDetectionPolicyReport.WorkspacesEntry
	(*v2.DestinationSelector)(nil),  // 5: common.gloo.solo.io.DestinationSelector
	(*v2.Status)(nil),               // 6: common.gloo.solo.io.Status
	(*v2.DestinationReference)(nil), // 7: common.gloo.solo.io.DestinationReference
	(*wrappers.UInt32Value)(nil),    // 8: google.protobuf.UInt32Value
	(*duration.Duration)(nil),       // 9: google.protobuf.Duration
	(*v2.Report)(nil),               // 10: common.gloo.solo.io.Report
}
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_depIdxs = []int32{
	5,  // 0: resilience.policy.gloo.solo.io.OutlierDetectionPolicySpec.apply_to_destinations:type_name -> common.gloo.solo.io.DestinationSelector
	3,  // 1: resilience.policy.gloo.solo.io.OutlierDetectionPolicySpec.config:type_name -> resilience.policy.gloo.solo.io.OutlierDetectionPolicySpec.Config
	6,  // 2: resilience.policy.gloo.solo.io.OutlierDetectionPolicyStatus.common:type_name -> common.gloo.solo.io.Status
	4,  // 3: resilience.policy.gloo.solo.io.OutlierDetectionPolicyReport.workspaces:type_name -> resilience.policy.gloo.solo.io.OutlierDetectionPolicyReport.WorkspacesEntry
	7,  // 4: resilience.policy.gloo.solo.io.OutlierDetectionPolicyReport.selected_destination_ports:type_name -> common.gloo.solo.io.DestinationReference
	8,  // 5: resilience.policy.gloo.solo.io.OutlierDetectionPolicySpec.Config.consecutive_errors:type_name -> google.protobuf.UInt32Value
	9,  // 6: resilience.policy.gloo.solo.io.OutlierDetectionPolicySpec.Config.interval:type_name -> google.protobuf.Duration
	9,  // 7: resilience.policy.gloo.solo.io.OutlierDetectionPolicySpec.Config.base_ejection_time:type_name -> google.protobuf.Duration
	8,  // 8: resilience.policy.gloo.solo.io.OutlierDetectionPolicySpec.Config.consecutive_gateway_errors:type_name -> google.protobuf.UInt32Value
	10, // 9: resilience.policy.gloo.solo.io.OutlierDetectionPolicyReport.WorkspacesEntry.value:type_name -> common.gloo.solo.io.Report
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_init()
}
func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_init() {
	if File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutlierDetectionPolicySpec); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutlierDetectionPolicyStatus); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutlierDetectionPolicyReport); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutlierDetectionPolicySpec_Config); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto = out.File
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_outlier_detection_policy_proto_depIdxs = nil
}
