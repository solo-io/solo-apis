// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/protocol_upgrade/protocol_upgrade.proto

package protocol_upgrade

import (
	reflect "reflect"
	sync "sync"

	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProtocolUpgradeConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to UpgradeType:
	//
	//	*ProtocolUpgradeConfig_Websocket
	UpgradeType isProtocolUpgradeConfig_UpgradeType `protobuf_oneof:"upgrade_type"`
}

func (x *ProtocolUpgradeConfig) Reset() {
	*x = ProtocolUpgradeConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolUpgradeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolUpgradeConfig) ProtoMessage() {}

func (x *ProtocolUpgradeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolUpgradeConfig.ProtoReflect.Descriptor instead.
func (*ProtocolUpgradeConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_rawDescGZIP(), []int{0}
}

func (m *ProtocolUpgradeConfig) GetUpgradeType() isProtocolUpgradeConfig_UpgradeType {
	if m != nil {
		return m.UpgradeType
	}
	return nil
}

func (x *ProtocolUpgradeConfig) GetWebsocket() *ProtocolUpgradeConfig_ProtocolUpgradeSpec {
	if x, ok := x.GetUpgradeType().(*ProtocolUpgradeConfig_Websocket); ok {
		return x.Websocket
	}
	return nil
}

type isProtocolUpgradeConfig_UpgradeType interface {
	isProtocolUpgradeConfig_UpgradeType()
}

type ProtocolUpgradeConfig_Websocket struct {
	// Specification for websocket upgrade requests.
	Websocket *ProtocolUpgradeConfig_ProtocolUpgradeSpec `protobuf:"bytes,1,opt,name=websocket,proto3,oneof"`
}

func (*ProtocolUpgradeConfig_Websocket) isProtocolUpgradeConfig_UpgradeType() {}

type ProtocolUpgradeConfig_ProtocolUpgradeSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the upgrade should be enabled. If left unset, Envoy will enable the protocol upgrade.
	Enabled *wrappers.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *ProtocolUpgradeConfig_ProtocolUpgradeSpec) Reset() {
	*x = ProtocolUpgradeConfig_ProtocolUpgradeSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolUpgradeConfig_ProtocolUpgradeSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolUpgradeConfig_ProtocolUpgradeSpec) ProtoMessage() {}

func (x *ProtocolUpgradeConfig_ProtocolUpgradeSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolUpgradeConfig_ProtocolUpgradeSpec.ProtoReflect.Descriptor instead.
func (*ProtocolUpgradeConfig_ProtocolUpgradeSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ProtocolUpgradeConfig_ProtocolUpgradeSpec) GetEnabled() *wrappers.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_rawDesc = []byte{
	0x0a, 0x5d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x5f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x5f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x01, 0x0a, 0x15, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x70, 0x0a, 0x09, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x5f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x55, 0x70,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x09, 0x77, 0x65, 0x62,
	0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x4b, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x34, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x42, 0x57, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0x5a, 0x4d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_goTypes = []interface{}{
	(*ProtocolUpgradeConfig)(nil),                     // 0: protocol_upgrade.options.gloo.solo.io.ProtocolUpgradeConfig
	(*ProtocolUpgradeConfig_ProtocolUpgradeSpec)(nil), // 1: protocol_upgrade.options.gloo.solo.io.ProtocolUpgradeConfig.ProtocolUpgradeSpec
	(*wrappers.BoolValue)(nil),                        // 2: google.protobuf.BoolValue
}
var file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_depIdxs = []int32{
	1, // 0: protocol_upgrade.options.gloo.solo.io.ProtocolUpgradeConfig.websocket:type_name -> protocol_upgrade.options.gloo.solo.io.ProtocolUpgradeConfig.ProtocolUpgradeSpec
	2, // 1: protocol_upgrade.options.gloo.solo.io.ProtocolUpgradeConfig.ProtocolUpgradeSpec.enabled:type_name -> google.protobuf.BoolValue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolUpgradeConfig); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolUpgradeConfig_ProtocolUpgradeSpec); i {
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
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ProtocolUpgradeConfig_Websocket)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_protocol_upgrade_protocol_upgrade_proto_depIdxs = nil
}
