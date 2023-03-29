// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/policy/v2/resilience/graphql_persisted_query_cache_policy.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/golang/protobuf/ptypes/wrappers"
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

// To improve network performance for large query strings, the GraphQL filter supports
// automatic persisted queries. A persisted query is a query string that's cached on the server side,
// along with it's hash. The client can then send the hash instead of the full query string, thus
// reducing request sizes (responses are unaffected).
// Persisted queries are especially effective when clients send queries as GET requests.
// This enables clients to take advantage of the browser cache and integrate with a CDN.
// If left empty, persisted query cache is not used.
type GraphQLPersistedQueryCachePolicySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Select the graphql routes where the policy will be applied.
	// If left empty, will apply to no graphql routes.
	// selection via Destination selectors will not work for this policy, and only
	// route selectors will work.
	// If more than one GraphQLPersistedQueryCachePolicy applies to a GraphQLRoute, the oldest
	// one will be applied.
	ApplyToRoutes []*v2.RouteSelector `protobuf:"bytes,1,rep,name=apply_to_routes,json=applyToRoutes,proto3" json:"apply_to_routes,omitempty"`
	// Persisted query cache configuration.
	Config *GraphQLPersistedQueryCachePolicySpec_Config `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *GraphQLPersistedQueryCachePolicySpec) Reset() {
	*x = GraphQLPersistedQueryCachePolicySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphQLPersistedQueryCachePolicySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphQLPersistedQueryCachePolicySpec) ProtoMessage() {}

func (x *GraphQLPersistedQueryCachePolicySpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphQLPersistedQueryCachePolicySpec.ProtoReflect.Descriptor instead.
func (*GraphQLPersistedQueryCachePolicySpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_rawDescGZIP(), []int{0}
}

func (x *GraphQLPersistedQueryCachePolicySpec) GetApplyToRoutes() []*v2.RouteSelector {
	if x != nil {
		return x.ApplyToRoutes
	}
	return nil
}

func (x *GraphQLPersistedQueryCachePolicySpec) GetConfig() *GraphQLPersistedQueryCachePolicySpec_Config {
	if x != nil {
		return x.Config
	}
	return nil
}

// reflects the status of the GraphQLPersistedQueryCachePolicy
type GraphQLPersistedQueryCachePolicyStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Global *v2.GenericGlobalStatus `protobuf:"bytes,1,opt,name=global,proto3" json:"global,omitempty"`
	// The status of the resource in each workspace that it exists in.
	Workspaces map[string]*v2.WorkspaceStatus `protobuf:"bytes,2,rep,name=workspaces,proto3" json:"workspaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A list of references to all routes selected by the policy.
	SelectedRoutes []*v2.RouteReference `protobuf:"bytes,3,rep,name=selected_routes,json=selectedRoutes,proto3" json:"selected_routes,omitempty"`
}

func (x *GraphQLPersistedQueryCachePolicyStatus) Reset() {
	*x = GraphQLPersistedQueryCachePolicyStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphQLPersistedQueryCachePolicyStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphQLPersistedQueryCachePolicyStatus) ProtoMessage() {}

