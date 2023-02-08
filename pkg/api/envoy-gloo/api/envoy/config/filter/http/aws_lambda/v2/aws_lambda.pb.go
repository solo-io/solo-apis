// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo-mesh/external/envoy-gloo/api/envoy/config/filter/http/aws_lambda/v2/aws_lambda.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	v3 "github.com/solo-io/solo-apis/api/gloo-mesh/external/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/solo-io/solo-apis/api/gloo-mesh/external/envoyproxy/protoc-gen-validate/validate"
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

// AWS Lambda contains the configuration necessary to perform transform regular
// http calls to AWS Lambda invocations.
type AWSLambdaPerRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the function
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The qualifier of the function (defaults to $LATEST if not specified)
	Qualifier string `protobuf:"bytes,2,opt,name=qualifier,proto3" json:"qualifier,omitempty"`
	// Invocation type - async or regular.
	Async bool `protobuf:"varint,3,opt,name=async,proto3" json:"async,omitempty"`
	// Optional default body if the body is empty. By default on default
	// body is used if the body empty, and an empty body will be sent upstream.
	EmptyBodyOverride *wrappers.StringValue `protobuf:"bytes,4,opt,name=empty_body_override,json=emptyBodyOverride,proto3" json:"empty_body_override,omitempty"`
	// Unwrap responses as AWS ALB does.
	// Expects json lambda responses to construct response.
	// Intended to ease migration when previously using alb to invoke Lambdas.
	// When set on a route the filter will not stream data on the encoding step.
	// Defaults to false.
	UnwrapAsAlb bool `protobuf:"varint,5,opt,name=unwrap_as_alb,json=unwrapAsAlb,proto3" json:"unwrap_as_alb,omitempty"`
	// This is a transformer config, as defined in api.envoy.config.filter.http.transformation.v2
	// used to process response data.
	// This config cannot be configured simultaneously with unwrap_as_alb
	TransformerConfig *v3.TypedExtensionConfig `protobuf:"bytes,6,opt,name=transformer_config,json=transformerConfig,proto3" json:"transformer_config,omitempty"`
}

func (x *AWSLambdaPerRoute) Reset() {
	*x = AWSLambdaPerRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSLambdaPerRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSLambdaPerRoute) ProtoMessage() {}

func (x *AWSLambdaPerRoute) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSLambdaPerRoute.ProtoReflect.Descriptor instead.
func (*AWSLambdaPerRoute) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_rawDescGZIP(), []int{0}
}

func (x *AWSLambdaPerRoute) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AWSLambdaPerRoute) GetQualifier() string {
	if x != nil {
		return x.Qualifier
	}
	return ""
}

func (x *AWSLambdaPerRoute) GetAsync() bool {
	if x != nil {
		return x.Async
	}
	return false
}

func (x *AWSLambdaPerRoute) GetEmptyBodyOverride() *wrappers.StringValue {
	if x != nil {
		return x.EmptyBodyOverride
	}
	return nil
}

func (x *AWSLambdaPerRoute) GetUnwrapAsAlb() bool {
	if x != nil {
		return x.UnwrapAsAlb
	}
	return false
}

func (x *AWSLambdaPerRoute) GetTransformerConfig() *v3.TypedExtensionConfig {
	if x != nil {
		return x.TransformerConfig
	}
	return nil
}

type AWSLambdaProtocolExtension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The host header for AWS this cluster
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// The region for this cluster
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	// The access_key for AWS this cluster
	AccessKey string `protobuf:"bytes,3,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	// The secret_key for AWS this cluster
	SecretKey string `protobuf:"bytes,4,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	// The session_token for AWS this cluster
	SessionToken string `protobuf:"bytes,5,opt,name=session_token,json=sessionToken,proto3" json:"session_token,omitempty"`
	// The role_arn to use when generating credentials for the mounted projected SA token
	RoleArn string `protobuf:"bytes,6,opt,name=role_arn,json=roleArn,proto3" json:"role_arn,omitempty"`
}

