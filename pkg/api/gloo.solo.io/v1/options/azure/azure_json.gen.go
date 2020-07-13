// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/azure/azure.proto

package azure

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for UpstreamSpec
func (this *UpstreamSpec) MarshalJSON() ([]byte, error) {
	str, err := AzureMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for UpstreamSpec
func (this *UpstreamSpec) UnmarshalJSON(b []byte) error {
	return AzureUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for UpstreamSpec_FunctionSpec
func (this *UpstreamSpec_FunctionSpec) MarshalJSON() ([]byte, error) {
	str, err := AzureMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for UpstreamSpec_FunctionSpec
func (this *UpstreamSpec_FunctionSpec) UnmarshalJSON(b []byte) error {
	return AzureUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for DestinationSpec
func (this *DestinationSpec) MarshalJSON() ([]byte, error) {
	str, err := AzureMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for DestinationSpec
func (this *DestinationSpec) UnmarshalJSON(b []byte) error {
	return AzureUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	AzureMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	AzureUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
