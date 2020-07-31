// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/headers/headers.proto

package headers

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// This plugin provides configuration options to append and remove headers from
// requests and responses
// HeaderManipulation can be specified on routes, virtual hosts, or weighted destinations
type HeaderManipulation struct {
	// Specifies a list of HTTP headers that should be added to each request
	// handled by this route or virtual host. For more information, including
	// details on header value syntax, see the
	// [Envoy documentation](https://www.envoyproxy.io/docs/envoy/latest/configuration/http_conn_man/headers#config-http-conn-man-headers-custom-request-headers) .
	RequestHeadersToAdd []*HeaderValueOption `protobuf:"bytes,1,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	// Specifies a list of HTTP headers that should be removed from each request
	// handled by this route or virtual host.
	RequestHeadersToRemove []string `protobuf:"bytes,2,rep,name=request_headers_to_remove,json=requestHeadersToRemove,proto3" json:"request_headers_to_remove,omitempty"`
	// Specifies a list of HTTP headers that should be added to each response
	// handled by this route or host. For more information, including
	// details on header value syntax, see the
	// [Envoy documentation](https://www.envoyproxy.io/docs/envoy/latest/configuration/http_conn_man/headers#config-http-conn-man-headers-custom-request-headers) .
	ResponseHeadersToAdd []*HeaderValueOption `protobuf:"bytes,3,rep,name=response_headers_to_add,json=responseHeadersToAdd,proto3" json:"response_headers_to_add,omitempty"`
	// Specifies a list of HTTP headers that should be removed from each response
	// handled by this route or virtual host.
	ResponseHeadersToRemove []string `protobuf:"bytes,4,rep,name=response_headers_to_remove,json=responseHeadersToRemove,proto3" json:"response_headers_to_remove,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *HeaderManipulation) Reset()         { *m = HeaderManipulation{} }
func (m *HeaderManipulation) String() string { return proto.CompactTextString(m) }
func (*HeaderManipulation) ProtoMessage()    {}
func (*HeaderManipulation) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e5bf0e2e363354d, []int{0}
}
func (m *HeaderManipulation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderManipulation.Unmarshal(m, b)
}
func (m *HeaderManipulation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderManipulation.Marshal(b, m, deterministic)
}
func (m *HeaderManipulation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderManipulation.Merge(m, src)
}
func (m *HeaderManipulation) XXX_Size() int {
	return xxx_messageInfo_HeaderManipulation.Size(m)
}
func (m *HeaderManipulation) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderManipulation.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderManipulation proto.InternalMessageInfo

func (m *HeaderManipulation) GetRequestHeadersToAdd() []*HeaderValueOption {
	if m != nil {
		return m.RequestHeadersToAdd
	}
	return nil
}

func (m *HeaderManipulation) GetRequestHeadersToRemove() []string {
	if m != nil {
		return m.RequestHeadersToRemove
	}
	return nil
}

func (m *HeaderManipulation) GetResponseHeadersToAdd() []*HeaderValueOption {
	if m != nil {
		return m.ResponseHeadersToAdd
	}
	return nil
}

func (m *HeaderManipulation) GetResponseHeadersToRemove() []string {
	if m != nil {
		return m.ResponseHeadersToRemove
	}
	return nil
}

