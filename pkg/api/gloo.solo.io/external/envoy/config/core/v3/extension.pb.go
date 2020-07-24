// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/config/core/v3/extension.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/udpa/annotations"
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

// Message type for extension configuration.
// [#next-major-version: revisit all existing typed_config that doesn't use this wrapper.].
type TypedExtensionConfig struct {
	// The name of an extension. This is not used to select the extension, instead
	// it serves the role of an opaque identifier.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The typed config for the extension. The type URL will be used to identify
	// the extension. In the case that the type URL is *udpa.type.v1.TypedStruct*,
	// the inner type URL of *TypedStruct* will be utilized. See the
	// :ref:`extension configuration overview
	// <config_overview_extension_configuration>` for further details.
	TypedConfig          *types.Any `protobuf:"bytes,2,opt,name=typed_config,json=typedConfig,proto3" json:"typed_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TypedExtensionConfig) Reset()         { *m = TypedExtensionConfig{} }
func (m *TypedExtensionConfig) String() string { return proto.CompactTextString(m) }
func (*TypedExtensionConfig) ProtoMessage()    {}
func (*TypedExtensionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c662bf2ca26fbd1a, []int{0}
}
func (m *TypedExtensionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypedExtensionConfig.Unmarshal(m, b)
}
func (m *TypedExtensionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypedExtensionConfig.Marshal(b, m, deterministic)
}
func (m *TypedExtensionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypedExtensionConfig.Merge(m, src)
}
func (m *TypedExtensionConfig) XXX_Size() int {
	return xxx_messageInfo_TypedExtensionConfig.Size(m)
}
func (m *TypedExtensionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TypedExtensionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TypedExtensionConfig proto.InternalMessageInfo

func (m *TypedExtensionConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TypedExtensionConfig) GetTypedConfig() *types.Any {
	if m != nil {
		return m.TypedConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*TypedExtensionConfig)(nil), "envoy.config.core.v3.TypedExtensionConfig")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/config/core/v3/extension.proto", fileDescriptor_c662bf2ca26fbd1a)
}

var fileDescriptor_c662bf2ca26fbd1a = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4a, 0x3b, 0x31,
	0x14, 0xc6, 0xc9, 0xd0, 0x7f, 0xff, 0x75, 0x2a, 0x52, 0xca, 0x80, 0xb5, 0xa2, 0x94, 0xae, 0xba,
	0xf1, 0x3d, 0xb0, 0x27, 0xe8, 0x88, 0x6b, 0x4b, 0x71, 0x21, 0x6e, 0x24, 0x6d, 0xd3, 0x18, 0x9c,
	0xe6, 0x85, 0x99, 0xcc, 0xd0, 0xd9, 0x79, 0x07, 0x6f, 0xe0, 0xca, 0x33, 0x78, 0x02, 0xb7, 0x5e,
	0xc1, 0x63, 0x74, 0x25, 0x93, 0x74, 0xba, 0x12, 0xdd, 0x24, 0x8f, 0xf7, 0xbe, 0x2f, 0xf9, 0xf1,
	0xbe, 0xf0, 0x4e, 0x2a, 0xfb, 0x98, 0xcf, 0x61, 0x41, 0x6b, 0xcc, 0x28, 0xa1, 0x0b, 0x45, 0xfe,
	0xe6, 0x46, 0x65, 0xc8, 0x8d, 0x42, 0x99, 0x10, 0xf9, 0x43, 0x6c, 0xac, 0x48, 0x35, 0x4f, 0x50,
	0xe8, 0x82, 0x4a, 0x5c, 0x90, 0x5e, 0x29, 0x89, 0x0b, 0x4a, 0x05, 0x16, 0x63, 0x37, 0xd5, 0x99,
	0x22, 0x0d, 0x26, 0x25, 0x4b, 0xdd, 0xc8, 0xa9, 0xc0, 0xab, 0xa0, 0x52, 0x41, 0x31, 0xee, 0x9f,
	0x48, 0x22, 0x99, 0x08, 0x74, 0x9a, 0x79, 0xbe, 0x42, 0xae, 0x4b, 0x6f, 0xe8, 0x9f, 0xe5, 0x4b,
	0xc3, 0x91, 0x6b, 0x4d, 0x96, 0x5b, 0x45, 0x3a, 0xc3, 0xcc, 0x72, 0x9b, 0x67, 0xbb, 0xf1, 0x71,
	0xc1, 0x13, 0xb5, 0xe4, 0x56, 0x60, 0x5d, 0xec, 0x06, 0x91, 0x24, 0x49, 0xae, 0xc4, 0xaa, 0xf2,
	0xdd, 0x61, 0x11, 0x46, 0xb7, 0xa5, 0x11, 0xcb, 0xeb, 0x1a, 0xeb, 0xca, 0x81, 0x74, 0x4f, 0xc3,
	0x86, 0xe6, 0x6b, 0xd1, 0x63, 0x03, 0x36, 0x3a, 0x88, 0xff, 0x6f, 0xe3, 0x46, 0x1a, 0x74, 0xd8,
	0xcc, 0x35, 0xbb, 0x93, 0xf0, 0xd0, 0x56, 0xa6, 0x07, 0x4f, 0xdd, 0x0b, 0x06, 0x6c, 0xd4, 0xbe,
	0x8c, 0xc0, 0x43, 0x43, 0x0d, 0x0d, 0x13, 0x5d, 0xc6, 0xad, 0x6d, 0xfc, 0xef, 0x95, 0x05, 0x2d,
	0x36, 0x6b, 0x3b, 0x8f, 0x7f, 0x3f, 0x7e, 0x61, 0x6f, 0x5f, 0xe7, 0xec, 0xfd, 0xf9, 0xe3, 0xb3,
	0x19, 0x74, 0x82, 0x70, 0xa8, 0x08, 0xdc, 0x22, 0x4c, 0x4a, 0x9b, 0x12, 0x7e, 0xda, 0x49, 0x7c,
	0xb4, 0x67, 0x9c, 0x56, 0x1f, 0x4c, 0xd9, 0xfd, 0xcd, 0xaf, 0xa9, 0x98, 0x27, 0xb9, 0x4f, 0x06,
	0xaa, 0x36, 0xa8, 0x3f, 0xc2, 0x99, 0x37, 0x1d, 0xfa, 0xf8, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x44,
	0xad, 0xf9, 0x60, 0xef, 0x01, 0x00, 0x00,
}

func (this *TypedExtensionConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TypedExtensionConfig)
	if !ok {
		that2, ok := that.(TypedExtensionConfig)
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
	if this.Name != that1.Name {
		return false
	}
	if !this.TypedConfig.Equal(that1.TypedConfig) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
