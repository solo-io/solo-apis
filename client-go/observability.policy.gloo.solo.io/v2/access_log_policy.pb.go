// {{% reuse "conrefs/snippets/policies/ov_access_logs.md" %}}
// You can use access log policies to configure log collection for workloads
// that have injected sidecars or are standalone proxies, such as gateways.
// AccessLogPolicies are applied at the *Workload* level.
//
// Note: Be sure to [enable access logging]({{< link path="/observability/tools/access-logs/" >}})
// by modifying your default Istio operator installation.
//
// ## Example
// This example filters access logs for the `reviews` service, so that only
// logs that contain the header `foo: bar` are recorded.
// ```yaml
// apiVersion: observability.policy.gloo.solo.io/v2
// kind: AccessLogPolicy
// metadata:
//   name: access-log-policy
//   namespace: bookinfo
// spec:
//   applyToWorkloads:
//   - selector:
//       cluster: cluster1
//       labels:
//         app: reviews
//       namespace: bookinfo
//   config:
//     filters:
//     #- statusCodeMatcher:
//     #    value: 200
//     #    comparator: EQ
//     - headerMatcher:
//         name: foo
//         value: bar
//         regex: false
//         invertMatch: false
//     #includedRequestHeaders:
//     #  - x-user-agent
//     #includedResponseHeaders:
//     #  - x-server
//     #includedResponseTrailers:
//     #  - x-expires
//     #includedFilterStateObjects:
// ```

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/policy/v2/observability/access_log_policy.proto

package v2

import (
	reflect "reflect"
	sync "sync"

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

// Specifications for the policy.
type AccessLogPolicySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Select the workloads where access logs will be collected.
	// If empty, the policy applies to all workloads in the workspace.
	ApplyToWorkloads []*v2.WorkloadSelector `protobuf:"bytes,1,rep,name=apply_to_workloads,json=applyToWorkloads,proto3" json:"apply_to_workloads,omitempty"`
	// Details of the policy specifying how to collect access logs from the selected workloads.
	Config *AccessLogPolicySpec_Config `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *AccessLogPolicySpec) Reset() {
	*x = AccessLogPolicySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessLogPolicySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessLogPolicySpec) ProtoMessage() {}

func (x *AccessLogPolicySpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessLogPolicySpec.ProtoReflect.Descriptor instead.
func (*AccessLogPolicySpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_rawDescGZIP(), []int{0}
}

func (x *AccessLogPolicySpec) GetApplyToWorkloads() []*v2.WorkloadSelector {
	if x != nil {
		return x.ApplyToWorkloads
	}
	return nil
}

func (x *AccessLogPolicySpec) GetConfig() *AccessLogPolicySpec_Config {
	if x != nil {
		return x.Config
	}
	return nil
}

// The status of the policy after it is applied to your Gloo environment.
type AccessLogPolicyStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The state and workspace conditions of the applied policy.
	Common *v2.Status `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	// The number of workloads selected by the policy.
	NumSelectedWorkloads uint32 `protobuf:"varint,2,opt,name=num_selected_workloads,json=numSelectedWorkloads,proto3" json:"num_selected_workloads,omitempty"`
}

func (x *AccessLogPolicyStatus) Reset() {
	*x = AccessLogPolicyStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessLogPolicyStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessLogPolicyStatus) ProtoMessage() {}

func (x *AccessLogPolicyStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessLogPolicyStatus.ProtoReflect.Descriptor instead.
func (*AccessLogPolicyStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_rawDescGZIP(), []int{1}
}

func (x *AccessLogPolicyStatus) GetCommon() *v2.Status {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *AccessLogPolicyStatus) GetNumSelectedWorkloads() uint32 {
	if x != nil {
		return x.NumSelectedWorkloads
	}
	return 0
}

// The report shows the resources that the policy selects after the policy is successfully applied.
type AccessLogPolicyReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of workspaces in which the policy can apply to workloads.
	Workspaces map[string]*v2.Report `protobuf:"bytes,1,rep,name=workspaces,proto3" json:"workspaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A list of workloads selected by the policy.
	SelectedWorkloads []*v2.WorkloadReference `protobuf:"bytes,2,rep,name=selected_workloads,json=selectedWorkloads,proto3" json:"selected_workloads,omitempty"`
}

func (x *AccessLogPolicyReport) Reset() {
	*x = AccessLogPolicyReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessLogPolicyReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessLogPolicyReport) ProtoMessage() {}

func (x *AccessLogPolicyReport) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessLogPolicyReport.ProtoReflect.Descriptor instead.
func (*AccessLogPolicyReport) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_rawDescGZIP(), []int{2}
}

func (x *AccessLogPolicyReport) GetWorkspaces() map[string]*v2.Report {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

func (x *AccessLogPolicyReport) GetSelectedWorkloads() []*v2.WorkloadReference {
	if x != nil {
		return x.SelectedWorkloads
	}
	return nil
}

// Details of the policy specifying how to collect access logs from the selected workloads.
type AccessLogPolicySpec_Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Criteria for determining which access logs are recorded for the workload.
	// The list is disjunctive, meaning that a request is logged if it matches any filter.
	// If empty, all access logs are recorded.
	Filters []*AccessLogPolicySpec_Config_Filter `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
	// Request headers to include in access logs.
	IncludedRequestHeaders []string `protobuf:"bytes,3,rep,name=included_request_headers,json=includedRequestHeaders,proto3" json:"included_request_headers,omitempty"`
	// Response headers to include in access logs.
	IncludedResponseHeaders []string `protobuf:"bytes,4,rep,name=included_response_headers,json=includedResponseHeaders,proto3" json:"included_response_headers,omitempty"`
	// Response trailers to include in access logs.
	IncludedResponseTrailers []string `protobuf:"bytes,5,rep,name=included_response_trailers,json=includedResponseTrailers,proto3" json:"included_response_trailers,omitempty"`
	// Filter state objects to include in access logs.
	IncludedFilterStateObjects []string `protobuf:"bytes,6,rep,name=included_filter_state_objects,json=includedFilterStateObjects,proto3" json:"included_filter_state_objects,omitempty"`
}

