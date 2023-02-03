// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo-fed/fed.rpc/v1/failover_scheme.proto

package v1

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	types "github.com/solo-io/solo-apis/pkg/api/fed.solo.io/v1/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type FailoverScheme struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *ObjectMeta                 `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec     *types.FailoverSchemeSpec   `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Status   *types.FailoverSchemeStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *FailoverScheme) Reset() {
	*x = FailoverScheme{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailoverScheme) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailoverScheme) ProtoMessage() {}

func (x *FailoverScheme) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailoverScheme.ProtoReflect.Descriptor instead.
func (*FailoverScheme) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_rawDescGZIP(), []int{0}
}

func (x *FailoverScheme) GetMetadata() *ObjectMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *FailoverScheme) GetSpec() *types.FailoverSchemeSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *FailoverScheme) GetStatus() *types.FailoverSchemeStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type GetFailoverSchemeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpstreamRef *v1.ClusterObjectRef `protobuf:"bytes,1,opt,name=upstream_ref,json=upstreamRef,proto3" json:"upstream_ref,omitempty"`
}

func (x *GetFailoverSchemeRequest) Reset() {
	*x = GetFailoverSchemeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFailoverSchemeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFailoverSchemeRequest) ProtoMessage() {}

func (x *GetFailoverSchemeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFailoverSchemeRequest.ProtoReflect.Descriptor instead.
func (*GetFailoverSchemeRequest) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_rawDescGZIP(), []int{1}
}

func (x *GetFailoverSchemeRequest) GetUpstreamRef() *v1.ClusterObjectRef {
	if x != nil {
		return x.UpstreamRef
	}
	return nil
}

type GetFailoverSchemeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FailoverScheme *FailoverScheme `protobuf:"bytes,1,opt,name=failover_scheme,json=failoverScheme,proto3" json:"failover_scheme,omitempty"`
}

func (x *GetFailoverSchemeResponse) Reset() {
	*x = GetFailoverSchemeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFailoverSchemeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFailoverSchemeResponse) ProtoMessage() {}

func (x *GetFailoverSchemeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFailoverSchemeResponse.ProtoReflect.Descriptor instead.
func (*GetFailoverSchemeResponse) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_rawDescGZIP(), []int{2}
}

func (x *GetFailoverSchemeResponse) GetFailoverScheme() *FailoverScheme {
	if x != nil {
		return x.FailoverScheme
	}
	return nil
}

type GetFailoverSchemeYamlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FailoverSchemeRef *v1.ObjectRef `protobuf:"bytes,1,opt,name=failover_scheme_ref,json=failoverSchemeRef,proto3" json:"failover_scheme_ref,omitempty"`
}

func (x *GetFailoverSchemeYamlRequest) Reset() {
	*x = GetFailoverSchemeYamlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFailoverSchemeYamlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFailoverSchemeYamlRequest) ProtoMessage() {}

func (x *GetFailoverSchemeYamlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFailoverSchemeYamlRequest.ProtoReflect.Descriptor instead.
func (*GetFailoverSchemeYamlRequest) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_rawDescGZIP(), []int{3}
}

func (x *GetFailoverSchemeYamlRequest) GetFailoverSchemeRef() *v1.ObjectRef {
	if x != nil {
		return x.FailoverSchemeRef
	}
	return nil
}

type GetFailoverSchemeYamlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	YamlData *ResourceYaml `protobuf:"bytes,1,opt,name=yaml_data,json=yamlData,proto3" json:"yaml_data,omitempty"`
}

func (x *GetFailoverSchemeYamlResponse) Reset() {
	*x = GetFailoverSchemeYamlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFailoverSchemeYamlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFailoverSchemeYamlResponse) ProtoMessage() {}

func (x *GetFailoverSchemeYamlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFailoverSchemeYamlResponse.ProtoReflect.Descriptor instead.
func (*GetFailoverSchemeYamlResponse) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_rawDescGZIP(), []int{4}
}

