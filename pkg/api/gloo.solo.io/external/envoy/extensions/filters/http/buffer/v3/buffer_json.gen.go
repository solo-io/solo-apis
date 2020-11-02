// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/extensions/filters/http/buffer/v3/buffer.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// MarshalJSON is a custom marshaler for Buffer
func (this *Buffer) MarshalJSON() ([]byte, error) {
	str, err := BufferMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Buffer
func (this *Buffer) UnmarshalJSON(b []byte) error {
	return BufferUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for BufferPerRoute
func (this *BufferPerRoute) MarshalJSON() ([]byte, error) {
	str, err := BufferMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for BufferPerRoute
func (this *BufferPerRoute) UnmarshalJSON(b []byte) error {
	return BufferUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	BufferMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	BufferUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
