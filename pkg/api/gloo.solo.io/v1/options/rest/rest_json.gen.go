// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/rest/rest.proto

package rest

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/extensions/transformation"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/transformation"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for ServiceSpec
func (this *ServiceSpec) MarshalJSON() ([]byte, error) {
	str, err := RestMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ServiceSpec
func (this *ServiceSpec) UnmarshalJSON(b []byte) error {
	return RestUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ServiceSpec_SwaggerInfo
func (this *ServiceSpec_SwaggerInfo) MarshalJSON() ([]byte, error) {
	str, err := RestMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ServiceSpec_SwaggerInfo
func (this *ServiceSpec_SwaggerInfo) UnmarshalJSON(b []byte) error {
	return RestUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for DestinationSpec
func (this *DestinationSpec) MarshalJSON() ([]byte, error) {
	str, err := RestMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for DestinationSpec
func (this *DestinationSpec) UnmarshalJSON(b []byte) error {
	return RestUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	RestMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{EnumsAsInts: true}
	RestUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