func (x *GetFailoverSchemeYamlResponse) GetYamlData() *ResourceYaml {
	if x != nil {
		return x.YamlData
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x66, 0x65, 0x64, 0x2f, 0x66, 0x65, 0x64, 0x2e,
	0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x5f,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x66, 0x65,
	0x64, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65,
	0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x66, 0x65, 0x64, 0x2f, 0x66, 0x65, 0x64,
	0x2f, 0x76, 0x31, 0x2f, 0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x66, 0x65, 0x64, 0x2f, 0x66, 0x65,
	0x64, 0x2e, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x0e, 0x46, 0x61, 0x69, 0x6c, 0x6f, 0x76,
	0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x65, 0x64,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x33, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x61,
	0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x62, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a,
	0x0c, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x52, 0x0b, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x66, 0x22, 0x65, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x46, 0x61, 0x69, 0x6c,
	0x6f, 0x76, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x48, 0x0a, 0x0f, 0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x65,
	0x64, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x61,
	0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x0e, 0x66, 0x61,
	0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x22, 0x6c, 0x0a, 0x1c,
	0x47, 0x65, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x65, 0x59, 0x61, 0x6d, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x13,
	0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x5f,
	0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x52, 0x11, 0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65,
	0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x66, 0x22, 0x5b, 0x0a, 0x1d, 0x47, 0x65,
	0x74, 0x46, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x59,
	0x61, 0x6d, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x79,
	0x61, 0x6d, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x59, 0x61, 0x6d, 0x6c, 0x52, 0x08, 0x79,
	0x61, 0x6d, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x32, 0xfb, 0x01, 0x0a, 0x11, 0x46, 0x61, 0x69, 0x6c,
	0x6f, 0x76, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x41, 0x70, 0x69, 0x12, 0x6c, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x65, 0x12, 0x29, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x78, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65,
	0x59, 0x61, 0x6d, 0x6c, 0x12, 0x2d, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x6f, 0x76,
	0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x59, 0x61, 0x6d, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65,
	0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x59, 0x61, 0x6d, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x39, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x65,
	0x64, 0x2e, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_goTypes = []interface{}{
	(*FailoverScheme)(nil),                // 0: fed.rpc.solo.io.FailoverScheme
	(*GetFailoverSchemeRequest)(nil),      // 1: fed.rpc.solo.io.GetFailoverSchemeRequest
	(*GetFailoverSchemeResponse)(nil),     // 2: fed.rpc.solo.io.GetFailoverSchemeResponse
	(*GetFailoverSchemeYamlRequest)(nil),  // 3: fed.rpc.solo.io.GetFailoverSchemeYamlRequest
	(*GetFailoverSchemeYamlResponse)(nil), // 4: fed.rpc.solo.io.GetFailoverSchemeYamlResponse
	(*ObjectMeta)(nil),                    // 5: fed.rpc.solo.io.ObjectMeta
	(*types.FailoverSchemeSpec)(nil),      // 6: fed.solo.io.FailoverSchemeSpec
	(*types.FailoverSchemeStatus)(nil),    // 7: fed.solo.io.FailoverSchemeStatus
	(*v1.ClusterObjectRef)(nil),           // 8: core.skv2.solo.io.ClusterObjectRef
	(*v1.ObjectRef)(nil),                  // 9: core.skv2.solo.io.ObjectRef
	(*ResourceYaml)(nil),                  // 10: fed.rpc.solo.io.ResourceYaml
}
var file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_depIdxs = []int32{
	5,  // 0: fed.rpc.solo.io.FailoverScheme.metadata:type_name -> fed.rpc.solo.io.ObjectMeta
	6,  // 1: fed.rpc.solo.io.FailoverScheme.spec:type_name -> fed.solo.io.FailoverSchemeSpec
	7,  // 2: fed.rpc.solo.io.FailoverScheme.status:type_name -> fed.solo.io.FailoverSchemeStatus
	8,  // 3: fed.rpc.solo.io.GetFailoverSchemeRequest.upstream_ref:type_name -> core.skv2.solo.io.ClusterObjectRef
	0,  // 4: fed.rpc.solo.io.GetFailoverSchemeResponse.failover_scheme:type_name -> fed.rpc.solo.io.FailoverScheme
	9,  // 5: fed.rpc.solo.io.GetFailoverSchemeYamlRequest.failover_scheme_ref:type_name -> core.skv2.solo.io.ObjectRef
	10, // 6: fed.rpc.solo.io.GetFailoverSchemeYamlResponse.yaml_data:type_name -> fed.rpc.solo.io.ResourceYaml
	1,  // 7: fed.rpc.solo.io.FailoverSchemeApi.GetFailoverScheme:input_type -> fed.rpc.solo.io.GetFailoverSchemeRequest
	3,  // 8: fed.rpc.solo.io.FailoverSchemeApi.GetFailoverSchemeYaml:input_type -> fed.rpc.solo.io.GetFailoverSchemeYamlRequest
	2,  // 9: fed.rpc.solo.io.FailoverSchemeApi.GetFailoverScheme:output_type -> fed.rpc.solo.io.GetFailoverSchemeResponse
	4,  // 10: fed.rpc.solo.io.FailoverSchemeApi.GetFailoverSchemeYaml:output_type -> fed.rpc.solo.io.GetFailoverSchemeYamlResponse
	9,  // [9:11] is the sub-list for method output_type
	7,  // [7:9] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_init() }
func file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto != nil {
		return
	}
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailoverScheme); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFailoverSchemeRequest); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFailoverSchemeResponse); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFailoverSchemeYamlRequest); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFailoverSchemeYamlResponse); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_failover_scheme_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FailoverSchemeApiClient is the client API for FailoverSchemeApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FailoverSchemeApiClient interface {
	GetFailoverScheme(ctx context.Context, in *GetFailoverSchemeRequest, opts ...grpc.CallOption) (*GetFailoverSchemeResponse, error)
	GetFailoverSchemeYaml(ctx context.Context, in *GetFailoverSchemeYamlRequest, opts ...grpc.CallOption) (*GetFailoverSchemeYamlResponse, error)
}

