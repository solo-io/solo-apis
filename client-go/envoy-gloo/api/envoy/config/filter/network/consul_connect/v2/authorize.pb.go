// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: envoy_codegen_imports/github.com/solo-io/envoy-gloo/api/envoy/config/filter/network/consul_connect/v2/authorize.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/golang/protobuf/ptypes/duration"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// [#proto-status: experimental]
type AuthorizePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target           string `protobuf:"bytes,1,opt,name=Target,proto3" json:"Target,omitempty"`
	ClientCertURI    string `protobuf:"bytes,2,opt,name=ClientCertURI,proto3" json:"ClientCertURI,omitempty"`
	ClientCertSerial string `protobuf:"bytes,3,opt,name=ClientCertSerial,proto3" json:"ClientCertSerial,omitempty"`
}

func (x *AuthorizePayload) Reset() {
	*x = AuthorizePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizePayload) ProtoMessage() {}

func (x *AuthorizePayload) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizePayload.ProtoReflect.Descriptor instead.
func (*AuthorizePayload) Descriptor() ([]byte, []int) {
	return file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorizePayload) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *AuthorizePayload) GetClientCertURI() string {
	if x != nil {
		return x.ClientCertURI
	}
	return ""
}

func (x *AuthorizePayload) GetClientCertSerial() string {
	if x != nil {
		return x.ClientCertSerial
	}
	return ""
}

// [#proto-status: experimental]
type AuthorizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authorized bool   `protobuf:"varint,1,opt,name=Authorized,proto3" json:"Authorized,omitempty"`
	Reason     string `protobuf:"bytes,2,opt,name=Reason,proto3" json:"Reason,omitempty"`
}

func (x *AuthorizeResponse) Reset() {
	*x = AuthorizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeResponse) ProtoMessage() {}

func (x *AuthorizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeResponse.ProtoReflect.Descriptor instead.
func (*AuthorizeResponse) Descriptor() ([]byte, []int) {
	return file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorizeResponse) GetAuthorized() bool {
	if x != nil {
		return x.Authorized
	}
	return false
}

func (x *AuthorizeResponse) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

var File_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto protoreflect.FileDescriptor

var file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_rawDesc = []byte{
	0x0a, 0x75, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x5f,
	0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2d, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5f, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x24, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x55, 0x52, 0x49,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65,
	0x72, 0x74, 0x55, 0x52, 0x49, 0x12, 0x2a, 0x0a, 0x10, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43,
	0x65, 0x72, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x22, 0x4b, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0xaa,
	0x01, 0x0a, 0x28, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2d, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_rawDescOnce sync.Once
	file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_rawDescData = file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_rawDesc
)

func file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_rawDescGZIP() []byte {
	file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_rawDescOnce.Do(func() {
		file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_rawDescData)
	})
	return file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_rawDescData
}

var file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_goTypes = []interface{}{
	(*AuthorizePayload)(nil),  // 0: agent.connect.authorize.v1.AuthorizePayload
	(*AuthorizeResponse)(nil), // 1: agent.connect.authorize.v1.AuthorizeResponse
}
var file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_init()
}
func file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_init() {
	if File_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizePayload); i {
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
		file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeResponse); i {
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
			RawDescriptor: file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_goTypes,
		DependencyIndexes: file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_depIdxs,
		MessageInfos:      file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_msgTypes,
	}.Build()
	File_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto = out.File
	file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_rawDesc = nil
	file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_goTypes = nil
	file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_authorize_proto_depIdxs = nil
}
