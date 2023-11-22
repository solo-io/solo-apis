// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/enterprise/options/tap/tap.proto

package tap

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	duration "github.com/golang/protobuf/ptypes/duration"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Tap filter: a filter that copies the contents of HTTP requests and responses
// to an external tap server. The full HTTP headers and bodies are reported in
// full to the configured address, and data can be reported using either over
// HTTP or GRPC.
type Tap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Sinks to which tap data should be output. Currently, only a single sink
	// is supported.
	Sinks []*Sink `protobuf:"bytes,1,rep,name=sinks,proto3" json:"sinks,omitempty"`
}

func (x *Tap) Reset() {
	*x = Tap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tap) ProtoMessage() {}

func (x *Tap) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tap.ProtoReflect.Descriptor instead.
func (*Tap) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_rawDescGZIP(), []int{0}
}

func (x *Tap) GetSinks() []*Sink {
	if x != nil {
		return x.Sinks
	}
	return nil
}

type Sink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the output sink to which tap data should be written
	//
	// Types that are assignable to SinkType:
	//	*Sink_GrpcService
	//	*Sink_HttpService
	SinkType isSink_SinkType `protobuf_oneof:"SinkType"`
}

func (x *Sink) Reset() {
	*x = Sink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sink) ProtoMessage() {}

func (x *Sink) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sink.ProtoReflect.Descriptor instead.
func (*Sink) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_rawDescGZIP(), []int{1}
}

func (m *Sink) GetSinkType() isSink_SinkType {
	if m != nil {
		return m.SinkType
	}
	return nil
}

func (x *Sink) GetGrpcService() *GrpcService {
	if x, ok := x.GetSinkType().(*Sink_GrpcService); ok {
		return x.GrpcService
	}
	return nil
}

func (x *Sink) GetHttpService() *HttpService {
	if x, ok := x.GetSinkType().(*Sink_HttpService); ok {
		return x.HttpService
	}
	return nil
}

type isSink_SinkType interface {
	isSink_SinkType()
}

type Sink_GrpcService struct {
	// Write tap data out to a GRPC service
	GrpcService *GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3,oneof"`
}

type Sink_HttpService struct {
	// Write tap data out to a HTTP service
	HttpService *HttpService `protobuf:"bytes,2,opt,name=http_service,json=httpService,proto3,oneof"`
}

func (*Sink_GrpcService) isSink_SinkType() {}

func (*Sink_HttpService) isSink_SinkType() {}

// A tap sink over a GRPC service
type GrpcService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Upstream reference to the tap server
	TapServer *core.ResourceRef `protobuf:"bytes,1,opt,name=tap_server,json=tapServer,proto3" json:"tap_server,omitempty"`
}

func (x *GrpcService) Reset() {
	*x = GrpcService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcService) ProtoMessage() {}

func (x *GrpcService) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcService.ProtoReflect.Descriptor instead.
func (*GrpcService) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_rawDescGZIP(), []int{2}
}

func (x *GrpcService) GetTapServer() *core.ResourceRef {
	if x != nil {
		return x.TapServer
	}
	return nil
}

// A tap sink over a HTTP service
type HttpService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Upstream reference to the tap server
	TapServer *core.ResourceRef `protobuf:"bytes,1,opt,name=tap_server,json=tapServer,proto3" json:"tap_server,omitempty"`
	// Connection timeout
	Timeout *duration.Duration `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *HttpService) Reset() {
	*x = HttpService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpService) ProtoMessage() {}

func (x *HttpService) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpService.ProtoReflect.Descriptor instead.
func (*HttpService) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_rawDescGZIP(), []int{3}
}

func (x *HttpService) GetTapServer() *core.ResourceRef {
	if x != nil {
		return x.TapServer
	}
	return nil
}

func (x *HttpService) GetTimeout() *duration.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_rawDesc = []byte{
	0x0a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x74, 0x61, 0x70, 0x2f, 0x74, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x74, 0x61, 0x70, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x47, 0x0a, 0x03, 0x54, 0x61, 0x70, 0x12, 0x40, 0x0a, 0x05, 0x73, 0x69, 0x6e, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x53, 0x69, 0x6e, 0x6b, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x92, 0x01, 0x04, 0x08,
	0x01, 0x10, 0x01, 0x52, 0x05, 0x73, 0x69, 0x6e, 0x6b, 0x73, 0x22, 0xaf, 0x01, 0x0a, 0x04, 0x53,
	0x69, 0x6e, 0x6b, 0x12, 0x4a, 0x0a, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x61, 0x70, 0x2e,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x48, 0x00, 0x52, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4a, 0x0a, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x48, 0x74, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0b,
	0x68, 0x74, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x0f, 0x0a, 0x08, 0x53,
	0x69, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x51, 0x0a, 0x0b,
	0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0a, 0x74,
	0x61, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x09, 0x74, 0x61, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22,
	0x90, 0x01, 0x0a, 0x0b, 0x48, 0x74, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x42, 0x0a, 0x0a, 0x74, 0x61, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x09, 0x74, 0x61, 0x70, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x74, 0x61, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_goTypes = []interface{}{
	(*Tap)(nil),               // 0: tap.options.gloo.solo.io.Tap
	(*Sink)(nil),              // 1: tap.options.gloo.solo.io.Sink
	(*GrpcService)(nil),       // 2: tap.options.gloo.solo.io.GrpcService
	(*HttpService)(nil),       // 3: tap.options.gloo.solo.io.HttpService
	(*core.ResourceRef)(nil),  // 4: core.solo.io.ResourceRef
	(*duration.Duration)(nil), // 5: google.protobuf.Duration
}
var file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_depIdxs = []int32{
	1, // 0: tap.options.gloo.solo.io.Tap.sinks:type_name -> tap.options.gloo.solo.io.Sink
	2, // 1: tap.options.gloo.solo.io.Sink.grpc_service:type_name -> tap.options.gloo.solo.io.GrpcService
	3, // 2: tap.options.gloo.solo.io.Sink.http_service:type_name -> tap.options.gloo.solo.io.HttpService
	4, // 3: tap.options.gloo.solo.io.GrpcService.tap_server:type_name -> core.solo.io.ResourceRef
	4, // 4: tap.options.gloo.solo.io.HttpService.tap_server:type_name -> core.solo.io.ResourceRef
	5, // 5: tap.options.gloo.solo.io.HttpService.timeout:type_name -> google.protobuf.Duration
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tap); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sink); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcService); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpService); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Sink_GrpcService)(nil),
		(*Sink_HttpService)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_enterprise_options_tap_tap_proto_depIdxs = nil
}
