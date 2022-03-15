// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/type/matcher/v3/string.proto

package v3

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/annotations"
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

// Specifies the way to match a string.
// [#next-free-field: 7]
type StringMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MatchPattern:
	//	*StringMatcher_Exact
	//	*StringMatcher_Prefix
	//	*StringMatcher_Suffix
	//	*StringMatcher_SafeRegex
	MatchPattern isStringMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	// If true, indicates the exact/prefix/suffix matching should be case insensitive. This has no
	// effect for the safe_regex match.
	// For example, the matcher *data* will match both input string *Data* and *data* if set to true.
	IgnoreCase bool `protobuf:"varint,6,opt,name=ignore_case,json=ignoreCase,proto3" json:"ignore_case,omitempty"`
}

func (x *StringMatcher) Reset() {
	*x = StringMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringMatcher) ProtoMessage() {}

func (x *StringMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringMatcher.ProtoReflect.Descriptor instead.
func (*StringMatcher) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_rawDescGZIP(), []int{0}
}

func (m *StringMatcher) GetMatchPattern() isStringMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (x *StringMatcher) GetExact() string {
	if x, ok := x.GetMatchPattern().(*StringMatcher_Exact); ok {
		return x.Exact
	}
	return ""
}

func (x *StringMatcher) GetPrefix() string {
	if x, ok := x.GetMatchPattern().(*StringMatcher_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (x *StringMatcher) GetSuffix() string {
	if x, ok := x.GetMatchPattern().(*StringMatcher_Suffix); ok {
		return x.Suffix
	}
	return ""
}

func (x *StringMatcher) GetSafeRegex() *RegexMatcher {
	if x, ok := x.GetMatchPattern().(*StringMatcher_SafeRegex); ok {
		return x.SafeRegex
	}
	return nil
}

func (x *StringMatcher) GetIgnoreCase() bool {
	if x != nil {
		return x.IgnoreCase
	}
	return false
}

type isStringMatcher_MatchPattern interface {
	isStringMatcher_MatchPattern()
}

type StringMatcher_Exact struct {
	// The input string must match exactly the string specified here.
	//
	// Examples:
	//
	// * *abc* only matches the value *abc*.
	Exact string `protobuf:"bytes,1,opt,name=exact,proto3,oneof"`
}

type StringMatcher_Prefix struct {
	// The input string must have the prefix specified here.
	// Note: empty prefix is not allowed, please use regex instead.
	//
	// Examples:
	//
	// * *abc* matches the value *abc.xyz*
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3,oneof"`
}

type StringMatcher_Suffix struct {
	// The input string must have the suffix specified here.
	// Note: empty prefix is not allowed, please use regex instead.
	//
	// Examples:
	//
	// * *abc* matches the value *xyz.abc*
	Suffix string `protobuf:"bytes,3,opt,name=suffix,proto3,oneof"`
}

type StringMatcher_SafeRegex struct {
	// The input string must match the regular expression specified here.
	SafeRegex *RegexMatcher `protobuf:"bytes,5,opt,name=safe_regex,json=safeRegex,proto3,oneof"`
}

func (*StringMatcher_Exact) isStringMatcher_MatchPattern() {}

func (*StringMatcher_Prefix) isStringMatcher_MatchPattern() {}

func (*StringMatcher_Suffix) isStringMatcher_MatchPattern() {}

func (*StringMatcher_SafeRegex) isStringMatcher_MatchPattern() {}

// Specifies a list of ways to match a string.
type ListStringMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Patterns []*StringMatcher `protobuf:"bytes,1,rep,name=patterns,proto3" json:"patterns,omitempty"`
}

func (x *ListStringMatcher) Reset() {
	*x = ListStringMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStringMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStringMatcher) ProtoMessage() {}

func (x *ListStringMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStringMatcher.ProtoReflect.Descriptor instead.
func (*ListStringMatcher) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_rawDescGZIP(), []int{1}
}

func (x *ListStringMatcher) GetPatterns() []*StringMatcher {
	if x != nil {
		return x.Patterns
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_rawDesc = []byte{
	0x0a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f,
	0x76, 0x33, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb3, 0x02, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x78, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x78, 0x61, 0x63, 0x74, 0x12, 0x21, 0x0a, 0x06,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x20, 0x01, 0x48, 0x00, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12,
	0x21, 0x0a, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x48, 0x00, 0x52, 0x06, 0x73, 0x75, 0x66, 0x66,
	0x69, 0x78, 0x12, 0x56, 0x0a, 0x0a, 0x73, 0x61, 0x66, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x67, 0x65, 0x78, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52,
	0x09, 0x73, 0x61, 0x66, 0x65, 0x52, 0x65, 0x67, 0x65, 0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x67,
	0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x73, 0x65, 0x3a, 0x28, 0x8a, 0xc8, 0xde,
	0x8e, 0x04, 0x22, 0x0a, 0x20, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0x14, 0x0a, 0x0d, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x70,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x4a, 0x04, 0x08, 0x04, 0x10,
	0x05, 0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x22, 0x95, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x52,
	0x0a, 0x08, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x08, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x6e, 0x73, 0x3a, 0x2c, 0x8a, 0xc8, 0xde, 0x8e, 0x04, 0x26, 0x0a, 0x24, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x42, 0xa2, 0x01, 0x0a, 0x2b, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33,
	0x42, 0x0b, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76,
	0x33, 0xe2, 0xb5, 0xdf, 0xcb, 0x07, 0x02, 0x10, 0x02, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04,
	0x01, 0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_goTypes = []interface{}{
	(*StringMatcher)(nil),     // 0: solo.io.envoy.type.matcher.v3.StringMatcher
	(*ListStringMatcher)(nil), // 1: solo.io.envoy.type.matcher.v3.ListStringMatcher
	(*RegexMatcher)(nil),      // 2: solo.io.envoy.type.matcher.v3.RegexMatcher
}
var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_depIdxs = []int32{
	2, // 0: solo.io.envoy.type.matcher.v3.StringMatcher.safe_regex:type_name -> solo.io.envoy.type.matcher.v3.RegexMatcher
	0, // 1: solo.io.envoy.type.matcher.v3.ListStringMatcher.patterns:type_name -> solo.io.envoy.type.matcher.v3.StringMatcher
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto != nil {
		return
	}
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_regex_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringMatcher); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStringMatcher); i {
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
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*StringMatcher_Exact)(nil),
		(*StringMatcher_Prefix)(nil),
		(*StringMatcher_Suffix)(nil),
		(*StringMatcher_SafeRegex)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_type_matcher_v3_string_proto_depIdxs = nil
}
