// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/extensions/jwt/solo_jwt_authn.proto

package jwt

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SoloJwtAuthnPerRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Requirement     string                                          `protobuf:"bytes,1,opt,name=requirement,proto3" json:"requirement,omitempty"`
	ClaimsToHeaders map[string]*SoloJwtAuthnPerRoute_ClaimToHeaders `protobuf:"bytes,2,rep,name=claims_to_headers,json=claimsToHeaders,proto3" json:"claims_to_headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// clear the route cache if claims were added to the header
	ClearRouteCache   bool   `protobuf:"varint,3,opt,name=clear_route_cache,json=clearRouteCache,proto3" json:"clear_route_cache,omitempty"`
	PayloadInMetadata string `protobuf:"bytes,4,opt,name=payload_in_metadata,json=payloadInMetadata,proto3" json:"payload_in_metadata,omitempty"`
}

func (x *SoloJwtAuthnPerRoute) Reset() {
	*x = SoloJwtAuthnPerRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SoloJwtAuthnPerRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SoloJwtAuthnPerRoute) ProtoMessage() {}

func (x *SoloJwtAuthnPerRoute) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SoloJwtAuthnPerRoute.ProtoReflect.Descriptor instead.
func (*SoloJwtAuthnPerRoute) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_rawDescGZIP(), []int{0}
}

func (x *SoloJwtAuthnPerRoute) GetRequirement() string {
	if x != nil {
		return x.Requirement
	}
	return ""
}

func (x *SoloJwtAuthnPerRoute) GetClaimsToHeaders() map[string]*SoloJwtAuthnPerRoute_ClaimToHeaders {
	if x != nil {
		return x.ClaimsToHeaders
	}
	return nil
}

func (x *SoloJwtAuthnPerRoute) GetClearRouteCache() bool {
	if x != nil {
		return x.ClearRouteCache
	}
	return false
}

func (x *SoloJwtAuthnPerRoute) GetPayloadInMetadata() string {
	if x != nil {
		return x.PayloadInMetadata
	}
	return ""
}

// If this is specified, one of the claims will be copied to a header
// and the route cache will be cleared.
type SoloJwtAuthnPerRoute_ClaimToHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Claim  string `protobuf:"bytes,1,opt,name=claim,proto3" json:"claim,omitempty"`
	Header string `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	Append bool   `protobuf:"varint,3,opt,name=append,proto3" json:"append,omitempty"`
}

func (x *SoloJwtAuthnPerRoute_ClaimToHeader) Reset() {
	*x = SoloJwtAuthnPerRoute_ClaimToHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SoloJwtAuthnPerRoute_ClaimToHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SoloJwtAuthnPerRoute_ClaimToHeader) ProtoMessage() {}

func (x *SoloJwtAuthnPerRoute_ClaimToHeader) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SoloJwtAuthnPerRoute_ClaimToHeader.ProtoReflect.Descriptor instead.
func (*SoloJwtAuthnPerRoute_ClaimToHeader) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SoloJwtAuthnPerRoute_ClaimToHeader) GetClaim() string {
	if x != nil {
		return x.Claim
	}
	return ""
}

func (x *SoloJwtAuthnPerRoute_ClaimToHeader) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (x *SoloJwtAuthnPerRoute_ClaimToHeader) GetAppend() bool {
	if x != nil {
		return x.Append
	}
	return false
}

type SoloJwtAuthnPerRoute_ClaimToHeaders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Claims []*SoloJwtAuthnPerRoute_ClaimToHeader `protobuf:"bytes,1,rep,name=claims,proto3" json:"claims,omitempty"`
}

func (x *SoloJwtAuthnPerRoute_ClaimToHeaders) Reset() {
	*x = SoloJwtAuthnPerRoute_ClaimToHeaders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SoloJwtAuthnPerRoute_ClaimToHeaders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SoloJwtAuthnPerRoute_ClaimToHeaders) ProtoMessage() {}

func (x *SoloJwtAuthnPerRoute_ClaimToHeaders) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SoloJwtAuthnPerRoute_ClaimToHeaders.ProtoReflect.Descriptor instead.
func (*SoloJwtAuthnPerRoute_ClaimToHeaders) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_rawDescGZIP(), []int{0, 1}
}

func (x *SoloJwtAuthnPerRoute_ClaimToHeaders) GetClaims() []*SoloJwtAuthnPerRoute_ClaimToHeader {
	if x != nil {
		return x.Claims
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_rawDesc = []byte{
	0x0a, 0x5d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6a, 0x77, 0x74, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x5f,
	0x6a, 0x77, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x2a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x5f, 0x6a,
	0x77, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x32, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xff, 0x04, 0x0a, 0x14, 0x53, 0x6f, 0x6c, 0x6f, 0x4a, 0x77, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x6e, 0x50, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x81, 0x01, 0x0a, 0x11, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x55, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x5f, 0x6a, 0x77, 0x74, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x6f, 0x6c, 0x6f, 0x4a, 0x77, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x6e, 0x50, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x43, 0x6c,
	0x61, 0x69, 0x6d, 0x73, 0x54, 0x6f, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x54, 0x6f, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x5f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f,
	0x63, 0x6c, 0x65, 0x61, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12,
	0x2e, 0x0a, 0x13, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x69, 0x6e, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x55, 0x0a, 0x0d, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x54, 0x6f, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x1a, 0x78, 0x0a, 0x0e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x54,
	0x6f, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x66, 0x0a, 0x06, 0x63, 0x6c, 0x61, 0x69,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x5f, 0x6a, 0x77, 0x74, 0x5f, 0x61, 0x75, 0x74,
	0x68, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x6f, 0x6c, 0x6f, 0x4a, 0x77, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x6e, 0x50, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d,
	0x54, 0x6f, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73,
	0x1a, 0x93, 0x01, 0x0a, 0x14, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x54, 0x6f, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x65, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x5f, 0x6a, 0x77, 0x74, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x6f, 0x6c, 0x6f, 0x4a, 0x77, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x6e, 0x50, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x43, 0x6c, 0x61,
	0x69, 0x6d, 0x54, 0x6f, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0xa0, 0x01, 0x0a, 0x38, 0x69, 0x6f, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x5f, 0x6a, 0x77, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6e,
	0x2e, 0x76, 0x32, 0x42, 0x11, 0x53, 0x6f, 0x6c, 0x6f, 0x4a, 0x77, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6a, 0x77, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_goTypes = []interface{}{
	(*SoloJwtAuthnPerRoute)(nil),                // 0: envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute
	(*SoloJwtAuthnPerRoute_ClaimToHeader)(nil),  // 1: envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader
	(*SoloJwtAuthnPerRoute_ClaimToHeaders)(nil), // 2: envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders
	nil, // 3: envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimsToHeadersEntry
}
var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_depIdxs = []int32{
	3, // 0: envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.claims_to_headers:type_name -> envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimsToHeadersEntry
	1, // 1: envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders.claims:type_name -> envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader
	2, // 2: envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimsToHeadersEntry.value:type_name -> envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SoloJwtAuthnPerRoute); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SoloJwtAuthnPerRoute_ClaimToHeader); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SoloJwtAuthnPerRoute_ClaimToHeaders); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_jwt_solo_jwt_authn_proto_depIdxs = nil
}