// Header name/value pair plus option to control append behavior.
type HeaderValueOption struct {
	// Header name/value pair that this option applies to.
	Header *HeaderValue `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// Should the value be appended? If true (default), the value is appended to
	// existing values.
	Append               *types.BoolValue `protobuf:"bytes,2,opt,name=append,proto3" json:"append,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *HeaderValueOption) Reset()         { *m = HeaderValueOption{} }
func (m *HeaderValueOption) String() string { return proto.CompactTextString(m) }
func (*HeaderValueOption) ProtoMessage()    {}
func (*HeaderValueOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e5bf0e2e363354d, []int{1}
}
func (m *HeaderValueOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderValueOption.Unmarshal(m, b)
}
func (m *HeaderValueOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderValueOption.Marshal(b, m, deterministic)
}
func (m *HeaderValueOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderValueOption.Merge(m, src)
}
func (m *HeaderValueOption) XXX_Size() int {
	return xxx_messageInfo_HeaderValueOption.Size(m)
}
func (m *HeaderValueOption) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderValueOption.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderValueOption proto.InternalMessageInfo

func (m *HeaderValueOption) GetHeader() *HeaderValue {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HeaderValueOption) GetAppend() *types.BoolValue {
	if m != nil {
		return m.Append
	}
	return nil
}

// Header name/value pair.
type HeaderValue struct {
	// Header name.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Header value.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeaderValue) Reset()         { *m = HeaderValue{} }
func (m *HeaderValue) String() string { return proto.CompactTextString(m) }
func (*HeaderValue) ProtoMessage()    {}
func (*HeaderValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e5bf0e2e363354d, []int{2}
}
func (m *HeaderValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderValue.Unmarshal(m, b)
}
func (m *HeaderValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderValue.Marshal(b, m, deterministic)
}
func (m *HeaderValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderValue.Merge(m, src)
}
func (m *HeaderValue) XXX_Size() int {
	return xxx_messageInfo_HeaderValue.Size(m)
}
func (m *HeaderValue) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderValue.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderValue proto.InternalMessageInfo

func (m *HeaderValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *HeaderValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*HeaderManipulation)(nil), "headers.options.gloo.solo.io.HeaderManipulation")
	proto.RegisterType((*HeaderValueOption)(nil), "headers.options.gloo.solo.io.HeaderValueOption")
	proto.RegisterType((*HeaderValue)(nil), "headers.options.gloo.solo.io.HeaderValue")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/headers/headers.proto", fileDescriptor_2e5bf0e2e363354d)
}

var fileDescriptor_2e5bf0e2e363354d = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x4e, 0xfa, 0x40,
	0x14, 0xc5, 0x53, 0xca, 0x9f, 0x84, 0x61, 0xf3, 0x77, 0x24, 0x50, 0x1b, 0x43, 0x08, 0x2b, 0x5c,
	0x38, 0x13, 0x31, 0x2e, 0x8c, 0x2b, 0x88, 0x0b, 0x13, 0xa3, 0x26, 0x8d, 0x71, 0xe1, 0x86, 0x0c,
	0x76, 0x28, 0x0d, 0xa5, 0x77, 0xec, 0x07, 0xe2, 0x2b, 0xf8, 0x24, 0x3e, 0x82, 0xcf, 0xe2, 0xd2,
	0x77, 0x70, 0x6f, 0xe6, 0xa3, 0x44, 0x81, 0x10, 0xe3, 0xa6, 0x9d, 0x99, 0x7b, 0xcf, 0x39, 0xbf,
	0xe4, 0x5e, 0x74, 0x19, 0x84, 0xd9, 0x24, 0x1f, 0x91, 0x07, 0x98, 0xd1, 0x14, 0x22, 0x38, 0x0c,
	0x41, 0xff, 0x99, 0x08, 0x53, 0xca, 0x44, 0x48, 0x83, 0x08, 0x40, 0x7f, 0xe6, 0x47, 0x14, 0x44,
	0x16, 0x42, 0x9c, 0xd2, 0x09, 0x67, 0x3e, 0x4f, 0x96, 0x7f, 0x22, 0x12, 0xc8, 0x00, 0xef, 0x17,
	0x57, 0xd3, 0x46, 0xa4, 0x8c, 0x48, 0x3b, 0x12, 0x82, 0x5b, 0x0f, 0x20, 0x00, 0xd5, 0x48, 0xe5,
	0x49, 0x6b, 0x5c, 0xcc, 0x17, 0x99, 0x7e, 0xe4, 0x8b, 0xcc, 0xbc, 0xb5, 0x02, 0x80, 0x20, 0xe2,
	0x54, 0xdd, 0x46, 0xf9, 0x98, 0x3e, 0x25, 0x4c, 0x88, 0x65, 0x4e, 0xe7, 0xbd, 0x84, 0xf0, 0x85,
	0x8a, 0xba, 0x62, 0x71, 0x28, 0xf2, 0x88, 0xc9, 0x38, 0xec, 0xa3, 0x46, 0xc2, 0x1f, 0x73, 0x9e,
	0x66, 0x43, 0x03, 0x32, 0xcc, 0x60, 0xc8, 0x7c, 0xdf, 0xb1, 0xda, 0x76, 0xb7, 0xd6, 0xa3, 0x64,
	0x1b, 0x1f, 0xd1, 0x8e, 0x77, 0x2c, 0xca, 0xf9, 0x8d, 0xaa, 0x7b, 0xbb, 0xc6, 0x4e, 0x57, 0xd2,
	0x5b, 0xe8, 0xfb, 0x3e, 0x3e, 0x45, 0x7b, 0x1b, 0x52, 0x12, 0x3e, 0x83, 0x39, 0x77, 0x4a, 0x6d,
	0xbb, 0x5b, 0xf5, 0x1a, 0xab, 0x3a, 0x4f, 0x55, 0xf1, 0x18, 0x35, 0x13, 0x9e, 0x0a, 0x88, 0x53,
	0xbe, 0x4a, 0x68, 0xff, 0x8d, 0xb0, 0x5e, 0xf8, 0xfd, 0x40, 0x3c, 0x43, 0xee, 0xa6, 0x1c, 0xc3,
	0x58, 0x56, 0x8c, 0xcd, 0x35, 0xa5, 0x86, 0xec, 0xbc, 0x58, 0x68, 0x67, 0x2d, 0x08, 0xf7, 0x51,
	0x45, 0x3b, 0x39, 0x56, 0xdb, 0xea, 0xd6, 0x7a, 0x07, 0xbf, 0x26, 0xf5, 0x8c, 0x10, 0xf7, 0x50,
	0x45, 0x4e, 0x31, 0xf6, 0x9d, 0x92, 0xb2, 0x70, 0x89, 0x1e, 0x33, 0x29, 0xc6, 0x4c, 0x06, 0x00,
	0x91, 0xd1, 0xe8, 0xce, 0xce, 0x09, 0xaa, 0x7d, 0xb3, 0xc2, 0xff, 0x91, 0x3d, 0xe5, 0xcf, 0x0a,
	0xa1, 0xea, 0xc9, 0x23, 0xae, 0xa3, 0x7f, 0x73, 0x59, 0x52, 0x9e, 0x55, 0x4f, 0x5f, 0x06, 0xd7,
	0x6f, 0x9f, 0x65, 0xeb, 0xf5, 0xa3, 0x65, 0xdd, 0x9f, 0x6f, 0xdd, 0x6f, 0x31, 0x0d, 0x96, 0x3b,
	0x5e, 0xb0, 0x6f, 0x58, 0xf3, 0x51, 0x45, 0x21, 0x1e, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xab,
	0xfd, 0xaf, 0x74, 0x2e, 0x03, 0x00, 0x00,
}

