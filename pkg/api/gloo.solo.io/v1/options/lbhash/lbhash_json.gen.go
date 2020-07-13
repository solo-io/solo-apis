// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/lbhash/lbhash.proto

package lbhash

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for RouteActionHashConfig
func (this *RouteActionHashConfig) MarshalJSON() ([]byte, error) {
	str, err := LbhashMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RouteActionHashConfig
func (this *RouteActionHashConfig) UnmarshalJSON(b []byte) error {
	return LbhashUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Cookie
func (this *Cookie) MarshalJSON() ([]byte, error) {
	str, err := LbhashMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Cookie
func (this *Cookie) UnmarshalJSON(b []byte) error {
	return LbhashUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HashPolicy
func (this *HashPolicy) MarshalJSON() ([]byte, error) {
	str, err := LbhashMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HashPolicy
func (this *HashPolicy) UnmarshalJSON(b []byte) error {
	return LbhashUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	LbhashMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	LbhashUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
