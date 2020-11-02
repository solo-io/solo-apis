// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-fed/fed.gateway/v1/gateway.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v11 "github.com/solo-io/solo-apis/pkg/api/fed.solo.io/core/v1"
	v1 "github.com/solo-io/solo-apis/pkg/api/gateway.solo.io/v1"
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

type FederatedGatewaySpec struct {
	Template             *FederatedGatewaySpec_Template `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
	Placement            *v1alpha1.Placement            `protobuf:"bytes,2,opt,name=placement,proto3" json:"placement,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *FederatedGatewaySpec) Reset()         { *m = FederatedGatewaySpec{} }
func (m *FederatedGatewaySpec) String() string { return proto.CompactTextString(m) }
func (*FederatedGatewaySpec) ProtoMessage()    {}
func (*FederatedGatewaySpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_c27d1ae31fe6011e, []int{0}
}
func (m *FederatedGatewaySpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FederatedGatewaySpec.Unmarshal(m, b)
}
func (m *FederatedGatewaySpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FederatedGatewaySpec.Marshal(b, m, deterministic)
}
func (m *FederatedGatewaySpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FederatedGatewaySpec.Merge(m, src)
}
func (m *FederatedGatewaySpec) XXX_Size() int {
	return xxx_messageInfo_FederatedGatewaySpec.Size(m)
}
func (m *FederatedGatewaySpec) XXX_DiscardUnknown() {
	xxx_messageInfo_FederatedGatewaySpec.DiscardUnknown(m)
}

var xxx_messageInfo_FederatedGatewaySpec proto.InternalMessageInfo

func (m *FederatedGatewaySpec) GetTemplate() *FederatedGatewaySpec_Template {
	if m != nil {
		return m.Template
	}
	return nil
}

func (m *FederatedGatewaySpec) GetPlacement() *v1alpha1.Placement {
	if m != nil {
		return m.Placement
	}
	return nil
}

type FederatedGatewaySpec_Template struct {
	Spec                 *v1.GatewaySpec       `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	Metadata             *v11.TemplateMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FederatedGatewaySpec_Template) Reset()         { *m = FederatedGatewaySpec_Template{} }
func (m *FederatedGatewaySpec_Template) String() string { return proto.CompactTextString(m) }
func (*FederatedGatewaySpec_Template) ProtoMessage()    {}
func (*FederatedGatewaySpec_Template) Descriptor() ([]byte, []int) {
	return fileDescriptor_c27d1ae31fe6011e, []int{0, 0}
}
func (m *FederatedGatewaySpec_Template) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FederatedGatewaySpec_Template.Unmarshal(m, b)
}
func (m *FederatedGatewaySpec_Template) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FederatedGatewaySpec_Template.Marshal(b, m, deterministic)
}
func (m *FederatedGatewaySpec_Template) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FederatedGatewaySpec_Template.Merge(m, src)
}
func (m *FederatedGatewaySpec_Template) XXX_Size() int {
	return xxx_messageInfo_FederatedGatewaySpec_Template.Size(m)
}
func (m *FederatedGatewaySpec_Template) XXX_DiscardUnknown() {
	xxx_messageInfo_FederatedGatewaySpec_Template.DiscardUnknown(m)
}

var xxx_messageInfo_FederatedGatewaySpec_Template proto.InternalMessageInfo

func (m *FederatedGatewaySpec_Template) GetSpec() *v1.GatewaySpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *FederatedGatewaySpec_Template) GetMetadata() *v11.TemplateMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type FederatedGatewayStatus struct {
	PlacementStatus      *v11.PlacementStatus `protobuf:"bytes,1,opt,name=placement_status,json=placementStatus,proto3" json:"placement_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FederatedGatewayStatus) Reset()         { *m = FederatedGatewayStatus{} }
func (m *FederatedGatewayStatus) String() string { return proto.CompactTextString(m) }
func (*FederatedGatewayStatus) ProtoMessage()    {}
func (*FederatedGatewayStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c27d1ae31fe6011e, []int{1}
}
func (m *FederatedGatewayStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FederatedGatewayStatus.Unmarshal(m, b)
}
func (m *FederatedGatewayStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FederatedGatewayStatus.Marshal(b, m, deterministic)
}
func (m *FederatedGatewayStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FederatedGatewayStatus.Merge(m, src)
}
func (m *FederatedGatewayStatus) XXX_Size() int {
	return xxx_messageInfo_FederatedGatewayStatus.Size(m)
}
func (m *FederatedGatewayStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_FederatedGatewayStatus.DiscardUnknown(m)
}

var xxx_messageInfo_FederatedGatewayStatus proto.InternalMessageInfo

func (m *FederatedGatewayStatus) GetPlacementStatus() *v11.PlacementStatus {
	if m != nil {
		return m.PlacementStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*FederatedGatewaySpec)(nil), "fed.gateway.solo.io.FederatedGatewaySpec")
	proto.RegisterType((*FederatedGatewaySpec_Template)(nil), "fed.gateway.solo.io.FederatedGatewaySpec.Template")
	proto.RegisterType((*FederatedGatewayStatus)(nil), "fed.gateway.solo.io.FederatedGatewayStatus")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo-fed/fed.gateway/v1/gateway.proto", fileDescriptor_c27d1ae31fe6011e)
}

var fileDescriptor_c27d1ae31fe6011e = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xbf, 0xee, 0xd3, 0x30,
	0x10, 0xc7, 0x95, 0xaa, 0x42, 0xc5, 0x0c, 0xa0, 0x50, 0xa1, 0x2a, 0x42, 0x05, 0x3a, 0x21, 0xa1,
	0xda, 0x34, 0x4c, 0x0c, 0x20, 0xc1, 0x50, 0x06, 0xfe, 0x08, 0x05, 0x26, 0x16, 0xe4, 0x26, 0x17,
	0x37, 0xaa, 0x53, 0x5b, 0xf6, 0xa5, 0x7f, 0x24, 0xde, 0x86, 0x85, 0x47, 0xe0, 0x79, 0x78, 0x07,
	0x76, 0x14, 0xc7, 0x09, 0xad, 0xd2, 0xfe, 0xa6, 0x38, 0xe7, 0xfb, 0x7c, 0xef, 0x7b, 0x77, 0x26,
	0x4b, 0x51, 0xe0, 0xba, 0x5a, 0xd1, 0x54, 0x95, 0xcc, 0x2a, 0xa9, 0xe6, 0x85, 0x6a, 0xbe, 0x5c,
	0x17, 0x96, 0x71, 0x5d, 0x30, 0x21, 0x95, 0x9a, 0xe7, 0x90, 0xb1, 0x1c, 0x32, 0x2a, 0x38, 0xc2,
	0x9e, 0x1f, 0xd9, 0x6e, 0xc1, 0xfc, 0x91, 0x6a, 0xa3, 0x50, 0x85, 0xf7, 0x4f, 0x6e, 0x69, 0x2d,
	0x40, 0x0b, 0x15, 0x8d, 0x85, 0x12, 0xca, 0xdd, 0xb3, 0xfa, 0xd4, 0xa4, 0x46, 0x21, 0x1c, 0xb0,
	0x09, 0xc2, 0x01, 0x7d, 0xec, 0xe5, 0x79, 0x4d, 0xbb, 0xd9, 0xc5, 0x73, 0xd8, 0x22, 0x18, 0x6d,
	0x0a, 0x0b, 0x6c, 0xb7, 0xe0, 0x52, 0xaf, 0xf9, 0x82, 0x95, 0x95, 0xc4, 0x22, 0x95, 0x95, 0x45,
	0x30, 0x1e, 0x8d, 0xaf, 0xdb, 0x65, 0xa9, 0x32, 0x35, 0xcc, 0xb4, 0xe4, 0x29, 0x94, 0xb0, 0x6d,
	0xcb, 0x3d, 0xeb, 0x33, 0xec, 0x5a, 0x6b, 0xd1, 0x54, 0x28, 0x25, 0x24, 0x30, 0xf7, 0xb7, 0xaa,
	0x72, 0xb6, 0x37, 0x5c, 0x6b, 0x30, 0xb6, 0xb9, 0x9f, 0xfd, 0x1c, 0x90, 0xf1, 0x12, 0x32, 0x30,
	0x1c, 0x21, 0x7b, 0xd7, 0xa0, 0x5f, 0x34, 0xa4, 0xe1, 0x27, 0x32, 0x42, 0x28, 0xb5, 0xe4, 0x08,
	0x93, 0xe0, 0x71, 0xf0, 0xf4, 0x4e, 0x1c, 0xd3, 0x0b, 0x63, 0xa2, 0x97, 0x60, 0xfa, 0xd5, 0x93,
	0x49, 0xa7, 0x11, 0xbe, 0x22, 0xb7, 0xbb, 0x46, 0x26, 0x03, 0x27, 0xf8, 0x88, 0x9e, 0x4d, 0xa4,
	0x55, 0xfc, 0xdc, 0xa6, 0x25, 0xff, 0x89, 0xe8, 0x07, 0x19, 0xb5, 0xa2, 0xe1, 0x73, 0x32, 0xb4,
	0x1a, 0x52, 0x6f, 0xeb, 0x61, 0xcf, 0xd2, 0x89, 0x93, 0xc4, 0x65, 0x86, 0xaf, 0xc9, 0xa8, 0x04,
	0xe4, 0x19, 0x47, 0xee, 0x6b, 0xcf, 0x68, 0x3d, 0x5e, 0xd7, 0x51, 0x8b, 0xb5, 0xfa, 0x1f, 0x7d,
	0x66, 0xd2, 0x31, 0xb3, 0x9c, 0x3c, 0xe8, 0xf5, 0x89, 0x1c, 0x2b, 0x1b, 0x7e, 0x20, 0xf7, 0x3a,
	0x93, 0xdf, 0xad, 0x8b, 0x79, 0x5f, 0x4f, 0xfa, 0x15, 0xba, 0xce, 0x1a, 0x38, 0xb9, 0xab, 0xcf,
	0x03, 0x6f, 0xdf, 0xff, 0xfe, 0x3b, 0x0c, 0x7e, 0xfd, 0x99, 0x06, 0xdf, 0xde, 0xdc, 0xf8, 0xb4,
	0xf5, 0x46, 0xb8, 0xdd, 0x5f, 0x58, 0x48, 0xfd, 0x04, 0xf0, 0xa8, 0xc1, 0xae, 0x6e, 0xb9, 0x0d,
	0xbf, 0xf8, 0x17, 0x00, 0x00, 0xff, 0xff, 0xea, 0xb4, 0x64, 0x07, 0x26, 0x03, 0x00, 0x00,
}

func (this *FederatedGatewaySpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FederatedGatewaySpec)
	if !ok {
		that2, ok := that.(FederatedGatewaySpec)
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
func (this *FederatedGatewaySpec_Template) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FederatedGatewaySpec_Template)
	if !ok {
		that2, ok := that.(FederatedGatewaySpec_Template)
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
func (this *FederatedGatewayStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FederatedGatewayStatus)
	if !ok {
		that2, ok := that.(FederatedGatewayStatus)
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
