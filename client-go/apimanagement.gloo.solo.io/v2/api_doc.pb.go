// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/apimanagement/v2/api_doc.proto

package v2

import (
	reflect "reflect"
	sync "sync"

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

// The ApiDoc resource represents an the schema of an API
// served by a Destination (Service, ExternalService).
// ApiDocs are typically created by Gloo Mesh Discovery running on the agent.
// They can also be created manually by users.
// The ApiDoc type is used to represent different types of API schema specification languages:
// - OpenAPI
// - gRPC
// - GraphQL
type ApiDocSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the type of schema this resource contains
	//
	// Types that are assignable to SchemaType:
	//
	//	*ApiDocSpec_Openapi
	//	*ApiDocSpec_Grpc
	//	*ApiDocSpec_Graphql
	SchemaType isApiDocSpec_SchemaType `protobuf_oneof:"schema_type"`
	// the destinations that serve this API, if any.
	// When manually creating an ApiDoc for a service that serves an OpenAPI or gRPC schema the destination selector
	// field is required.
	// A stitched OpenAPI schema for a set of routes exposed by a route table
	// will have a reference to its corresponding route table.
	ServedBy []*ApiDocSpec_ServedBy `protobuf:"bytes,4,rep,name=served_by,json=servedBy,proto3" json:"served_by,omitempty"`
}

func (x *ApiDocSpec) Reset() {
	*x = ApiDocSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiDocSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiDocSpec) ProtoMessage() {}

func (x *ApiDocSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiDocSpec.ProtoReflect.Descriptor instead.
func (*ApiDocSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDescGZIP(), []int{0}
}

func (m *ApiDocSpec) GetSchemaType() isApiDocSpec_SchemaType {
	if m != nil {
		return m.SchemaType
	}
	return nil
}

func (x *ApiDocSpec) GetOpenapi() *ApiDocSpec_OpenAPISchema {
	if x, ok := x.GetSchemaType().(*ApiDocSpec_Openapi); ok {
		return x.Openapi
	}
	return nil
}

func (x *ApiDocSpec) GetGrpc() *ApiDocSpec_GrpcSchema {
	if x, ok := x.GetSchemaType().(*ApiDocSpec_Grpc); ok {
		return x.Grpc
	}
	return nil
}

func (x *ApiDocSpec) GetGraphql() *ApiDocSpec_GraphQLSchema {
	if x, ok := x.GetSchemaType().(*ApiDocSpec_Graphql); ok {
		return x.Graphql
	}
	return nil
}

func (x *ApiDocSpec) GetServedBy() []*ApiDocSpec_ServedBy {
	if x != nil {
		return x.ServedBy
	}
	return nil
}

type isApiDocSpec_SchemaType interface {
	isApiDocSpec_SchemaType()
}

type ApiDocSpec_Openapi struct {
	// OpenAPI schema
	Openapi *ApiDocSpec_OpenAPISchema `protobuf:"bytes,1,opt,name=openapi,proto3,oneof"`
}

type ApiDocSpec_Grpc struct {
	// gRPC schema
	Grpc *ApiDocSpec_GrpcSchema `protobuf:"bytes,2,opt,name=grpc,proto3,oneof"`
}

type ApiDocSpec_Graphql struct {
	// GraphQL schema
	Graphql *ApiDocSpec_GraphQLSchema `protobuf:"bytes,3,opt,name=graphql,proto3,oneof"`
}

func (*ApiDocSpec_Openapi) isApiDocSpec_SchemaType() {}

func (*ApiDocSpec_Grpc) isApiDocSpec_SchemaType() {}

func (*ApiDocSpec_Graphql) isApiDocSpec_SchemaType() {}

type ApiDocStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Common *v2.Status `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	// The name of workspace that owns the GraphQL API schema.
	OwnerWorkspace string `protobuf:"bytes,2,opt,name=owner_workspace,json=ownerWorkspace,proto3" json:"owner_workspace,omitempty"`
	// Count of the destinations serving the API.
	SelectedServingDestinations uint32 `protobuf:"varint,3,opt,name=selected_serving_destinations,json=selectedServingDestinations,proto3" json:"selected_serving_destinations,omitempty"`
}

func (x *ApiDocStatus) Reset() {
	*x = ApiDocStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiDocStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiDocStatus) ProtoMessage() {}

func (x *ApiDocStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiDocStatus.ProtoReflect.Descriptor instead.
func (*ApiDocStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDescGZIP(), []int{1}
}

func (x *ApiDocStatus) GetCommon() *v2.Status {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *ApiDocStatus) GetOwnerWorkspace() string {
	if x != nil {
		return x.OwnerWorkspace
	}
	return ""
}

func (x *ApiDocStatus) GetSelectedServingDestinations() uint32 {
	if x != nil {
		return x.SelectedServingDestinations
	}
	return 0
}

type ApiDocReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspaces map[string]*v2.Report `protobuf:"bytes,1,rep,name=workspaces,proto3" json:"workspaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The name of the workspace that owns the Graqphql API schema.
	OwnerWorkspace string `protobuf:"bytes,2,opt,name=owner_workspace,json=ownerWorkspace,proto3" json:"owner_workspace,omitempty"`
	// A list of destinations serving the API.
	ServingDestinations []*v2.DestinationReference `protobuf:"bytes,3,rep,name=serving_destinations,json=servingDestinations,proto3" json:"serving_destinations,omitempty"`
}

