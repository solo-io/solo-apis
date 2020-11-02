// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/subset.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

type Subset struct {
	Values               map[string]string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Subset) Reset()         { *m = Subset{} }
func (m *Subset) String() string { return proto.CompactTextString(m) }
func (*Subset) ProtoMessage()    {}
func (*Subset) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c1ff6166a3be1c8, []int{0}
}
func (m *Subset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subset.Unmarshal(m, b)
}
func (m *Subset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subset.Marshal(b, m, deterministic)
}
func (m *Subset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subset.Merge(m, src)
}
func (m *Subset) XXX_Size() int {
	return xxx_messageInfo_Subset.Size(m)
}
func (m *Subset) XXX_DiscardUnknown() {
	xxx_messageInfo_Subset.DiscardUnknown(m)
}

var xxx_messageInfo_Subset proto.InternalMessageInfo

func (m *Subset) GetValues() map[string]string {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*Subset)(nil), "gloo.solo.io.Subset")
	proto.RegisterMapType((map[string]string)(nil), "gloo.solo.io.Subset.ValuesEntry")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo/gloo/v1/subset.proto", fileDescriptor_1c1ff6166a3be1c8)
}

var fileDescriptor_1c1ff6166a3be1c8 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcf, 0xc9, 0xd7, 0xcd, 0xcc, 0x87, 0xd0,
	0x89, 0x05, 0x99, 0xc5, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0xe9, 0x39, 0xf9, 0xf9, 0x10, 0xa2, 0xcc,
	0x50, 0xbf, 0xb8, 0x34, 0xa9, 0x38, 0xb5, 0x44, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x88, 0x07,
	0x24, 0xaa, 0x07, 0x52, 0xad, 0x97, 0x99, 0x2f, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x96, 0xd0,
	0x07, 0xb1, 0x20, 0x6a, 0xa4, 0x84, 0x52, 0x2b, 0x4a, 0x20, 0x82, 0xa9, 0x15, 0x50, 0x7d, 0x4a,
	0xb5, 0x5c, 0x6c, 0xc1, 0x60, 0x73, 0x84, 0x2c, 0xb8, 0xd8, 0xca, 0x12, 0x73, 0x4a, 0x53, 0x8b,
	0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x14, 0xf4, 0x90, 0x8d, 0xd4, 0x83, 0xa8, 0xd2, 0x0b,
	0x03, 0x2b, 0x71, 0xcd, 0x2b, 0x29, 0xaa, 0x0c, 0x82, 0xaa, 0x97, 0xb2, 0xe4, 0xe2, 0x46, 0x12,
	0x16, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31,
	0x85, 0x44, 0xb8, 0x58, 0xc1, 0x4a, 0x25, 0x98, 0xc0, 0x62, 0x10, 0x8e, 0x15, 0x93, 0x05, 0xa3,
	0x93, 0xdd, 0x8e, 0xaf, 0x2c, 0x8c, 0x2b, 0x1e, 0xc9, 0x31, 0x46, 0x99, 0xe0, 0xf5, 0x7c, 0x41,
	0x76, 0x3a, 0x3c, 0x00, 0x60, 0x4e, 0xd2, 0x2f, 0x33, 0x4c, 0x62, 0x03, 0xfb, 0xc2, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x52, 0x6d, 0x7b, 0xbc, 0x3b, 0x01, 0x00, 0x00,
}

func (this *Subset) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Subset)
	if !ok {
		that2, ok := that.(Subset)
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
	if len(this.Values) != len(that1.Values) {
		return false
	}
	for i := range this.Values {
		if this.Values[i] != that1.Values[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
