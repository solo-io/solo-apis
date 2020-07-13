// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/pipe/pipe.proto

package pipe

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for UpstreamSpec
func (this *UpstreamSpec) MarshalJSON() ([]byte, error) {
	str, err := PipeMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for UpstreamSpec
func (this *UpstreamSpec) UnmarshalJSON(b []byte) error {
	return PipeUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	PipeMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	PipeUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
