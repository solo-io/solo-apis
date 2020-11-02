// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/tracing/tracing.proto

package tracing

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// Contains settings for configuring Envoy's tracing capabilities at the listener level.
// See here for additional information on Envoy's tracing capabilities: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/observability/tracing.html
// See here for additional information about configuring tracing with Gloo: https://gloo.solo.io/user_guides/setup_options/observability/#tracing
type ListenerTracingSettings struct {
	// Optional. If specified, Envoy will include the headers and header values for any matching request headers.
	RequestHeadersForTags []string `protobuf:"bytes,1,rep,name=request_headers_for_tags,json=requestHeadersForTags,proto3" json:"request_headers_for_tags,omitempty"`
	// Optional. If true, Envoy will include logs for streaming events. Default: false.
	Verbose bool `protobuf:"varint,2,opt,name=verbose,proto3" json:"verbose,omitempty"`
	// Requests can produce traces by random sampling or when the `x-client-trace-id` header is provided.
	// TracePercentages defines the limits for random, forced, and overall tracing percentages.
	TracePercentages     *TracePercentages `protobuf:"bytes,3,opt,name=trace_percentages,json=tracePercentages,proto3" json:"trace_percentages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListenerTracingSettings) Reset()         { *m = ListenerTracingSettings{} }
func (m *ListenerTracingSettings) String() string { return proto.CompactTextString(m) }
func (*ListenerTracingSettings) ProtoMessage()    {}
func (*ListenerTracingSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_863e36444a76f832, []int{0}
}
func (m *ListenerTracingSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerTracingSettings.Unmarshal(m, b)
}
func (m *ListenerTracingSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerTracingSettings.Marshal(b, m, deterministic)
}
func (m *ListenerTracingSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerTracingSettings.Merge(m, src)
}
func (m *ListenerTracingSettings) XXX_Size() int {
	return xxx_messageInfo_ListenerTracingSettings.Size(m)
}
func (m *ListenerTracingSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerTracingSettings.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerTracingSettings proto.InternalMessageInfo

func (m *ListenerTracingSettings) GetRequestHeadersForTags() []string {
	if m != nil {
		return m.RequestHeadersForTags
	}
	return nil
}

func (m *ListenerTracingSettings) GetVerbose() bool {
	if m != nil {
		return m.Verbose
	}
	return false
}

func (m *ListenerTracingSettings) GetTracePercentages() *TracePercentages {
	if m != nil {
		return m.TracePercentages
	}
	return nil
}

// Contains settings for configuring Envoy's tracing capabilities at the route level.
// Note: must also specify ListenerTracingSettings for the associated listener.
// See here for additional information on Envoy's tracing capabilities: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/observability/tracing.html
// See here for additional information about configuring tracing with Gloo: https://gloo.solo.io/user_guides/setup_options/observability/#tracing
type RouteTracingSettings struct {
	// Optional. If set, will be used to identify the route that produced the trace.
	// Note that this value will be overridden if the "x-envoy-decorator-operation" header is passed.
	RouteDescriptor string `protobuf:"bytes,1,opt,name=route_descriptor,json=routeDescriptor,proto3" json:"route_descriptor,omitempty"`
	// Requests can produce traces by random sampling or when the `x-client-trace-id` header is provided.
	// TracePercentages defines the limits for random, forced, and overall tracing percentages.
	TracePercentages     *TracePercentages `protobuf:"bytes,2,opt,name=trace_percentages,json=tracePercentages,proto3" json:"trace_percentages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RouteTracingSettings) Reset()         { *m = RouteTracingSettings{} }
func (m *RouteTracingSettings) String() string { return proto.CompactTextString(m) }
func (*RouteTracingSettings) ProtoMessage()    {}
func (*RouteTracingSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_863e36444a76f832, []int{1}
}
func (m *RouteTracingSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteTracingSettings.Unmarshal(m, b)
}
func (m *RouteTracingSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteTracingSettings.Marshal(b, m, deterministic)
}
func (m *RouteTracingSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteTracingSettings.Merge(m, src)
}
func (m *RouteTracingSettings) XXX_Size() int {
	return xxx_messageInfo_RouteTracingSettings.Size(m)
}
func (m *RouteTracingSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteTracingSettings.DiscardUnknown(m)
}

var xxx_messageInfo_RouteTracingSettings proto.InternalMessageInfo

func (m *RouteTracingSettings) GetRouteDescriptor() string {
	if m != nil {
		return m.RouteDescriptor
	}
	return ""
}

func (m *RouteTracingSettings) GetTracePercentages() *TracePercentages {
	if m != nil {
		return m.TracePercentages
	}
	return nil
}

// Requests can produce traces by random sampling or when the `x-client-trace-id` header is provided.
// TracePercentages defines the limits for random, forced, and overall tracing percentages.
type TracePercentages struct {
	// Percentage of requests that should produce traces when the `x-client-trace-id` header is provided.
	// optional, defaults to 100.0
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	ClientSamplePercentage *types.FloatValue `protobuf:"bytes,1,opt,name=client_sample_percentage,json=clientSamplePercentage,proto3" json:"client_sample_percentage,omitempty"`
	// Percentage of requests that should produce traces by random sampling.
	// optional, defaults to 100.0
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	RandomSamplePercentage *types.FloatValue `protobuf:"bytes,2,opt,name=random_sample_percentage,json=randomSamplePercentage,proto3" json:"random_sample_percentage,omitempty"`
	// Overall percentage of requests that should produce traces.
	// optional, defaults to 100.0
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	OverallSamplePercentage *types.FloatValue `protobuf:"bytes,3,opt,name=overall_sample_percentage,json=overallSamplePercentage,proto3" json:"overall_sample_percentage,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}          `json:"-"`
	XXX_unrecognized        []byte            `json:"-"`
	XXX_sizecache           int32             `json:"-"`
}

func (m *TracePercentages) Reset()         { *m = TracePercentages{} }
func (m *TracePercentages) String() string { return proto.CompactTextString(m) }
func (*TracePercentages) ProtoMessage()    {}
func (*TracePercentages) Descriptor() ([]byte, []int) {
	return fileDescriptor_863e36444a76f832, []int{2}
}
func (m *TracePercentages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TracePercentages.Unmarshal(m, b)
}
func (m *TracePercentages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TracePercentages.Marshal(b, m, deterministic)
}
func (m *TracePercentages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TracePercentages.Merge(m, src)
}
func (m *TracePercentages) XXX_Size() int {
	return xxx_messageInfo_TracePercentages.Size(m)
}
func (m *TracePercentages) XXX_DiscardUnknown() {
	xxx_messageInfo_TracePercentages.DiscardUnknown(m)
}

var xxx_messageInfo_TracePercentages proto.InternalMessageInfo

func (m *TracePercentages) GetClientSamplePercentage() *types.FloatValue {
	if m != nil {
		return m.ClientSamplePercentage
	}
	return nil
}

func (m *TracePercentages) GetRandomSamplePercentage() *types.FloatValue {
	if m != nil {
		return m.RandomSamplePercentage
	}
	return nil
}

func (m *TracePercentages) GetOverallSamplePercentage() *types.FloatValue {
	if m != nil {
		return m.OverallSamplePercentage
	}
	return nil
}

func init() {
	proto.RegisterType((*ListenerTracingSettings)(nil), "tracing.options.gloo.solo.io.ListenerTracingSettings")
	proto.RegisterType((*RouteTracingSettings)(nil), "tracing.options.gloo.solo.io.RouteTracingSettings")
	proto.RegisterType((*TracePercentages)(nil), "tracing.options.gloo.solo.io.TracePercentages")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/tracing/tracing.proto", fileDescriptor_863e36444a76f832)
}

var fileDescriptor_863e36444a76f832 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xe5, 0x04, 0x01, 0x75, 0x0f, 0x84, 0x55, 0xa1, 0x4b, 0x41, 0x55, 0xd4, 0x53, 0x38,
	0xe0, 0x15, 0xe5, 0xc0, 0x1d, 0x55, 0x15, 0x12, 0x08, 0xa1, 0x6d, 0x01, 0x09, 0x0e, 0x2b, 0x67,
	0x33, 0x75, 0x2d, 0x9c, 0x1d, 0x33, 0x9e, 0x0d, 0x7d, 0x15, 0x2e, 0x9c, 0x79, 0x04, 0x5e, 0x81,
	0xd7, 0xe0, 0x1d, 0xb8, 0x23, 0xaf, 0xb3, 0x2d, 0x0a, 0xa1, 0x70, 0xe0, 0xb2, 0xde, 0xf9, 0xf3,
	0xfd, 0xfc, 0xd9, 0xd6, 0xc8, 0x67, 0xc6, 0xf2, 0x69, 0x3b, 0x55, 0x35, 0xce, 0x8b, 0x80, 0x0e,
	0x1f, 0x58, 0x4c, 0xab, 0xf6, 0x36, 0x14, 0xda, 0xdb, 0xc2, 0x38, 0xc4, 0xf4, 0x59, 0x3c, 0x2c,
	0xd0, 0xb3, 0xc5, 0x26, 0x14, 0x4c, 0xba, 0xb6, 0x8d, 0xe9, 0x57, 0xe5, 0x09, 0x19, 0xb3, 0x7b,
	0x7d, 0xb8, 0x6c, 0x53, 0x51, 0xa6, 0x22, 0x4e, 0x59, 0xdc, 0xd9, 0x32, 0x68, 0xb0, 0x6b, 0x2c,
	0xe2, 0x5f, 0xd2, 0xec, 0xec, 0x1a, 0x44, 0xe3, 0xa0, 0xe8, 0xa2, 0x69, 0x7b, 0x52, 0x7c, 0x24,
	0xed, 0x3d, 0x50, 0xf8, 0x53, 0x7d, 0xd6, 0x92, 0x8e, 0xf4, 0x65, 0x3d, 0x83, 0x33, 0x4e, 0x50,
	0x38, 0xe3, 0x94, 0xdb, 0xfb, 0x26, 0xe4, 0xf6, 0x73, 0x1b, 0x18, 0x1a, 0xa0, 0xe3, 0x64, 0xe9,
	0x08, 0x98, 0x6d, 0x63, 0x42, 0xf6, 0x58, 0xe6, 0x04, 0x1f, 0x5a, 0x08, 0x5c, 0x9d, 0x82, 0x9e,
	0x01, 0x85, 0xea, 0x04, 0xa9, 0x62, 0x6d, 0x42, 0x2e, 0xc6, 0xc3, 0xc9, 0x46, 0x79, 0x6b, 0x59,
	0x7f, 0x9a, 0xca, 0x87, 0x48, 0xc7, 0xda, 0x84, 0x2c, 0x97, 0xd7, 0x16, 0x40, 0x53, 0x0c, 0x90,
	0x0f, 0xc6, 0x62, 0x72, 0xbd, 0xec, 0xc3, 0xec, 0x9d, 0xbc, 0x19, 0x0f, 0x0e, 0x95, 0x07, 0xaa,
	0xa1, 0x61, 0x6d, 0x20, 0xe4, 0xc3, 0xb1, 0x98, 0x6c, 0xee, 0x2b, 0x75, 0xd9, 0x95, 0xa8, 0x68,
	0x0e, 0x5e, 0x5e, 0xa8, 0xca, 0x11, 0xaf, 0x64, 0xf6, 0x3e, 0x0b, 0xb9, 0x55, 0x62, 0xcb, 0xb0,
	0x7a, 0x90, 0xfb, 0x72, 0x44, 0x31, 0x5f, 0xcd, 0x20, 0xd4, 0x64, 0x3d, 0x23, 0xe5, 0x62, 0x2c,
	0x26, 0x1b, 0xe5, 0x8d, 0x2e, 0x7f, 0x70, 0x9e, 0x5e, 0x6f, 0x70, 0xf0, 0x9f, 0x0c, 0x7e, 0x1a,
	0xc8, 0xd1, 0x6a, 0x5b, 0xf6, 0x4a, 0xe6, 0xb5, 0xb3, 0xd0, 0x70, 0x15, 0xf4, 0xdc, 0xbb, 0x5f,
	0x77, 0xee, 0x4c, 0x6e, 0xee, 0xdf, 0x55, 0xe9, 0x61, 0x55, 0xff, 0xb0, 0xea, 0xd0, 0xa1, 0xe6,
	0xd7, 0xda, 0xb5, 0x50, 0xde, 0x4e, 0xe2, 0xa3, 0x4e, 0x7b, 0xc1, 0x8d, 0x58, 0xd2, 0xcd, 0x0c,
	0xe7, 0x6b, 0xb0, 0x83, 0x7f, 0xc0, 0x26, 0xf1, 0x6f, 0xd8, 0x37, 0xf2, 0x0e, 0x2e, 0x80, 0xb4,
	0x73, 0x6b, 0xb8, 0xc3, 0xbf, 0x73, 0xb7, 0x97, 0xea, 0x55, 0xf0, 0x93, 0x17, 0x5f, 0x7f, 0x5c,
	0x11, 0x5f, 0xbe, 0xef, 0x8a, 0xb7, 0x07, 0x97, 0xce, 0x99, 0x7f, 0x6f, 0xce, 0x67, 0xad, 0xbf,
	0xfb, 0x35, 0xe3, 0x36, 0xbd, 0xda, 0xed, 0xfe, 0xe8, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x14,
	0xbd, 0xfb, 0xf8, 0xb6, 0x03, 0x00, 0x00,
}

func (this *ListenerTracingSettings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListenerTracingSettings)
	if !ok {
		that2, ok := that.(ListenerTracingSettings)
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
	if len(this.RequestHeadersForTags) != len(that1.RequestHeadersForTags) {
		return false
	}
	for i := range this.RequestHeadersForTags {
		if this.RequestHeadersForTags[i] != that1.RequestHeadersForTags[i] {
			return false
		}
	}
	if this.Verbose != that1.Verbose {
		return false
	}
	if !this.TracePercentages.Equal(that1.TracePercentages) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RouteTracingSettings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteTracingSettings)
	if !ok {
		that2, ok := that.(RouteTracingSettings)
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
	if this.RouteDescriptor != that1.RouteDescriptor {
		return false
	}
	if !this.TracePercentages.Equal(that1.TracePercentages) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TracePercentages) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TracePercentages)
	if !ok {
		that2, ok := that.(TracePercentages)
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
	if !this.ClientSamplePercentage.Equal(that1.ClientSamplePercentage) {
		return false
	}
	if !this.RandomSamplePercentage.Equal(that1.RandomSamplePercentage) {
		return false
	}
	if !this.OverallSamplePercentage.Equal(that1.OverallSamplePercentage) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
