// The ApiDoc resource represents the schema of an API served by a Destination (Service, ExternalService).
// ApiDocs are typically created by Gloo Platform discovery running on the agent
// in registered workload clusters, but you can also manually create an `ApiDoc` CR.
// The ApiDoc type is used to represent different types of API schema specification languages:
// - OpenAPI
// - gRPC
// - GraphQL

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

// The destinations that serve this API, if any.
type ServedBy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specify the backing destination for your app, by label or by name.
	// This destination matches the destinations that you later route to.
	// Supported destinations are Kubernetes services, Gloo virtual destinations, and
	// Gloo external services.
	// This field is required when you manually create an ApiDoc for a service that
	// serves an OpenAPI or gRPC schema.
	DestinationSelector *v2.DestinationSelector `protobuf:"bytes,1,opt,name=destination_selector,json=destinationSelector,proto3" json:"destination_selector,omitempty"`
}

func (x *ServedBy) Reset() {
	*x = ServedBy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServedBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServedBy) ProtoMessage() {}

func (x *ServedBy) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ServedBy.ProtoReflect.Descriptor instead.
func (*ServedBy) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDescGZIP(), []int{0}
}

func (x *ServedBy) GetDestinationSelector() *v2.DestinationSelector {
	if x != nil {
		return x.DestinationSelector
	}
	return nil
}

// Specifications for the ApiDoc.
type ApiDocSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The API schema specification language.
	//
	// Types that are assignable to SchemaType:
	//
	//	*ApiDocSpec_Openapi
	//	*ApiDocSpec_Grpc
	//	*ApiDocSpec_Graphql
	SchemaType isApiDocSpec_SchemaType `protobuf_oneof:"schema_type"`
	// The destinations that serve this API, if any.
	ServedBy []*ServedBy `protobuf:"bytes,4,rep,name=served_by,json=servedBy,proto3" json:"served_by,omitempty"`
}

func (x *ApiDocSpec) Reset() {
	*x = ApiDocSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiDocSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiDocSpec) ProtoMessage() {}

func (x *ApiDocSpec) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ApiDocSpec.ProtoReflect.Descriptor instead.
func (*ApiDocSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDescGZIP(), []int{1}
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

func (x *ApiDocSpec) GetServedBy() []*ServedBy {
	if x != nil {
		return x.ServedBy
	}
	return nil
}

type isApiDocSpec_SchemaType interface {
	isApiDocSpec_SchemaType()
}

type ApiDocSpec_Openapi struct {
	// The OpenAPI schema specification language. Specify only one schema type.
	Openapi *ApiDocSpec_OpenAPISchema `protobuf:"bytes,1,opt,name=openapi,proto3,oneof"`
}

type ApiDocSpec_Grpc struct {
	// The gRPC schema specification language. Specify only one schema type.
	Grpc *ApiDocSpec_GrpcSchema `protobuf:"bytes,2,opt,name=grpc,proto3,oneof"`
}

type ApiDocSpec_Graphql struct {
	// The graphQL schema specification language. Specify only one schema type.
	Graphql *ApiDocSpec_GraphQLSchema `protobuf:"bytes,3,opt,name=graphql,proto3,oneof"`
}

func (*ApiDocSpec_Openapi) isApiDocSpec_SchemaType() {}

func (*ApiDocSpec_Grpc) isApiDocSpec_SchemaType() {}

func (*ApiDocSpec_Graphql) isApiDocSpec_SchemaType() {}

// The status of the ApiDoc after it is applied to your Gloo environment.
type ApiDocStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The state and workspace conditions of the applied resource.
	Common *v2.Status `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	// The name of workspace that owns the APIDoc.
	OwnerWorkspace string `protobuf:"bytes,2,opt,name=owner_workspace,json=ownerWorkspace,proto3" json:"owner_workspace,omitempty"`
	// Count of the destinations serving the API.
	SelectedServingDestinations uint32 `protobuf:"varint,3,opt,name=selected_serving_destinations,json=selectedServingDestinations,proto3" json:"selected_serving_destinations,omitempty"`
}

func (x *ApiDocStatus) Reset() {
	*x = ApiDocStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiDocStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiDocStatus) ProtoMessage() {}

func (x *ApiDocStatus) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ApiDocStatus.ProtoReflect.Descriptor instead.
func (*ApiDocStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDescGZIP(), []int{2}
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

// The report shows the resources that the ApiDoc selects after the it is successfully applied.
type ApiDocReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of workspaces in which the ApiDoc can be applied.
	Workspaces map[string]*v2.Report `protobuf:"bytes,1,rep,name=workspaces,proto3" json:"workspaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The name of the workspace that owns the Graqphql API schema.
	OwnerWorkspace string `protobuf:"bytes,2,opt,name=owner_workspace,json=ownerWorkspace,proto3" json:"owner_workspace,omitempty"`
	// A list of destinations serving the API.
	ServingDestinations []*v2.DestinationReference `protobuf:"bytes,3,rep,name=serving_destinations,json=servingDestinations,proto3" json:"serving_destinations,omitempty"`
}

