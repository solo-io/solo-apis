// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo-mesh/external/envoyproxy/data-plane-api/envoy/type/matcher/v3/number.proto

package v3

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/solo-io/solo-apis/api/gloo-mesh/external/cncf/udpa/udpa/annotations"
	v3 "github.com/solo-io/solo-apis/api/gloo-mesh/external/envoyproxy/go-control-plane/envoy/type/v3"
	_ "github.com/solo-io/solo-apis/api/gloo-mesh/external/envoyproxy/protoc-gen-validate/validate"
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

// Specifies the way to match a double value.
type DoubleMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MatchPattern:
	//	*DoubleMatcher_Range
	//	*DoubleMatcher_Exact
	MatchPattern isDoubleMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
}

func (x *DoubleMatcher) Reset() {
	*x = DoubleMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoubleMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoubleMatcher) ProtoMessage() {}

func (x *DoubleMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoubleMatcher.ProtoReflect.Descriptor instead.
func (*DoubleMatcher) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_rawDescGZIP(), []int{0}
}

func (m *DoubleMatcher) GetMatchPattern() isDoubleMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (x *DoubleMatcher) GetRange() *v3.DoubleRange {
	if x, ok := x.GetMatchPattern().(*DoubleMatcher_Range); ok {
		return x.Range
	}
	return nil
}

func (x *DoubleMatcher) GetExact() float64 {
	if x, ok := x.GetMatchPattern().(*DoubleMatcher_Exact); ok {
		return x.Exact
	}
	return 0
}

type isDoubleMatcher_MatchPattern interface {
	isDoubleMatcher_MatchPattern()
}

type DoubleMatcher_Range struct {
	// If specified, the input double value must be in the range specified here.
	// Note: The range is using half-open interval semantics [start, end).
	Range *v3.DoubleRange `protobuf:"bytes,1,opt,name=range,proto3,oneof"`
}

type DoubleMatcher_Exact struct {
	// If specified, the input double value must be equal to the value specified here.
	Exact float64 `protobuf:"fixed64,2,opt,name=exact,proto3,oneof"`
}

func (*DoubleMatcher_Range) isDoubleMatcher_MatchPattern() {}

func (*DoubleMatcher_Exact) isDoubleMatcher_MatchPattern() {}

var File_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_rawDesc = []byte{
	0x0a, 0x70, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x67, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x5b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6e, 0x63, 0x66, 0x2f, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x5f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6e, 0x63, 0x66, 0x2f, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x6a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a,
	0x0d, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x32,
	0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x00, 0x52, 0x05, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x78, 0x61, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x48, 0x00, 0x52, 0x05, 0x65, 0x78, 0x61, 0x63, 0x74, 0x3a, 0x27, 0x9a, 0xc5, 0x88, 0x1e,
	0x22, 0x0a, 0x20, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x42, 0x14, 0x0a, 0x0d, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x70, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x6e, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0xa3, 0x01, 0x0a, 0x23, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76,
	0x33, 0x42, 0x0b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x65, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_goTypes = []interface{}{
	(*DoubleMatcher)(nil),  // 0: envoy.type.matcher.v3.DoubleMatcher
	(*v3.DoubleRange)(nil), // 1: envoy.type.v3.DoubleRange
}
var file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_depIdxs = []int32{
	1, // 0: envoy.type.matcher.v3.DoubleMatcher.range:type_name -> envoy.type.v3.DoubleRange
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoubleMatcher); i {
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
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DoubleMatcher_Range)(nil),
		(*DoubleMatcher_Exact)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoyproxy_data_plane_api_envoy_type_matcher_v3_number_proto_depIdxs = nil
}
