// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/policy/v2/resilience/retry_timeout_policy.proto

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

// RetryTimeoutPolicy is used to add automatic retries and timeouts to requests matching selected routes.
// RetryTimeoutPolicies are applied at the *Route* level.
type RetryTimeoutPolicySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// select the routes where the policy will be applied
	// if left empty, will apply to all routes in the workspace.
	ApplyToRoutes []*v2.RouteSelector `protobuf:"bytes,1,rep,name=apply_to_routes,json=applyToRoutes,proto3" json:"apply_to_routes,omitempty"`
	// The details of the retry/timeout policy to apply to the selected routes.
	Config *RetryTimeoutPolicySpec_Config `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *RetryTimeoutPolicySpec) Reset() {
	*x = RetryTimeoutPolicySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryTimeoutPolicySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryTimeoutPolicySpec) ProtoMessage() {}

func (x *RetryTimeoutPolicySpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryTimeoutPolicySpec.ProtoReflect.Descriptor instead.
func (*RetryTimeoutPolicySpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDescGZIP(), []int{0}
}

func (x *RetryTimeoutPolicySpec) GetApplyToRoutes() []*v2.RouteSelector {
	if x != nil {
		return x.ApplyToRoutes
	}
	return nil
}

func (x *RetryTimeoutPolicySpec) GetConfig() *RetryTimeoutPolicySpec_Config {
	if x != nil {
		return x.Config
	}
	return nil
}

// reflects the status of the RetryTimeoutPolicy
type RetryTimeoutPolicyStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Global *v2.GenericGlobalStatus `protobuf:"bytes,1,opt,name=global,proto3" json:"global,omitempty"`
	// The status of the resource in each workspace that it exists in.
	Workspaces map[string]*v2.WorkspaceStatus `protobuf:"bytes,2,rep,name=workspaces,proto3" json:"workspaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Routes selected by the policy
	SelectedRoutes []*v2.RouteReference `protobuf:"bytes,3,rep,name=selected_routes,json=selectedRoutes,proto3" json:"selected_routes,omitempty"`
}

func (x *RetryTimeoutPolicyStatus) Reset() {
	*x = RetryTimeoutPolicyStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryTimeoutPolicyStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryTimeoutPolicyStatus) ProtoMessage() {}

func (x *RetryTimeoutPolicyStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryTimeoutPolicyStatus.ProtoReflect.Descriptor instead.
func (*RetryTimeoutPolicyStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDescGZIP(), []int{1}
}

func (x *RetryTimeoutPolicyStatus) GetGlobal() *v2.GenericGlobalStatus {
	if x != nil {
		return x.Global
	}
	return nil
}

func (x *RetryTimeoutPolicyStatus) GetWorkspaces() map[string]*v2.WorkspaceStatus {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

func (x *RetryTimeoutPolicyStatus) GetSelectedRoutes() []*v2.RouteReference {
	if x != nil {
		return x.SelectedRoutes
	}
	return nil
}

type RetryTimeoutPolicySpec_Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Set a retry policy on requests matched on the selected routes.
	Retries *RetryTimeoutPolicySpec_Config_RetryPolicy `protobuf:"bytes,3,opt,name=retries,proto3" json:"retries,omitempty"`
	// Set a timeout on requests matched on the selected routes.
	// For information about the value format, see the [Google protocol buffer documentation](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/duration).
	RequestTimeout *duration.Duration `protobuf:"bytes,4,opt,name=request_timeout,json=requestTimeout,proto3" json:"request_timeout,omitempty"`
}

func (x *RetryTimeoutPolicySpec_Config) Reset() {
	*x = RetryTimeoutPolicySpec_Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryTimeoutPolicySpec_Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryTimeoutPolicySpec_Config) ProtoMessage() {}

func (x *RetryTimeoutPolicySpec_Config) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryTimeoutPolicySpec_Config.ProtoReflect.Descriptor instead.
func (*RetryTimeoutPolicySpec_Config) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RetryTimeoutPolicySpec_Config) GetRetries() *RetryTimeoutPolicySpec_Config_RetryPolicy {
	if x != nil {
		return x.Retries
	}
	return nil
}

