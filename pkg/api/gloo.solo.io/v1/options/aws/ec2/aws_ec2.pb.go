// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/aws/ec2/aws_ec2.proto

package ec2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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

// Upstream Spec for AWS Lambda Upstreams
// AWS Upstreams represent a collection of Lambda Functions for a particular AWS Account (IAM Role or User account)
// in a particular region
type UpstreamSpec struct {
	// The AWS Region where the desired EC2 instances exist
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// Optional, if not set, Gloo will try to use the default AWS secret specified by environment variables.
	// If a secret is not provided, the environment must specify both the AWS access key and secret.
	// The environment variables used to indicate the AWS account can be:
	// - for the access key: "AWS_ACCESS_KEY_ID" or "AWS_ACCESS_KEY"
	// - for the secret: "AWS_SECRET_ACCESS_KEY" or "AWS_SECRET_KEY"
	// If set, a [Gloo Secret Ref](https://gloo.solo.io/introduction/concepts/#Secrets) to an AWS Secret
	// AWS Secrets can be created with `glooctl secret create aws ...`
	// If the secret is created manually, it must conform to the following structure:
	//  ```
	//  access_key: <aws access key>
	//  secret_key: <aws secret key>
	//  ```
	// Gloo will create an EC2 API client with this credential. You may choose to use a credential with limited access
	// in conjunction with a list of Roles, specified by their Amazon Resource Number (ARN).
	SecretRef *core.ResourceRef `protobuf:"bytes,2,opt,name=secret_ref,json=secretRef,proto3" json:"secret_ref,omitempty"`
	// Optional, Amazon Resource Number (ARN) referring to IAM Role that should be assumed when the Upstream
	// queries for eligible EC2 instances.
	// If provided, Gloo will create an EC2 API client with the provided role. If not provided, Gloo will not assume
	// a role.
	RoleArn string `protobuf:"bytes,7,opt,name=role_arn,json=roleArn,proto3" json:"role_arn,omitempty"`
	// List of tag filters for selecting instances
	// An instance must match all the filters in order to be selected
	// Filter keys are not case-sensitive
	Filters []*TagFilter `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	// If set, will use the EC2 public IP address. Defaults to the private IP address.
	PublicIp bool `protobuf:"varint,4,opt,name=public_ip,json=publicIp,proto3" json:"public_ip,omitempty"`
	// If set, will use this port on EC2 instances. Defaults to port 80.
	Port                 uint32   `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_02f25168f7bdc45f, []int{0}
}
func (m *UpstreamSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSpec.Unmarshal(m, b)
}
func (m *UpstreamSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSpec.Marshal(b, m, deterministic)
}
func (m *UpstreamSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSpec.Merge(m, src)
}
func (m *UpstreamSpec) XXX_Size() int {
	return xxx_messageInfo_UpstreamSpec.Size(m)
}
func (m *UpstreamSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSpec proto.InternalMessageInfo

func (m *UpstreamSpec) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *UpstreamSpec) GetSecretRef() *core.ResourceRef {
	if m != nil {
		return m.SecretRef
	}
	return nil
}

func (m *UpstreamSpec) GetRoleArn() string {
	if m != nil {
		return m.RoleArn
	}
	return ""
}

func (m *UpstreamSpec) GetFilters() []*TagFilter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *UpstreamSpec) GetPublicIp() bool {
	if m != nil {
		return m.PublicIp
	}
	return false
}

func (m *UpstreamSpec) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type TagFilter struct {
	// Types that are valid to be assigned to Spec:
	//	*TagFilter_Key
	//	*TagFilter_KvPair_
	Spec                 isTagFilter_Spec `protobuf_oneof:"spec"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TagFilter) Reset()         { *m = TagFilter{} }
func (m *TagFilter) String() string { return proto.CompactTextString(m) }
func (*TagFilter) ProtoMessage()    {}
func (*TagFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_02f25168f7bdc45f, []int{1}
}
func (m *TagFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagFilter.Unmarshal(m, b)
}
func (m *TagFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagFilter.Marshal(b, m, deterministic)
}
func (m *TagFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagFilter.Merge(m, src)
}
func (m *TagFilter) XXX_Size() int {
	return xxx_messageInfo_TagFilter.Size(m)
}
func (m *TagFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_TagFilter.DiscardUnknown(m)
}

var xxx_messageInfo_TagFilter proto.InternalMessageInfo

type isTagFilter_Spec interface {
	isTagFilter_Spec()
	Equal(interface{}) bool
}

type TagFilter_Key struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3,oneof" json:"key,omitempty"`
}
type TagFilter_KvPair_ struct {
	KvPair *TagFilter_KvPair `protobuf:"bytes,2,opt,name=kv_pair,json=kvPair,proto3,oneof" json:"kv_pair,omitempty"`
}

