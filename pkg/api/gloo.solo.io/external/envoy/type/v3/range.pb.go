// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/type/v3/range.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Specifies the int64 start and end of the range using half-open interval semantics [start,
// end).
type Int64Range struct {
	// start of the range (inclusive)
	Start int64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	// end of the range (exclusive)
	End                  int64    `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64Range) Reset()         { *m = Int64Range{} }
func (m *Int64Range) String() string { return proto.CompactTextString(m) }
func (*Int64Range) ProtoMessage()    {}
func (*Int64Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3cb43aa215b7b78, []int{0}
}
func (m *Int64Range) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64Range.Unmarshal(m, b)
}
func (m *Int64Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64Range.Marshal(b, m, deterministic)
}
func (m *Int64Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64Range.Merge(m, src)
}
func (m *Int64Range) XXX_Size() int {
	return xxx_messageInfo_Int64Range.Size(m)
}
func (m *Int64Range) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64Range.DiscardUnknown(m)
}

var xxx_messageInfo_Int64Range proto.InternalMessageInfo

func (m *Int64Range) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Int64Range) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

// Specifies the int32 start and end of the range using half-open interval semantics [start,
// end).
type Int32Range struct {
	// start of the range (inclusive)
	Start int32 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	// end of the range (exclusive)
	End                  int32    `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int32Range) Reset()         { *m = Int32Range{} }
func (m *Int32Range) String() string { return proto.CompactTextString(m) }
func (*Int32Range) ProtoMessage()    {}
func (*Int32Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3cb43aa215b7b78, []int{1}
}
func (m *Int32Range) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int32Range.Unmarshal(m, b)
}
func (m *Int32Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int32Range.Marshal(b, m, deterministic)
}
func (m *Int32Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int32Range.Merge(m, src)
}
func (m *Int32Range) XXX_Size() int {
	return xxx_messageInfo_Int32Range.Size(m)
}
func (m *Int32Range) XXX_DiscardUnknown() {
	xxx_messageInfo_Int32Range.DiscardUnknown(m)
}

var xxx_messageInfo_Int32Range proto.InternalMessageInfo

func (m *Int32Range) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Int32Range) GetEnd() int32 {
	if m != nil {
		return m.End
	}
	return 0
}

