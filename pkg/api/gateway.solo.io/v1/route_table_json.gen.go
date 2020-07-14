// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gateway/v1/route_table.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for RouteTable
func (this *RouteTable) MarshalJSON() ([]byte, error) {
	str, err := RouteTableMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RouteTable
func (this *RouteTable) UnmarshalJSON(b []byte) error {
	return RouteTableUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	RouteTableMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	RouteTableUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
