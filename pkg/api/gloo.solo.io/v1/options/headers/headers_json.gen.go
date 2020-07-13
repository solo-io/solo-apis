// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/headers/headers.proto

package headers

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

// MarshalJSON is a custom marshaler for HeaderManipulation
func (this *HeaderManipulation) MarshalJSON() ([]byte, error) {
	str, err := HeadersMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HeaderManipulation
func (this *HeaderManipulation) UnmarshalJSON(b []byte) error {
	return HeadersUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HeaderValueOption
func (this *HeaderValueOption) MarshalJSON() ([]byte, error) {
	str, err := HeadersMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HeaderValueOption
func (this *HeaderValueOption) UnmarshalJSON(b []byte) error {
	return HeadersUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HeaderValue
func (this *HeaderValue) MarshalJSON() ([]byte, error) {
	str, err := HeadersMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HeaderValue
func (this *HeaderValue) UnmarshalJSON(b []byte) error {
	return HeadersUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	HeadersMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	HeadersUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