func (x *ApiDocReport) Reset() {
	*x = ApiDocReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiDocReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiDocReport) ProtoMessage() {}

func (x *ApiDocReport) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ApiDocReport.ProtoReflect.Descriptor instead.
func (*ApiDocReport) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDescGZIP(), []int{3}
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

// The YAML- or JSON-formatted OpenAPI v2 or v3 schema string to use for your API.
//
// ## Example
// For detailed information about the settings in this example, see
// [Create your APIs](https://docs.solo.io/gloo-mesh-gateway/2.4.x/portal/guides/apis/apidocs/) in the Gloo Portal documentation.
// ```yaml
// apiVersion: apimanagement.gloo.solo.io/v2
// kind: ApiDoc
// metadata:
//
//	annotations:
//	  cluster.solo.io/cluster: ""
//	name: customers-api-schema
//	namespace: default
//
// spec:
//
//	openapi:
//	  inlineString: '{ "info": { "title": "Gloo Platform Portal API", "version": "1.0.0", "description": "Review the following reference documentation for the Gloo Platform portal APIs. Use these endpoints to manage user access to both the developer portal and the API resources exposed by the portal." }, "openapi": "3.0.0", "servers": [ { "url": "https://api.gloo-platform-portal.com/v1" } ], "paths": { "/login": { "get": { "description": "Logs user into the developer portal. This is the path that should be used as the callbackPath in the ExtAuthPolicy's OIDC configuration.", "operationId": "login", "security": [ { "identityToken": [ ] } ], "responses": { "200": { "description": "Successfully logged in" } }, "summary": "Logs user into the developer portal", "tags": [ "User" ] } }'
//	servedBy:
//	- destinationSelector:
//	    port:
//	      number: 8080
//	    selector:
//	      cluster: $CLUSTER_NAME
//	      name: app
//	      namespace: app
//
// ```
type ApiDocSpec_OpenAPISchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The inline, YAML- or JSON-formatted, OpenAPI v2 or v3 schema.
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
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ApiDocSpec_OpenAPISchema) GetInlineString() string {
	if x != nil {
		return x.InlineString
	}
	return ""
}

// A complete gRPC schema describing the API.
//
// ## Example
// In this gRPC example for a basic user service app,
// the base64-encoded descriptor includes a set of fields that are defined
// for various queries, such as `UserSearch` and `UserByCountry`. For detailed
// information about the settings in this example, see
// [gRPC schema](https://docs.solo.io/gloo-mesh-gateway/2.4.x/graphql/resolvers/resolved/resolver_grpc/)
// in the GraphQL integration documentation.
// ```yaml
// {{% readfile file="static/content/examples/generated/int/graphql_routes/cluster-1/api-doc_bookinfo_grpc-schema.yaml" %}}
// ```
type ApiDocSpec_GrpcSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Protobuf descriptors that represent the gRPC services provided by your API, encoded in base64.
	// For more information, see the
	// [protobuf reference for `FileDescriptorSet`](https://developers.google.com/protocol-buffers/docs/reference/java/com/google/protobuf/DescriptorProtos.FileDescriptorSet).
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
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDescGZIP(), []int{1, 1}
}

func (x *ApiDocSpec_GrpcSchema) GetDescriptors() []byte {
	if x != nil {
		return x.Descriptors
	}
	return nil
}

