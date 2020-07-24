// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/type/v3/semantic_version.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/udpa/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for SemanticVersion
func (this *SemanticVersion) MarshalJSON() ([]byte, error) {
	str, err := SemanticVersionMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for SemanticVersion
func (this *SemanticVersion) UnmarshalJSON(b []byte) error {
	return SemanticVersionUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	SemanticVersionMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	SemanticVersionUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
