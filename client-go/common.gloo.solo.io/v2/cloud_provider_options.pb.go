// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/common/v2/cloud_provider_options.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LambdaOptions_RequestTransformation int32

const (
	// Default value. include headers, querystring, request path, and request method in the event payload sent to aws lambda.
	LambdaOptions_REQUEST_DEFAULT LambdaOptions_RequestTransformation = 0
	// Disable all transformations for the request to AWS Lambda.
	LambdaOptions_REQUEST_DISABLE LambdaOptions_RequestTransformation = 1
)

// Enum value maps for LambdaOptions_RequestTransformation.
var (
	LambdaOptions_RequestTransformation_name = map[int32]string{
		0: "REQUEST_DEFAULT",
		1: "REQUEST_DISABLE",
	}
	LambdaOptions_RequestTransformation_value = map[string]int32{
		"REQUEST_DEFAULT": 0,
		"REQUEST_DISABLE": 1,
	}
)

func (x LambdaOptions_RequestTransformation) Enum() *LambdaOptions_RequestTransformation {
	p := new(LambdaOptions_RequestTransformation)
	*p = x
	return p
}

func (x LambdaOptions_RequestTransformation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LambdaOptions_RequestTransformation) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_enumTypes[0].Descriptor()
}

func (LambdaOptions_RequestTransformation) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_enumTypes[0]
}

func (x LambdaOptions_RequestTransformation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LambdaOptions_RequestTransformation.Descriptor instead.
func (LambdaOptions_RequestTransformation) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescGZIP(), []int{0, 0}
}

type LambdaOptions_ResponseTransformation int32

const (
	// Default value. Unwrap the response as if the proxy was an AWS API Gateway. de-jsonify response bodies returned
	// from aws lambda, sets response status code, and sets response headers from the JSON payload.
	LambdaOptions_RESPONSE_DEFAULT LambdaOptions_ResponseTransformation = 0
	// Disable all transformations for the response from AWS Lambda.
	LambdaOptions_RESPONSE_DISABLE LambdaOptions_ResponseTransformation = 1
)

// Enum value maps for LambdaOptions_ResponseTransformation.
var (
	LambdaOptions_ResponseTransformation_name = map[int32]string{
		0: "RESPONSE_DEFAULT",
		1: "RESPONSE_DISABLE",
	}
	LambdaOptions_ResponseTransformation_value = map[string]int32{
		"RESPONSE_DEFAULT": 0,
		"RESPONSE_DISABLE": 1,
	}
)

func (x LambdaOptions_ResponseTransformation) Enum() *LambdaOptions_ResponseTransformation {
	p := new(LambdaOptions_ResponseTransformation)
	*p = x
	return p
}

func (x LambdaOptions_ResponseTransformation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LambdaOptions_ResponseTransformation) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_enumTypes[1].Descriptor()
}

func (LambdaOptions_ResponseTransformation) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_enumTypes[1]
}

func (x LambdaOptions_ResponseTransformation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LambdaOptions_ResponseTransformation.Descriptor instead.
func (LambdaOptions_ResponseTransformation) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescGZIP(), []int{0, 1}
}

type LambdaOptions_InvocationStyle int32

const (
	LambdaOptions_SYNC  LambdaOptions_InvocationStyle = 0
	LambdaOptions_ASYNC LambdaOptions_InvocationStyle = 1
)

// Enum value maps for LambdaOptions_InvocationStyle.
var (
	LambdaOptions_InvocationStyle_name = map[int32]string{
		0: "SYNC",
		1: "ASYNC",
	}
	LambdaOptions_InvocationStyle_value = map[string]int32{
		"SYNC":  0,
		"ASYNC": 1,
	}
)

func (x LambdaOptions_InvocationStyle) Enum() *LambdaOptions_InvocationStyle {
	p := new(LambdaOptions_InvocationStyle)
	*p = x
	return p
}

func (x LambdaOptions_InvocationStyle) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LambdaOptions_InvocationStyle) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_enumTypes[2].Descriptor()
}

func (LambdaOptions_InvocationStyle) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_enumTypes[2]
}