func (this *HeaderManipulation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HeaderManipulation)
	if !ok {
		that2, ok := that.(HeaderManipulation)
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
	if len(this.RequestHeadersToAdd) != len(that1.RequestHeadersToAdd) {
		return false
	}
	for i := range this.RequestHeadersToAdd {
		if !this.RequestHeadersToAdd[i].Equal(that1.RequestHeadersToAdd[i]) {
			return false
		}
	}
	if len(this.RequestHeadersToRemove) != len(that1.RequestHeadersToRemove) {
		return false
	}
	for i := range this.RequestHeadersToRemove {
		if this.RequestHeadersToRemove[i] != that1.RequestHeadersToRemove[i] {
			return false
		}
	}
	if len(this.ResponseHeadersToAdd) != len(that1.ResponseHeadersToAdd) {
		return false
	}
	for i := range this.ResponseHeadersToAdd {
		if !this.ResponseHeadersToAdd[i].Equal(that1.ResponseHeadersToAdd[i]) {
			return false
		}
	}
	if len(this.ResponseHeadersToRemove) != len(that1.ResponseHeadersToRemove) {
		return false
	}
	for i := range this.ResponseHeadersToRemove {
		if this.ResponseHeadersToRemove[i] != that1.ResponseHeadersToRemove[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *HeaderValueOption) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HeaderValueOption)
	if !ok {
		that2, ok := that.(HeaderValueOption)
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
	if !this.Header.Equal(that1.Header) {
		return false
	}
	if !this.Append.Equal(that1.Append) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *HeaderValue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HeaderValue)
	if !ok {
		that2, ok := that.(HeaderValue)
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
