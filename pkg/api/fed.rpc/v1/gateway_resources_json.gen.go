// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-fed/fed.rpc/v1/gateway_resources.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	_ "github.com/solo-io/solo-apis/pkg/api/gateway.solo.io/v1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for Gateway
func (this *Gateway) MarshalJSON() ([]byte, error) {
	str, err := GatewayResourcesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Gateway
func (this *Gateway) UnmarshalJSON(b []byte) error {
	return GatewayResourcesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for VirtualService
func (this *VirtualService) MarshalJSON() ([]byte, error) {
	str, err := GatewayResourcesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualService
func (this *VirtualService) UnmarshalJSON(b []byte) error {
	return GatewayResourcesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RouteTable
func (this *RouteTable) MarshalJSON() ([]byte, error) {
	str, err := GatewayResourcesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RouteTable
func (this *RouteTable) UnmarshalJSON(b []byte) error {
	return GatewayResourcesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ListGatewaysRequest
func (this *ListGatewaysRequest) MarshalJSON() ([]byte, error) {
	str, err := GatewayResourcesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ListGatewaysRequest
func (this *ListGatewaysRequest) UnmarshalJSON(b []byte) error {
	return GatewayResourcesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ListGatewaysResponse
func (this *ListGatewaysResponse) MarshalJSON() ([]byte, error) {
	str, err := GatewayResourcesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ListGatewaysResponse
func (this *ListGatewaysResponse) UnmarshalJSON(b []byte) error {
	return GatewayResourcesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GetGatewayYamlRequest
func (this *GetGatewayYamlRequest) MarshalJSON() ([]byte, error) {
	str, err := GatewayResourcesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GetGatewayYamlRequest
func (this *GetGatewayYamlRequest) UnmarshalJSON(b []byte) error {
	return GatewayResourcesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GetGatewayYamlResponse
func (this *GetGatewayYamlResponse) MarshalJSON() ([]byte, error) {
	str, err := GatewayResourcesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GetGatewayYamlResponse
func (this *GetGatewayYamlResponse) UnmarshalJSON(b []byte) error {
	return GatewayResourcesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ListVirtualServicesRequest
func (this *ListVirtualServicesRequest) MarshalJSON() ([]byte, error) {
	str, err := GatewayResourcesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ListVirtualServicesRequest
func (this *ListVirtualServicesRequest) UnmarshalJSON(b []byte) error {
	return GatewayResourcesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ListVirtualServicesResponse
func (this *ListVirtualServicesResponse) MarshalJSON() ([]byte, error) {
	str, err := GatewayResourcesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ListVirtualServicesResponse
func (this *ListVirtualServicesResponse) UnmarshalJSON(b []byte) error {
	return GatewayResourcesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GetVirtualServiceYamlRequest
func (this *GetVirtualServiceYamlRequest) MarshalJSON() ([]byte, error) {
	str, err := GatewayResourcesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GetVirtualServiceYamlRequest
func (this *GetVirtualServiceYamlRequest) UnmarshalJSON(b []byte) error {
	return GatewayResourcesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GetVirtualServiceYamlResponse
func (this *GetVirtualServiceYamlResponse) MarshalJSON() ([]byte, error) {
	str, err := GatewayResourcesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GetVirtualServiceYamlResponse
func (this *GetVirtualServiceYamlResponse) UnmarshalJSON(b []byte) error {
	return GatewayResourcesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ListRouteTablesRequest
func (this *ListRouteTablesRequest) MarshalJSON() ([]byte, error) {
	str, err := GatewayResourcesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ListRouteTablesRequest
func (this *ListRouteTablesRequest) UnmarshalJSON(b []byte) error {
	return GatewayResourcesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ListRouteTablesResponse
func (this *ListRouteTablesResponse) MarshalJSON() ([]byte, error) {
	str, err := GatewayResourcesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ListRouteTablesResponse
func (this *ListRouteTablesResponse) UnmarshalJSON(b []byte) error {
	return GatewayResourcesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GetRouteTableYamlRequest
func (this *GetRouteTableYamlRequest) MarshalJSON() ([]byte, error) {
	str, err := GatewayResourcesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GetRouteTableYamlRequest
func (this *GetRouteTableYamlRequest) UnmarshalJSON(b []byte) error {
	return GatewayResourcesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GetRouteTableYamlResponse
func (this *GetRouteTableYamlResponse) MarshalJSON() ([]byte, error) {
	str, err := GatewayResourcesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GetRouteTableYamlResponse
func (this *GetRouteTableYamlResponse) UnmarshalJSON(b []byte) error {
	return GatewayResourcesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	GatewayResourcesMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	GatewayResourcesUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
