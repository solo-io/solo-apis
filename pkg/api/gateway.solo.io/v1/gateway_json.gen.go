// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gateway/v1/gateway.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for GatewaySpec
func (this *GatewaySpec) MarshalJSON() ([]byte, error) {
	str, err := GatewayMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GatewaySpec
func (this *GatewaySpec) UnmarshalJSON(b []byte) error {
	return GatewayUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HttpGateway
func (this *HttpGateway) MarshalJSON() ([]byte, error) {
	str, err := GatewayMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HttpGateway
func (this *HttpGateway) UnmarshalJSON(b []byte) error {
	return GatewayUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TcpGateway
func (this *TcpGateway) MarshalJSON() ([]byte, error) {
	str, err := GatewayMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TcpGateway
func (this *TcpGateway) UnmarshalJSON(b []byte) error {
	return GatewayUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GatewayStatus
func (this *GatewayStatus) MarshalJSON() ([]byte, error) {
	str, err := GatewayMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GatewayStatus
func (this *GatewayStatus) UnmarshalJSON(b []byte) error {
	return GatewayUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	GatewayMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{EnumsAsInts: true}
	GatewayUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
