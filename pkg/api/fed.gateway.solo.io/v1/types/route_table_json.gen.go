// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-fed/fed.gateway/v1/route_table.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-apis/pkg/api/fed.solo.io/core/v1"
	_ "github.com/solo-io/solo-apis/pkg/api/gateway.solo.io/v1"
	_ "github.com/solo-io/solo-apis/pkg/api/multicluster.solo.io/v1alpha1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for FederatedRouteTableSpec
func (this *FederatedRouteTableSpec) MarshalJSON() ([]byte, error) {
	str, err := RouteTableMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for FederatedRouteTableSpec
func (this *FederatedRouteTableSpec) UnmarshalJSON(b []byte) error {
	return RouteTableUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for FederatedRouteTableSpec_Template
func (this *FederatedRouteTableSpec_Template) MarshalJSON() ([]byte, error) {
	str, err := RouteTableMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for FederatedRouteTableSpec_Template
func (this *FederatedRouteTableSpec_Template) UnmarshalJSON(b []byte) error {
	return RouteTableUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for FederatedRouteTableStatus
func (this *FederatedRouteTableStatus) MarshalJSON() ([]byte, error) {
	str, err := RouteTableMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for FederatedRouteTableStatus
func (this *FederatedRouteTableStatus) UnmarshalJSON(b []byte) error {
	return RouteTableUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	RouteTableMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	RouteTableUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
