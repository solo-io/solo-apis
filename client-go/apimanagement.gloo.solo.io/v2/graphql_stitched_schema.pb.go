// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/apimanagement/v2/graphql_stitched_schema.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/cue/encoding/protobuf/cue"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
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

type GraphQLStitchedSchemaSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of GraphQL Subschemas that compose this GraphQL stitched schema.
	Subschemas []*GraphQLStitchedSchemaSpec_Subschema `protobuf:"bytes,1,rep,name=subschemas,proto3" json:"subschemas,omitempty"`
	Options    *GraphQLStitchedSchemaSpec_Options     `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *GraphQLStitchedSchemaSpec) Reset() {
	*x = GraphQLStitchedSchemaSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphQLStitchedSchemaSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphQLStitchedSchemaSpec) ProtoMessage() {}

func (x *GraphQLStitchedSchemaSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphQLStitchedSchemaSpec.ProtoReflect.Descriptor instead.
func (*GraphQLStitchedSchemaSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_rawDescGZIP(), []int{0}
}

func (x *GraphQLStitchedSchemaSpec) GetSubschemas() []*GraphQLStitchedSchemaSpec_Subschema {
	if x != nil {
		return x.Subschemas
	}
	return nil
}

func (x *GraphQLStitchedSchemaSpec) GetOptions() *GraphQLStitchedSchemaSpec_Options {
	if x != nil {
		return x.Options
	}
	return nil
}

type GraphQLStitchedSchemaStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The state and workspace conditions of the applied resource.
	Common *v2.Status `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	// The name of the workspace that owns the GraphQL stitched schema.
	OwnedByWorkspace string `protobuf:"bytes,4,opt,name=owned_by_workspace,json=ownedByWorkspace,proto3" json:"owned_by_workspace,omitempty"`
}

func (x *GraphQLStitchedSchemaStatus) Reset() {
	*x = GraphQLStitchedSchemaStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphQLStitchedSchemaStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphQLStitchedSchemaStatus) ProtoMessage() {}

func (x *GraphQLStitchedSchemaStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphQLStitchedSchemaStatus.ProtoReflect.Descriptor instead.
func (*GraphQLStitchedSchemaStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_rawDescGZIP(), []int{1}
}

func (x *GraphQLStitchedSchemaStatus) GetCommon() *v2.Status {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *GraphQLStitchedSchemaStatus) GetOwnedByWorkspace() string {
	if x != nil {
		return x.OwnedByWorkspace
	}
	return ""
}

type GraphQLStitchedSchemaReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspaces map[string]*v2.Report `protobuf:"bytes,1,rep,name=workspaces,proto3" json:"workspaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The name of the workspace that owns the GraphQL stitched schema.
	OwnerWorkspace string `protobuf:"bytes,2,opt,name=owner_workspace,json=ownerWorkspace,proto3" json:"owner_workspace,omitempty"`
}

func (x *GraphQLStitchedSchemaReport) Reset() {
	*x = GraphQLStitchedSchemaReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphQLStitchedSchemaReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphQLStitchedSchemaReport) ProtoMessage() {}

func (x *GraphQLStitchedSchemaReport) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphQLStitchedSchemaReport.ProtoReflect.Descriptor instead.
func (*GraphQLStitchedSchemaReport) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_rawDescGZIP(), []int{2}
}

func (x *GraphQLStitchedSchemaReport) GetWorkspaces() map[string]*v2.Report {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

func (x *GraphQLStitchedSchemaReport) GetOwnerWorkspace() string {
	if x != nil {
		return x.OwnerWorkspace
	}
	return ""
}

type GraphQLStitchedSchemaSpec_Subschema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to GraphqlSchema:
	//
	//	*GraphQLStitchedSchemaSpec_Subschema_Schema
	//	*GraphQLStitchedSchemaSpec_Subschema_StitchedSchema
	GraphqlSchema isGraphQLStitchedSchemaSpec_Subschema_GraphqlSchema `protobuf_oneof:"graphql_schema"`
	// Type merge configuration for this subschema. Let's say this subschema is a Users service schema
	// and provides the User type (with a query to fetch a user given the username)
	// ```gql
	// type Query {
	// GetUser(username: String): User
	// }
	// type User {
	// username: String
	// firstName: String
	// lastName: String
	// }
	// ```
	// and another subschema, e.g. Reviews schema, may have a partial User type:
	// ```gql
	// type Review {
	// author: User
	// }
	// type User {
	// username: String
	// }
	// ```
	// We want to provide the relevant information from this Users service schema,
	// so that another API that can give us a partial User type (with the username) will then
	// be able to have access to the full user type. With the correct type merging config under the Users subschema, e.g.:
	// ```yaml
	// type_merge:
	// User:
	// selection_set: '{ username }'
	// query_name: 'GetUser'
	// args:
	// username: username
	// ```
	// the stitched schema will now be able to provide the full user type to all types that require it. In this case,
	// we can now get the first name of an author from the Review.author field even though the Reviews schema doesn't
	// provide the full User type. Note: Type merging can be used for all GraphQL types except Mutations and Subscriptions.
	TypeMerge map[string]*GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig `protobuf:"bytes,3,rep,name=type_merge,json=typeMerge,proto3" json:"type_merge,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GraphQLStitchedSchemaSpec_Subschema) Reset() {
	*x = GraphQLStitchedSchemaSpec_Subschema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphQLStitchedSchemaSpec_Subschema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphQLStitchedSchemaSpec_Subschema) ProtoMessage() {}

func (x *GraphQLStitchedSchemaSpec_Subschema) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphQLStitchedSchemaSpec_Subschema.ProtoReflect.Descriptor instead.
func (*GraphQLStitchedSchemaSpec_Subschema) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_rawDescGZIP(), []int{0, 0}
}

func (m *GraphQLStitchedSchemaSpec_Subschema) GetGraphqlSchema() isGraphQLStitchedSchemaSpec_Subschema_GraphqlSchema {
	if m != nil {
		return m.GraphqlSchema
	}
	return nil
}

func (x *GraphQLStitchedSchemaSpec_Subschema) GetSchema() *v1.ClusterObjectRef {
	if x, ok := x.GetGraphqlSchema().(*GraphQLStitchedSchemaSpec_Subschema_Schema); ok {
		return x.Schema
	}
	return nil
}

func (x *GraphQLStitchedSchemaSpec_Subschema) GetStitchedSchema() *v1.ClusterObjectRef {
	if x, ok := x.GetGraphqlSchema().(*GraphQLStitchedSchemaSpec_Subschema_StitchedSchema); ok {
		return x.StitchedSchema
	}
	return nil
}

func (x *GraphQLStitchedSchemaSpec_Subschema) GetTypeMerge() map[string]*GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig {
	if x != nil {
		return x.TypeMerge
	}
	return nil
}

type isGraphQLStitchedSchemaSpec_Subschema_GraphqlSchema interface {
	isGraphQLStitchedSchemaSpec_Subschema_GraphqlSchema()
}

type GraphQLStitchedSchemaSpec_Subschema_Schema struct {
	// Reference to a GraphQLSchema resource that contains the configuration for this subschema.
	Schema *v1.ClusterObjectRef `protobuf:"bytes,1,opt,name=schema,proto3,oneof"`
}

type GraphQLStitchedSchemaSpec_Subschema_StitchedSchema struct {
	// Reference to a GraphQLStitchedSchema resource that contains the configuration for this subschema.
	StitchedSchema *v1.ClusterObjectRef `protobuf:"bytes,2,opt,name=stitched_schema,json=stitchedSchema,proto3,oneof"`
}

func (*GraphQLStitchedSchemaSpec_Subschema_Schema) isGraphQLStitchedSchemaSpec_Subschema_GraphqlSchema() {
}

func (*GraphQLStitchedSchemaSpec_Subschema_StitchedSchema) isGraphQLStitchedSchemaSpec_Subschema_GraphqlSchema() {
}

type GraphQLStitchedSchemaSpec_Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Enable introspection queries on this GraphQL Stitched API.
	// Introspection queries are used by GraphQL developers to understand the schema of the GraphQL API and create
	// queries that are valid against the schema. Introspection is disabled by default, and should be disabled for
	// production environments.
	EnableIntrospection bool `protobuf:"varint,1,opt,name=enable_introspection,json=enableIntrospection,proto3" json:"enable_introspection,omitempty"`
}

func (x *GraphQLStitchedSchemaSpec_Options) Reset() {
	*x = GraphQLStitchedSchemaSpec_Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphQLStitchedSchemaSpec_Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphQLStitchedSchemaSpec_Options) ProtoMessage() {}

func (x *GraphQLStitchedSchemaSpec_Options) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphQLStitchedSchemaSpec_Options.ProtoReflect.Descriptor instead.
func (*GraphQLStitchedSchemaSpec_Options) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_rawDescGZIP(), []int{0, 1}
}

func (x *GraphQLStitchedSchemaSpec_Options) GetEnableIntrospection() bool {
	if x != nil {
		return x.EnableIntrospection
	}
	return false
}

type GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This specifies one or more key fields required from other services to perform this query.
	// Query planning will automatically resolve these fields from other subschemas in dependency order.
	// This is a graphql selection set specified as a string
	// e.g. '{ username }'
	SelectionSet string `protobuf:"bytes,1,opt,name=selection_set,json=selectionSet,proto3" json:"selection_set,omitempty"`
	// specifies the root field from this subschema used to request the local type
	QueryName string            `protobuf:"bytes,2,opt,name=query_name,json=queryName,proto3" json:"query_name,omitempty"`
	Args      map[string]string `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig) Reset() {
	*x = GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig) ProtoMessage() {}

