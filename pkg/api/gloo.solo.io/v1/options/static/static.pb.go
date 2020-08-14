// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/static/static.proto

package static

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	options "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options"
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

// Static upstreams are used to route request to services listening at fixed IP/Host & Port pairs.
// Static upstreams can be used to proxy any kind of service, and therefore contain a ServiceSpec
// for additional service-specific configuration.
// Unlike upstreams created by service discovery, Static Upstreams must be created manually by users
type UpstreamSpec struct {
	// A list of addresses and ports
	// at least one must be specified
	Hosts []*Host `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	// Attempt to use outbound TLS
	// Gloo will automatically set this to true for port 443
	UseTls bool `protobuf:"varint,3,opt,name=use_tls,json=useTls,proto3" json:"use_tls,omitempty"`
	// An optional Service Spec describing the service listening at this address
	ServiceSpec          *options.ServiceSpec `protobuf:"bytes,5,opt,name=service_spec,json=serviceSpec,proto3" json:"service_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f106f4f05c7d485, []int{0}
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

func (m *UpstreamSpec) GetHosts() []*Host {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *UpstreamSpec) GetUseTls() bool {
	if m != nil {
		return m.UseTls
	}
	return false
}

func (m *UpstreamSpec) GetServiceSpec() *options.ServiceSpec {
	if m != nil {
		return m.ServiceSpec
	}
	return nil
}

// Represents a single instance of an upstream
type Host struct {
	// Address (hostname or IP)
	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	// Port the instance is listening on
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// (Enterprise Only): Host specific health checking configuration.
	HealthCheckConfig    *Host_HealthCheckConfig `protobuf:"bytes,3,opt,name=health_check_config,json=healthCheckConfig,proto3" json:"health_check_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Host) Reset()         { *m = Host{} }
func (m *Host) String() string { return proto.CompactTextString(m) }
func (*Host) ProtoMessage()    {}
func (*Host) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f106f4f05c7d485, []int{1}
}
func (m *Host) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Host.Unmarshal(m, b)
}
func (m *Host) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Host.Marshal(b, m, deterministic)
}
func (m *Host) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Host.Merge(m, src)
}
func (m *Host) XXX_Size() int {
	return xxx_messageInfo_Host.Size(m)
}
func (m *Host) XXX_DiscardUnknown() {
	xxx_messageInfo_Host.DiscardUnknown(m)
}

var xxx_messageInfo_Host proto.InternalMessageInfo

func (m *Host) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Host) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Host) GetHealthCheckConfig() *Host_HealthCheckConfig {
	if m != nil {
		return m.HealthCheckConfig
	}
	return nil
}

type Host_HealthCheckConfig struct {
	// (Enterprise Only): Path to use when health checking this specific host.
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Host_HealthCheckConfig) Reset()         { *m = Host_HealthCheckConfig{} }
func (m *Host_HealthCheckConfig) String() string { return proto.CompactTextString(m) }
func (*Host_HealthCheckConfig) ProtoMessage()    {}
func (*Host_HealthCheckConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f106f4f05c7d485, []int{1, 0}
}
func (m *Host_HealthCheckConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Host_HealthCheckConfig.Unmarshal(m, b)
}
func (m *Host_HealthCheckConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Host_HealthCheckConfig.Marshal(b, m, deterministic)
}
func (m *Host_HealthCheckConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Host_HealthCheckConfig.Merge(m, src)
}
func (m *Host_HealthCheckConfig) XXX_Size() int {
	return xxx_messageInfo_Host_HealthCheckConfig.Size(m)
}
func (m *Host_HealthCheckConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_Host_HealthCheckConfig.DiscardUnknown(m)
}

var xxx_messageInfo_Host_HealthCheckConfig proto.InternalMessageInfo

func (m *Host_HealthCheckConfig) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "static.options.gloo.solo.io.UpstreamSpec")
	proto.RegisterType((*Host)(nil), "static.options.gloo.solo.io.Host")
	proto.RegisterType((*Host_HealthCheckConfig)(nil), "static.options.gloo.solo.io.Host.HealthCheckConfig")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/static/static.proto", fileDescriptor_8f106f4f05c7d485)
}

var fileDescriptor_8f106f4f05c7d485 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x8e, 0xd3, 0x30,
	0x14, 0x85, 0x65, 0xfa, 0x03, 0xb8, 0x65, 0x51, 0x83, 0x44, 0x54, 0xa4, 0x2a, 0x74, 0x43, 0x36,
	0xd8, 0xa2, 0x15, 0x62, 0x4f, 0x59, 0x94, 0x05, 0x9b, 0x14, 0x36, 0x6c, 0x22, 0xd7, 0x75, 0x1d,
	0xab, 0x69, 0xaf, 0x15, 0x3b, 0xa5, 0x4f, 0x84, 0x78, 0x04, 0x16, 0xf3, 0x34, 0xf3, 0x0e, 0xb3,
	0x1f, 0xd9, 0x49, 0x67, 0x46, 0xd3, 0x51, 0x67, 0x36, 0xf1, 0xbd, 0x27, 0xe7, 0x5c, 0x7f, 0x96,
	0x8d, 0xbf, 0x2b, 0xed, 0xf2, 0x6a, 0x49, 0x05, 0x6c, 0x99, 0x85, 0x02, 0x3e, 0x6a, 0xa8, 0x57,
	0x6e, 0xb4, 0x65, 0xdc, 0x68, 0xa6, 0x0a, 0x80, 0xfa, 0xb3, 0xff, 0xc4, 0xc0, 0x38, 0x0d, 0x3b,
	0xcb, 0xac, 0xe3, 0x4e, 0x8b, 0x66, 0xa1, 0xa6, 0x04, 0x07, 0xe4, 0x5d, 0xd3, 0x35, 0x1e, 0xea,
	0x33, 0xd4, 0xcf, 0xa2, 0x1a, 0x86, 0x6f, 0x14, 0x28, 0x08, 0x3e, 0xe6, 0xab, 0x3a, 0x32, 0x24,
	0xf2, 0xe0, 0x6a, 0x51, 0x1e, 0x5c, 0xa3, 0x8d, 0x14, 0x80, 0x2a, 0x24, 0x0b, 0xdd, 0xb2, 0x5a,
	0xb3, 0x3f, 0x25, 0x37, 0x46, 0x96, 0xb6, 0xf9, 0xff, 0xf9, 0x29, 0x78, 0xb2, 0xdc, 0x6b, 0x21,
	0x33, 0x6b, 0x64, 0x43, 0x37, 0xfe, 0x8b, 0x70, 0xff, 0x97, 0xb1, 0xae, 0x94, 0x7c, 0xbb, 0x30,
	0x52, 0x90, 0x2f, 0xb8, 0x93, 0x83, 0x75, 0x36, 0x42, 0x71, 0x2b, 0xe9, 0x4d, 0xde, 0xd3, 0x33,
	0xf8, 0x74, 0x0e, 0xd6, 0xa5, 0xb5, 0x9f, 0xbc, 0xc5, 0xcf, 0x2b, 0x2b, 0x33, 0x57, 0xd8, 0xa8,
	0x15, 0xa3, 0xe4, 0x45, 0xda, 0xad, 0xac, 0xfc, 0x59, 0x58, 0xf2, 0x0d, 0xf7, 0xef, 0x6e, 0x1c,
	0x75, 0x62, 0x14, 0x06, 0x3f, 0x38, 0x71, 0x51, 0x3b, 0x3d, 0x4a, 0xda, 0xb3, 0xb7, 0xcd, 0xf8,
	0x02, 0xe1, 0xb6, 0xdf, 0x8e, 0x10, 0xdc, 0xe6, 0xab, 0x55, 0x19, 0xa1, 0x18, 0x25, 0x2f, 0xd3,
	0x50, 0x7b, 0xcd, 0x40, 0xe9, 0xa2, 0x67, 0x31, 0x4a, 0x5e, 0xa5, 0xa1, 0x26, 0x02, 0xbf, 0xce,
	0x25, 0x2f, 0x5c, 0x9e, 0x89, 0x5c, 0x8a, 0x4d, 0x26, 0x60, 0xb7, 0xd6, 0x2a, 0xb0, 0xf5, 0x26,
	0xd3, 0x47, 0x8f, 0x45, 0xe7, 0x21, 0x3c, 0xf3, 0xd9, 0x59, 0x88, 0xa6, 0x83, 0xfc, 0xbe, 0x34,
	0xfc, 0x80, 0x07, 0x27, 0xbe, 0x40, 0xc3, 0x5d, 0x7e, 0x24, 0xf4, 0xf5, 0xd7, 0x1f, 0xff, 0xaf,
	0xda, 0xe8, 0xdf, 0xe5, 0x08, 0xfd, 0x9e, 0x9d, 0x7d, 0x5a, 0x66, 0xa3, 0x6e, 0xee, 0xef, 0xc8,
	0x73, 0xfa, 0xc2, 0x96, 0xdd, 0x70, 0x7b, 0xd3, 0xeb, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0x54,
	0x1b, 0xae, 0xa8, 0x02, 0x00, 0x00,
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
	if len(this.Hosts) != len(that1.Hosts) {
		return false
	}
	for i := range this.Hosts {
		if !this.Hosts[i].Equal(that1.Hosts[i]) {
			return false
		}
	}
	if this.UseTls != that1.UseTls {
		return false
	}
	if !this.ServiceSpec.Equal(that1.ServiceSpec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Host) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Host)
	if !ok {
		that2, ok := that.(Host)
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
	if this.Addr != that1.Addr {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if !this.HealthCheckConfig.Equal(that1.HealthCheckConfig) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Host_HealthCheckConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Host_HealthCheckConfig)
	if !ok {
		that2, ok := that.(Host_HealthCheckConfig)
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
	if this.Path != that1.Path {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