func (x *AWSLambdaProtocolExtension) Reset() {
	*x = AWSLambdaProtocolExtension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSLambdaProtocolExtension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSLambdaProtocolExtension) ProtoMessage() {}

func (x *AWSLambdaProtocolExtension) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSLambdaProtocolExtension.ProtoReflect.Descriptor instead.
func (*AWSLambdaProtocolExtension) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_rawDescGZIP(), []int{1}
}

func (x *AWSLambdaProtocolExtension) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *AWSLambdaProtocolExtension) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *AWSLambdaProtocolExtension) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *AWSLambdaProtocolExtension) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *AWSLambdaProtocolExtension) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

func (x *AWSLambdaProtocolExtension) GetRoleArn() string {
	if x != nil {
		return x.RoleArn
	}
	return ""
}

type AWSLambdaConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to CredentialsFetcher:
	//
	//	*AWSLambdaConfig_UseDefaultCredentials
	//	*AWSLambdaConfig_ServiceAccountCredentials_
	CredentialsFetcher isAWSLambdaConfig_CredentialsFetcher `protobuf_oneof:"credentials_fetcher"`
	// Send downstream path and method as `x-envoy-original-path` and
	// `x-envoy-original-method` headers on the request to AWS lambda.
	// Defaults to false.
	PropagateOriginalRouting bool `protobuf:"varint,3,opt,name=propagate_original_routing,json=propagateOriginalRouting,proto3" json:"propagate_original_routing,omitempty"`
	// Sets cadence for refreshing credentials for Service Account.
	// Does nothing if Service account is not set.
	// Does not affect the default filewatch for service account only augments it.
	// Defaults to not refreshing on time period. Suggested is 15 minutes.
	CredentialRefreshDelay *duration.Duration `protobuf:"bytes,4,opt,name=credential_refresh_delay,json=credentialRefreshDelay,proto3" json:"credential_refresh_delay,omitempty"`
}

func (x *AWSLambdaConfig) Reset() {
	*x = AWSLambdaConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSLambdaConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSLambdaConfig) ProtoMessage() {}