func (x *GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig.ProtoReflect.Descriptor instead.
func (*GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig) GetSelectionSet() string {
	if x != nil {
		return x.SelectionSet
	}
	return ""
}

func (x *GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig) GetQueryName() string {
	if x != nil {
		return x.QueryName
	}
	return ""
}

func (x *GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig) GetArgs() map[string]string {
	if x != nil {
		return x.Args
	}
	return nil
}

var File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_rawDesc = []byte{
	0x0a, 0x67, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x71, 0x6c, 0x5f, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x61, 0x70, 0x69, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x1f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x75, 0x65, 0x2f, 0x63, 0x75, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x64, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72,
	0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc1, 0x07, 0x0a, 0x19, 0x47, 0x72, 0x61, 0x70, 0x68, 0x51, 0x4c, 0x53, 0x74, 0x69,
	0x74, 0x63, 0x68, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x5f, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x51, 0x4c, 0x53, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x64,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73,
	0x12, 0x57, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3d, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x51, 0x4c, 0x53, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x64, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0xab, 0x05, 0x0a, 0x09, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73,
	0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x48, 0x00, 0x52, 0x06,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x4e, 0x0a, 0x0f, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68,
	0x65, 0x64, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x66, 0x48, 0x00, 0x52, 0x0e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x64,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x6d, 0x0a, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6d,
	0x65, 0x72, 0x67, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x61, 0x70, 0x69,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x51, 0x4c, 0x53,
	0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x53, 0x70, 0x65,
	0x63, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x4d, 0x65, 0x72, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x74, 0x79, 0x70, 0x65,
	0x4d, 0x65, 0x72, 0x67, 0x65, 0x1a, 0x8d, 0x01, 0x0a, 0x0e, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x65,
	0x72, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x65, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4f, 0x2e, 0x61, 0x70, 0x69, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x51, 0x4c, 0x53, 0x74,
	0x69, 0x74, 0x63, 0x68, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x53, 0x70, 0x65, 0x63,
	0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x4d,
	0x65, 0x72, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xfd, 0x01, 0x0a, 0x0f, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x65,
	0x72, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x6d, 0x0a,
	0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x59, 0x2e, 0x61, 0x70,
	0x69, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x51, 0x4c,
	0x53, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x53, 0x70,
	0x65, 0x63, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x72, 0x67,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x1a, 0x37, 0x0a, 0x09,
	0x41, 0x72, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c,
	0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x3c, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x31, 0x0a, 0x14, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x74,
	0x72, 0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x13, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x80, 0x01, 0x0a, 0x1b, 0x47, 0x72, 0x61, 0x70, 0x68, 0x51,
	0x4c, 0x53, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x12, 0x6f, 0x77,
	0x6e, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6f, 0x77, 0x6e, 0x65, 0x64, 0x42, 0x79, 0x57,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x8b, 0x02, 0x0a, 0x1b, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x51, 0x4c, 0x53, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x67, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x61,
	0x70, 0x69, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x51,
	0x4c, 0x53, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x1a, 0x5a, 0x0a, 0x0f, 0x57, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x5b, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73,
	0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x32, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0,
	0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_rawDescData = file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_goTypes = []interface{}{
	(*GraphQLStitchedSchemaSpec)(nil),           // 0: apimanagement.gloo.solo.io.GraphQLStitchedSchemaSpec
	(*GraphQLStitchedSchemaStatus)(nil),         // 1: apimanagement.gloo.solo.io.GraphQLStitchedSchemaStatus
	(*GraphQLStitchedSchemaReport)(nil),         // 2: apimanagement.gloo.solo.io.GraphQLStitchedSchemaReport
	(*GraphQLStitchedSchemaSpec_Subschema)(nil), // 3: apimanagement.gloo.solo.io.GraphQLStitchedSchemaSpec.Subschema
	(*GraphQLStitchedSchemaSpec_Options)(nil),   // 4: apimanagement.gloo.solo.io.GraphQLStitchedSchemaSpec.Options
	nil, // 5: apimanagement.gloo.solo.io.GraphQLStitchedSchemaSpec.Subschema.TypeMergeEntry
	(*GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig)(nil), // 6: apimanagement.gloo.solo.io.GraphQLStitchedSchemaSpec.Subschema.TypeMergeConfig
	nil,                         // 7: apimanagement.gloo.solo.io.GraphQLStitchedSchemaSpec.Subschema.TypeMergeConfig.ArgsEntry
	nil,                         // 8: apimanagement.gloo.solo.io.GraphQLStitchedSchemaReport.WorkspacesEntry
	(*v2.Status)(nil),           // 9: common.gloo.solo.io.Status
	(*v1.ClusterObjectRef)(nil), // 10: core.skv2.solo.io.ClusterObjectRef
	(*v2.Report)(nil),           // 11: common.gloo.solo.io.Report
}
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_depIdxs = []int32{
	3,  // 0: apimanagement.gloo.solo.io.GraphQLStitchedSchemaSpec.subschemas:type_name -> apimanagement.gloo.solo.io.GraphQLStitchedSchemaSpec.Subschema
	4,  // 1: apimanagement.gloo.solo.io.GraphQLStitchedSchemaSpec.options:type_name -> apimanagement.gloo.solo.io.GraphQLStitchedSchemaSpec.Options
	9,  // 2: apimanagement.gloo.solo.io.GraphQLStitchedSchemaStatus.common:type_name -> common.gloo.solo.io.Status
	8,  // 3: apimanagement.gloo.solo.io.GraphQLStitchedSchemaReport.workspaces:type_name -> apimanagement.gloo.solo.io.GraphQLStitchedSchemaReport.WorkspacesEntry
	10, // 4: apimanagement.gloo.solo.io.GraphQLStitchedSchemaSpec.Subschema.schema:type_name -> core.skv2.solo.io.ClusterObjectRef
	10, // 5: apimanagement.gloo.solo.io.GraphQLStitchedSchemaSpec.Subschema.stitched_schema:type_name -> core.skv2.solo.io.ClusterObjectRef
	5,  // 6: apimanagement.gloo.solo.io.GraphQLStitchedSchemaSpec.Subschema.type_merge:type_name -> apimanagement.gloo.solo.io.GraphQLStitchedSchemaSpec.Subschema.TypeMergeEntry
	6,  // 7: apimanagement.gloo.solo.io.GraphQLStitchedSchemaSpec.Subschema.TypeMergeEntry.value:type_name -> apimanagement.gloo.solo.io.GraphQLStitchedSchemaSpec.Subschema.TypeMergeConfig
	7,  // 8: apimanagement.gloo.solo.io.GraphQLStitchedSchemaSpec.Subschema.TypeMergeConfig.args:type_name -> apimanagement.gloo.solo.io.GraphQLStitchedSchemaSpec.Subschema.TypeMergeConfig.ArgsEntry
	11, // 9: apimanagement.gloo.solo.io.GraphQLStitchedSchemaReport.WorkspacesEntry.value:type_name -> common.gloo.solo.io.Report
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_init()
}
func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_init() {
	if File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto != nil {
		return
	}
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_resolver_map_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphQLStitchedSchemaSpec); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphQLStitchedSchemaStatus); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphQLStitchedSchemaReport); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphQLStitchedSchemaSpec_Subschema); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphQLStitchedSchemaSpec_Options); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig); i {
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
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*GraphQLStitchedSchemaSpec_Subschema_Schema)(nil),
		(*GraphQLStitchedSchemaSpec_Subschema_StitchedSchema)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto = out.File
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_graphql_stitched_schema_proto_depIdxs = nil
}
