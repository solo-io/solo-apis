// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/enterprise/options/rbac/rbac.proto

package rbac

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for Settings
func (this *Settings) MarshalJSON() ([]byte, error) {
	str, err := RbacMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Settings
func (this *Settings) UnmarshalJSON(b []byte) error {
	return RbacUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ExtensionSettings
func (this *ExtensionSettings) MarshalJSON() ([]byte, error) {
	str, err := RbacMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ExtensionSettings
func (this *ExtensionSettings) UnmarshalJSON(b []byte) error {
	return RbacUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Policy
func (this *Policy) MarshalJSON() ([]byte, error) {
	str, err := RbacMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Policy
func (this *Policy) UnmarshalJSON(b []byte) error {
	return RbacUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Principal
func (this *Principal) MarshalJSON() ([]byte, error) {
	str, err := RbacMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Principal
func (this *Principal) UnmarshalJSON(b []byte) error {
	return RbacUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for JWTPrincipal
func (this *JWTPrincipal) MarshalJSON() ([]byte, error) {
	str, err := RbacMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for JWTPrincipal
func (this *JWTPrincipal) UnmarshalJSON(b []byte) error {
	return RbacUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Permissions
func (this *Permissions) MarshalJSON() ([]byte, error) {
	str, err := RbacMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Permissions
func (this *Permissions) UnmarshalJSON(b []byte) error {
	return RbacUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	RbacMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{EnumsAsInts: true}
	RbacUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
