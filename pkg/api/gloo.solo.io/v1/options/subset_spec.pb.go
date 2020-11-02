// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/subset_spec.proto

package options

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

type SubsetSpec struct {
	Selectors            []*Selector `protobuf:"bytes,1,rep,name=selectors,proto3" json:"selectors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SubsetSpec) Reset()         { *m = SubsetSpec{} }
func (m *SubsetSpec) String() string { return proto.CompactTextString(m) }
func (*SubsetSpec) ProtoMessage()    {}
func (*SubsetSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_694a393eae582d53, []int{0}
}
func (m *SubsetSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubsetSpec.Unmarshal(m, b)
}
func (m *SubsetSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubsetSpec.Marshal(b, m, deterministic)
}
func (m *SubsetSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubsetSpec.Merge(m, src)
}
func (m *SubsetSpec) XXX_Size() int {
	return xxx_messageInfo_SubsetSpec.Size(m)
}
func (m *SubsetSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_SubsetSpec.DiscardUnknown(m)
}

var xxx_messageInfo_SubsetSpec proto.InternalMessageInfo

func (m *SubsetSpec) GetSelectors() []*Selector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

type Selector struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Selector) Reset()         { *m = Selector{} }
func (m *Selector) String() string { return proto.CompactTextString(m) }
func (*Selector) ProtoMessage()    {}
func (*Selector) Descriptor() ([]byte, []int) {
	return fileDescriptor_694a393eae582d53, []int{1}
}
func (m *Selector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Selector.Unmarshal(m, b)
}
func (m *Selector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Selector.Marshal(b, m, deterministic)
}
func (m *Selector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Selector.Merge(m, src)
}
func (m *Selector) XXX_Size() int {
	return xxx_messageInfo_Selector.Size(m)
}
func (m *Selector) XXX_DiscardUnknown() {
	xxx_messageInfo_Selector.DiscardUnknown(m)
}

var xxx_messageInfo_Selector proto.InternalMessageInfo

func (m *Selector) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func init() {
	proto.RegisterType((*SubsetSpec)(nil), "options.gloo.solo.io.SubsetSpec")
	proto.RegisterType((*Selector)(nil), "options.gloo.solo.io.Selector")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/subset_spec.proto", fileDescriptor_694a393eae582d53)
}

var fileDescriptor_694a393eae582d53 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcf, 0xc9, 0xd7, 0xcd, 0xcc, 0x87, 0xd0,
	0x89, 0x05, 0x99, 0xc5, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0xe9, 0x39, 0xf9, 0xf9, 0x10, 0xa2, 0xcc,
	0x50, 0x3f, 0xbf, 0xa0, 0x24, 0x33, 0x3f, 0xaf, 0x58, 0xbf, 0xb8, 0x34, 0xa9, 0x38, 0xb5, 0x24,
	0xbe, 0xb8, 0x20, 0x35, 0x59, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x48, 0x04, 0x2a, 0xa5, 0x07,
	0x52, 0xaa, 0x07, 0x32, 0x42, 0x2f, 0x33, 0x5f, 0x4a, 0x24, 0x3d, 0x3f, 0x3d, 0x1f, 0xac, 0x40,
	0x1f, 0xc4, 0x82, 0xa8, 0x95, 0x12, 0x4a, 0xad, 0x28, 0x81, 0x08, 0xa6, 0x56, 0x94, 0x40, 0xc4,
	0x94, 0xbc, 0xb8, 0xb8, 0x82, 0xc1, 0x86, 0x06, 0x17, 0xa4, 0x26, 0x0b, 0xd9, 0x70, 0x71, 0x16,
	0xa7, 0xe6, 0xa4, 0x26, 0x97, 0xe4, 0x17, 0x15, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0xc9,
	0xe9, 0x61, 0xb3, 0x41, 0x2f, 0x18, 0xaa, 0x2c, 0x08, 0xa1, 0x41, 0x49, 0x8e, 0x8b, 0x03, 0x26,
	0x2c, 0x24, 0xc4, 0xc5, 0x92, 0x9d, 0x5a, 0x09, 0x31, 0x84, 0x33, 0x08, 0xcc, 0x76, 0x72, 0xdb,
	0xf1, 0x95, 0x85, 0x71, 0xc5, 0x23, 0x39, 0xc6, 0x28, 0x1b, 0xbc, 0xde, 0x2f, 0xc8, 0x4e, 0x87,
	0x07, 0x01, 0xcc, 0x42, 0xa4, 0x50, 0x48, 0x62, 0x03, 0x3b, 0xdd, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0xa4, 0x43, 0x3d, 0xf8, 0x45, 0x01, 0x00, 0x00,
}

func (this *SubsetSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SubsetSpec)
	if !ok {
		that2, ok := that.(SubsetSpec)
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
	if len(this.Selectors) != len(that1.Selectors) {
		return false
	}
	for i := range this.Selectors {
		if !this.Selectors[i].Equal(that1.Selectors[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Selector) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Selector)
	if !ok {
		that2, ok := that.(Selector)
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
	if len(this.Keys) != len(that1.Keys) {
		return false
	}
	for i := range this.Keys {
		if this.Keys[i] != that1.Keys[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
