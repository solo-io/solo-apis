// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/service_spec.proto

package options

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	grpc "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/grpc"
	rest "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/rest"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Describes APIs and application-level information for services
// Gloo routes to. ServiceSpec is contained within the UpstreamSpec for certain types
// of upstreams, including Kubernetes, Consul, and Static.
// ServiceSpec configuration is opaque to Gloo and handled by Service Options.
type ServiceSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Note to developers: new Service plugins must be added to this oneof field
	// to be usable by Gloo. (plugins currently need to be compiled into Gloo)
	//
	// Types that are assignable to PluginType:
	//	*ServiceSpec_Rest
	//	*ServiceSpec_Grpc
	PluginType isServiceSpec_PluginType `protobuf_oneof:"plugin_type"`
}

func (x *ServiceSpec) Reset() {
	*x = ServiceSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceSpec) ProtoMessage() {}

func (x *ServiceSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceSpec.ProtoReflect.Descriptor instead.
func (*ServiceSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_rawDescGZIP(), []int{0}
}

func (m *ServiceSpec) GetPluginType() isServiceSpec_PluginType {
	if m != nil {
		return m.PluginType
	}
	return nil
}

func (x *ServiceSpec) GetRest() *rest.ServiceSpec {
	if x, ok := x.GetPluginType().(*ServiceSpec_Rest); ok {
		return x.Rest
	}
	return nil
}

func (x *ServiceSpec) GetGrpc() *grpc.ServiceSpec {
	if x, ok := x.GetPluginType().(*ServiceSpec_Grpc); ok {
		return x.Grpc
	}
	return nil
}

type isServiceSpec_PluginType interface {
	isServiceSpec_PluginType()
}

type ServiceSpec_Rest struct {
	Rest *rest.ServiceSpec `protobuf:"bytes,1,opt,name=rest,proto3,oneof"`
}

type ServiceSpec_Grpc struct {
	Grpc *grpc.ServiceSpec `protobuf:"bytes,2,opt,name=grpc,proto3,oneof"`
}

func (*ServiceSpec_Rest) isServiceSpec_PluginType() {}

func (*ServiceSpec_Grpc) isServiceSpec_PluginType() {}

var File_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_rawDesc = []byte{
	0x0a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x1a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x3c, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x04, 0x72, 0x65, 0x73, 0x74,
	0x12, 0x3c, 0x0a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x42, 0x0d,
	0x0a, 0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x46, 0x5a,
	0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xb8, 0xf5, 0x04,
	0x01, 0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_goTypes = []interface{}{
	(*ServiceSpec)(nil),      // 0: options.gloo.solo.io.ServiceSpec
	(*rest.ServiceSpec)(nil), // 1: rest.options.gloo.solo.io.ServiceSpec
	(*grpc.ServiceSpec)(nil), // 2: grpc.options.gloo.solo.io.ServiceSpec
}
var file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_depIdxs = []int32{
	1, // 0: options.gloo.solo.io.ServiceSpec.rest:type_name -> rest.options.gloo.solo.io.ServiceSpec
	2, // 1: options.gloo.solo.io.ServiceSpec.grpc:type_name -> grpc.options.gloo.solo.io.ServiceSpec
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_init() }
func file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceSpec); i {
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
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ServiceSpec_Rest)(nil),
		(*ServiceSpec_Grpc)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_proto_depIdxs = nil
}