func (x *RetryTimeoutPolicySpec_Config) GetRequestTimeout() *duration.Duration {
	if x != nil {
		return x.RequestTimeout
	}
	return nil
}

// Specify retries for failed requests.
type RetryTimeoutPolicySpec_Config_RetryPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of retries for a given request
	// The interval between retries will be determined automatically (25ms+). When request
	// `timeout` of the [HTTP route](https://istio.io/docs/reference/config/networking/virtual-service/#HTTPRoute)
	// or `per_try_timeout` is configured, the actual number of retries attempted also depends on
	// the specified request `timeout` and `per_try_timeout` values.
	// Defaults to 2 retries.
	// For information about the value format, see the [Google protocol buffer documentation](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/int32-value).
	Attempts *wrappers.Int32Value `protobuf:"bytes,1,opt,name=attempts,proto3" json:"attempts,omitempty"`
	// Timeout per retry attempt for a given request. Format: `1h`/`1m`/`1s`/`1ms`. *Must be >= 1ms*.
	// Default is same value as request `timeout` of
	// the [HTTP route](https://istio.io/docs/reference/config/networking/virtual-service/#HTTPRoute),
	// which means no timeout.
	// For information about the value format, see the [Google protocol buffer documentation](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/duration).
	PerTryTimeout *duration.Duration `protobuf:"bytes,2,opt,name=per_try_timeout,json=perTryTimeout,proto3" json:"per_try_timeout,omitempty"`
	// Specifies the conditions under which retry takes place.
	// One or more policies can be specified using a ‘,’ delimited list.
	// If `retry_on` specifies a valid HTTP status, it will be added to retriable_status_codes retry policy.
	// See the [retry policies](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/router_filter#x-envoy-retry-on)
	// and [gRPC retry policies](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/router_filter#x-envoy-retry-grpc-on) for more details.
	RetryOn string `protobuf:"bytes,3,opt,name=retry_on,json=retryOn,proto3" json:"retry_on,omitempty"`
	// Flag to specify whether the retries should retry to other localities.
	// See the [retry plugin configuration](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/http/http_connection_management#retry-plugin-configuration) for more details.
	// Defaults to false.
	RetryRemoteLocalities *wrappers.BoolValue `protobuf:"bytes,4,opt,name=retry_remote_localities,json=retryRemoteLocalities,proto3" json:"retry_remote_localities,omitempty"`
}

func (x *RetryTimeoutPolicySpec_Config_RetryPolicy) Reset() {
	*x = RetryTimeoutPolicySpec_Config_RetryPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryTimeoutPolicySpec_Config_RetryPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryTimeoutPolicySpec_Config_RetryPolicy) ProtoMessage() {}

func (x *RetryTimeoutPolicySpec_Config_RetryPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryTimeoutPolicySpec_Config_RetryPolicy.ProtoReflect.Descriptor instead.
func (*RetryTimeoutPolicySpec_Config_RetryPolicy) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *RetryTimeoutPolicySpec_Config_RetryPolicy) GetAttempts() *wrappers.Int32Value {
	if x != nil {
		return x.Attempts
	}
	return nil
}

func (x *RetryTimeoutPolicySpec_Config_RetryPolicy) GetPerTryTimeout() *duration.Duration {
	if x != nil {
		return x.PerTryTimeout
	}
	return nil
}

func (x *RetryTimeoutPolicySpec_Config_RetryPolicy) GetRetryOn() string {
	if x != nil {
		return x.RetryOn
	}
	return ""
}

