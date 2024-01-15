// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: envoy_codegen_imports/github.com/solo-io/envoy-gloo-ee/api/envoy/config/matching/inputs/http_inputs/v3/http_inputs.proto

package v3

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/metadata/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Provides the route name as an input
// [#extension: io.solo.matching.inputs.http_inputs]
type RouteNameInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RouteNameInput) Reset() {
	*x = RouteNameInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteNameInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteNameInput) ProtoMessage() {}

func (x *RouteNameInput) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteNameInput.ProtoReflect.Descriptor instead.
func (*RouteNameInput) Descriptor() ([]byte, []int) {
	return file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_rawDescGZIP(), []int{0}
}

// Metadata for input. Must be string type.
type MetadataInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetadataKey *v3.MetadataKey `protobuf:"bytes,1,opt,name=metadata_key,json=metadataKey,proto3" json:"metadata_key,omitempty"`
	// value to return if the path exists but not a string
	OnValueNotString *wrappers.StringValue `protobuf:"bytes,2,opt,name=on_value_not_string,json=onValueNotString,proto3" json:"on_value_not_string,omitempty"`
	OnMissingValue   *wrappers.StringValue `protobuf:"bytes,3,opt,name=on_missing_value,json=onMissingValue,proto3" json:"on_missing_value,omitempty"`
}

func (x *MetadataInput) Reset() {
	*x = MetadataInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataInput) ProtoMessage() {}

func (x *MetadataInput) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataInput.ProtoReflect.Descriptor instead.
func (*MetadataInput) Descriptor() ([]byte, []int) {
	return file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_rawDescGZIP(), []int{1}
}

func (x *MetadataInput) GetMetadataKey() *v3.MetadataKey {
	if x != nil {
		return x.MetadataKey
	}
	return nil
}

func (x *MetadataInput) GetOnValueNotString() *wrappers.StringValue {
	if x != nil {
		return x.OnValueNotString
	}
	return nil
}

func (x *MetadataInput) GetOnMissingValue() *wrappers.StringValue {
	if x != nil {
		return x.OnMissingValue
	}
	return nil
}

var File_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto protoreflect.FileDescriptor

var file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_rawDesc = []byte{
	0x0a, 0x78, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x5f,
	0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2d, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x65, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2b, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e,
	0x67, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x33, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0xec, 0x01, 0x0a, 0x0d, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x46, 0x0a, 0x0c, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x4b, 0x65, 0x79, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4b,
	0x65, 0x79, 0x12, 0x4b, 0x0a, 0x13, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6e,
	0x6f, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x6f,
	0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4e, 0x6f, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x46, 0x0a, 0x10, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x6f, 0x6e, 0x4d, 0x69, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0xc9, 0x01, 0x0a, 0x39, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x73, 0x2e, 0x76, 0x33, 0x42, 0x13, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2d, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x65, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x68, 0x74, 0x74,
	0x70, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06,
	0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_rawDescOnce sync.Once
	file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_rawDescData = file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_rawDesc
)

func file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_rawDescGZIP() []byte {
	file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_rawDescOnce.Do(func() {
		file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_rawDescData)
	})
	return file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_rawDescData
}

var file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_goTypes = []interface{}{
	(*RouteNameInput)(nil),       // 0: envoy.config.matching.inputs.http_inputs.v3.RouteNameInput
	(*MetadataInput)(nil),        // 1: envoy.config.matching.inputs.http_inputs.v3.MetadataInput
	(*v3.MetadataKey)(nil),       // 2: envoy.type.metadata.v3.MetadataKey
	(*wrappers.StringValue)(nil), // 3: google.protobuf.StringValue
}
var file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_depIdxs = []int32{
	2, // 0: envoy.config.matching.inputs.http_inputs.v3.MetadataInput.metadata_key:type_name -> envoy.type.metadata.v3.MetadataKey
	3, // 1: envoy.config.matching.inputs.http_inputs.v3.MetadataInput.on_value_not_string:type_name -> google.protobuf.StringValue
	3, // 2: envoy.config.matching.inputs.http_inputs.v3.MetadataInput.on_missing_value:type_name -> google.protobuf.StringValue
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_init()
}
func file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_init() {
	if File_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteNameInput); i {
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
		file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataInput); i {
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
			RawDescriptor: file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_goTypes,
		DependencyIndexes: file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_depIdxs,
		MessageInfos:      file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_msgTypes,
	}.Build()
	File_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto = out.File
	file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_rawDesc = nil
	file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_goTypes = nil
	file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_ee_api_envoy_config_matching_inputs_http_inputs_v3_http_inputs_proto_depIdxs = nil
}