func (x *AccessLogPolicySpec_Config) Reset() {
	*x = AccessLogPolicySpec_Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessLogPolicySpec_Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessLogPolicySpec_Config) ProtoMessage() {}

func (x *AccessLogPolicySpec_Config) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessLogPolicySpec_Config.ProtoReflect.Descriptor instead.
func (*AccessLogPolicySpec_Config) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AccessLogPolicySpec_Config) GetFilters() []*AccessLogPolicySpec_Config_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *AccessLogPolicySpec_Config) GetIncludedRequestHeaders() []string {
	if x != nil {
		return x.IncludedRequestHeaders
	}
	return nil
}

func (x *AccessLogPolicySpec_Config) GetIncludedResponseHeaders() []string {
	if x != nil {
		return x.IncludedResponseHeaders
	}
	return nil
}

func (x *AccessLogPolicySpec_Config) GetIncludedResponseTrailers() []string {
	if x != nil {
		return x.IncludedResponseTrailers
	}
	return nil
}

func (x *AccessLogPolicySpec_Config) GetIncludedFilterStateObjects() []string {
	if x != nil {
		return x.IncludedFilterStateObjects
	}
	return nil
}

// Criteria for recording access logs. A request must match all specified criteria in the filter to be recorded.
type AccessLogPolicySpec_Config_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Criteria type.
	//
	// Types that are assignable to Type:
	//
	//	*AccessLogPolicySpec_Config_Filter_StatusCodeMatcher
	//	*AccessLogPolicySpec_Config_Filter_HeaderMatcher
	Type isAccessLogPolicySpec_Config_Filter_Type `protobuf_oneof:"type"`
}

func (x *AccessLogPolicySpec_Config_Filter) Reset() {
	*x = AccessLogPolicySpec_Config_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessLogPolicySpec_Config_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessLogPolicySpec_Config_Filter) ProtoMessage() {}

func (x *AccessLogPolicySpec_Config_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessLogPolicySpec_Config_Filter.ProtoReflect.Descriptor instead.
func (*AccessLogPolicySpec_Config_Filter) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (m *AccessLogPolicySpec_Config_Filter) GetType() isAccessLogPolicySpec_Config_Filter_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *AccessLogPolicySpec_Config_Filter) GetStatusCodeMatcher() *v2.StatusCodeMatcher {
	if x, ok := x.GetType().(*AccessLogPolicySpec_Config_Filter_StatusCodeMatcher); ok {
		return x.StatusCodeMatcher
	}
	return nil
}

func (x *AccessLogPolicySpec_Config_Filter) GetHeaderMatcher() *v2.HeaderMatcher {
	if x, ok := x.GetType().(*AccessLogPolicySpec_Config_Filter_HeaderMatcher); ok {
		return x.HeaderMatcher
	}
	return nil
}

type isAccessLogPolicySpec_Config_Filter_Type interface {
	isAccessLogPolicySpec_Config_Filter_Type()
}

type AccessLogPolicySpec_Config_Filter_StatusCodeMatcher struct {
	// Matches against a response HTTP status code. Omit to match any status code.
	StatusCodeMatcher *v2.StatusCodeMatcher `protobuf:"bytes,1,opt,name=status_code_matcher,json=statusCodeMatcher,proto3,oneof"`
}

type AccessLogPolicySpec_Config_Filter_HeaderMatcher struct {
	// Matches against a request or response HTTP header. Omit to match any headers.
	HeaderMatcher *v2.HeaderMatcher `protobuf:"bytes,2,opt,name=header_matcher,json=headerMatcher,proto3,oneof"`
}

func (*AccessLogPolicySpec_Config_Filter_StatusCodeMatcher) isAccessLogPolicySpec_Config_Filter_Type() {
}

func (*AccessLogPolicySpec_Config_Filter_HeaderMatcher) isAccessLogPolicySpec_Config_Filter_Type() {}