func (x *RetryTimeoutPolicySpec_Config_RetryPolicy) GetRetryRemoteLocalities() *wrappers.BoolValue {
	if x != nil {
		return x.RetryRemoteLocalities
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDesc = []byte{
	0x0a, 0x67, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x72,
	0x65, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x72, 0x65, 0x73, 0x69, 0x6c,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x51, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32,
	0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xea, 0x04, 0x0a, 0x16, 0x52, 0x65, 0x74, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4a, 0x0a, 0x0f, 0x61,
	0x70, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x54,
	0x6f, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x55, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0xac,
	0x03, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x63, 0x0a, 0x07, 0x72, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x72, 0x65, 0x73,
	0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x74, 0x72,
	0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70,
	0x65, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x07, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x42,
	0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x1a, 0xf8, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x37, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x08, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x12, 0x41, 0x0a, 0x0f, 0x70,
	0x65, 0x72, 0x5f, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0d, 0x70, 0x65, 0x72, 0x54, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x72, 0x79, 0x4f, 0x6e, 0x12, 0x52, 0x0a, 0x17, 0x72, 0x65, 0x74,
	0x72, 0x79, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x15, 0x72, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0xf9, 0x02,
	0x0a, 0x18, 0x52, 0x65, 0x74, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x40, 0x0a, 0x06, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x12, 0x68, 0x0a, 0x0a,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x48, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x4c, 0x0a, 0x0f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x0e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x1a, 0x63, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3a, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x54, 0x5a, 0x46, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2f, 0x76, 0x32, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_goTypes = []interface{}{
	(*RetryTimeoutPolicySpec)(nil),                    // 0: resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec
	(*RetryTimeoutPolicyStatus)(nil),                  // 1: resilience.policy.gloo.solo.io.RetryTimeoutPolicyStatus
	(*RetryTimeoutPolicySpec_Config)(nil),             // 2: resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec.Config
	(*RetryTimeoutPolicySpec_Config_RetryPolicy)(nil), // 3: resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec.Config.RetryPolicy
	nil,                            // 4: resilience.policy.gloo.solo.io.RetryTimeoutPolicyStatus.WorkspacesEntry
	(*v2.RouteSelector)(nil),       // 5: common.gloo.solo.io.RouteSelector
	(*v2.GenericGlobalStatus)(nil), // 6: common.gloo.solo.io.GenericGlobalStatus
	(*v2.RouteReference)(nil),      // 7: common.gloo.solo.io.RouteReference
	(*duration.Duration)(nil),      // 8: google.protobuf.Duration
	(*wrappers.Int32Value)(nil),    // 9: google.protobuf.Int32Value
	(*wrappers.BoolValue)(nil),     // 10: google.protobuf.BoolValue
	(*v2.WorkspaceStatus)(nil),     // 11: common.gloo.solo.io.WorkspaceStatus
}
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_depIdxs = []int32{
	5,  // 0: resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec.apply_to_routes:type_name -> common.gloo.solo.io.RouteSelector
	2,  // 1: resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec.config:type_name -> resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec.Config
	6,  // 2: resilience.policy.gloo.solo.io.RetryTimeoutPolicyStatus.global:type_name -> common.gloo.solo.io.GenericGlobalStatus
	4,  // 3: resilience.policy.gloo.solo.io.RetryTimeoutPolicyStatus.workspaces:type_name -> resilience.policy.gloo.solo.io.RetryTimeoutPolicyStatus.WorkspacesEntry
	7,  // 4: resilience.policy.gloo.solo.io.RetryTimeoutPolicyStatus.selected_routes:type_name -> common.gloo.solo.io.RouteReference
	3,  // 5: resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec.Config.retries:type_name -> resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec.Config.RetryPolicy
	8,  // 6: resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec.Config.request_timeout:type_name -> google.protobuf.Duration
	9,  // 7: resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec.Config.RetryPolicy.attempts:type_name -> google.protobuf.Int32Value
	8,  // 8: resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec.Config.RetryPolicy.per_try_timeout:type_name -> google.protobuf.Duration
	10, // 9: resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec.Config.RetryPolicy.retry_remote_localities:type_name -> google.protobuf.BoolValue
	11, // 10: resilience.policy.gloo.solo.io.RetryTimeoutPolicyStatus.WorkspacesEntry.value:type_name -> common.gloo.solo.io.WorkspaceStatus
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetryTimeoutPolicySpec); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetryTimeoutPolicyStatus); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetryTimeoutPolicySpec_Config); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetryTimeoutPolicySpec_Config_RetryPolicy); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_depIdxs = nil
}