func (*TagFilter_Key) isTagFilter_Spec()     {}
func (*TagFilter_KvPair_) isTagFilter_Spec() {}

func (m *TagFilter) GetSpec() isTagFilter_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *TagFilter) GetKey() string {
	if x, ok := m.GetSpec().(*TagFilter_Key); ok {
		return x.Key
	}
	return ""
}

func (m *TagFilter) GetKvPair() *TagFilter_KvPair {
	if x, ok := m.GetSpec().(*TagFilter_KvPair_); ok {
		return x.KvPair
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TagFilter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TagFilter_Key)(nil),
		(*TagFilter_KvPair_)(nil),
	}
}

type TagFilter_KvPair struct {
	// keys are not case-sensitive, as with AWS Condition Keys
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// values are case-sensitive
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TagFilter_KvPair) Reset()         { *m = TagFilter_KvPair{} }
func (m *TagFilter_KvPair) String() string { return proto.CompactTextString(m) }
func (*TagFilter_KvPair) ProtoMessage()    {}
func (*TagFilter_KvPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_02f25168f7bdc45f, []int{1, 0}
}
func (m *TagFilter_KvPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagFilter_KvPair.Unmarshal(m, b)
}
func (m *TagFilter_KvPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagFilter_KvPair.Marshal(b, m, deterministic)
}
func (m *TagFilter_KvPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagFilter_KvPair.Merge(m, src)
}
func (m *TagFilter_KvPair) XXX_Size() int {
	return xxx_messageInfo_TagFilter_KvPair.Size(m)
}
func (m *TagFilter_KvPair) XXX_DiscardUnknown() {
	xxx_messageInfo_TagFilter_KvPair.DiscardUnknown(m)
}

var xxx_messageInfo_TagFilter_KvPair proto.InternalMessageInfo

func (m *TagFilter_KvPair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *TagFilter_KvPair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "aws_ec2.options.gloo.solo.io.UpstreamSpec")
	proto.RegisterType((*TagFilter)(nil), "aws_ec2.options.gloo.solo.io.TagFilter")
	proto.RegisterType((*TagFilter_KvPair)(nil), "aws_ec2.options.gloo.solo.io.TagFilter.KvPair")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/aws/ec2/aws_ec2.proto", fileDescriptor_02f25168f7bdc45f)
}