// Provide a schema definition in GraphQL SDL format.
// The GraphQL schema also has logging options for logging sensitive
// request-related information, and schema extension configuration such as
// custom type definitions. For more information about the different schema features,
// see the [GraphQL documentation](https://graphql.org/learn/schema/).
//
// ## Example
// In this GraphQL example for the Bookinfo sample app, a query type and object
// types are defined. For detailed information about the settings in this example, see
// [Example GraphQL ApiDoc](https://docs.solo.io/gloo-mesh-gateway/2.4.x/graphql/apidoc/#example-graphql-apidoc)
// in the GraphQL integration documentation.
// ```yaml
// {{% readfile file="static/content/examples/generated/int/graphql_proxied_introspection/cluster-1/api-doc_bookinfo_music-schema.yaml" %}}
// ```
type ApiDocSpec_GraphQLSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required: The GraphQL schema definition. Root-level query and mutation
	// types are supported, and you must define at least a query type.
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
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_rawDescGZIP(), []int{1, 2}
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
	0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x57, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72,
	0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x61,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73,
	0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a,
	0x08, 0x53, 0x65, 0x72, 0x76, 0x65, 0x64, 0x42, 0x79, 0x12, 0x5b, 0x0a, 0x14, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x13, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0xef, 0x03, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x44, 0x6f,
	0x63, 0x53, 0x70, 0x65, 0x63, 0x12, 0x50, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x44, 0x6f, 0x63, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x4f,
	0x70, 0x65, 0x6e, 0x41, 0x50, 0x49, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x48, 0x00, 0x52, 0x07,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x12, 0x47, 0x0a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x44, 0x6f, 0x63, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x47, 0x72,
	0x70, 0x63, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x48, 0x00, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63,
	0x12, 0x50, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41,
	0x70, 0x69, 0x44, 0x6f, 0x63, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x51,
	0x4c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x48, 0x00, 0x52, 0x07, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x71, 0x6c, 0x12, 0x41, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x64, 0x42, 0x79, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x64, 0x42, 0x79, 0x1a, 0x34, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x50, 0x49,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x2e, 0x0a, 0x0a, 0x47,
	0x72, 0x70, 0x63, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x1a, 0x3c, 0x0a, 0x0d, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x51, 0x4c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x2b, 0x0a, 0x11,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x44,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xb0, 0x01, 0x0a, 0x0c, 0x41, 0x70, 0x69,
	0x44, 0x6f, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x27,
	0x0a, 0x0f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x57, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x1d, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1b,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xcb, 0x02, 0x0a, 0x0c,
	0x41, 0x70, 0x69, 0x44, 0x6f, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x58, 0x0a, 0x0a,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x38, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x70,
	0x69, 0x44, 0x6f, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x5c, 0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x5a, 0x0a,
	0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x5b, 0x5a, 0x4d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70,
	0x72, 0x69, 0x73, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x32, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5,
	0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(*ServedBy)(nil),                 // 0: apimanagement.gloo.solo.io.ServedBy
	(*ApiDocSpec)(nil),               // 1: apimanagement.gloo.solo.io.ApiDocSpec
	(*ApiDocStatus)(nil),             // 2: apimanagement.gloo.solo.io.ApiDocStatus
	(*ApiDocReport)(nil),             // 3: apimanagement.gloo.solo.io.ApiDocReport
	(*ApiDocSpec_OpenAPISchema)(nil), // 4: apimanagement.gloo.solo.io.ApiDocSpec.OpenAPISchema
	(*ApiDocSpec_GrpcSchema)(nil),    // 5: apimanagement.gloo.solo.io.ApiDocSpec.GrpcSchema
	(*ApiDocSpec_GraphQLSchema)(nil), // 6: apimanagement.gloo.solo.io.ApiDocSpec.GraphQLSchema
	nil,                              // 7: apimanagement.gloo.solo.io.ApiDocReport.WorkspacesEntry
	(*v2.DestinationSelector)(nil),   // 8: common.gloo.solo.io.DestinationSelector
	(*v2.Status)(nil),                // 9: common.gloo.solo.io.Status
	(*v2.DestinationReference)(nil),  // 10: common.gloo.solo.io.DestinationReference
	(*v2.Report)(nil),                // 11: common.gloo.solo.io.Report
}
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_depIdxs = []int32{
	8,  // 0: apimanagement.gloo.solo.io.ServedBy.destination_selector:type_name -> common.gloo.solo.io.DestinationSelector
	4,  // 1: apimanagement.gloo.solo.io.ApiDocSpec.openapi:type_name -> apimanagement.gloo.solo.io.ApiDocSpec.OpenAPISchema
	5,  // 2: apimanagement.gloo.solo.io.ApiDocSpec.grpc:type_name -> apimanagement.gloo.solo.io.ApiDocSpec.GrpcSchema
	6,  // 3: apimanagement.gloo.solo.io.ApiDocSpec.graphql:type_name -> apimanagement.gloo.solo.io.ApiDocSpec.GraphQLSchema
	0,  // 4: apimanagement.gloo.solo.io.ApiDocSpec.served_by:type_name -> apimanagement.gloo.solo.io.ServedBy
	9,  // 5: apimanagement.gloo.solo.io.ApiDocStatus.common:type_name -> common.gloo.solo.io.Status
	7,  // 6: apimanagement.gloo.solo.io.ApiDocReport.workspaces:type_name -> apimanagement.gloo.solo.io.ApiDocReport.WorkspacesEntry
	10, // 7: apimanagement.gloo.solo.io.ApiDocReport.serving_destinations:type_name -> common.gloo.solo.io.DestinationReference
	11, // 8: apimanagement.gloo.solo.io.ApiDocReport.WorkspacesEntry.value:type_name -> common.gloo.solo.io.Report
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
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
			switch v := v.(*ServedBy); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_apimanagement_v2_api_doc_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ApiDocSpec_Openapi)(nil),
		(*ApiDocSpec_Grpc)(nil),
		(*ApiDocSpec_Graphql)(nil),
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