var File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_rawDesc = []byte{
	0x0a, 0x6b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67,
	0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x6f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x59, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65,
	0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x32,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x68, 0x74, 0x74, 0x70,
	0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65,
	0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x52,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xdd, 0x05, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x53, 0x0a, 0x12, 0x61, 0x70,
	0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12,
	0x55, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3d, 0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x99, 0x04, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x5e, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x44, 0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x12, 0x38, 0x0a, 0x18, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x16, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x3a, 0x0a, 0x19, 0x69,
	0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x17,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x3c, 0x0a, 0x1a, 0x69, 0x6e, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x72, 0x61,
	0x69, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x18, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x72, 0x61,
	0x69, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x41, 0x0a, 0x1d, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x1a, 0x69, 0x6e,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x1a, 0xb7, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x12, 0x58, 0x0a, 0x13, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x48, 0x00, 0x52, 0x11, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x4b, 0x0a,
	0x0e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0d, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x15, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x12, 0x34, 0x0a, 0x16, 0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x14, 0x6e, 0x75, 0x6d, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x57, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x22, 0xb4, 0x02, 0x0a, 0x15, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4c, 0x6f, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x68, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4c, 0x6f, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x55, 0x0a, 0x12, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x11, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x73, 0x1a, 0x5a, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x65,
	0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x32, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04,
	0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_rawDescData = file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_goTypes = []interface{}{
	(*AccessLogPolicySpec)(nil),               // 0: observability.policy.gloo.solo.io.AccessLogPolicySpec
	(*AccessLogPolicyStatus)(nil),             // 1: observability.policy.gloo.solo.io.AccessLogPolicyStatus
	(*AccessLogPolicyReport)(nil),             // 2: observability.policy.gloo.solo.io.AccessLogPolicyReport
	(*AccessLogPolicySpec_Config)(nil),        // 3: observability.policy.gloo.solo.io.AccessLogPolicySpec.Config
	(*AccessLogPolicySpec_Config_Filter)(nil), // 4: observability.policy.gloo.solo.io.AccessLogPolicySpec.Config.Filter
	nil,                          // 5: observability.policy.gloo.solo.io.AccessLogPolicyReport.WorkspacesEntry
	(*v2.WorkloadSelector)(nil),  // 6: common.gloo.solo.io.WorkloadSelector
	(*v2.Status)(nil),            // 7: common.gloo.solo.io.Status
	(*v2.WorkloadReference)(nil), // 8: common.gloo.solo.io.WorkloadReference
	(*v2.StatusCodeMatcher)(nil), // 9: common.gloo.solo.io.StatusCodeMatcher
	(*v2.HeaderMatcher)(nil),     // 10: common.gloo.solo.io.HeaderMatcher
	(*v2.Report)(nil),            // 11: common.gloo.solo.io.Report
}
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_depIdxs = []int32{
	6,  // 0: observability.policy.gloo.solo.io.AccessLogPolicySpec.apply_to_workloads:type_name -> common.gloo.solo.io.WorkloadSelector
	3,  // 1: observability.policy.gloo.solo.io.AccessLogPolicySpec.config:type_name -> observability.policy.gloo.solo.io.AccessLogPolicySpec.Config
	7,  // 2: observability.policy.gloo.solo.io.AccessLogPolicyStatus.common:type_name -> common.gloo.solo.io.Status
	5,  // 3: observability.policy.gloo.solo.io.AccessLogPolicyReport.workspaces:type_name -> observability.policy.gloo.solo.io.AccessLogPolicyReport.WorkspacesEntry
	8,  // 4: observability.policy.gloo.solo.io.AccessLogPolicyReport.selected_workloads:type_name -> common.gloo.solo.io.WorkloadReference
	4,  // 5: observability.policy.gloo.solo.io.AccessLogPolicySpec.Config.filters:type_name -> observability.policy.gloo.solo.io.AccessLogPolicySpec.Config.Filter
	9,  // 6: observability.policy.gloo.solo.io.AccessLogPolicySpec.Config.Filter.status_code_matcher:type_name -> common.gloo.solo.io.StatusCodeMatcher
	10, // 7: observability.policy.gloo.solo.io.AccessLogPolicySpec.Config.Filter.header_matcher:type_name -> common.gloo.solo.io.HeaderMatcher
	11, // 8: observability.policy.gloo.solo.io.AccessLogPolicyReport.WorkspacesEntry.value:type_name -> common.gloo.solo.io.Report
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_init()
}
func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_init() {
	if File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessLogPolicySpec); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessLogPolicyStatus); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessLogPolicyReport); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessLogPolicySpec_Config); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessLogPolicySpec_Config_Filter); i {
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
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*AccessLogPolicySpec_Config_Filter_StatusCodeMatcher)(nil),
		(*AccessLogPolicySpec_Config_Filter_HeaderMatcher)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto = out.File
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_observability_access_log_policy_proto_depIdxs = nil
}
