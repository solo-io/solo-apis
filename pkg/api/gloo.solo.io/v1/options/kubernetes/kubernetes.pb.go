// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/kubernetes/kubernetes.proto

package kubernetes

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	options "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options"
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

// Kubernetes Upstreams represent a set of one or more addressable pods for a Kubernetes Service
// the Gloo Kubernetes Upstream maps to a single service port. Because Kubernetes Services support multiple ports,
// Gloo requires that a different upstream be created for each port
// Kubernetes Upstreams are typically generated automatically by Gloo from the Kubernetes API
type UpstreamSpec struct {
	// The name of the Kubernetes Service
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The namespace where the Service lives
	ServiceNamespace string `protobuf:"bytes,2,opt,name=service_namespace,json=serviceNamespace,proto3" json:"service_namespace,omitempty"`
	// The access port of the kubernetes service is listening. This port is used by Gloo to look up the corresponding
	// port on the pod for routing.
	ServicePort uint32 `protobuf:"varint,3,opt,name=service_port,json=servicePort,proto3" json:"service_port,omitempty"`
	// Allows finer-grained filtering of pods for the Upstream. Gloo will select pods based on their labels if
	// any are provided here.
	// (see [Kubernetes labels and selectors](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)
	Selector map[string]string `protobuf:"bytes,4,rep,name=selector,proto3" json:"selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// An optional Service Spec describing the service listening at this address
	ServiceSpec *options.ServiceSpec `protobuf:"bytes,5,opt,name=service_spec,json=serviceSpec,proto3" json:"service_spec,omitempty"`
	// Subset configuration. For discovery sources that has labels (like kubernetes). this
	// configuration allows you to partition the upstream to a set of subsets.
	// for each unique set of keys and values, a subset will be created.
	SubsetSpec           *options.SubsetSpec `protobuf:"bytes,6,opt,name=subset_spec,json=subsetSpec,proto3" json:"subset_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_69b43d2b114ab83a, []int{0}
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

func (m *UpstreamSpec) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *UpstreamSpec) GetServiceNamespace() string {
	if m != nil {
		return m.ServiceNamespace
	}
	return ""
}

func (m *UpstreamSpec) GetServicePort() uint32 {
	if m != nil {
		return m.ServicePort
	}
	return 0
}

func (m *UpstreamSpec) GetSelector() map[string]string {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *UpstreamSpec) GetServiceSpec() *options.ServiceSpec {
	if m != nil {
		return m.ServiceSpec
	}
	return nil
}

func (m *UpstreamSpec) GetSubsetSpec() *options.SubsetSpec {
	if m != nil {
		return m.SubsetSpec
	}
	return nil
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "kubernetes.options.gloo.solo.io.UpstreamSpec")
	proto.RegisterMapType((map[string]string)(nil), "kubernetes.options.gloo.solo.io.UpstreamSpec.SelectorEntry")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/kubernetes/kubernetes.proto", fileDescriptor_69b43d2b114ab83a)
}