var fileDescriptor_02f25168f7bdc45f = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0xad, 0xd9, 0x6d, 0x76, 0xe3, 0x82, 0x84, 0xac, 0x0a, 0x65, 0x17, 0x84, 0xa2, 0x5e, 0xc8,
	0x05, 0x9b, 0x2e, 0x17, 0xae, 0xad, 0x10, 0x6a, 0x55, 0x09, 0x21, 0x03, 0x17, 0x2e, 0x91, 0xd7,
	0x9a, 0x04, 0x2b, 0xe9, 0x8e, 0x65, 0x7b, 0x43, 0xf9, 0x1f, 0x0e, 0x7c, 0x02, 0xdf, 0xc3, 0x0f,
	0x70, 0xe2, 0x8e, 0x62, 0xef, 0x56, 0x1c, 0x50, 0xc5, 0x25, 0x33, 0xef, 0xe5, 0xcd, 0xcc, 0x1b,
	0x79, 0xe8, 0x55, 0x6b, 0xc2, 0xe7, 0xed, 0x9a, 0x6b, 0xbc, 0x16, 0x1e, 0x7b, 0x7c, 0x6e, 0x30,
	0x45, 0x65, 0x8d, 0x17, 0xca, 0x1a, 0xd1, 0xf6, 0x88, 0xe9, 0x33, 0x9c, 0x0a, 0xb4, 0xc1, 0xe0,
	0xc6, 0x0b, 0xf5, 0xc5, 0x0b, 0xd0, 0xab, 0x31, 0xd6, 0xa0, 0x57, 0xdc, 0x3a, 0x0c, 0xc8, 0x9e,
	0xec, 0xe1, 0x4e, 0xc6, 0xc7, 0x32, 0x3e, 0xb6, 0xe3, 0x06, 0x97, 0xc7, 0x2d, 0xb6, 0x18, 0x85,
	0x62, 0xcc, 0x52, 0xcd, 0x92, 0xc1, 0x4d, 0x48, 0x24, 0xdc, 0x84, 0x1d, 0xb7, 0x88, 0x0e, 0x3a,
	0x13, 0xa2, 0x81, 0xe1, 0x54, 0x38, 0x68, 0xd2, 0xaf, 0x93, 0x5f, 0x84, 0xde, 0xff, 0x68, 0x7d,
	0x70, 0xa0, 0xae, 0xdf, 0x5b, 0xd0, 0xec, 0x11, 0xcd, 0x1c, 0xb4, 0x06, 0x37, 0x05, 0x29, 0x49,
	0x95, 0xcb, 0x1d, 0x62, 0xaf, 0x28, 0xf5, 0xa0, 0x1d, 0x84, 0xda, 0x41, 0x53, 0xdc, 0x2b, 0x49,
	0x75, 0xb4, 0x5a, 0x70, 0x8d, 0x0e, 0xf6, 0x86, 0xb8, 0x04, 0x8f, 0x5b, 0xa7, 0x41, 0x42, 0x23,
	0xf3, 0x24, 0x96, 0xd0, 0xb0, 0x05, 0x9d, 0x3b, 0xec, 0xa1, 0x56, 0x6e, 0x53, 0xcc, 0x62, 0xcf,
	0xd9, 0x88, 0xcf, 0xdc, 0x86, 0x9d, 0xd1, 0x59, 0x63, 0xfa, 0x00, 0xce, 0x17, 0x93, 0x72, 0x52,
	0x1d, 0xad, 0x9e, 0xf1, 0xbb, 0x56, 0xe6, 0x1f, 0x54, 0xfb, 0x26, 0xea, 0xe5, 0xbe, 0x8e, 0x3d,
	0xa6, 0xb9, 0xdd, 0xae, 0x7b, 0xa3, 0x6b, 0x63, 0x8b, 0x69, 0x49, 0xaa, 0xb9, 0x9c, 0x27, 0xe2,
	0xd2, 0x32, 0x46, 0xa7, 0x16, 0x5d, 0x28, 0x0e, 0x4b, 0x52, 0x3d, 0x90, 0x31, 0x3f, 0xf9, 0x46,
	0x68, 0x7e, 0xdb, 0x87, 0x31, 0x3a, 0xe9, 0xe0, 0x6b, 0xda, 0xf5, 0xe2, 0x40, 0x8e, 0x80, 0x5d,
	0xd2, 0x59, 0x37, 0xd4, 0x56, 0x19, 0xb7, 0xdb, 0x93, 0xff, 0xa7, 0x2b, 0x7e, 0x35, 0xbc, 0x53,
	0xc6, 0x5d, 0x1c, 0xc8, 0xac, 0x8b, 0xd9, 0xf2, 0x05, 0xcd, 0x12, 0xc7, 0x1e, 0xfe, 0x35, 0x28,
	0x8d, 0x39, 0xa6, 0x87, 0x83, 0xea, 0xb7, 0x10, 0x87, 0xe4, 0x32, 0x81, 0xf3, 0x8c, 0x4e, 0xbd,
	0x05, 0x7d, 0xfe, 0xf6, 0xc7, 0xef, 0x29, 0xf9, 0xfe, 0xf3, 0x29, 0xf9, 0xf4, 0xfa, 0xce, 0x93,
	0xb2, 0x5d, 0x7b, 0x7b, 0x56, 0x7b, 0x47, 0xff, 0xb8, 0xac, 0x75, 0x16, 0xdf, 0xfb, 0xe5, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x87, 0xbb, 0x67, 0x64, 0xa1, 0x02, 0x00, 0x00,
}

func (this *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Region != that1.Region {
		return false
	}
	if !this.SecretRef.Equal(that1.SecretRef) {
		return false
	}
	if this.RoleArn != that1.RoleArn {
		return false
	}
	if len(this.Filters) != len(that1.Filters) {
		return false
	}
	for i := range this.Filters {
		if !this.Filters[i].Equal(that1.Filters[i]) {
			return false
		}
	}
	if this.PublicIp != that1.PublicIp {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TagFilter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TagFilter)
	if !ok {
		that2, ok := that.(TagFilter)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Spec == nil {
		if this.Spec != nil {
			return false
		}
	} else if this.Spec == nil {
		return false
	} else if !this.Spec.Equal(that1.Spec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TagFilter_Key) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TagFilter_Key)
	if !ok {
		that2, ok := that.(TagFilter_Key)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	return true
}
func (this *TagFilter_KvPair_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TagFilter_KvPair_)
	if !ok {
		that2, ok := that.(TagFilter_KvPair_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.KvPair.Equal(that1.KvPair) {
		return false
	}
	return true
}
func (this *TagFilter_KvPair) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TagFilter_KvPair)
	if !ok {
		that2, ok := that.(TagFilter_KvPair)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