func (x LambdaOptions_InvocationStyle) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LambdaOptions_InvocationStyle.Descriptor instead.
func (LambdaOptions_InvocationStyle) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescGZIP(), []int{0, 2}
}

type LambdaOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestTransformation  LambdaOptions_RequestTransformation  `protobuf:"varint,1,opt,name=request_transformation,json=requestTransformation,proto3,enum=common.gloo.solo.io.LambdaOptions_RequestTransformation" json:"request_transformation,omitempty"`
	ResponseTransformation LambdaOptions_ResponseTransformation `protobuf:"varint,2,opt,name=response_transformation,json=responseTransformation,proto3,enum=common.gloo.solo.io.LambdaOptions_ResponseTransformation" json:"response_transformation,omitempty"`
	// Can be either Sync or Async. See [AWS Invoke](https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html)
	// for more details.
	InvocationStyle LambdaOptions_InvocationStyle `protobuf:"varint,3,opt,name=invocation_style,json=invocationStyle,proto3,enum=common.gloo.solo.io.LambdaOptions_InvocationStyle" json:"invocation_style,omitempty"`
}

func (x *LambdaOptions) Reset() {
	*x = LambdaOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LambdaOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LambdaOptions) ProtoMessage() {}

func (x *LambdaOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LambdaOptions.ProtoReflect.Descriptor instead.
func (*LambdaOptions) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescGZIP(), []int{0}
}

func (x *LambdaOptions) GetRequestTransformation() LambdaOptions_RequestTransformation {
	if x != nil {
		return x.RequestTransformation
	}
	return LambdaOptions_REQUEST_DEFAULT
}

func (x *LambdaOptions) GetResponseTransformation() LambdaOptions_ResponseTransformation {
	if x != nil {
		return x.ResponseTransformation
	}
	return LambdaOptions_RESPONSE_DEFAULT
}

func (x *LambdaOptions) GetInvocationStyle() LambdaOptions_InvocationStyle {
	if x != nil {
		return x.InvocationStyle
	}
	return LambdaOptions_SYNC
}

type CloudProviderFilterOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ProviderOptions:
	//
	//	*CloudProviderFilterOptions_Aws
	ProviderOptions isCloudProviderFilterOptions_ProviderOptions `protobuf_oneof:"provider_options"`
}

func (x *CloudProviderFilterOptions) Reset() {
	*x = CloudProviderFilterOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudProviderFilterOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudProviderFilterOptions) ProtoMessage() {}

func (x *CloudProviderFilterOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudProviderFilterOptions.ProtoReflect.Descriptor instead.
func (*CloudProviderFilterOptions) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescGZIP(), []int{1}
}

func (m *CloudProviderFilterOptions) GetProviderOptions() isCloudProviderFilterOptions_ProviderOptions {
	if m != nil {
		return m.ProviderOptions
	}
	return nil
}

func (x *CloudProviderFilterOptions) GetAws() *AWSFilterOptions {
	if x, ok := x.GetProviderOptions().(*CloudProviderFilterOptions_Aws); ok {
		return x.Aws
	}
	return nil
}

type isCloudProviderFilterOptions_ProviderOptions interface {
	isCloudProviderFilterOptions_ProviderOptions()
}

type CloudProviderFilterOptions_Aws struct {
	// Optional: filter out route tables that use AWS functionality, if provided. Each evaluated route must match at
	// at least one of every repeated field provided, if not empty.
	Aws *AWSFilterOptions `protobuf:"bytes,1,opt,name=aws,proto3,oneof"`
}

func (*CloudProviderFilterOptions_Aws) isCloudProviderFilterOptions_ProviderOptions() {}

type AWSFilterOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional: filter out route tables that use an AWS account ID which does not match the selector, if provided.
	AccountIds []string `protobuf:"bytes,1,rep,name=account_ids,json=accountIds,proto3" json:"account_ids,omitempty"`
	// Optional: filter out route tables that use IAM invoke roles which do not match the selector, if provided.
	// Regex supported.
	IamRoles []string `protobuf:"bytes,2,rep,name=iam_roles,json=iamRoles,proto3" json:"iam_roles,omitempty"`
	// Optional: filter out route tables that use regions which do not match the selector, if provided.
	Regions []string `protobuf:"bytes,3,rep,name=regions,proto3" json:"regions,omitempty"`
	// Optional: filter out route tables that use backend Lambda functions that do no match the selector,
	// if provided. Regex supported.
	LambdaFunctions []string `protobuf:"bytes,4,rep,name=lambda_functions,json=lambdaFunctions,proto3" json:"lambda_functions,omitempty"`
}

