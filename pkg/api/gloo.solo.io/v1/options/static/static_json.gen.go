// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/static/static.proto

package static

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for UpstreamSpec
func (this *UpstreamSpec) MarshalJSON() ([]byte, error) {
	str, err := StaticMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for UpstreamSpec
func (this *UpstreamSpec) UnmarshalJSON(b []byte) error {
	return StaticUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Host
func (this *Host) MarshalJSON() ([]byte, error) {
	str, err := StaticMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Host
func (this *Host) UnmarshalJSON(b []byte) error {
	return StaticUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Host_HealthCheckConfig
func (this *Host_HealthCheckConfig) MarshalJSON() ([]byte, error) {
	str, err := StaticMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Host_HealthCheckConfig
func (this *Host_HealthCheckConfig) UnmarshalJSON(b []byte) error {
	return StaticUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	StaticMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{EnumsAsInts: true}
	StaticUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
