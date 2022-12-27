// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/common/v2/cloud_provider_options.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

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
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_enumTypes[0].Descriptor()
}

func (LambdaOptions_InvocationStyle) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_enumTypes[0]
}

func (x LambdaOptions_InvocationStyle) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LambdaOptions_InvocationStyle.Descriptor instead.
func (LambdaOptions_InvocationStyle) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescGZIP(), []int{1, 0}
}

// FunctionDestinationSpec specifies the name and options to use when calling a serverless function.
type FunctionDestinationSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The logical name used to call the specified Cloud Provider.
	LogicalName string `protobuf:"bytes,1,opt,name=logical_name,json=logicalName,proto3" json:"logical_name,omitempty"`
	// Provider-specific options that can be used when routing to cloud functions.
	//
	// Types that are assignable to ProviderOptions:
	//
	//	*FunctionDestinationSpec_AwsLambda
	ProviderOptions isFunctionDestinationSpec_ProviderOptions `protobuf_oneof:"provider_options"`
}

func (x *FunctionDestinationSpec) Reset() {
	*x = FunctionDestinationSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionDestinationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionDestinationSpec) ProtoMessage() {}

func (x *FunctionDestinationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionDestinationSpec.ProtoReflect.Descriptor instead.
func (*FunctionDestinationSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescGZIP(), []int{0}
}

func (x *FunctionDestinationSpec) GetLogicalName() string {
	if x != nil {
		return x.LogicalName
	}
	return ""
}

func (m *FunctionDestinationSpec) GetProviderOptions() isFunctionDestinationSpec_ProviderOptions {
	if m != nil {
		return m.ProviderOptions
	}
	return nil
}

func (x *FunctionDestinationSpec) GetAwsLambda() *LambdaOptions {
	if x, ok := x.GetProviderOptions().(*FunctionDestinationSpec_AwsLambda); ok {
		return x.AwsLambda
	}
	return nil
}

type isFunctionDestinationSpec_ProviderOptions interface {
	isFunctionDestinationSpec_ProviderOptions()
}

type FunctionDestinationSpec_AwsLambda struct {
	AwsLambda *LambdaOptions `protobuf:"bytes,2,opt,name=aws_lambda,json=awsLambda,proto3,oneof"`
}

func (*FunctionDestinationSpec_AwsLambda) isFunctionDestinationSpec_ProviderOptions() {}

type LambdaOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unwrap the response as if the proxy was an ALB.
	// Intended to ease migration when previously using alb to invoke Lambdas.
	// For further information see below link for the expected format when true.
	// https://docs.aws.amazon.com/elasticloadbalancing/latest/application/lambda-functions.html
	// Only one of `unwrapAsAlb` or `unwrapAsApiGateway` should be provided.
	// If more than one is provided only one will be checked with priority unwrapAsAlb, unwrapAsApiGateway
	UnwrapAsAlb bool `protobuf:"varint,1,opt,name=unwrapAsAlb,proto3" json:"unwrapAsAlb,omitempty"`
	// Unwrap the response as if the proxy was an AWS API Gateway.
	// Intended to ease migration when previously using API Gateway to invoke Lambdas.
	// Only one of `unwrapAsAlb` or `unwrapAsApiGateway` should be provided.
	// If more than one is provided only one will be checked with priority unwrapAsAlb, unwrapAsApiGateway
	UnwrapAsApiGateway bool `protobuf:"varint,3,opt,name=unwrapAsApiGateway,proto3" json:"unwrapAsApiGateway,omitempty"`
	// Can be either Sync or Async. See [AWS Invoke](https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html)
	// for more details.
	InvocationStyle LambdaOptions_InvocationStyle `protobuf:"varint,2,opt,name=invocation_style,json=invocationStyle,proto3,enum=common.gloo.solo.io.LambdaOptions_InvocationStyle" json:"invocation_style,omitempty"`
}

func (x *LambdaOptions) Reset() {
	*x = LambdaOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LambdaOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LambdaOptions) ProtoMessage() {}

func (x *LambdaOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[1]
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
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescGZIP(), []int{1}
}

func (x *LambdaOptions) GetUnwrapAsAlb() bool {
	if x != nil {
		return x.UnwrapAsAlb
	}
	return false
}