func (x *ApiDocReport) Reset() {
	*x = ApiDocReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiDocReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiDocReport) ProtoMessage() {}

func (x *ApiDocReport) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiDocReport.ProtoReflect.Descriptor instead.
func (*ApiDocReport) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDescGZIP(), []int{2}
}

func (x *ApiDocReport) GetWorkspaces() map[string]*v2.Report {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

func (x *ApiDocReport) GetOwnerWorkspace() string {
	if x != nil {
		return x.OwnerWorkspace
	}
	return ""
}

func (x *ApiDocReport) GetServingDestinations() []*v2.DestinationReference {
	if x != nil {
		return x.ServingDestinations
	}
	return nil
}

type ApiDocSpec_ServedBy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ServedByType:
	//
	//	*ApiDocSpec_ServedBy_DestinationSelector
	//	*ApiDocSpec_ServedBy_RouteTable
	ServedByType isApiDocSpec_ServedBy_ServedByType `protobuf_oneof:"served_by_type"`
}

func (x *ApiDocSpec_ServedBy) Reset() {
	*x = ApiDocSpec_ServedBy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiDocSpec_ServedBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiDocSpec_ServedBy) ProtoMessage() {}

func (x *ApiDocSpec_ServedBy) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiDocSpec_ServedBy.ProtoReflect.Descriptor instead.
func (*ApiDocSpec_ServedBy) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDescGZIP(), []int{0, 0}
}

func (m *ApiDocSpec_ServedBy) GetServedByType() isApiDocSpec_ServedBy_ServedByType {
	if m != nil {
		return m.ServedByType
	}
	return nil
}

func (x *ApiDocSpec_ServedBy) GetDestinationSelector() *v2.DestinationSelector {
	if x, ok := x.GetServedByType().(*ApiDocSpec_ServedBy_DestinationSelector); ok {
		return x.DestinationSelector
	}
	return nil
}

func (x *ApiDocSpec_ServedBy) GetRouteTable() *v2.ObjectReference {
	if x, ok := x.GetServedByType().(*ApiDocSpec_ServedBy_RouteTable); ok {
		return x.RouteTable
	}
	return nil
}

type isApiDocSpec_ServedBy_ServedByType interface {
	isApiDocSpec_ServedBy_ServedByType()
}

type ApiDocSpec_ServedBy_DestinationSelector struct {
	DestinationSelector *v2.DestinationSelector `protobuf:"bytes,1,opt,name=destination_selector,json=destinationSelector,proto3,oneof"`
}

type ApiDocSpec_ServedBy_RouteTable struct {
	RouteTable *v2.ObjectReference `protobuf:"bytes,2,opt,name=route_table,json=routeTable,proto3,oneof"`
}

func (*ApiDocSpec_ServedBy_DestinationSelector) isApiDocSpec_ServedBy_ServedByType() {}

func (*ApiDocSpec_ServedBy_RouteTable) isApiDocSpec_ServedBy_ServedByType() {}

// a complete openapi schema describing the API
type ApiDocSpec_OpenAPISchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// inline string containing the OpenAPI schema
	InlineString string `protobuf:"bytes,1,opt,name=inline_string,json=inlineString,proto3" json:"inline_string,omitempty"`
}

