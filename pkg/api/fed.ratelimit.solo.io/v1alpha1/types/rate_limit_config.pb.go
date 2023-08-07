// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo-fed/fed.ratelimit/v1alpha1/rate_limit_config.proto

package types

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1 "github.com/solo-io/solo-apis/pkg/api/fed.solo.io/core/v1"
	types "github.com/solo-io/solo-apis/pkg/api/multicluster.solo.io/v1alpha1/types"
	v1alpha1 "github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The Federated[Resource] CRDs are a gloo-fed wrapper around Gloo Edge CRDs, with a Placement field indicating which
// clusters and namespaces to federate the object to.
// For more, see: https://docs.solo.io/gloo-edge/latest/guides/gloo_federation/federated_configuration/
type FederatedRateLimitConfigSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Template  *FederatedRateLimitConfigSpec_Template `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
	Placement *types.Placement                       `protobuf:"bytes,2,opt,name=placement,proto3" json:"placement,omitempty"`
}

func (x *FederatedRateLimitConfigSpec) Reset() {
	*x = FederatedRateLimitConfigSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FederatedRateLimitConfigSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederatedRateLimitConfigSpec) ProtoMessage() {}

func (x *FederatedRateLimitConfigSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederatedRateLimitConfigSpec.ProtoReflect.Descriptor instead.
func (*FederatedRateLimitConfigSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_rawDescGZIP(), []int{0}
}

func (x *FederatedRateLimitConfigSpec) GetTemplate() *FederatedRateLimitConfigSpec_Template {
	if x != nil {
		return x.Template
	}
	return nil
}

func (x *FederatedRateLimitConfigSpec) GetPlacement() *types.Placement {
	if x != nil {
		return x.Placement
	}
	return nil
}

type FederatedRateLimitConfigStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlacementStatus             *v1.PlacementStatus            `protobuf:"bytes,1,opt,name=placement_status,json=placementStatus,proto3" json:"placement_status,omitempty"`
	NamespacedPlacementStatuses map[string]*v1.PlacementStatus `protobuf:"bytes,2,rep,name=namespaced_placement_statuses,json=namespacedPlacementStatuses,proto3" json:"namespaced_placement_statuses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FederatedRateLimitConfigStatus) Reset() {
	*x = FederatedRateLimitConfigStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FederatedRateLimitConfigStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederatedRateLimitConfigStatus) ProtoMessage() {}

func (x *FederatedRateLimitConfigStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederatedRateLimitConfigStatus.ProtoReflect.Descriptor instead.
func (*FederatedRateLimitConfigStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_rawDescGZIP(), []int{1}
}

func (x *FederatedRateLimitConfigStatus) GetPlacementStatus() *v1.PlacementStatus {
	if x != nil {
		return x.PlacementStatus
	}
	return nil
}

func (x *FederatedRateLimitConfigStatus) GetNamespacedPlacementStatuses() map[string]*v1.PlacementStatus {
	if x != nil {
		return x.NamespacedPlacementStatuses
	}
	return nil
}

type FederatedRateLimitConfigSpec_Template struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spec     *v1alpha1.RateLimitConfigSpec `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	Metadata *v1.TemplateMetadata          `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *FederatedRateLimitConfigSpec_Template) Reset() {
	*x = FederatedRateLimitConfigSpec_Template{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FederatedRateLimitConfigSpec_Template) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederatedRateLimitConfigSpec_Template) ProtoMessage() {}

func (x *FederatedRateLimitConfigSpec_Template) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederatedRateLimitConfigSpec_Template.ProtoReflect.Descriptor instead.
func (*FederatedRateLimitConfigSpec_Template) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *FederatedRateLimitConfigSpec_Template) GetSpec() *v1alpha1.RateLimitConfigSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *FederatedRateLimitConfigSpec_Template) GetMetadata() *v1.TemplateMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_rawDesc = []byte{
	0x0a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x66, 0x65, 0x64, 0x2f, 0x66, 0x65, 0x64, 0x2e,
	0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x66, 0x65, 0x64, 0x2e,
	0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x66, 0x65, 0x64,
	0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2d, 0x66, 0x65, 0x64, 0x2f, 0x66, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x2d, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x02, 0x0a, 0x1c, 0x46, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x12, 0x58, 0x0a, 0x08, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x66, 0x65,
	0x64, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63,
	0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x1a, 0x8a, 0x01, 0x0a, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12,
	0x3e, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12,
	0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22,
	0xfe, 0x02, 0x0a, 0x1e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x4c, 0x0a, 0x10, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x0f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x9a, 0x01, 0x0a, 0x1d, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x5f,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x56, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x72,
	0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x1b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x1a, 0x71, 0x0a,
	0x20, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x37, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x53, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_goTypes = []interface{}{
	(*FederatedRateLimitConfigSpec)(nil),          // 0: fed.ratelimit.solo.io.FederatedRateLimitConfigSpec
	(*FederatedRateLimitConfigStatus)(nil),        // 1: fed.ratelimit.solo.io.FederatedRateLimitConfigStatus
	(*FederatedRateLimitConfigSpec_Template)(nil), // 2: fed.ratelimit.solo.io.FederatedRateLimitConfigSpec.Template
	nil,                                  // 3: fed.ratelimit.solo.io.FederatedRateLimitConfigStatus.NamespacedPlacementStatusesEntry
	(*types.Placement)(nil),              // 4: multicluster.solo.io.Placement
	(*v1.PlacementStatus)(nil),           // 5: core.fed.solo.io.PlacementStatus
	(*v1alpha1.RateLimitConfigSpec)(nil), // 6: ratelimit.api.solo.io.RateLimitConfigSpec
	(*v1.TemplateMetadata)(nil),          // 7: core.fed.solo.io.TemplateMetadata
}
var file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_depIdxs = []int32{
	2, // 0: fed.ratelimit.solo.io.FederatedRateLimitConfigSpec.template:type_name -> fed.ratelimit.solo.io.FederatedRateLimitConfigSpec.Template
	4, // 1: fed.ratelimit.solo.io.FederatedRateLimitConfigSpec.placement:type_name -> multicluster.solo.io.Placement
	5, // 2: fed.ratelimit.solo.io.FederatedRateLimitConfigStatus.placement_status:type_name -> core.fed.solo.io.PlacementStatus
	3, // 3: fed.ratelimit.solo.io.FederatedRateLimitConfigStatus.namespaced_placement_statuses:type_name -> fed.ratelimit.solo.io.FederatedRateLimitConfigStatus.NamespacedPlacementStatusesEntry
	6, // 4: fed.ratelimit.solo.io.FederatedRateLimitConfigSpec.Template.spec:type_name -> ratelimit.api.solo.io.RateLimitConfigSpec
	7, // 5: fed.ratelimit.solo.io.FederatedRateLimitConfigSpec.Template.metadata:type_name -> core.fed.solo.io.TemplateMetadata
	5, // 6: fed.ratelimit.solo.io.FederatedRateLimitConfigStatus.NamespacedPlacementStatusesEntry.value:type_name -> core.fed.solo.io.PlacementStatus
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FederatedRateLimitConfigSpec); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FederatedRateLimitConfigStatus); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FederatedRateLimitConfigSpec_Template); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_ratelimit_v1alpha1_rate_limit_config_proto_depIdxs = nil
}
