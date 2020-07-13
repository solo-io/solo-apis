// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/extensions/jwt/solo_jwt_authn.proto

package jwt

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for SoloJwtAuthnPerRoute
func (this *SoloJwtAuthnPerRoute) MarshalJSON() ([]byte, error) {
	str, err := SoloJwtAuthnMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for SoloJwtAuthnPerRoute
func (this *SoloJwtAuthnPerRoute) UnmarshalJSON(b []byte) error {
	return SoloJwtAuthnUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for SoloJwtAuthnPerRoute_ClaimToHeader
func (this *SoloJwtAuthnPerRoute_ClaimToHeader) MarshalJSON() ([]byte, error) {
	str, err := SoloJwtAuthnMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for SoloJwtAuthnPerRoute_ClaimToHeader
func (this *SoloJwtAuthnPerRoute_ClaimToHeader) UnmarshalJSON(b []byte) error {
	return SoloJwtAuthnUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for SoloJwtAuthnPerRoute_ClaimToHeaders
func (this *SoloJwtAuthnPerRoute_ClaimToHeaders) MarshalJSON() ([]byte, error) {
	str, err := SoloJwtAuthnMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for SoloJwtAuthnPerRoute_ClaimToHeaders
func (this *SoloJwtAuthnPerRoute_ClaimToHeaders) UnmarshalJSON(b []byte) error {
	return SoloJwtAuthnUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	SoloJwtAuthnMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	SoloJwtAuthnUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