func (x *ApiDocSpec_OpenAPISchema) Reset() {
	*x = ApiDocSpec_OpenAPISchema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiDocSpec_OpenAPISchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiDocSpec_OpenAPISchema) ProtoMessage() {}

func (x *ApiDocSpec_OpenAPISchema) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiDocSpec_OpenAPISchema.ProtoReflect.Descriptor instead.
func (*ApiDocSpec_OpenAPISchema) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ApiDocSpec_OpenAPISchema) GetInlineString() string {
	if x != nil {
		return x.InlineString
	}
	return ""
}

// a complete grpc schema describing the API
// @exclude TODO(Api-team)- unimplemented in Gloo Mesh 2.1
type ApiDocSpec_GrpcSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Protobuf Descriptors that represent the gRPC services provided by the API.
	// this is a serialized base64-encoded google.protobuf.FileDescriptorSet
	Descriptors []byte `protobuf:"bytes,1,opt,name=descriptors,proto3" json:"descriptors,omitempty"`
}

func (x *ApiDocSpec_GrpcSchema) Reset() {
	*x = ApiDocSpec_GrpcSchema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiDocSpec_GrpcSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiDocSpec_GrpcSchema) ProtoMessage() {}

func (x *ApiDocSpec_GrpcSchema) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiDocSpec_GrpcSchema.ProtoReflect.Descriptor instead.
func (*ApiDocSpec_GrpcSchema) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDescGZIP(), []int{0, 2}
}

func (x *ApiDocSpec_GrpcSchema) GetDescriptors() []byte {
	if x != nil {
		return x.Descriptors
	}
	return nil
}

// The GraphQLSchema resource is responsible for providing the schema definition
// in GraphQL SDL format. It also has logging options for logging sensitive request-related information.
// And schema extension configuration such as custom type definitions.
type ApiDocSpec_GraphQLSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required: GraphQL Schema Definition. Root-level Query and Mutation types are supported, while the Subscription type is not
	// yet supported.
	SchemaDefinition string `protobuf:"bytes,1,opt,name=schema_definition,json=schemaDefinition,proto3" json:"schema_definition,omitempty"`
}

func (x *ApiDocSpec_GraphQLSchema) Reset() {
	*x = ApiDocSpec_GraphQLSchema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiDocSpec_GraphQLSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiDocSpec_GraphQLSchema) ProtoMessage() {}

func (x *ApiDocSpec_GraphQLSchema) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiDocSpec_GraphQLSchema.ProtoReflect.Descriptor instead.
func (*ApiDocSpec_GraphQLSchema) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDescGZIP(), []int{0, 3}
}

func (x *ApiDocSpec_GraphQLSchema) GetSchemaDefinition() string {
	if x != nil {
		return x.SchemaDefinition
	}
	return ""
}

