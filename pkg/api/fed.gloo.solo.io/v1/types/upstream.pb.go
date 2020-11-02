// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-fed/fed.gloo/v1/upstream.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v11 "github.com/solo-io/solo-apis/pkg/api/fed.solo.io/core/v1"
	v1 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1"
	v1alpha1 "github.com/solo-io/solo-apis/pkg/api/multicluster.solo.io/v1alpha1"
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

type FederatedUpstreamSpec struct {
	Template             *FederatedUpstreamSpec_Template `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
	Placement            *v1alpha1.Placement             `protobuf:"bytes,2,opt,name=placement,proto3" json:"placement,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *FederatedUpstreamSpec) Reset()         { *m = FederatedUpstreamSpec{} }
func (m *FederatedUpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*FederatedUpstreamSpec) ProtoMessage()    {}
func (*FederatedUpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_18e980e6de8364bd, []int{0}
}
func (m *FederatedUpstreamSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FederatedUpstreamSpec.Unmarshal(m, b)
}
func (m *FederatedUpstreamSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FederatedUpstreamSpec.Marshal(b, m, deterministic)
}
func (m *FederatedUpstreamSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FederatedUpstreamSpec.Merge(m, src)
}
func (m *FederatedUpstreamSpec) XXX_Size() int {
	return xxx_messageInfo_FederatedUpstreamSpec.Size(m)
}
func (m *FederatedUpstreamSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_FederatedUpstreamSpec.DiscardUnknown(m)
}

var xxx_messageInfo_FederatedUpstreamSpec proto.InternalMessageInfo

func (m *FederatedUpstreamSpec) GetTemplate() *FederatedUpstreamSpec_Template {
	if m != nil {
		return m.Template
	}
	return nil
}

func (m *FederatedUpstreamSpec) GetPlacement() *v1alpha1.Placement {
	if m != nil {
		return m.Placement
	}
	return nil
}

type FederatedUpstreamSpec_Template struct {
	Spec                 *v1.UpstreamSpec      `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	Metadata             *v11.TemplateMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FederatedUpstreamSpec_Template) Reset()         { *m = FederatedUpstreamSpec_Template{} }
func (m *FederatedUpstreamSpec_Template) String() string { return proto.CompactTextString(m) }
func (*FederatedUpstreamSpec_Template) ProtoMessage()    {}
func (*FederatedUpstreamSpec_Template) Descriptor() ([]byte, []int) {
	return fileDescriptor_18e980e6de8364bd, []int{0, 0}
}
func (m *FederatedUpstreamSpec_Template) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FederatedUpstreamSpec_Template.Unmarshal(m, b)
}
func (m *FederatedUpstreamSpec_Template) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FederatedUpstreamSpec_Template.Marshal(b, m, deterministic)
}
func (m *FederatedUpstreamSpec_Template) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FederatedUpstreamSpec_Template.Merge(m, src)
}
func (m *FederatedUpstreamSpec_Template) XXX_Size() int {
	return xxx_messageInfo_FederatedUpstreamSpec_Template.Size(m)
}
func (m *FederatedUpstreamSpec_Template) XXX_DiscardUnknown() {
	xxx_messageInfo_FederatedUpstreamSpec_Template.DiscardUnknown(m)
}

var xxx_messageInfo_FederatedUpstreamSpec_Template proto.InternalMessageInfo

func (m *FederatedUpstreamSpec_Template) GetSpec() *v1.UpstreamSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *FederatedUpstreamSpec_Template) GetMetadata() *v11.TemplateMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type FederatedUpstreamStatus struct {
	PlacementStatus      *v11.PlacementStatus `protobuf:"bytes,1,opt,name=placement_status,json=placementStatus,proto3" json:"placement_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FederatedUpstreamStatus) Reset()         { *m = FederatedUpstreamStatus{} }
func (m *FederatedUpstreamStatus) String() string { return proto.CompactTextString(m) }
func (*FederatedUpstreamStatus) ProtoMessage()    {}
func (*FederatedUpstreamStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_18e980e6de8364bd, []int{1}
}
func (m *FederatedUpstreamStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FederatedUpstreamStatus.Unmarshal(m, b)
}
func (m *FederatedUpstreamStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FederatedUpstreamStatus.Marshal(b, m, deterministic)
}
func (m *FederatedUpstreamStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FederatedUpstreamStatus.Merge(m, src)
}
func (m *FederatedUpstreamStatus) XXX_Size() int {
	return xxx_messageInfo_FederatedUpstreamStatus.Size(m)
}
func (m *FederatedUpstreamStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_FederatedUpstreamStatus.DiscardUnknown(m)
}

var xxx_messageInfo_FederatedUpstreamStatus proto.InternalMessageInfo

func (m *FederatedUpstreamStatus) GetPlacementStatus() *v11.PlacementStatus {
	if m != nil {
		return m.PlacementStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*FederatedUpstreamSpec)(nil), "fed.gloo.solo.io.FederatedUpstreamSpec")
	proto.RegisterType((*FederatedUpstreamSpec_Template)(nil), "fed.gloo.solo.io.FederatedUpstreamSpec.Template")
	proto.RegisterType((*FederatedUpstreamStatus)(nil), "fed.gloo.solo.io.FederatedUpstreamStatus")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo-fed/fed.gloo/v1/upstream.proto", fileDescriptor_18e980e6de8364bd)
}

var fileDescriptor_18e980e6de8364bd = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xbf, 0x4e, 0xeb, 0x30,
	0x14, 0xc6, 0x95, 0xaa, 0xba, 0xea, 0xf5, 0x1d, 0x6e, 0x15, 0x81, 0xa8, 0x32, 0xf0, 0xa7, 0x13,
	0x0c, 0xb5, 0x69, 0x99, 0x18, 0xe8, 0x80, 0x10, 0x62, 0x28, 0x12, 0x0a, 0xb0, 0xb0, 0x20, 0x37,
	0x39, 0x4d, 0xa3, 0x26, 0xb5, 0x15, 0x9f, 0x54, 0x85, 0xa7, 0x61, 0xe4, 0x11, 0x78, 0x1e, 0xde,
	0x81, 0x1d, 0xc5, 0xb1, 0x03, 0x69, 0x0b, 0x53, 0x92, 0xe3, 0xf3, 0xfb, 0xce, 0xf7, 0xf9, 0x84,
	0x5c, 0x44, 0x31, 0x4e, 0xf3, 0x31, 0x0d, 0x44, 0xca, 0x94, 0x48, 0x44, 0x2f, 0x16, 0xe5, 0x93,
	0xcb, 0x58, 0x31, 0x2e, 0x63, 0x16, 0x25, 0x42, 0xf4, 0x26, 0x10, 0xb2, 0x09, 0x84, 0xb4, 0xf8,
	0x60, 0x8b, 0x3e, 0xcb, 0xa5, 0xc2, 0x0c, 0x78, 0x4a, 0x65, 0x26, 0x50, 0xb8, 0x6d, 0x7b, 0x46,
	0x0b, 0x96, 0xc6, 0xc2, 0xdb, 0x8a, 0x44, 0x24, 0xf4, 0x21, 0x2b, 0xde, 0xca, 0x3e, 0xcf, 0x85,
	0x25, 0x96, 0x45, 0x58, 0xa2, 0xa9, 0x9d, 0xd6, 0xc7, 0xa9, 0xd9, 0x62, 0xd0, 0x83, 0x39, 0x42,
	0x26, 0xb3, 0x58, 0x01, 0x5b, 0xf4, 0x79, 0x22, 0xa7, 0xbc, 0xcf, 0xd2, 0x3c, 0xc1, 0x38, 0x48,
	0x72, 0x85, 0x90, 0x19, 0x74, 0xf0, 0xb3, 0x53, 0x16, 0x88, 0xac, 0x80, 0x99, 0x4c, 0x78, 0x00,
	0x29, 0xcc, 0xed, 0xb8, 0xa3, 0x75, 0x86, 0x6d, 0x4e, 0xd5, 0x7d, 0x69, 0x90, 0xed, 0x4b, 0x08,
	0x21, 0xe3, 0x08, 0xe1, 0xbd, 0x39, 0xbb, 0x95, 0x10, 0xb8, 0x23, 0xd2, 0x42, 0x48, 0x65, 0xc2,
	0x11, 0x3a, 0xce, 0xbe, 0x73, 0xf8, 0x6f, 0x70, 0x4c, 0x57, 0xaf, 0x80, 0x6e, 0x44, 0xe9, 0x9d,
	0xe1, 0xfc, 0x4a, 0xc1, 0x3d, 0x23, 0x7f, 0x2b, 0x97, 0x9d, 0x86, 0x96, 0xdb, 0xa3, 0xb5, 0xb8,
	0x56, 0xf2, 0xc6, 0xb6, 0xf9, 0x5f, 0x84, 0xf7, 0x4c, 0x5a, 0x56, 0xd4, 0xa5, 0xa4, 0xa9, 0x24,
	0x04, 0xc6, 0x94, 0x57, 0x37, 0xf4, 0xdd, 0x87, 0xaf, 0xfb, 0xdc, 0x21, 0x69, 0xa5, 0x80, 0x3c,
	0xe4, 0xc8, 0xcd, 0xe4, 0x2e, 0x2d, 0x6e, 0x4e, 0xa7, 0xb1, 0x9c, 0x55, 0xbf, 0x36, 0x9d, 0x7e,
	0xc5, 0x74, 0x23, 0xb2, 0xb3, 0x1e, 0x13, 0x39, 0xe6, 0xca, 0x1d, 0x91, 0x76, 0xe5, 0xf1, 0x51,
	0xe9, 0x9a, 0xb1, 0x75, 0xb0, 0x3e, 0xa2, 0x0a, 0x56, 0xc2, 0xfe, 0x7f, 0x59, 0x2f, 0x9c, 0x5f,
	0xbd, 0x7d, 0x34, 0x9d, 0xd7, 0xf7, 0x5d, 0xe7, 0x61, 0xf8, 0xeb, 0x1f, 0x2b, 0x67, 0x91, 0xde,
	0xeb, 0xea, 0x36, 0x8a, 0xf5, 0xe2, 0x93, 0x04, 0x35, 0xfe, 0xa3, 0x97, 0x7b, 0xf2, 0x19, 0x00,
	0x00, 0xff, 0xff, 0x26, 0x27, 0x63, 0xcb, 0xfa, 0x02, 0x00, 0x00,
}

func (this *FederatedUpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FederatedUpstreamSpec)
	if !ok {
		that2, ok := that.(FederatedUpstreamSpec)
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
	if !this.Template.Equal(that1.Template) {
		return false
	}
	if !this.Placement.Equal(that1.Placement) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *FederatedUpstreamSpec_Template) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FederatedUpstreamSpec_Template)
	if !ok {
		that2, ok := that.(FederatedUpstreamSpec_Template)
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
	if !this.Spec.Equal(that1.Spec) {
		return false
	}
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *FederatedUpstreamStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FederatedUpstreamStatus)
	if !ok {
		that2, ok := that.(FederatedUpstreamStatus)
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
	if !this.PlacementStatus.Equal(that1.PlacementStatus) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
