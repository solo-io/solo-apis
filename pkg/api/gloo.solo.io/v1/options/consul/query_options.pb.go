// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/consul/query_options.proto

package consul

import (
	_ "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

// These are the same consistency modes offered by Consul. For more information please review https://www.consul.io/api-docs/features/consistency.
// and https://pkg.go.dev/github.com/hashicorp/consul/api#QueryOptions.
//
// Note: Gloo handles staleness well (as it runs update loops ~ once/second) but makes many requests
// to get consul endpoints so users may want to opt into stale reads once the implications are understood.
type ConsulConsistencyModes int32

const (
	// default - Consul HTTP API queries use default mode by default.
	// It is strongly consistent in almost all cases. However, there is a small
	// window in which a new leader may be elected during which the old leader may
	// respond with stale values. The trade-off is fast reads but potentially stale
	// values. The condition resulting in stale reads is hard to trigger, and most
	// clients should not need to worry about this case. Also, note that this race
	// condition only applies to reads, not writes.
	//
	// For more, see https://www.consul.io/api-docs/features/consistency#consul-http-api-queries
	ConsulConsistencyModes_DefaultMode ConsulConsistencyModes = 0
	// AllowStale allows any Consul server (non-leader) to service
	// a read. This allows for lower latency and higher throughput
	//
	// stale - Consul DNS queries use stale mode by default. This mode allows any
	// server to handle the read regardless of whether it is the leader. The trade-off
	// is very fast and scalable reads with a higher likelihood of stale values. Results
	// are generally consistent to within 50 milliseconds of the leader, though there is
	// no upper limit on this staleness. Since this mode allows reads without a leader, a
	// cluster that is unavailable (no quorum) can still respond to queries.
	//
	// Gloo handles staleness well (as it runs update loops ~ once/second) so users may want
	// to opt into this.
	ConsulConsistencyModes_StaleMode ConsulConsistencyModes = 1
	// RequireConsistent forces the read to be fully consistent.
	// This is more expensive but prevents ever performing a stale
	// read.
	//
	// consistent - This mode is strongly consistent without caveats. It requires that a
	// leader verify with a quorum of peers that it is still leader. This introduces an
	// additional round-trip to all server nodes. The trade-off is increased latency due
	// to an extra round trip. Most clients should not use this unless they cannot tolerate
	// a stale read.
	//
	// Gloo handles staleness well (as it runs update loops ~ once/second) so users may want
	// to strongly consider their requirements before enabling this.
	ConsulConsistencyModes_ConsistentMode ConsulConsistencyModes = 2
)

// Enum value maps for ConsulConsistencyModes.
var (
	ConsulConsistencyModes_name = map[int32]string{
		0: "DefaultMode",
		1: "StaleMode",
		2: "ConsistentMode",
	}
	ConsulConsistencyModes_value = map[string]int32{
		"DefaultMode":    0,
		"StaleMode":      1,
		"ConsistentMode": 2,
	}
)

func (x ConsulConsistencyModes) Enum() *ConsulConsistencyModes {
	p := new(ConsulConsistencyModes)
	*p = x
	return p
}

func (x ConsulConsistencyModes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConsulConsistencyModes) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_enumTypes[0].Descriptor()
}

func (ConsulConsistencyModes) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_enumTypes[0]
}

func (x ConsulConsistencyModes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConsulConsistencyModes.Descriptor instead.
func (ConsulConsistencyModes) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_rawDescGZIP(), []int{0}
}

// somewhat mirrors client query options struct in consul catalog api
// only has options that we know we want configurable at both upstream and settings-wide discovery levels
type QueryOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UseCache requests that the agent cache results locally. See
	// https://www.consul.io/api/features/caching.html for more details on the
	// semantics.
	//
	// Defaults to true.
	// opts users into background refresh caching https://www.consul.io/api-docs/features/caching#background-refresh-caching
	UseCache *wrappers.BoolValue `protobuf:"bytes,1,opt,name=use_cache,json=useCache,proto3" json:"use_cache,omitempty"`
}

func (x *QueryOptions) Reset() {
	*x = QueryOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryOptions) ProtoMessage() {}

func (x *QueryOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryOptions.ProtoReflect.Descriptor instead.
func (*QueryOptions) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_rawDescGZIP(), []int{0}
}

func (x *QueryOptions) GetUseCache() *wrappers.BoolValue {
	if x != nil {
		return x.UseCache
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_rawDesc = []byte{
	0x0a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x37, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x08, 0x75, 0x73, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x2a, 0x4c, 0x0a, 0x16,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63,
	0x79, 0x4d, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x4d, 0x6f, 0x64, 0x65, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x6c, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x10, 0x02, 0x42, 0x51, 0xb8, 0xf5, 0x04, 0x01,
	0xc0, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_goTypes = []interface{}{
	(ConsulConsistencyModes)(0), // 0: consul.options.gloo.solo.io.ConsulConsistencyModes
	(*QueryOptions)(nil),        // 1: consul.options.gloo.solo.io.QueryOptions
	(*wrappers.BoolValue)(nil),  // 2: google.protobuf.BoolValue
}
var file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_depIdxs = []int32{
	2, // 0: consul.options.gloo.solo.io.QueryOptions.use_cache:type_name -> google.protobuf.BoolValue
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryOptions); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_consul_query_options_proto_depIdxs = nil
}
