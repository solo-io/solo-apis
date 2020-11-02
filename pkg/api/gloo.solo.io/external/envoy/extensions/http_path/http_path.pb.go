// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/extensions/http_path/http_path.proto

package http_path

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v3 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/udpa/annotations"
	math "math"
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

// Same as HTTP health checker, but allows a custom path per endpoint
// The http path to use can be overriden using endpoint metadata. The endpoint specific path should
// be in the "io.solo.health_checkers.http_path" namespace, under a string value named "path".
type HttpPath struct {
	// Http health check.
	HttpHealthCheck      *v3.HealthCheck_HttpHealthCheck `protobuf:"bytes,1,opt,name=http_health_check,json=httpHealthCheck,proto3" json:"http_health_check,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *HttpPath) Reset()         { *m = HttpPath{} }
func (m *HttpPath) String() string { return proto.CompactTextString(m) }
func (*HttpPath) ProtoMessage()    {}
func (*HttpPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_957d442127254af9, []int{0}
}
func (m *HttpPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpPath.Unmarshal(m, b)
}
func (m *HttpPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpPath.Marshal(b, m, deterministic)
}
func (m *HttpPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpPath.Merge(m, src)
}
func (m *HttpPath) XXX_Size() int {
	return xxx_messageInfo_HttpPath.Size(m)
}
func (m *HttpPath) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpPath.DiscardUnknown(m)
}

var xxx_messageInfo_HttpPath proto.InternalMessageInfo

func (m *HttpPath) GetHttpHealthCheck() *v3.HealthCheck_HttpHealthCheck {
	if m != nil {
		return m.HttpHealthCheck
	}
	return nil
}

func init() {
	proto.RegisterType((*HttpPath)(nil), "envoy.config.health_checker.http_path.v2.HttpPath")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/extensions/http_path/http_path.proto", fileDescriptor_957d442127254af9)
}

var fileDescriptor_957d442127254af9 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4a, 0x03, 0x41,
	0x10, 0x86, 0xb9, 0x14, 0x41, 0xce, 0x42, 0x0d, 0x16, 0x12, 0x50, 0xc4, 0xc6, 0x34, 0xce, 0x62,
	0x02, 0x3e, 0x40, 0x6c, 0x52, 0x06, 0xc1, 0x46, 0xd0, 0x30, 0x39, 0xd7, 0xdd, 0x25, 0xe7, 0xce,
	0xb2, 0x3b, 0x77, 0x24, 0x9d, 0x6f, 0xa3, 0xcf, 0xe0, 0x13, 0xd8, 0xfa, 0x0a, 0x3e, 0x89, 0xdc,
	0xae, 0x24, 0x97, 0x46, 0xd2, 0xdc, 0xcd, 0x0c, 0xff, 0x7e, 0xff, 0x30, 0x7f, 0xfe, 0xa4, 0x0c,
	0xeb, 0x6a, 0x0e, 0x05, 0xbd, 0x8a, 0x40, 0x25, 0x5d, 0x19, 0x4a, 0x7f, 0x74, 0x26, 0x08, 0x74,
	0x46, 0xa8, 0x92, 0x28, 0x7d, 0xe4, 0x92, 0xa5, 0xb7, 0x58, 0x0a, 0x69, 0x6b, 0x5a, 0xc5, 0xd6,
	0x06, 0x43, 0x36, 0x08, 0xcd, 0xec, 0x66, 0x0e, 0x59, 0x6f, 0x2a, 0x70, 0x9e, 0x98, 0x7a, 0x83,
	0xa8, 0x85, 0x82, 0xec, 0x8b, 0x51, 0xa0, 0x25, 0x96, 0xac, 0x67, 0x85, 0x96, 0xc5, 0x42, 0x7a,
	0xd8, 0x88, 0xeb, 0x61, 0xff, 0xb4, 0x7a, 0x76, 0x28, 0xd0, 0x5a, 0x62, 0xe4, 0x48, 0x0d, 0x8c,
	0x5c, 0x85, 0x04, 0xea, 0x5f, 0x26, 0xd3, 0x04, 0x12, 0x05, 0x79, 0x29, 0xea, 0x91, 0x68, 0x03,
	0xff, 0x84, 0xc7, 0x8a, 0x14, 0xc5, 0x52, 0x34, 0x55, 0x9a, 0x5e, 0x98, 0x7c, 0x6f, 0xc2, 0xec,
	0xa6, 0xc8, 0xba, 0xf7, 0x98, 0x1f, 0x45, 0xe7, 0xf6, 0xe3, 0x93, 0xec, 0x3c, 0x1b, 0xec, 0x0f,
	0xaf, 0x61, 0x6b, 0xdf, 0xc6, 0x06, 0xea, 0x11, 0x4c, 0xa2, 0xf2, 0x36, 0xba, 0x34, 0x98, 0x56,
	0x7f, 0x77, 0xa0, 0xb7, 0x07, 0xe3, 0xf7, 0xec, 0xe3, 0xe7, 0x2c, 0xfb, 0x7c, 0xfb, 0xfa, 0xee,
	0x76, 0x0e, 0x3b, 0xf9, 0x8d, 0xa1, 0x04, 0x75, 0x9e, 0x96, 0x2b, 0xd8, 0xf5, 0x1e, 0xe3, 0xf5,
	0xbe, 0xd3, 0xec, 0xe1, 0xfe, 0xdf, 0x94, 0xdc, 0x42, 0xad, 0x93, 0x82, 0x66, 0x0c, 0x66, 0xa7,
	0xb0, 0xe6, 0xdd, 0x78, 0x9b, 0xd1, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x92, 0x58, 0x0b,
	0x05, 0x02, 0x00, 0x00,
}

func (this *HttpPath) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HttpPath)
	if !ok {
		that2, ok := that.(HttpPath)
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
	if !this.HttpHealthCheck.Equal(that1.HttpHealthCheck) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