func (x *AWSLambdaConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSLambdaConfig.ProtoReflect.Descriptor instead.
func (*AWSLambdaConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_rawDescGZIP(), []int{2}
}

func (m *AWSLambdaConfig) GetCredentialsFetcher() isAWSLambdaConfig_CredentialsFetcher {
	if m != nil {
		return m.CredentialsFetcher
	}
	return nil
}

func (x *AWSLambdaConfig) GetUseDefaultCredentials() *wrappers.BoolValue {
	if x, ok := x.GetCredentialsFetcher().(*AWSLambdaConfig_UseDefaultCredentials); ok {
		return x.UseDefaultCredentials
	}
	return nil
}

func (x *AWSLambdaConfig) GetServiceAccountCredentials() *AWSLambdaConfig_ServiceAccountCredentials {
	if x, ok := x.GetCredentialsFetcher().(*AWSLambdaConfig_ServiceAccountCredentials_); ok {
		return x.ServiceAccountCredentials
	}
	return nil
}

func (x *AWSLambdaConfig) GetPropagateOriginalRouting() bool {
	if x != nil {
		return x.PropagateOriginalRouting
	}
	return false
}

func (x *AWSLambdaConfig) GetCredentialRefreshDelay() *duration.Duration {
	if x != nil {
		return x.CredentialRefreshDelay
	}
	return nil
}

type isAWSLambdaConfig_CredentialsFetcher interface {
	isAWSLambdaConfig_CredentialsFetcher()
}

type AWSLambdaConfig_UseDefaultCredentials struct {
	// Use AWS default credentials chain to get credentials.
	// This will search environment variables, ECS metadata and instance metadata
	// to get the credentials. credentials will be rotated automatically.
	//
	// If credentials are provided on the cluster (using the
	// AWSLambdaProtocolExtension), it will override these credentials. This
	// defaults to false, but may change in the future to true.
	UseDefaultCredentials *wrappers.BoolValue `protobuf:"bytes,1,opt,name=use_default_credentials,json=useDefaultCredentials,proto3,oneof"`
}

type AWSLambdaConfig_ServiceAccountCredentials_ struct {
	// Use projected service account token, and role arn to create reate temporary
	// credentials with which to authenticate lambda requests.
	// This functionality is meant to work along side EKS service account to IAM
	// binding functionality as outlined here:
	// https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.htmll
	//
	// If the following environment values are not present, this option cannot be used.
	//  1. AWS_WEB_IDENTITY_TOKEN_FILE
	//  2. AWS_ROLE_ARN
	//
	// The role arn may also be specified in the `AWSLambdaProtocolExtension` on the cluster level,
	// to override the environment variable.
	ServiceAccountCredentials *AWSLambdaConfig_ServiceAccountCredentials `protobuf:"bytes,2,opt,name=service_account_credentials,json=serviceAccountCredentials,proto3,oneof"`
}

func (*AWSLambdaConfig_UseDefaultCredentials) isAWSLambdaConfig_CredentialsFetcher() {}

func (*AWSLambdaConfig_ServiceAccountCredentials_) isAWSLambdaConfig_CredentialsFetcher() {}

// In order to specify the aws sts endpoint, both the cluster and uri must be set.
// This is due to an envoy limitation which cannot infer the host or path from the cluster,
// and therefore must be explicitly specified via the uri
type AWSLambdaConfig_ServiceAccountCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the envoy cluster which represents the desired aws sts endpoint
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// The full uri of the aws sts endpoint
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	// timeout for the request
	Timeout *duration.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *AWSLambdaConfig_ServiceAccountCredentials) Reset() {
	*x = AWSLambdaConfig_ServiceAccountCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSLambdaConfig_ServiceAccountCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSLambdaConfig_ServiceAccountCredentials) ProtoMessage() {}

func (x *AWSLambdaConfig_ServiceAccountCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSLambdaConfig_ServiceAccountCredentials.ProtoReflect.Descriptor instead.
func (*AWSLambdaConfig_ServiceAccountCredentials) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_rawDescGZIP(), []int{2, 0}
}

