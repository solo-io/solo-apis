// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/config/filter/listener/tls_cipher_inspector/v3/tls_cipher_inspector.proto

package v3

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/udpa/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TlsCipherInspector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TlsCipherInspector) Reset() {
	*x = TlsCipherInspector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TlsCipherInspector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TlsCipherInspector) ProtoMessage() {}

func (x *TlsCipherInspector) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TlsCipherInspector.ProtoReflect.Descriptor instead.
func (*TlsCipherInspector) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_rawDescGZIP(), []int{0}
}

var File_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_rawDesc = []byte{
	0x0a, 0x83, 0x01, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x2f, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x5f, 0x69,
	0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x74, 0x6c, 0x73, 0x5f,
	0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x34, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x6c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x2e, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x1d, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x14, 0x0a, 0x12, 0x54, 0x6c, 0x73, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x70,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x42, 0xe4, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01,
	0xd0, 0xf5, 0x04, 0x01, 0xe2, 0xb5, 0xdf, 0xcb, 0x07, 0x02, 0x10, 0x02, 0x0a, 0x42, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x69, 0x70,
	0x68, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x33,
	0x42, 0x17, 0x54, 0x6c, 0x73, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x2f, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_goTypes = []interface{}{
	(*TlsCipherInspector)(nil), // 0: envoy.config.filter.listener.tls_cipher_inspector.v3.TlsCipherInspector
}
var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TlsCipherInspector); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_filter_listener_tls_cipher_inspector_v3_tls_cipher_inspector_proto_depIdxs = nil
}
