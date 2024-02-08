// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: envoy_codegen_imports/github.com/solo-io/envoy-gloo/api/envoy/config/filter/network/consul_connect/v2/consul_connect.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	duration "github.com/golang/protobuf/ptypes/duration"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ConsulConnect contains the configuration necessary to perform client
// certificate-based authorization using a REST call to the Authorize endpoint
// of Consul Connect.
type ConsulConnect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the service who owns this proxy
	// Target must be delivered by the filter as part of the authorize request
	// payload
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// the hostname of the authorization REST service
	AuthorizeHostname string `protobuf:"bytes,2,opt,name=authorize_hostname,json=authorizeHostname,proto3" json:"authorize_hostname,omitempty"`
	// the name of the Envoy Cluster representing the authorization REST service
	AuthorizeClusterName string `protobuf:"bytes,3,opt,name=authorize_cluster_name,json=authorizeClusterName,proto3" json:"authorize_cluster_name,omitempty"`
	// Connection Timeout tells the filter to set a timeout for unresponsive
	// connections created to this upstream. If not provided by the user, it will
	// set to a default value
	RequestTimeout *duration.Duration `protobuf:"bytes,4,opt,name=request_timeout,json=requestTimeout,proto3" json:"request_timeout,omitempty"`
}

func (x *ConsulConnect) Reset() {
	*x = ConsulConnect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsulConnect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsulConnect) ProtoMessage() {}

func (x *ConsulConnect) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsulConnect.ProtoReflect.Descriptor instead.
func (*ConsulConnect) Descriptor() ([]byte, []int) {
	return file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_rawDescGZIP(), []int{0}
}

func (x *ConsulConnect) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *ConsulConnect) GetAuthorizeHostname() string {
	if x != nil {
		return x.AuthorizeHostname
	}
	return ""
}

func (x *ConsulConnect) GetAuthorizeClusterName() string {
	if x != nil {
		return x.AuthorizeClusterName
	}
	return ""
}

func (x *ConsulConnect) GetRequestTimeout() *duration.Duration {
	if x != nil {
		return x.RequestTimeout
	}
	return nil
}

var File_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto protoreflect.FileDescriptor

var file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_rawDesc = []byte{
	0x0a, 0x7a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x5f,
	0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2d, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5f, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x01, 0x0a, 0x0d, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x1f, 0x0a, 0x06,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x36, 0x0a,
	0x12, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x20, 0x01, 0x52, 0x11, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x48, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x16, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x14,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x04, 0x98, 0xdf, 0x1f, 0x01, 0x52, 0x0e,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42, 0xc4,
	0x01, 0x0a, 0x3b, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x32, 0x42, 0x12,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65,
	0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x32,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2d, 0x67,
	0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_rawDescOnce sync.Once
	file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_rawDescData = file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_rawDesc
)

func file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_rawDescGZIP() []byte {
	file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_rawDescOnce.Do(func() {
		file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_rawDescData)
	})
	return file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_rawDescData
}

var file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_goTypes = []interface{}{
	(*ConsulConnect)(nil),     // 0: envoy.config.filter.network.consul_connect.v2.ConsulConnect
	(*duration.Duration)(nil), // 1: google.protobuf.Duration
}
var file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_depIdxs = []int32{
	1, // 0: envoy.config.filter.network.consul_connect.v2.ConsulConnect.request_timeout:type_name -> google.protobuf.Duration
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_init()
}
func file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_init() {
	if File_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsulConnect); i {
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
			RawDescriptor: file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_goTypes,
		DependencyIndexes: file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_depIdxs,
		MessageInfos:      file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_msgTypes,
	}.Build()
	File_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto = out.File
	file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_rawDesc = nil
	file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_goTypes = nil
	file_envoy_codegen_imports_github_com_solo_io_envoy_gloo_api_envoy_config_filter_network_consul_connect_v2_consul_connect_proto_depIdxs = nil
}