func (x *AWSLambdaConfig_ServiceAccountCredentials) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *AWSLambdaConfig_ServiceAccountCredentials) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *AWSLambdaConfig_ServiceAccountCredentials) GetTimeout() *duration.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_rawDesc = []byte{
	0x0a, 0x7a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2d, 0x67, 0x6c, 0x6f, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x61, 0x77,
	0x73, 0x5f, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x77, 0x73, 0x5f,
	0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x77, 0x73, 0x5f, 0x6c, 0x61, 0x6d, 0x62, 0x64,
	0x61, 0x2e, 0x76, 0x32, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x6a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x72, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x02, 0x0a, 0x11, 0x41, 0x57, 0x53, 0x4c, 0x61, 0x6d, 0x62,
	0x64, 0x61, 0x50, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20,
	0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x61, 0x6c, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75, 0x61, 0x6c,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x4c, 0x0a, 0x13, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x6f, 0x64,
	0x79, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x75, 0x6e, 0x77,
	0x72, 0x61, 0x70, 0x5f, 0x61, 0x73, 0x5f, 0x61, 0x6c, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x75, 0x6e, 0x77, 0x72, 0x61, 0x70, 0x41, 0x73, 0x41, 0x6c, 0x62, 0x12, 0x59, 0x0a,
	0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xd8, 0x01, 0x0a, 0x1a, 0x41, 0x57, 0x53,
	0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x4b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65,
	0x5f, 0x61, 0x72, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6c, 0x65,
	0x41, 0x72, 0x6e, 0x22, 0xb8, 0x04, 0x0a, 0x0f, 0x41, 0x57, 0x53, 0x4c, 0x61, 0x6d, 0x62, 0x64,
	0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x54, 0x0a, 0x17, 0x75, 0x73, 0x65, 0x5f, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x15, 0x75, 0x73, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x93, 0x01,
	0x0a, 0x1b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x51, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61,
	0x77, 0x73, 0x5f, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x57, 0x53,
	0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x48, 0x00, 0x52, 0x19, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x12, 0x3c, 0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65,
	0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x70, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x53, 0x0a, 0x18, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x16,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x1a, 0x8e, 0x01, 0x0a, 0x19, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x12, 0x21, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x07,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x03, 0x75,
	0x72, 0x69, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42, 0x15, 0x0a, 0x13, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0xa4,
	0x01, 0x0a, 0x34, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x77, 0x73, 0x5f, 0x6c, 0x61,
	0x6d, 0x62, 0x64, 0x61, 0x2e, 0x76, 0x32, 0x42, 0x0e, 0x41, 0x77, 0x73, 0x4c, 0x61, 0x6d, 0x62,
	0x64, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2d, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x61, 0x77, 0x73, 0x5f, 0x6c, 0x61, 0x6d, 0x62,
	0x64, 0x61, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_goTypes = []interface{}{
	(*AWSLambdaPerRoute)(nil),                         // 0: envoy.config.filter.http.aws_lambda.v2.AWSLambdaPerRoute
	(*AWSLambdaProtocolExtension)(nil),                // 1: envoy.config.filter.http.aws_lambda.v2.AWSLambdaProtocolExtension
	(*AWSLambdaConfig)(nil),                           // 2: envoy.config.filter.http.aws_lambda.v2.AWSLambdaConfig
	(*AWSLambdaConfig_ServiceAccountCredentials)(nil), // 3: envoy.config.filter.http.aws_lambda.v2.AWSLambdaConfig.ServiceAccountCredentials
	(*wrappers.StringValue)(nil),                      // 4: google.protobuf.StringValue
	(*v3.TypedExtensionConfig)(nil),                   // 5: envoy.config.core.v3.TypedExtensionConfig
	(*wrappers.BoolValue)(nil),                        // 6: google.protobuf.BoolValue
	(*duration.Duration)(nil),                         // 7: google.protobuf.Duration
}
var file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_depIdxs = []int32{
	4, // 0: envoy.config.filter.http.aws_lambda.v2.AWSLambdaPerRoute.empty_body_override:type_name -> google.protobuf.StringValue
	5, // 1: envoy.config.filter.http.aws_lambda.v2.AWSLambdaPerRoute.transformer_config:type_name -> envoy.config.core.v3.TypedExtensionConfig
	6, // 2: envoy.config.filter.http.aws_lambda.v2.AWSLambdaConfig.use_default_credentials:type_name -> google.protobuf.BoolValue
	3, // 3: envoy.config.filter.http.aws_lambda.v2.AWSLambdaConfig.service_account_credentials:type_name -> envoy.config.filter.http.aws_lambda.v2.AWSLambdaConfig.ServiceAccountCredentials
	7, // 4: envoy.config.filter.http.aws_lambda.v2.AWSLambdaConfig.credential_refresh_delay:type_name -> google.protobuf.Duration
	7, // 5: envoy.config.filter.http.aws_lambda.v2.AWSLambdaConfig.ServiceAccountCredentials.timeout:type_name -> google.protobuf.Duration
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AWSLambdaPerRoute); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AWSLambdaProtocolExtension); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AWSLambdaConfig); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AWSLambdaConfig_ServiceAccountCredentials); i {
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
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*AWSLambdaConfig_UseDefaultCredentials)(nil),
		(*AWSLambdaConfig_ServiceAccountCredentials_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_api_envoy_config_filter_http_aws_lambda_v2_aws_lambda_proto_depIdxs = nil
}
