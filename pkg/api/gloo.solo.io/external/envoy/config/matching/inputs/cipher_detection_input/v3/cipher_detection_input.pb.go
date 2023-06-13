// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/config/matching/inputs/cipher_detection_input/v3/cipher_detection_input.proto

package v3

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Specifies that matching should be performed by the destination IP address.
// [#extension: io.solo.matching.inputs.cipher_detection_input]
type CipherDetectionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of TLS ciphers to send to the passthrough cipher chain. Ciphers must
	// be provided using their 16-bit value. These numbers of IANA standardised
	// values for all possible cipher suites and can be found here:
	// https://www.iana.org/assignments/tls-parameters/tls-parameters.xhtml
	//
	// example: to select the following cipher
	//
	//	0x00,0x3c TLS_RSA_WITH_AES_128_CBC_SHA256
	//
	// this field should be set to 0x003c
	PassthroughCiphers []uint32 `protobuf:"varint,1,rep,packed,name=passthrough_ciphers,json=passthroughCiphers,proto3" json:"passthrough_ciphers,omitempty"`
	// Terminating ciphers are those that should be considered as supported.
	// If not specified, defaults to whatever the current envoy implementation
	// finds to be the default set of ciphers.
	TerminatingCiphers []uint32 `protobuf:"varint,2,rep,packed,name=terminating_ciphers,json=terminatingCiphers,proto3" json:"terminating_ciphers,omitempty"`
}

func (x *CipherDetectionInput) Reset() {
	*x = CipherDetectionInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CipherDetectionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CipherDetectionInput) ProtoMessage() {}

func (x *CipherDetectionInput) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CipherDetectionInput.ProtoReflect.Descriptor instead.
func (*CipherDetectionInput) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_rawDescGZIP(), []int{0}
}

func (x *CipherDetectionInput) GetPassthroughCiphers() []uint32 {
	if x != nil {
		return x.PassthroughCiphers
	}
	return nil
}

func (x *CipherDetectionInput) GetTerminatingCiphers() []uint32 {
	if x != nil {
		return x.TerminatingCiphers
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_rawDesc = []byte{
	0x0a, 0x87, 0x01, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x73, 0x2f, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x69,
	0x70, 0x68, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e,
	0x67, 0x2e, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x33, 0x1a, 0x1d, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x14, 0x43, 0x69, 0x70, 0x68, 0x65,
	0x72, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x2f, 0x0a, 0x13, 0x70, 0x61, 0x73, 0x73, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x5f, 0x63,
	0x69, 0x70, 0x68, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x70, 0x61,
	0x73, 0x73, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x73,
	0x12, 0x2f, 0x0a, 0x13, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72,
	0x73, 0x42, 0xdc, 0x01, 0x0a, 0x3d, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72,
	0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x2e, 0x76, 0x33, 0x42, 0x12, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x71, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x73, 0x2f, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2f, 0x76, 0x33, 0xe2, 0xb5, 0xdf, 0xcb,
	0x07, 0x02, 0x10, 0x02, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_goTypes = []interface{}{
	(*CipherDetectionInput)(nil), // 0: envoy.config.matching.cipher_detection_input.v3.CipherDetectionInput
}
var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CipherDetectionInput); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_config_matching_inputs_cipher_detection_input_v3_cipher_detection_input_proto_depIdxs = nil
}