var fileDescriptor_69b43d2b114ab83a = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6a, 0xdb, 0x40,
	0x14, 0x45, 0x91, 0x65, 0x9b, 0x76, 0x64, 0x83, 0x3b, 0x78, 0x21, 0xbc, 0x68, 0xe5, 0xae, 0x04,
	0xa5, 0x23, 0xea, 0xb6, 0x50, 0xea, 0x55, 0x4b, 0x4b, 0x76, 0x21, 0xb6, 0x09, 0x81, 0x6c, 0x82,
	0x24, 0x1e, 0x8a, 0xb0, 0xa4, 0x37, 0xcc, 0x8c, 0x8c, 0xfd, 0x37, 0x59, 0xe6, 0x13, 0xf2, 0x3d,
	0xf9, 0x87, 0xec, 0x83, 0x46, 0xb2, 0x32, 0x01, 0x27, 0xf1, 0xc6, 0x7e, 0x7a, 0xdc, 0x7b, 0x74,
	0x75, 0x67, 0xc8, 0x22, 0x49, 0xd5, 0x75, 0x19, 0xb1, 0x18, 0xf3, 0x40, 0x62, 0x86, 0x5f, 0x53,
	0xac, 0xff, 0x43, 0x9e, 0xca, 0x20, 0xe4, 0x69, 0x90, 0x64, 0x88, 0xf5, 0xcf, 0xe6, 0x5b, 0x80,
	0x5c, 0xa5, 0x58, 0xc8, 0x60, 0x5d, 0x46, 0x20, 0x0a, 0x50, 0x60, 0x8e, 0x8c, 0x0b, 0x54, 0x48,
	0x3f, 0x19, 0x9b, 0x46, 0xcf, 0x2a, 0x3f, 0xab, 0xb8, 0x2c, 0xc5, 0xc9, 0x38, 0xc1, 0x04, 0xb5,
	0x36, 0xa8, 0xa6, 0xda, 0x36, 0xa1, 0xb0, 0x55, 0xf5, 0x12, 0xb6, 0xaa, 0xd9, 0xfd, 0x3c, 0x22,
	0x8a, 0x04, 0xb1, 0x49, 0x63, 0xb8, 0x92, 0x1c, 0xe2, 0xc6, 0xf6, 0xe3, 0x18, 0x5b, 0x19, 0x49,
	0x50, 0x86, 0xeb, 0xf3, 0x8d, 0x4d, 0x06, 0xe7, 0x5c, 0x2a, 0x01, 0x61, 0xbe, 0xe2, 0x10, 0xd3,
	0x29, 0x19, 0xec, 0xe1, 0x45, 0x98, 0x83, 0x6b, 0x79, 0x96, 0xff, 0x7e, 0xe9, 0x34, 0xbb, 0xd3,
	0x30, 0x07, 0xfa, 0x85, 0x7c, 0x30, 0x25, 0x92, 0x87, 0x31, 0xb8, 0x1d, 0xad, 0x1b, 0x19, 0x3a,
	0xbd, 0x37, 0x79, 0x1c, 0x85, 0x72, 0x6d, 0xcf, 0xf2, 0x87, 0x2d, 0xef, 0x0c, 0x85, 0xa2, 0x17,
	0xe4, 0x9d, 0x84, 0x0c, 0x62, 0x85, 0xc2, 0xed, 0x7a, 0xb6, 0xef, 0xcc, 0xe6, 0xec, 0x8d, 0x3a,
	0x99, 0x99, 0x99, 0xad, 0x1a, 0xf7, 0xff, 0x42, 0x89, 0xdd, 0xb2, 0x85, 0xd1, 0x7f, 0x4f, 0xef,
	0xae, 0x3e, 0xd9, 0xed, 0x79, 0x96, 0xef, 0xcc, 0xa6, 0x87, 0x89, 0xab, 0x5a, 0x59, 0x01, 0xdb,
	0x78, 0xba, 0x91, 0x3f, 0xc4, 0x31, 0x7a, 0x73, 0xfb, 0x1a, 0xe2, 0xbd, 0x00, 0xd1, 0x42, 0xcd,
	0x20, 0xb2, 0x9d, 0x27, 0x73, 0x32, 0x7c, 0x96, 0x91, 0x8e, 0x88, 0xbd, 0x86, 0x5d, 0x53, 0x6e,
	0x35, 0xd2, 0x31, 0xe9, 0x6d, 0xc2, 0xac, 0xdc, 0x17, 0x59, 0x3f, 0xfc, 0xee, 0xfc, 0xb2, 0xfe,
	0x2e, 0xee, 0x1e, 0xba, 0xd6, 0xed, 0xfd, 0x47, 0xeb, 0xf2, 0xe4, 0xd5, 0x7b, 0xcb, 0xd7, 0x49,
	0x7b, 0xf2, 0xfb, 0x40, 0x87, 0xaf, 0x6f, 0xd4, 0xd7, 0x87, 0xff, 0xfd, 0x31, 0x00, 0x00, 0xff,
	0xff, 0x8d, 0x09, 0xd0, 0x3f, 0x09, 0x03, 0x00, 0x00,
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
	if this.ServiceName != that1.ServiceName {
		return false
	}
	if this.ServiceNamespace != that1.ServiceNamespace {
		return false
	}
	if this.ServicePort != that1.ServicePort {
		return false
	}
	if len(this.Selector) != len(that1.Selector) {
		return false
	}
	for i := range this.Selector {
		if this.Selector[i] != that1.Selector[i] {
			return false
		}
	}
	if !this.ServiceSpec.Equal(that1.ServiceSpec) {
		return false
	}
	if !this.SubsetSpec.Equal(that1.SubsetSpec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
