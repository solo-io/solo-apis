// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo-mesh/external/envoy-gloo-ee/api/envoy/config/filter/http/cache/v2/cache_filter.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	v3 "github.com/solo-io/solo-apis/api/gloo-mesh/external/envoyproxy/go-control-plane/envoy/config/core/v3"
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

// [#proto-status: experimental]
type Cache struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The max content length of in-memory cached data for each worker thread.
	// The content length is calculated as the sum of lengths of both keys and
	// values.
	InMemoryMaxContentLength *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=in_memory_max_content_length,json=inMemoryMaxContentLength,proto3" json:"in_memory_max_content_length,omitempty"`
	GrpcService              *v3.GrpcService       `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	GrpcTimeout              *duration.Duration    `protobuf:"bytes,3,opt,name=grpc_timeout,json=grpcTimeout,proto3" json:"grpc_timeout,omitempty"`
	RedisCluster             string                `protobuf:"bytes,4,opt,name=redis_cluster,json=redisCluster,proto3" json:"redis_cluster,omitempty"`
	RedisStatPrefix          string                `protobuf:"bytes,5,opt,name=redis_stat_prefix,json=redisStatPrefix,proto3" json:"redis_stat_prefix,omitempty"`
	RedisOpTimeout           *duration.Duration    `protobuf:"bytes,6,opt,name=redis_op_timeout,json=redisOpTimeout,proto3" json:"redis_op_timeout,omitempty"`
	// [#not-implemented-hide:]
	MaxResponseContentLength *wrappers.UInt32Value `protobuf:"bytes,7,opt,name=max_response_content_length,json=maxResponseContentLength,proto3" json:"max_response_content_length,omitempty"`
}

func (x *Cache) Reset() {
	*x = Cache{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cache) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cache) ProtoMessage() {}

func (x *Cache) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cache.ProtoReflect.Descriptor instead.
func (*Cache) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_rawDescGZIP(), []int{0}
}

func (x *Cache) GetInMemoryMaxContentLength() *wrappers.UInt32Value {
	if x != nil {
		return x.InMemoryMaxContentLength
	}
	return nil
}

func (x *Cache) GetGrpcService() *v3.GrpcService {
	if x != nil {
		return x.GrpcService
	}
	return nil
}

func (x *Cache) GetGrpcTimeout() *duration.Duration {
	if x != nil {
		return x.GrpcTimeout
	}
	return nil
}

func (x *Cache) GetRedisCluster() string {
	if x != nil {
		return x.RedisCluster
	}
	return ""
}

func (x *Cache) GetRedisStatPrefix() string {
	if x != nil {
		return x.RedisStatPrefix
	}
	return ""
}

func (x *Cache) GetRedisOpTimeout() *duration.Duration {
	if x != nil {
		return x.RedisOpTimeout
	}
	return nil
}

func (x *Cache) GetMaxResponseContentLength() *wrappers.UInt32Value {
	if x != nil {
		return x.MaxResponseContentLength
	}
	return nil
}

type CacheFilterRouteConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsCacheable bool `protobuf:"varint,1,opt,name=is_cacheable,json=isCacheable,proto3" json:"is_cacheable,omitempty"`
}

func (x *CacheFilterRouteConfig) Reset() {
	*x = CacheFilterRouteConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheFilterRouteConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheFilterRouteConfig) ProtoMessage() {}

func (x *CacheFilterRouteConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheFilterRouteConfig.ProtoReflect.Descriptor instead.
func (*CacheFilterRouteConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_rawDescGZIP(), []int{1}
}

func (x *CacheFilterRouteConfig) GetIsCacheable() bool {
	if x != nil {
		return x.IsCacheable
	}
	return false
}

var File_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_rawDesc = []byte{
	0x0a, 0x7a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2d, 0x67, 0x6c, 0x6f, 0x6f,
	0x2d, 0x65, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x68, 0x74, 0x74, 0x70,
	0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x32, 0x1a,
	0x75, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x33, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x03, 0x0a, 0x05, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x12, 0x5c, 0x0a, 0x1c, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x61,
	0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x18, 0x69, 0x6e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4d, 0x61,
	0x78, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x44,
	0x0a, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x72, 0x70, 0x63,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x64, 0x69, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x73,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x64, 0x69, 0x73,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x53, 0x74, 0x61, 0x74, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x12, 0x43, 0x0a, 0x10, 0x72, 0x65, 0x64, 0x69, 0x73, 0x5f, 0x6f, 0x70, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x4f,
	0x70, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x5b, 0x0a, 0x1b, 0x6d, 0x61, 0x78, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x18, 0x6d, 0x61, 0x78,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x3b, 0x0a, 0x16, 0x43, 0x61, 0x63, 0x68, 0x65, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x61, 0x62,
	0x6c, 0x65, 0x42, 0x9f, 0x01, 0x0a, 0x2f, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x2e, 0x76, 0x32, 0x42, 0x10, 0x43, 0x61, 0x63, 0x68, 0x65, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2d, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x65, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_goTypes = []interface{}{
	(*Cache)(nil),                  // 0: envoy.config.filter.http.cache.v2.Cache
	(*CacheFilterRouteConfig)(nil), // 1: envoy.config.filter.http.cache.v2.CacheFilterRouteConfig
	(*wrappers.UInt32Value)(nil),   // 2: google.protobuf.UInt32Value
	(*v3.GrpcService)(nil),         // 3: envoy.config.core.v3.GrpcService
	(*duration.Duration)(nil),      // 4: google.protobuf.Duration
}
var file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_depIdxs = []int32{
	2, // 0: envoy.config.filter.http.cache.v2.Cache.in_memory_max_content_length:type_name -> google.protobuf.UInt32Value
	3, // 1: envoy.config.filter.http.cache.v2.Cache.grpc_service:type_name -> envoy.config.core.v3.GrpcService
	4, // 2: envoy.config.filter.http.cache.v2.Cache.grpc_timeout:type_name -> google.protobuf.Duration
	4, // 3: envoy.config.filter.http.cache.v2.Cache.redis_op_timeout:type_name -> google.protobuf.Duration
	2, // 4: envoy.config.filter.http.cache.v2.Cache.max_response_content_length:type_name -> google.protobuf.UInt32Value
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cache); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheFilterRouteConfig); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_envoy_gloo_ee_api_envoy_config_filter_http_cache_v2_cache_filter_proto_depIdxs = nil
}
