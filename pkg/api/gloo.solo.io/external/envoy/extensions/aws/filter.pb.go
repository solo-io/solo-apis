// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/extensions/aws/filter.proto

package aws

import (
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// AWS Lambda contains the configuration necessary to perform transform regular
// http calls to AWS Lambda invocations.
type AWSLambdaPerRoute struct {
	// The name of the function
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The qualifier of the function (defaults to $LATEST if not specified)
	Qualifier string `protobuf:"bytes,2,opt,name=qualifier,proto3" json:"qualifier,omitempty"`
	// Invocation type - async or regular.
	Async bool `protobuf:"varint,3,opt,name=async,proto3" json:"async,omitempty"`
	// Optional default body if the body is empty. By default on default
	// body is used if the body empty, and an empty body will be sent upstream.
	EmptyBodyOverride    *types.StringValue `protobuf:"bytes,4,opt,name=empty_body_override,json=emptyBodyOverride,proto3" json:"empty_body_override,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AWSLambdaPerRoute) Reset()         { *m = AWSLambdaPerRoute{} }
func (m *AWSLambdaPerRoute) String() string { return proto.CompactTextString(m) }
func (*AWSLambdaPerRoute) ProtoMessage()    {}
func (*AWSLambdaPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2816ebd70617e89, []int{0}
}
func (m *AWSLambdaPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AWSLambdaPerRoute.Unmarshal(m, b)
}
func (m *AWSLambdaPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AWSLambdaPerRoute.Marshal(b, m, deterministic)
}
func (m *AWSLambdaPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AWSLambdaPerRoute.Merge(m, src)
}
func (m *AWSLambdaPerRoute) XXX_Size() int {
	return xxx_messageInfo_AWSLambdaPerRoute.Size(m)
}
func (m *AWSLambdaPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_AWSLambdaPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_AWSLambdaPerRoute proto.InternalMessageInfo

func (m *AWSLambdaPerRoute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AWSLambdaPerRoute) GetQualifier() string {
	if m != nil {
		return m.Qualifier
	}
	return ""
}

func (m *AWSLambdaPerRoute) GetAsync() bool {
	if m != nil {
		return m.Async
	}
	return false
}

func (m *AWSLambdaPerRoute) GetEmptyBodyOverride() *types.StringValue {
	if m != nil {
		return m.EmptyBodyOverride
	}
	return nil
}

type AWSLambdaProtocolExtension struct {
	// The host header for AWS this cluster
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// The region for this cluster
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	// The access_key for AWS this cluster
	AccessKey string `protobuf:"bytes,3,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	// The secret_key for AWS this cluster
	SecretKey            string   `protobuf:"bytes,4,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AWSLambdaProtocolExtension) Reset()         { *m = AWSLambdaProtocolExtension{} }
func (m *AWSLambdaProtocolExtension) String() string { return proto.CompactTextString(m) }
func (*AWSLambdaProtocolExtension) ProtoMessage()    {}
func (*AWSLambdaProtocolExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2816ebd70617e89, []int{1}
}
func (m *AWSLambdaProtocolExtension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AWSLambdaProtocolExtension.Unmarshal(m, b)
}
func (m *AWSLambdaProtocolExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AWSLambdaProtocolExtension.Marshal(b, m, deterministic)
}
func (m *AWSLambdaProtocolExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AWSLambdaProtocolExtension.Merge(m, src)
}
func (m *AWSLambdaProtocolExtension) XXX_Size() int {
	return xxx_messageInfo_AWSLambdaProtocolExtension.Size(m)
}
func (m *AWSLambdaProtocolExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_AWSLambdaProtocolExtension.DiscardUnknown(m)
}

var xxx_messageInfo_AWSLambdaProtocolExtension proto.InternalMessageInfo

func (m *AWSLambdaProtocolExtension) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *AWSLambdaProtocolExtension) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *AWSLambdaProtocolExtension) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *AWSLambdaProtocolExtension) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

type AWSLambdaConfig struct {
	// Use AWS default credentials chain to get credentials.
	// This will search environment variables, ECS metadata and instance metadata
	// to get the credentials. credentials will be rotated automatically.
	//
	// If credentials are provided on the cluster (using the
	// AWSLambdaProtocolExtension), it will override these credentials. This
	// defaults to false, but may change in the future to true.
	UseDefaultCredentials *types.BoolValue `protobuf:"bytes,1,opt,name=use_default_credentials,json=useDefaultCredentials,proto3" json:"use_default_credentials,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}         `json:"-"`
	XXX_unrecognized      []byte           `json:"-"`
	XXX_sizecache         int32            `json:"-"`
}

func (m *AWSLambdaConfig) Reset()         { *m = AWSLambdaConfig{} }
func (m *AWSLambdaConfig) String() string { return proto.CompactTextString(m) }
func (*AWSLambdaConfig) ProtoMessage()    {}
func (*AWSLambdaConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2816ebd70617e89, []int{2}
}
func (m *AWSLambdaConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AWSLambdaConfig.Unmarshal(m, b)
}
func (m *AWSLambdaConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AWSLambdaConfig.Marshal(b, m, deterministic)
}
func (m *AWSLambdaConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AWSLambdaConfig.Merge(m, src)
}
func (m *AWSLambdaConfig) XXX_Size() int {
	return xxx_messageInfo_AWSLambdaConfig.Size(m)
}
func (m *AWSLambdaConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AWSLambdaConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AWSLambdaConfig proto.InternalMessageInfo

func (m *AWSLambdaConfig) GetUseDefaultCredentials() *types.BoolValue {
	if m != nil {
		return m.UseDefaultCredentials
	}
	return nil
}

func init() {
	proto.RegisterType((*AWSLambdaPerRoute)(nil), "envoy.config.filter.http.aws_lambda.v2.AWSLambdaPerRoute")
	proto.RegisterType((*AWSLambdaProtocolExtension)(nil), "envoy.config.filter.http.aws_lambda.v2.AWSLambdaProtocolExtension")
	proto.RegisterType((*AWSLambdaConfig)(nil), "envoy.config.filter.http.aws_lambda.v2.AWSLambdaConfig")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/extensions/aws/filter.proto", fileDescriptor_f2816ebd70617e89)
}

var fileDescriptor_f2816ebd70617e89 = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xd1, 0x6e, 0xd3, 0x30,
	0x18, 0x85, 0x95, 0x51, 0x06, 0xf5, 0x24, 0xd0, 0x02, 0x68, 0x55, 0x19, 0x50, 0xf5, 0x02, 0xf5,
	0x06, 0x5b, 0x2a, 0xbc, 0xc0, 0x32, 0xb8, 0xda, 0xa4, 0x4d, 0x99, 0x00, 0x89, 0x9b, 0xc8, 0x49,
	0xfe, 0xa4, 0xd6, 0x5c, 0xff, 0xc6, 0x76, 0xda, 0xf9, 0x4d, 0xe0, 0x35, 0x78, 0x3c, 0xae, 0x50,
	0xec, 0x76, 0x45, 0x9a, 0x04, 0xbb, 0x49, 0xe2, 0x73, 0x7c, 0xac, 0xef, 0x77, 0x0e, 0xf9, 0xdc,
	0x0a, 0xb7, 0xe8, 0x4a, 0x5a, 0xe1, 0x92, 0x59, 0x94, 0xf8, 0x4e, 0x60, 0x7c, 0x73, 0x2d, 0x2c,
	0xe3, 0x5a, 0xb0, 0x56, 0x22, 0xc6, 0x07, 0xdc, 0x38, 0x30, 0x8a, 0x4b, 0x06, 0x6a, 0x85, 0x3e,
	0x2c, 0x95, 0x15, 0xa8, 0x2c, 0xe3, 0x6b, 0xcb, 0x1a, 0x21, 0x1d, 0x18, 0xaa, 0x0d, 0x3a, 0x4c,
	0xdf, 0x86, 0x2d, 0xb4, 0x42, 0xd5, 0x88, 0x96, 0x6e, 0xac, 0x85, 0x73, 0x9a, 0xf2, 0xb5, 0x2d,
	0x24, 0x5f, 0x96, 0x35, 0xa7, 0xab, 0xf9, 0xf8, 0x75, 0x8b, 0xd8, 0x4a, 0x60, 0x21, 0x55, 0x76,
	0x0d, 0x5b, 0x1b, 0xae, 0x35, 0x18, 0x1b, 0xcf, 0x19, 0x1f, 0xad, 0xb8, 0x14, 0x35, 0x77, 0xc0,
	0xb6, 0x1f, 0xd1, 0x98, 0xfe, 0x4a, 0xc8, 0xe1, 0xc9, 0xd7, 0xab, 0xf3, 0x70, 0xd2, 0x25, 0x98,
	0x1c, 0x3b, 0x07, 0xe9, 0x4b, 0x32, 0x50, 0x7c, 0x09, 0xa3, 0x64, 0x92, 0xcc, 0x86, 0xd9, 0xa3,
	0xdf, 0xd9, 0xc0, 0xec, 0x4d, 0x92, 0x3c, 0x88, 0xe9, 0x31, 0x19, 0x7e, 0xef, 0xb8, 0x14, 0x8d,
	0x00, 0x33, 0xda, 0xeb, 0x77, 0xe4, 0x3b, 0x21, 0x7d, 0x4e, 0x1e, 0x72, 0xeb, 0x55, 0x35, 0x7a,
	0x30, 0x49, 0x66, 0x8f, 0xf3, 0xb8, 0x48, 0xcf, 0xc9, 0x33, 0x58, 0x6a, 0xe7, 0x8b, 0x12, 0x6b,
	0x5f, 0xe0, 0x0a, 0x8c, 0x11, 0x35, 0x8c, 0x06, 0x93, 0x64, 0x76, 0x30, 0x3f, 0xa6, 0x91, 0x9e,
	0x6e, 0xe9, 0xe9, 0x95, 0x33, 0x42, 0xb5, 0x5f, 0xb8, 0xec, 0x20, 0x3f, 0x0c, 0xc1, 0x0c, 0x6b,
	0x7f, 0xb1, 0x89, 0x4d, 0x7f, 0x24, 0x64, 0xbc, 0x83, 0xee, 0x43, 0x15, 0xca, 0x4f, 0xdb, 0x6b,
	0xec, 0xe9, 0x17, 0x68, 0xdd, 0x1d, 0xfa, 0x5e, 0x4c, 0xdf, 0x90, 0x7d, 0x03, 0xad, 0x40, 0x15,
	0xd1, 0x77, 0xf6, 0x46, 0x4e, 0x5f, 0x11, 0xc2, 0xab, 0x0a, 0xac, 0x2d, 0xae, 0xc1, 0x87, 0x29,
	0x86, 0xf9, 0x30, 0x2a, 0x67, 0xe0, 0x7b, 0xdb, 0x42, 0x65, 0xc0, 0x05, 0x7b, 0x10, 0xed, 0xa8,
	0x9c, 0x81, 0x9f, 0x02, 0x79, 0x7a, 0x4b, 0x76, 0x1a, 0xfe, 0x5a, 0x9a, 0x93, 0xa3, 0xce, 0x42,
	0x51, 0x43, 0xc3, 0x3b, 0xe9, 0x8a, 0xca, 0x40, 0x0d, 0xca, 0x09, 0x2e, 0x6d, 0x20, 0x3c, 0x98,
	0x8f, 0xef, 0xcc, 0x9f, 0x21, 0xca, 0x38, 0xfd, 0x8b, 0xce, 0xc2, 0xc7, 0x98, 0x3c, 0xdd, 0x05,
	0xb3, 0x9f, 0x09, 0xf9, 0x20, 0x90, 0x86, 0x76, 0x68, 0x83, 0x37, 0x9e, 0xde, 0xaf, 0x28, 0xd9,
	0x93, 0x93, 0xb5, 0xfd, 0xeb, 0xde, 0x2e, 0x93, 0x6f, 0x17, 0xff, 0x6c, 0xae, 0xbe, 0x6e, 0x6f,
	0xdb, 0x4b, 0x7b, 0x99, 0x8a, 0xff, 0x14, 0xb8, 0xdc, 0x0f, 0x63, 0xbc, 0xff, 0x13, 0x00, 0x00,
	0xff, 0xff, 0xdb, 0x3e, 0xa2, 0xb4, 0x13, 0x03, 0x00, 0x00,
}
