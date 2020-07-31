// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/proxy.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/core/matchers"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for ProxySpec
func (this *ProxySpec) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ProxySpec
func (this *ProxySpec) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Listener
func (this *Listener) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Listener
func (this *Listener) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TcpListener
func (this *TcpListener) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TcpListener
func (this *TcpListener) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TcpHost
func (this *TcpHost) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TcpHost
func (this *TcpHost) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TcpHost_TcpAction
func (this *TcpHost_TcpAction) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TcpHost_TcpAction
func (this *TcpHost_TcpAction) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HttpListener
func (this *HttpListener) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HttpListener
func (this *HttpListener) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for VirtualHost
func (this *VirtualHost) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualHost
func (this *VirtualHost) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Route
func (this *Route) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Route
func (this *Route) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RouteAction
func (this *RouteAction) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RouteAction
func (this *RouteAction) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Destination
func (this *Destination) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Destination
func (this *Destination) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for KubernetesServiceDestination
func (this *KubernetesServiceDestination) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for KubernetesServiceDestination
func (this *KubernetesServiceDestination) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ConsulServiceDestination
func (this *ConsulServiceDestination) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ConsulServiceDestination
func (this *ConsulServiceDestination) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for UpstreamGroupSpec
func (this *UpstreamGroupSpec) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for UpstreamGroupSpec
func (this *UpstreamGroupSpec) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MultiDestination
func (this *MultiDestination) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MultiDestination
func (this *MultiDestination) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for WeightedDestination
func (this *WeightedDestination) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for WeightedDestination
func (this *WeightedDestination) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RedirectAction
func (this *RedirectAction) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RedirectAction
func (this *RedirectAction) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for DirectResponseAction
func (this *DirectResponseAction) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for DirectResponseAction
func (this *DirectResponseAction) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for UpstreamGroupStatus
func (this *UpstreamGroupStatus) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for UpstreamGroupStatus
func (this *UpstreamGroupStatus) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ProxyStatus
func (this *ProxyStatus) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ProxyStatus
func (this *ProxyStatus) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	ProxyMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{EnumsAsInts: true}
	ProxyUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