func (x *GraphQLPersistedQueryCachePolicyStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphQLPersistedQueryCachePolicyStatus.ProtoReflect.Descriptor instead.
func (*GraphQLPersistedQueryCachePolicyStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_rawDescGZIP(), []int{1}
}

func (x *GraphQLPersistedQueryCachePolicyStatus) GetGlobal() *v2.GenericGlobalStatus {
	if x != nil {
		return x.Global
	}
	return nil
}

func (x *GraphQLPersistedQueryCachePolicyStatus) GetWorkspaces() map[string]*v2.WorkspaceStatus {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

func (x *GraphQLPersistedQueryCachePolicyStatus) GetSelectedRoutes() []*v2.RouteReference {
	if x != nil {
		return x.SelectedRoutes
	}
	return nil
}

// $hide_from_docs
type GraphQLPersistedQueryCachePolicyNewStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Common *v2.Status `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	// The number of routes selected by the policy.
	SelectedRoutes uint32 `protobuf:"varint,2,opt,name=selected_routes,json=selectedRoutes,proto3" json:"selected_routes,omitempty"`
}

func (x *GraphQLPersistedQueryCachePolicyNewStatus) Reset() {
	*x = GraphQLPersistedQueryCachePolicyNewStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphQLPersistedQueryCachePolicyNewStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphQLPersistedQueryCachePolicyNewStatus) ProtoMessage() {}

func (x *GraphQLPersistedQueryCachePolicyNewStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphQLPersistedQueryCachePolicyNewStatus.ProtoReflect.Descriptor instead.
func (*GraphQLPersistedQueryCachePolicyNewStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_rawDescGZIP(), []int{2}
}

func (x *GraphQLPersistedQueryCachePolicyNewStatus) GetCommon() *v2.Status {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *GraphQLPersistedQueryCachePolicyNewStatus) GetSelectedRoutes() uint32 {
	if x != nil {
		return x.SelectedRoutes
	}
	return 0
}

// $hide_from_docs
type GraphQLPersistedQueryCachePolicyReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspaces map[string]*v2.Report `protobuf:"bytes,1,rep,name=workspaces,proto3" json:"workspaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A list of references to all routes selected by the policy.
	SelectedRoutes []*v2.RouteReference `protobuf:"bytes,2,rep,name=selected_routes,json=selectedRoutes,proto3" json:"selected_routes,omitempty"`
}

func (x *GraphQLPersistedQueryCachePolicyReport) Reset() {
	*x = GraphQLPersistedQueryCachePolicyReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphQLPersistedQueryCachePolicyReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphQLPersistedQueryCachePolicyReport) ProtoMessage() {}

func (x *GraphQLPersistedQueryCachePolicyReport) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphQLPersistedQueryCachePolicyReport.ProtoReflect.Descriptor instead.
func (*GraphQLPersistedQueryCachePolicyReport) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_rawDescGZIP(), []int{3}
}

func (x *GraphQLPersistedQueryCachePolicyReport) GetWorkspaces() map[string]*v2.Report {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

func (x *GraphQLPersistedQueryCachePolicyReport) GetSelectedRoutes() []*v2.RouteReference {
	if x != nil {
		return x.SelectedRoutes
	}
	return nil
}

type GraphQLPersistedQueryCachePolicySpec_Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of queries to store in the persisted query cache, defaults to 1000.
	CacheSize uint32 `protobuf:"varint,1,opt,name=cache_size,json=cacheSize,proto3" json:"cache_size,omitempty"`
}

func (x *GraphQLPersistedQueryCachePolicySpec_Config) Reset() {
	*x = GraphQLPersistedQueryCachePolicySpec_Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphQLPersistedQueryCachePolicySpec_Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphQLPersistedQueryCachePolicySpec_Config) ProtoMessage() {}

func (x *GraphQLPersistedQueryCachePolicySpec_Config) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphQLPersistedQueryCachePolicySpec_Config.ProtoReflect.Descriptor instead.
func (*GraphQLPersistedQueryCachePolicySpec_Config) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GraphQLPersistedQueryCachePolicySpec_Config) GetCacheSize() uint32 {
	if x != nil {
		return x.CacheSize
	}
	return 0
}

