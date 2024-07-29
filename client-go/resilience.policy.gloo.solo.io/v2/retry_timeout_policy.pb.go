// {{% reuse "conrefs/snippets/policies/ov_retry_timeout.md" %}}
//
// **Retries**: A retry specifies the maximum number of times an Envoy proxy attempts
// to connect to a service if the initial call fails. Retries can enhance
// service availability and application performance by making sure
// that calls don’t fail permanently because of transient problems,
// such as a temporarily overloaded service or network.
// The interval between retries (25ms+) is variable and determined automatically by Istio
// to prevent the called service from being overwhelmed with requests.
// The default retry behavior for HTTP requests is to retry twice before returning the error.
//
// Like timeouts, Istio’s default retry behavior might not suit your application needs
// in terms of latency or availability. For example, too many retries to a failed service
// can slow things down. Also like timeouts, you can adjust your retry settings
// on a per-route basis.
// For more information, see the [Istio documentation](https://istio.io/latest/docs/concepts/traffic-management/#retries).
//
// **Timeouts**: A timeout is the amount of time that an Envoy proxy waits for replies
// from a service, ensuring that services don’t hang around waiting for replies forever.
// This allows calls to succeed or fail within a predictable timeframe.
// By default, the Envoy timeout for HTTP requests is disabled in Istio.
// For some applications and services, Istio’s default timeout might not be appropriate.
// For example, a timeout that is too long can result in excessive latency from waiting
// for replies from failing services. On the other hand, a timeout that is too short
// can result in calls failing unnecessarily while waiting for an operation
// that needs responses from multiple services.
//
// To find and use your optimal timeout settings, you can set timeouts dynamically per route.
// For more information, see the [Istio documentation](https://istio.io/latest/docs/concepts/traffic-management/#timeouts).
//
// ## Examples
//
// **Retry only**:
// ```yaml
// apiVersion: resilience.policy.gloo.solo.io/v2
// kind: RetryTimeoutPolicy
// metadata:
//   name: retry-only
//   namespace: bookinfo
//   annotations:
//     cluster.solo.io/cluster: $REMOTE_CLUSTER1
// spec:
//   applyToRoutes:
//     - route:
//         labels:
//           route: ratings # matches on route table route's labels
//   config:
//     retries:
//       attempts: 5 # optional (default is 2)
//       perTryTimeout: 2s
//       # retryOn specifies the conditions under which retry takes place.
//       # One or more policies can be specified using a ‘,’ delimited list.
//       retryOn: "connect-failure,refused-stream,unavailable,cancelled,retriable-status-codes"
//       # retryRemoteLocalities specifies whether the retries should retry to other localities, will default to false
//       retryRemoteLocalities: true
// ```
//
// **Timeout**:
// ```yaml
// apiVersion: resilience.policy.gloo.solo.io/v2
// kind: RetryTimeoutPolicy
// metadata:
//   name: retry-timeout
//   namespace: bookinfo
//   annotations:
//     cluster.solo.io/cluster: $REMOTE_CLUSTER1
// spec:
//   applyToRoutes:
//     - route:
//         labels:
//           route: ratings # matches on route table route's labels
//   config:
//     requestTimeout: 2s
// ```

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/policy/v2/resilience/retry_timeout_policy.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	v2 "github.com/solo-io/solo-apis/client-go/common.gloo.solo.io/v2"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RetryTimeoutPolicy is used to add automatic retries and timeouts to requests matching selected routes.
// RetryTimeoutPolicies are applied at the *Route* level.
type RetryTimeoutPolicySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Routes to apply the policy to.
	// If empty, the policy applies to all workloads in the workspace.
	ApplyToRoutes []*v2.RouteSelector `protobuf:"bytes,1,rep,name=apply_to_routes,json=applyToRoutes,proto3" json:"apply_to_routes,omitempty"`
	// The details of the retry/timeout policy to apply to the selected routes.
	Config *RetryTimeoutPolicySpec_Config `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *RetryTimeoutPolicySpec) Reset() {
	*x = RetryTimeoutPolicySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryTimeoutPolicySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryTimeoutPolicySpec) ProtoMessage() {}

func (x *RetryTimeoutPolicySpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[0]
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
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDescGZIP(), []int{0}
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

// The status of the policy after it is applied to your Gloo environment.
type RetryTimeoutPolicyStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The state and workspace conditions of the applied resource.
	Common *v2.Status `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	// The number of routes selected by the policy.
	NumSelectedRoutes uint32 `protobuf:"varint,2,opt,name=num_selected_routes,json=numSelectedRoutes,proto3" json:"num_selected_routes,omitempty"`
}

func (x *RetryTimeoutPolicyStatus) Reset() {
	*x = RetryTimeoutPolicyStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryTimeoutPolicyStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryTimeoutPolicyStatus) ProtoMessage() {}

func (x *RetryTimeoutPolicyStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[1]
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
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDescGZIP(), []int{1}
}

func (x *RetryTimeoutPolicyStatus) GetCommon() *v2.Status {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *RetryTimeoutPolicyStatus) GetNumSelectedRoutes() uint32 {
	if x != nil {
		return x.NumSelectedRoutes
	}
	return 0
}

type RetryTimeoutPolicyReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of workspaces in which the policy can apply to workloads.
	Workspaces map[string]*v2.Report `protobuf:"bytes,1,rep,name=workspaces,proto3" json:"workspaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A list of references to all routes selected by the policy.
	SelectedRoutes []*v2.RouteReference `protobuf:"bytes,2,rep,name=selected_routes,json=selectedRoutes,proto3" json:"selected_routes,omitempty"`
}

func (x *RetryTimeoutPolicyReport) Reset() {
	*x = RetryTimeoutPolicyReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryTimeoutPolicyReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryTimeoutPolicyReport) ProtoMessage() {}

func (x *RetryTimeoutPolicyReport) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryTimeoutPolicyReport.ProtoReflect.Descriptor instead.
func (*RetryTimeoutPolicyReport) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDescGZIP(), []int{2}
}

func (x *RetryTimeoutPolicyReport) GetWorkspaces() map[string]*v2.Report {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

func (x *RetryTimeoutPolicyReport) GetSelectedRoutes() []*v2.RouteReference {
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
	// For information about the value format, see the [Google protocol buffer documentation](https://protobuf.dev/reference/protobuf/google.protobuf/#duration).
	RequestTimeout *durationpb.Duration `protobuf:"bytes,4,opt,name=request_timeout,json=requestTimeout,proto3" json:"request_timeout,omitempty"`
}

func (x *RetryTimeoutPolicySpec_Config) Reset() {
	*x = RetryTimeoutPolicySpec_Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryTimeoutPolicySpec_Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryTimeoutPolicySpec_Config) ProtoMessage() {}

func (x *RetryTimeoutPolicySpec_Config) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[3]
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
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RetryTimeoutPolicySpec_Config) GetRetries() *RetryTimeoutPolicySpec_Config_RetryPolicy {
	if x != nil {
		return x.Retries
	}
	return nil
}

func (x *RetryTimeoutPolicySpec_Config) GetRequestTimeout() *durationpb.Duration {
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
	// For information about the value format, see the [Google protocol buffer documentation](https://protobuf.dev/reference/protobuf/google.protobuf/#int32-value).
	Attempts *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=attempts,proto3" json:"attempts,omitempty"`
	// Timeout per retry attempt for a given request. Format: `1h`/`1m`/`1s`/`1ms`. *Must be >= 1ms*.
	// Default is same value as request `timeout` of
	// the [HTTP route](https://istio.io/docs/reference/config/networking/virtual-service/#HTTPRoute),
	// which means no timeout.
	// For information about the value format, see the [Google protocol buffer documentation](https://protobuf.dev/reference/protobuf/google.protobuf/#duration).
	PerTryTimeout *durationpb.Duration `protobuf:"bytes,2,opt,name=per_try_timeout,json=perTryTimeout,proto3" json:"per_try_timeout,omitempty"`
	// Specifies the conditions under which retry takes place.
	// One or more policies can be specified using a ‘,’ delimited list.
	// If `retry_on` specifies a valid HTTP status, it will be added to retriable_status_codes retry policy.
	// See the [retry policies](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/router_filter#x-envoy-retry-on)
	// and [gRPC retry policies](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/router_filter#x-envoy-retry-grpc-on) for more details.
	RetryOn string `protobuf:"bytes,3,opt,name=retry_on,json=retryOn,proto3" json:"retry_on,omitempty"`
	// Flag to specify whether the retries should retry to other localities.
	// See the [retry plugin configuration](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/http/http_connection_management#retry-plugin-configuration) for more details.
	// Defaults to false.
	RetryRemoteLocalities *wrapperspb.BoolValue `protobuf:"bytes,4,opt,name=retry_remote_localities,json=retryRemoteLocalities,proto3" json:"retry_remote_localities,omitempty"`
}

func (x *RetryTimeoutPolicySpec_Config_RetryPolicy) Reset() {
	*x = RetryTimeoutPolicySpec_Config_RetryPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryTimeoutPolicySpec_Config_RetryPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryTimeoutPolicySpec_Config_RetryPolicy) ProtoMessage() {}

func (x *RetryTimeoutPolicySpec_Config_RetryPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[4]
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
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *RetryTimeoutPolicySpec_Config_RetryPolicy) GetAttempts() *wrapperspb.Int32Value {
	if x != nil {
		return x.Attempts
	}
	return nil
}

