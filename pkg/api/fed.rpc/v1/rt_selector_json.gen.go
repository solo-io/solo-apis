// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-fed/fed.rpc/v1/rt_selector.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/core/matchers"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for SubRouteTableRow
func (this *SubRouteTableRow) MarshalJSON() ([]byte, error) {
	str, err := RtSelectorMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for SubRouteTableRow
func (this *SubRouteTableRow) UnmarshalJSON(b []byte) error {
	return RtSelectorUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GetVirtualServiceRoutesRequest
func (this *GetVirtualServiceRoutesRequest) MarshalJSON() ([]byte, error) {
	str, err := RtSelectorMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GetVirtualServiceRoutesRequest
func (this *GetVirtualServiceRoutesRequest) UnmarshalJSON(b []byte) error {
	return RtSelectorUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GetVirtualServiceRoutesResponse
func (this *GetVirtualServiceRoutesResponse) MarshalJSON() ([]byte, error) {
	str, err := RtSelectorMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GetVirtualServiceRoutesResponse
func (this *GetVirtualServiceRoutesResponse) UnmarshalJSON(b []byte) error {
	return RtSelectorUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	RtSelectorMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	RtSelectorUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
