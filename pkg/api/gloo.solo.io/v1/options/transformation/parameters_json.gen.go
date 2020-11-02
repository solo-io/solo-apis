// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/transformation/parameters.proto

package transformation

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

// MarshalJSON is a custom marshaler for Parameters
func (this *Parameters) MarshalJSON() ([]byte, error) {
	str, err := ParametersMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Parameters
func (this *Parameters) UnmarshalJSON(b []byte) error {
	return ParametersUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	ParametersMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{EnumsAsInts: true}
	ParametersUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