type failoverSchemeApiClient struct {
	cc grpc.ClientConnInterface
}

func NewFailoverSchemeApiClient(cc grpc.ClientConnInterface) FailoverSchemeApiClient {
	return &failoverSchemeApiClient{cc}
}

func (c *failoverSchemeApiClient) GetFailoverScheme(ctx context.Context, in *GetFailoverSchemeRequest, opts ...grpc.CallOption) (*GetFailoverSchemeResponse, error) {
	out := new(GetFailoverSchemeResponse)
	err := c.cc.Invoke(ctx, "/fed.rpc.solo.io.FailoverSchemeApi/GetFailoverScheme", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *failoverSchemeApiClient) GetFailoverSchemeYaml(ctx context.Context, in *GetFailoverSchemeYamlRequest, opts ...grpc.CallOption) (*GetFailoverSchemeYamlResponse, error) {
	out := new(GetFailoverSchemeYamlResponse)
	err := c.cc.Invoke(ctx, "/fed.rpc.solo.io.FailoverSchemeApi/GetFailoverSchemeYaml", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FailoverSchemeApiServer is the server API for FailoverSchemeApi service.
type FailoverSchemeApiServer interface {
	GetFailoverScheme(context.Context, *GetFailoverSchemeRequest) (*GetFailoverSchemeResponse, error)
	GetFailoverSchemeYaml(context.Context, *GetFailoverSchemeYamlRequest) (*GetFailoverSchemeYamlResponse, error)
}

// UnimplementedFailoverSchemeApiServer can be embedded to have forward compatible implementations.
type UnimplementedFailoverSchemeApiServer struct {
}

func (*UnimplementedFailoverSchemeApiServer) GetFailoverScheme(context.Context, *GetFailoverSchemeRequest) (*GetFailoverSchemeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFailoverScheme not implemented")
}
func (*UnimplementedFailoverSchemeApiServer) GetFailoverSchemeYaml(context.Context, *GetFailoverSchemeYamlRequest) (*GetFailoverSchemeYamlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFailoverSchemeYaml not implemented")
}

func RegisterFailoverSchemeApiServer(s *grpc.Server, srv FailoverSchemeApiServer) {
	s.RegisterService(&_FailoverSchemeApi_serviceDesc, srv)
}

func _FailoverSchemeApi_GetFailoverScheme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFailoverSchemeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FailoverSchemeApiServer).GetFailoverScheme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fed.rpc.solo.io.FailoverSchemeApi/GetFailoverScheme",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FailoverSchemeApiServer).GetFailoverScheme(ctx, req.(*GetFailoverSchemeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FailoverSchemeApi_GetFailoverSchemeYaml_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFailoverSchemeYamlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FailoverSchemeApiServer).GetFailoverSchemeYaml(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fed.rpc.solo.io.FailoverSchemeApi/GetFailoverSchemeYaml",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FailoverSchemeApiServer).GetFailoverSchemeYaml(ctx, req.(*GetFailoverSchemeYamlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FailoverSchemeApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fed.rpc.solo.io.FailoverSchemeApi",
	HandlerType: (*FailoverSchemeApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFailoverScheme",
			Handler:    _FailoverSchemeApi_GetFailoverScheme_Handler,
		},
		{
			MethodName: "GetFailoverSchemeYaml",
			Handler:    _FailoverSchemeApi_GetFailoverSchemeYaml_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/solo-io/solo-apis/api/gloo-fed/fed.rpc/v1/failover_scheme.proto",
}