var File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDesc = []byte{
	0x0a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x5f,
	0x64, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x61, 0x70, 0x69, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x1f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x75, 0x65, 0x2f, 0x63, 0x75, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x53, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72,
	0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73,
	0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65,
	0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61,
	0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x53, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x32, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc1, 0x05, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x44, 0x6f, 0x63, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x50, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x34, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x41, 0x70, 0x69, 0x44, 0x6f, 0x63, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x41,
	0x50, 0x49, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x48, 0x00, 0x52, 0x07, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x12, 0x47, 0x0a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41,
	0x70, 0x69, 0x44, 0x6f, 0x63, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x48, 0x00, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x12, 0x50, 0x0a, 0x07,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e,
	0x61, 0x70, 0x69, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x44, 0x6f,
	0x63, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x51, 0x4c, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x48, 0x00, 0x52, 0x07, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x12, 0x4c,
	0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41,
	0x70, 0x69, 0x44, 0x6f, 0x63, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x64,
	0x42, 0x79, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x42, 0x79, 0x1a, 0xc4, 0x01, 0x0a,
	0x08, 0x53, 0x65, 0x72, 0x76, 0x65, 0x64, 0x42, 0x79, 0x12, 0x5d, 0x0a, 0x14, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x48, 0x00, 0x52, 0x13, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x47, 0x0a, 0x0b, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x42, 0x10, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x1a, 0x34, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x50, 0x49, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x2e, 0x0a, 0x0a, 0x47, 0x72, 0x70,
	0x63, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x1a, 0x3c, 0x0a, 0x0d, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x51, 0x4c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x44, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xb0, 0x01, 0x0a, 0x0c, 0x41, 0x70, 0x69, 0x44, 0x6f,
	0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x1d, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1b, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xcb, 0x02, 0x0a, 0x0c, 0x41, 0x70,
	0x69, 0x44, 0x6f, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x58, 0x0a, 0x0a, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38,
	0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x44,
	0x6f, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x5c, 0x0a,
	0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x5a, 0x0a, 0x0f, 0x57,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x5b, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69,
	0x73, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x32, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01,
	0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDescData = file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_goTypes = []interface{}{
	(*ApiDocSpec)(nil),               // 0: apimanagement.gloo.solo.io.ApiDocSpec
	(*ApiDocStatus)(nil),             // 1: apimanagement.gloo.solo.io.ApiDocStatus
	(*ApiDocReport)(nil),             // 2: apimanagement.gloo.solo.io.ApiDocReport
	(*ApiDocSpec_ServedBy)(nil),      // 3: apimanagement.gloo.solo.io.ApiDocSpec.ServedBy
	(*ApiDocSpec_OpenAPISchema)(nil), // 4: apimanagement.gloo.solo.io.ApiDocSpec.OpenAPISchema
	(*ApiDocSpec_GrpcSchema)(nil),    // 5: apimanagement.gloo.solo.io.ApiDocSpec.GrpcSchema
	(*ApiDocSpec_GraphQLSchema)(nil), // 6: apimanagement.gloo.solo.io.ApiDocSpec.GraphQLSchema
	nil,                              // 7: apimanagement.gloo.solo.io.ApiDocReport.WorkspacesEntry
	(*v2.Status)(nil),                // 8: common.gloo.solo.io.Status
	(*v2.DestinationReference)(nil),  // 9: common.gloo.solo.io.DestinationReference
	(*v2.DestinationSelector)(nil),   // 10: common.gloo.solo.io.DestinationSelector
	(*v2.ObjectReference)(nil),       // 11: common.gloo.solo.io.ObjectReference
	(*v2.Report)(nil),                // 12: common.gloo.solo.io.Report
}
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_depIdxs = []int32{
	4,  // 0: apimanagement.gloo.solo.io.ApiDocSpec.openapi:type_name -> apimanagement.gloo.solo.io.ApiDocSpec.OpenAPISchema
	5,  // 1: apimanagement.gloo.solo.io.ApiDocSpec.grpc:type_name -> apimanagement.gloo.solo.io.ApiDocSpec.GrpcSchema
	6,  // 2: apimanagement.gloo.solo.io.ApiDocSpec.graphql:type_name -> apimanagement.gloo.solo.io.ApiDocSpec.GraphQLSchema
	3,  // 3: apimanagement.gloo.solo.io.ApiDocSpec.served_by:type_name -> apimanagement.gloo.solo.io.ApiDocSpec.ServedBy
	8,  // 4: apimanagement.gloo.solo.io.ApiDocStatus.common:type_name -> common.gloo.solo.io.Status
	7,  // 5: apimanagement.gloo.solo.io.ApiDocReport.workspaces:type_name -> apimanagement.gloo.solo.io.ApiDocReport.WorkspacesEntry
	9,  // 6: apimanagement.gloo.solo.io.ApiDocReport.serving_destinations:type_name -> common.gloo.solo.io.DestinationReference
	10, // 7: apimanagement.gloo.solo.io.ApiDocSpec.ServedBy.destination_selector:type_name -> common.gloo.solo.io.DestinationSelector
	11, // 8: apimanagement.gloo.solo.io.ApiDocSpec.ServedBy.route_table:type_name -> common.gloo.solo.io.ObjectReference
	12, // 9: apimanagement.gloo.solo.io.ApiDocReport.WorkspacesEntry.value:type_name -> common.gloo.solo.io.Report
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_init()
}
func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_init() {
	if File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiDocSpec); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiDocStatus); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiDocReport); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiDocSpec_ServedBy); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiDocSpec_OpenAPISchema); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiDocSpec_GrpcSchema); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiDocSpec_GraphQLSchema); i {
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
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ApiDocSpec_Openapi)(nil),
		(*ApiDocSpec_Grpc)(nil),
		(*ApiDocSpec_Graphql)(nil),
	}
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ApiDocSpec_ServedBy_DestinationSelector)(nil),
		(*ApiDocSpec_ServedBy_RouteTable)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto = out.File
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_depIdxs = nil
}