func (x *AWSFilterOptions) Reset() {
	*x = AWSFilterOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSFilterOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSFilterOptions) ProtoMessage() {}

func (x *AWSFilterOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSFilterOptions.ProtoReflect.Descriptor instead.
func (*AWSFilterOptions) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescGZIP(), []int{2}
}

func (x *AWSFilterOptions) GetAccountIds() []string {
	if x != nil {
		return x.AccountIds
	}
	return nil
}

func (x *AWSFilterOptions) GetIamRoles() []string {
	if x != nil {
		return x.IamRoles
	}
	return nil
}

func (x *AWSFilterOptions) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *AWSFilterOptions) GetLambdaFunctions() []string {
	if x != nil {
		return x.LambdaFunctions
	}
	return nil
}

var File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDesc = []byte{
	0x0a, 0x5f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x04, 0x0a, 0x0d, 0x4c,
	0x61, 0x6d, 0x62, 0x64, 0x61, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x6f, 0x0a, 0x16,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x15, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x72, 0x0a,
	0x17, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x16, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x5d, 0x0a, 0x10, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x52,
	0x0f, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x79, 0x6c, 0x65,
	0x22, 0x41, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x13,
	0x0a, 0x0f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x01, 0x22, 0x44, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x10, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c,
	0x54, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f,
	0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x22, 0x26, 0x0a, 0x0f, 0x49, 0x6e, 0x76,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x08, 0x0a, 0x04,
	0x53, 0x59, 0x4e, 0x43, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x53, 0x59, 0x4e, 0x43, 0x10,
	0x01, 0x22, 0x6b, 0x0a, 0x1a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x39, 0x0a, 0x03, 0x61, 0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x41, 0x57, 0x53, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x03, 0x61, 0x77, 0x73, 0x42, 0x12, 0x0a, 0x10, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x95,
	0x01, 0x0a, 0x10, 0x41, 0x57, 0x53, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x61, 0x6d, 0x5f, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x69, 0x61, 0x6d, 0x52, 0x6f, 0x6c, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x6c,
	0x61, 0x6d, 0x62, 0x64, 0x61, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x54, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73,
	0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x32,
	0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescData = file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_goTypes = []interface{}{
	(LambdaOptions_RequestTransformation)(0),  // 0: common.gloo.solo.io.LambdaOptions.RequestTransformation
	(LambdaOptions_ResponseTransformation)(0), // 1: common.gloo.solo.io.LambdaOptions.ResponseTransformation
	(LambdaOptions_InvocationStyle)(0),        // 2: common.gloo.solo.io.LambdaOptions.InvocationStyle
	(*LambdaOptions)(nil),                     // 3: common.gloo.solo.io.LambdaOptions
	(*CloudProviderFilterOptions)(nil),        // 4: common.gloo.solo.io.CloudProviderFilterOptions
	(*AWSFilterOptions)(nil),                  // 5: common.gloo.solo.io.AWSFilterOptions
}
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_depIdxs = []int32{
	0, // 0: common.gloo.solo.io.LambdaOptions.request_transformation:type_name -> common.gloo.solo.io.LambdaOptions.RequestTransformation
	1, // 1: common.gloo.solo.io.LambdaOptions.response_transformation:type_name -> common.gloo.solo.io.LambdaOptions.ResponseTransformation
	2, // 2: common.gloo.solo.io.LambdaOptions.invocation_style:type_name -> common.gloo.solo.io.LambdaOptions.InvocationStyle
	5, // 3: common.gloo.solo.io.CloudProviderFilterOptions.aws:type_name -> common.gloo.solo.io.AWSFilterOptions
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_init()
}
func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_init() {
	if File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LambdaOptions); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudProviderFilterOptions); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AWSFilterOptions); i {
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
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CloudProviderFilterOptions_Aws)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto = out.File
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_depIdxs = nil
}