func (x *RetryTimeoutPolicySpec_Config_RetryPolicy) GetPerTryTimeout() *durationpb.Duration {
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

func (x *RetryTimeoutPolicySpec_Config_RetryPolicy) GetRetryRemoteLocalities() *wrapperspb.BoolValue {
	if x != nil {
		return x.RetryRemoteLocalities
	}
	return nil
}

var File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDesc = []byte{
	0x0a, 0x68, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2f,
	0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x72, 0x65, 0x73, 0x69,
	0x6c, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x52,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x32, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68,
	0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xea, 0x04, 0x0a, 0x16, 0x52, 0x65, 0x74, 0x72, 0x79, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4a,
	0x0a, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0d, 0x61, 0x70, 0x70,
	0x6c, 0x79, 0x54, 0x6f, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x55, 0x0a, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x72, 0x65, 0x73,
	0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x74, 0x72,
	0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70,
	0x65, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x1a, 0xac, 0x03, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x63, 0x0a, 0x07,
	0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x49, 0x2e,
	0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52,
	0x65, 0x74, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x65, 0x74,
	0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x07, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x42, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x1a, 0xf8, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x74, 0x72, 0x79, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x37, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x12, 0x41,
	0x0a, 0x0f, 0x70, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0d, 0x70, 0x65, 0x72, 0x54, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x74, 0x72, 0x79, 0x4f, 0x6e, 0x12, 0x52, 0x0a, 0x17,
	0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x15, 0x72, 0x65, 0x74, 0x72, 0x79,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x22, 0x7f, 0x0a, 0x18, 0x52, 0x65, 0x74, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11,
	0x6e, 0x75, 0x6d, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x22, 0xae, 0x02, 0x0a, 0x18, 0x52, 0x65, 0x74, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x68,
	0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x48, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x4c, 0x0a, 0x0f, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x5a, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x5f, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65,
	0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x32, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0,
	0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDescData = file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_goTypes = []interface{}{
	(*RetryTimeoutPolicySpec)(nil),                    // 0: resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec
	(*RetryTimeoutPolicyStatus)(nil),                  // 1: resilience.policy.gloo.solo.io.RetryTimeoutPolicyStatus
	(*RetryTimeoutPolicyReport)(nil),                  // 2: resilience.policy.gloo.solo.io.RetryTimeoutPolicyReport
	(*RetryTimeoutPolicySpec_Config)(nil),             // 3: resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec.Config
	(*RetryTimeoutPolicySpec_Config_RetryPolicy)(nil), // 4: resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec.Config.RetryPolicy
	nil,                           // 5: resilience.policy.gloo.solo.io.RetryTimeoutPolicyReport.WorkspacesEntry
	(*v2.RouteSelector)(nil),      // 6: common.gloo.solo.io.RouteSelector
	(*v2.Status)(nil),             // 7: common.gloo.solo.io.Status
	(*v2.RouteReference)(nil),     // 8: common.gloo.solo.io.RouteReference
	(*durationpb.Duration)(nil),   // 9: google.protobuf.Duration
	(*wrapperspb.Int32Value)(nil), // 10: google.protobuf.Int32Value
	(*wrapperspb.BoolValue)(nil),  // 11: google.protobuf.BoolValue
	(*v2.Report)(nil),             // 12: common.gloo.solo.io.Report
}
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_depIdxs = []int32{
	6,  // 0: resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec.apply_to_routes:type_name -> common.gloo.solo.io.RouteSelector
	3,  // 1: resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec.config:type_name -> resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec.Config
	7,  // 2: resilience.policy.gloo.solo.io.RetryTimeoutPolicyStatus.common:type_name -> common.gloo.solo.io.Status
	5,  // 3: resilience.policy.gloo.solo.io.RetryTimeoutPolicyReport.workspaces:type_name -> resilience.policy.gloo.solo.io.RetryTimeoutPolicyReport.WorkspacesEntry
	8,  // 4: resilience.policy.gloo.solo.io.RetryTimeoutPolicyReport.selected_routes:type_name -> common.gloo.solo.io.RouteReference
	4,  // 5: resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec.Config.retries:type_name -> resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec.Config.RetryPolicy
	9,  // 6: resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec.Config.request_timeout:type_name -> google.protobuf.Duration
	10, // 7: resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec.Config.RetryPolicy.attempts:type_name -> google.protobuf.Int32Value
	9,  // 8: resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec.Config.RetryPolicy.per_try_timeout:type_name -> google.protobuf.Duration
	11, // 9: resilience.policy.gloo.solo.io.RetryTimeoutPolicySpec.Config.RetryPolicy.retry_remote_localities:type_name -> google.protobuf.BoolValue
	12, // 10: resilience.policy.gloo.solo.io.RetryTimeoutPolicyReport.WorkspacesEntry.value:type_name -> common.gloo.solo.io.Report
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_init()
}
func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_init() {
	if File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetryTimeoutPolicyReport); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto = out.File
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_resilience_retry_timeout_policy_proto_depIdxs = nil
}