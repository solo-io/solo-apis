// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/secret.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-apis/pkg/api/enterprise.gloo.solo.io/v1"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for Secret
func (this *Secret) MarshalJSON() ([]byte, error) {
	str, err := SecretMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Secret
func (this *Secret) UnmarshalJSON(b []byte) error {
	return SecretUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for AwsSecret
func (this *AwsSecret) MarshalJSON() ([]byte, error) {
	str, err := SecretMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AwsSecret
func (this *AwsSecret) UnmarshalJSON(b []byte) error {
	return SecretUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for AzureSecret
func (this *AzureSecret) MarshalJSON() ([]byte, error) {
	str, err := SecretMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AzureSecret
func (this *AzureSecret) UnmarshalJSON(b []byte) error {
	return SecretUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TlsSecret
func (this *TlsSecret) MarshalJSON() ([]byte, error) {
	str, err := SecretMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TlsSecret
func (this *TlsSecret) UnmarshalJSON(b []byte) error {
	return SecretUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HeaderSecret
func (this *HeaderSecret) MarshalJSON() ([]byte, error) {
	str, err := SecretMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HeaderSecret
func (this *HeaderSecret) UnmarshalJSON(b []byte) error {
	return SecretUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	SecretMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{EnumsAsInts: true}
	SecretUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