func (x *LambdaOptions) GetUnwrapAsApiGateway() bool {
	if x != nil {
		return x.UnwrapAsApiGateway
	}
	return false
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
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudProviderFilterOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudProviderFilterOptions) ProtoMessage() {}

func (x *CloudProviderFilterOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[2]
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
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescGZIP(), []int{2}
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
	// Optional: filter out route tables that use AWS functionality, if provided
	Aws *AWSFilterOptions `protobuf:"bytes,1,opt,name=aws,proto3,oneof"`
}

func (*CloudProviderFilterOptions_Aws) isCloudProviderFilterOptions_ProviderOptions() {}

type AWSFilterOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional: filter out route tables that use an AWS account ID which does not match the selector, if provided.
	AccountIds []string `protobuf:"bytes,1,rep,name=account_ids,json=accountIds,proto3" json:"account_ids,omitempty"`
	// Optional: filter out route tables that use IAM roles which do not match the selector, if provided.
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
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSFilterOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSFilterOptions) ProtoMessage() {}

func (x *AWSFilterOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[3]
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
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescGZIP(), []int{3}
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

var File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDesc = []byte{
	0x0a, 0x5e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x17, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x6f, 0x67,
	0x69, 0x63, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x61, 0x77, 0x73, 0x5f,
	0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x48, 0x00, 0x52, 0x09, 0x61, 0x77, 0x73, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x42, 0x12, 0x0a,
	0x10, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0xe8, 0x01, 0x0a, 0x0d, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x77, 0x72, 0x61, 0x70, 0x41, 0x73, 0x41,
	0x6c, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x75, 0x6e, 0x77, 0x72, 0x61, 0x70,
	0x41, 0x73, 0x41, 0x6c, 0x62, 0x12, 0x2e, 0x0a, 0x12, 0x75, 0x6e, 0x77, 0x72, 0x61, 0x70, 0x41,
	0x73, 0x41, 0x70, 0x69, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x12, 0x75, 0x6e, 0x77, 0x72, 0x61, 0x70, 0x41, 0x73, 0x41, 0x70, 0x69, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x5d, 0x0a, 0x10, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x79, 0x6c, 0x65, 0x52, 0x0f, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x79, 0x6c, 0x65, 0x22, 0x26, 0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x59, 0x4e, 0x43, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x53, 0x59, 0x4e, 0x43, 0x10, 0x01, 0x22, 0x6b, 0x0a, 0x1a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x39, 0x0a, 0x03, 0x61, 0x77,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x57,
	0x53, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00,
	0x52, 0x03, 0x61, 0x77, 0x73, 0x42, 0x12, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x95, 0x01, 0x0a, 0x10, 0x41, 0x57,
	0x53, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x61, 0x6d, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x61, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61,
	0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0f, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0x49, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x32,
	0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_goTypes = []interface{}{
	(LambdaOptions_InvocationStyle)(0), // 0: common.gloo.solo.io.LambdaOptions.InvocationStyle
	(*FunctionDestinationSpec)(nil),    // 1: common.gloo.solo.io.FunctionDestinationSpec
	(*LambdaOptions)(nil),              // 2: common.gloo.solo.io.LambdaOptions
	(*CloudProviderFilterOptions)(nil), // 3: common.gloo.solo.io.CloudProviderFilterOptions
	(*AWSFilterOptions)(nil),           // 4: common.gloo.solo.io.AWSFilterOptions
}
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_depIdxs = []int32{
	2, // 0: common.gloo.solo.io.FunctionDestinationSpec.aws_lambda:type_name -> common.gloo.solo.io.LambdaOptions
	0, // 1: common.gloo.solo.io.LambdaOptions.invocation_style:type_name -> common.gloo.solo.io.LambdaOptions.InvocationStyle
	4, // 2: common.gloo.solo.io.CloudProviderFilterOptions.aws:type_name -> common.gloo.solo.io.AWSFilterOptions
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionDestinationSpec); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FunctionDestinationSpec_AwsLambda)(nil),
	}
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*CloudProviderFilterOptions_Aws)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_cloud_provider_options_proto_depIdxs = nil
}