var File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_rawDesc = []byte{
	0x0a, 0x77, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x64,
	0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x72, 0x65, 0x73, 0x69, 0x6c,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x51, 0x67,
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
	0x1a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x02, 0x0a, 0x24, 0x47, 0x72, 0x61, 0x70, 0x68, 0x51, 0x4c,
	0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4a, 0x0a,
	0x0f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c,
	0x79, 0x54, 0x6f, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x63, 0x0a, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x72, 0x65, 0x73, 0x69,
	0x6c, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68,
	0x51, 0x4c, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x27,
	0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x95, 0x03, 0x0a, 0x26, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x51, 0x4c, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x43, 0x61, 0x63, 0x68, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x40, 0x0a, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x12, 0x76, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x56, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x6c,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x51,
	0x4c, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x4c, 0x0a, 0x0f,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0e, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x63, 0x0a, 0x0f, 0x57, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x3a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x89, 0x01, 0x0a, 0x29, 0x47, 0x72, 0x61, 0x70, 0x68, 0x51, 0x4c, 0x50, 0x65, 0x72, 0x73, 0x69,
	0x73, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x61, 0x63, 0x68, 0x65, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x4e, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a,
	0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0xca, 0x02, 0x0a, 0x26,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x51, 0x4c, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x64,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x61, 0x63, 0x68, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x76, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x56, 0x2e, 0x72, 0x65, 0x73,
	0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x51, 0x4c, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x43, 0x61, 0x63, 0x68, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x4c,
	0x0a, 0x0f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0e, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x5a, 0x0a, 0x0f,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x54, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f,
	0x76, 0x32, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_goTypes = []interface{}{
	(*GraphQLPersistedQueryCachePolicySpec)(nil),        // 0: resilience.policy.gloo.solo.io.GraphQLPersistedQueryCachePolicySpec
	(*GraphQLPersistedQueryCachePolicyStatus)(nil),      // 1: resilience.policy.gloo.solo.io.GraphQLPersistedQueryCachePolicyStatus
	(*GraphQLPersistedQueryCachePolicyNewStatus)(nil),   // 2: resilience.policy.gloo.solo.io.GraphQLPersistedQueryCachePolicyNewStatus
	(*GraphQLPersistedQueryCachePolicyReport)(nil),      // 3: resilience.policy.gloo.solo.io.GraphQLPersistedQueryCachePolicyReport
	(*GraphQLPersistedQueryCachePolicySpec_Config)(nil), // 4: resilience.policy.gloo.solo.io.GraphQLPersistedQueryCachePolicySpec.Config
	nil,                            // 5: resilience.policy.gloo.solo.io.GraphQLPersistedQueryCachePolicyStatus.WorkspacesEntry
	nil,                            // 6: resilience.policy.gloo.solo.io.GraphQLPersistedQueryCachePolicyReport.WorkspacesEntry
	(*v2.RouteSelector)(nil),       // 7: common.gloo.solo.io.RouteSelector
	(*v2.GenericGlobalStatus)(nil), // 8: common.gloo.solo.io.GenericGlobalStatus
	(*v2.RouteReference)(nil),      // 9: common.gloo.solo.io.RouteReference
	(*v2.Status)(nil),              // 10: common.gloo.solo.io.Status
	(*v2.WorkspaceStatus)(nil),     // 11: common.gloo.solo.io.WorkspaceStatus
	(*v2.Report)(nil),              // 12: common.gloo.solo.io.Report
}
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_depIdxs = []int32{
	7,  // 0: resilience.policy.gloo.solo.io.GraphQLPersistedQueryCachePolicySpec.apply_to_routes:type_name -> common.gloo.solo.io.RouteSelector
	4,  // 1: resilience.policy.gloo.solo.io.GraphQLPersistedQueryCachePolicySpec.config:type_name -> resilience.policy.gloo.solo.io.GraphQLPersistedQueryCachePolicySpec.Config
	8,  // 2: resilience.policy.gloo.solo.io.GraphQLPersistedQueryCachePolicyStatus.global:type_name -> common.gloo.solo.io.GenericGlobalStatus
	5,  // 3: resilience.policy.gloo.solo.io.GraphQLPersistedQueryCachePolicyStatus.workspaces:type_name -> resilience.policy.gloo.solo.io.GraphQLPersistedQueryCachePolicyStatus.WorkspacesEntry
	9,  // 4: resilience.policy.gloo.solo.io.GraphQLPersistedQueryCachePolicyStatus.selected_routes:type_name -> common.gloo.solo.io.RouteReference
	10, // 5: resilience.policy.gloo.solo.io.GraphQLPersistedQueryCachePolicyNewStatus.common:type_name -> common.gloo.solo.io.Status
	6,  // 6: resilience.policy.gloo.solo.io.GraphQLPersistedQueryCachePolicyReport.workspaces:type_name -> resilience.policy.gloo.solo.io.GraphQLPersistedQueryCachePolicyReport.WorkspacesEntry
	9,  // 7: resilience.policy.gloo.solo.io.GraphQLPersistedQueryCachePolicyReport.selected_routes:type_name -> common.gloo.solo.io.RouteReference
	11, // 8: resilience.policy.gloo.solo.io.GraphQLPersistedQueryCachePolicyStatus.WorkspacesEntry.value:type_name -> common.gloo.solo.io.WorkspaceStatus
	12, // 9: resilience.policy.gloo.solo.io.GraphQLPersistedQueryCachePolicyReport.WorkspacesEntry.value:type_name -> common.gloo.solo.io.Report
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphQLPersistedQueryCachePolicySpec); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphQLPersistedQueryCachePolicyStatus); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphQLPersistedQueryCachePolicyNewStatus); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphQLPersistedQueryCachePolicyReport); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphQLPersistedQueryCachePolicySpec_Config); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_policy_v2_resilience_graphql_persisted_query_cache_policy_proto_depIdxs = nil
}