// Specifies the double start and end of the range using half-open interval semantics [start,
// end).
type DoubleRange struct {
	// start of the range (inclusive)
	Start float64 `protobuf:"fixed64,1,opt,name=start,proto3" json:"start,omitempty"`
	// end of the range (exclusive)
	End                  float64  `protobuf:"fixed64,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoubleRange) Reset()         { *m = DoubleRange{} }
func (m *DoubleRange) String() string { return proto.CompactTextString(m) }
func (*DoubleRange) ProtoMessage()    {}
func (*DoubleRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3cb43aa215b7b78, []int{2}
}
func (m *DoubleRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoubleRange.Unmarshal(m, b)
}
func (m *DoubleRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoubleRange.Marshal(b, m, deterministic)
}
func (m *DoubleRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleRange.Merge(m, src)
}
func (m *DoubleRange) XXX_Size() int {
	return xxx_messageInfo_DoubleRange.Size(m)
}
func (m *DoubleRange) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleRange.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleRange proto.InternalMessageInfo

func (m *DoubleRange) GetStart() float64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *DoubleRange) GetEnd() float64 {
	if m != nil {
		return m.End
	}
	return 0
}

func init() {
	proto.RegisterType((*Int64Range)(nil), "envoy.type.v3.Int64Range")
	proto.RegisterType((*Int32Range)(nil), "envoy.type.v3.Int32Range")
	proto.RegisterType((*DoubleRange)(nil), "envoy.type.v3.DoubleRange")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/type/v3/range.proto", fileDescriptor_d3cb43aa215b7b78)
}

var fileDescriptor_d3cb43aa215b7b78 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x31, 0x4b, 0x33, 0x31,
	0x1c, 0xc6, 0x49, 0x4b, 0x3b, 0xe4, 0xe5, 0x85, 0x52, 0xaa, 0x96, 0x6a, 0x8b, 0x76, 0x72, 0x31,
	0x01, 0x4f, 0x1c, 0x1c, 0x8b, 0x83, 0x0e, 0x42, 0x29, 0x4e, 0x6e, 0x39, 0x1b, 0x62, 0xf0, 0xcc,
	0x3f, 0x24, 0xb9, 0xa3, 0xb7, 0xe9, 0xe6, 0x67, 0xf0, 0x13, 0xf8, 0x19, 0xdc, 0x05, 0x57, 0xbf,
	0x82, 0x9f, 0x44, 0x92, 0x1c, 0x7a, 0xe2, 0xd9, 0xe5, 0xee, 0x9f, 0xe7, 0x79, 0xf2, 0x4b, 0xc8,
	0x83, 0x2f, 0x84, 0x74, 0x37, 0x79, 0x4a, 0xae, 0xe1, 0x8e, 0x5a, 0xc8, 0xe0, 0x40, 0x42, 0xfc,
	0x33, 0x2d, 0x2d, 0x65, 0x5a, 0x52, 0x91, 0x01, 0xc4, 0x0f, 0x5f, 0x39, 0x6e, 0x14, 0xcb, 0x28,
	0x57, 0x05, 0x94, 0xd4, 0x95, 0x9a, 0xd3, 0x22, 0xa1, 0x86, 0x29, 0xc1, 0x89, 0x36, 0xe0, 0xa0,
	0xff, 0x3f, 0x58, 0xc4, 0x5b, 0xa4, 0x48, 0x46, 0xe3, 0x7c, 0xa9, 0x19, 0x65, 0x4a, 0x81, 0x63,
	0x4e, 0x82, 0xb2, 0xd4, 0x3a, 0xe6, 0x72, 0x1b, 0xd3, 0xa3, 0xbd, 0x5f, 0x76, 0xc1, 0x8d, 0x95,
	0xa0, 0xa4, 0x12, 0x55, 0x64, 0x20, 0x40, 0x40, 0x18, 0xa9, 0x9f, 0xa2, 0x3a, 0x5d, 0x60, 0x7c,
	0xae, 0xdc, 0xf1, 0xd1, 0xc2, 0x1f, 0xdd, 0x1f, 0xe0, 0x8e, 0x75, 0xcc, 0xb8, 0x21, 0xda, 0x45,
	0xfb, 0xed, 0x45, 0x5c, 0xf4, 0x7b, 0xb8, 0xcd, 0xd5, 0x72, 0xd8, 0x0a, 0x9a, 0x1f, 0x4f, 0x76,
	0x9e, 0x5e, 0x1f, 0x27, 0x5b, 0x78, 0xa3, 0x76, 0xc7, 0x6f, 0x4a, 0xc5, 0x4c, 0x0e, 0x1b, 0x98,
	0x9d, 0x06, 0x66, 0x67, 0x1d, 0xb3, 0xa2, 0x4c, 0x2f, 0xf1, 0xbf, 0x53, 0xc8, 0xd3, 0x8c, 0x37,
	0x40, 0x51, 0x03, 0x14, 0x45, 0xe8, 0xd8, 0x43, 0x87, 0x78, 0xb3, 0x06, 0xad, 0x61, 0x66, 0x0f,
	0xe8, 0xf9, 0x63, 0x82, 0x5e, 0xee, 0xdf, 0xde, 0xbb, 0xad, 0x5e, 0x0b, 0x6f, 0x4b, 0x20, 0x21,
	0xa9, 0x0d, 0xac, 0x4a, 0xf2, 0xa3, 0x81, 0x19, 0x0e, 0x5b, 0xe6, 0xfe, 0xd5, 0xe6, 0xe8, 0xea,
	0x6c, 0x6d, 0xdb, 0xfa, 0x56, 0x7c, 0x35, 0x4e, 0xbc, 0x4c, 0xe4, 0x5f, 0xa5, 0xa7, 0xdd, 0x50,
	0x44, 0xf2, 0x19, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x66, 0x2b, 0x96, 0x40, 0x02, 0x00, 0x00,
}

func (this *Int64Range) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Int64Range)
	if !ok {
		that2, ok := that.(Int64Range)
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
	if this.Start != that1.Start {
		return false
	}
	if this.End != that1.End {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Int32Range) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Int32Range)
	if !ok {
		that2, ok := that.(Int32Range)
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
	if this.Start != that1.Start {
		return false
	}
	if this.End != that1.End {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DoubleRange) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DoubleRange)
	if !ok {
		that2, ok := that.(DoubleRange)
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
	if this.Start != that1.Start {
		return false
	}
	if this.End != that1.End {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
