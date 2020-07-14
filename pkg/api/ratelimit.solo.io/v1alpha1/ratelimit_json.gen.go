// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/rate-limiter/v1alpha1/ratelimit.proto

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for RateLimitConfigSpec
func (this *RateLimitConfigSpec) MarshalJSON() ([]byte, error) {
	str, err := RatelimitMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RateLimitConfigSpec
func (this *RateLimitConfigSpec) UnmarshalJSON(b []byte) error {
	return RatelimitUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RateLimitConfigSpec_Raw
func (this *RateLimitConfigSpec_Raw) MarshalJSON() ([]byte, error) {
	str, err := RatelimitMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RateLimitConfigSpec_Raw
func (this *RateLimitConfigSpec_Raw) UnmarshalJSON(b []byte) error {
	return RatelimitUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RateLimitConfigStatus
func (this *RateLimitConfigStatus) MarshalJSON() ([]byte, error) {
	str, err := RatelimitMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RateLimitConfigStatus
func (this *RateLimitConfigStatus) UnmarshalJSON(b []byte) error {
	return RatelimitUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Descriptor
func (this *Descriptor) MarshalJSON() ([]byte, error) {
	str, err := RatelimitMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Descriptor
func (this *Descriptor) UnmarshalJSON(b []byte) error {
	return RatelimitUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RateLimitActions
func (this *RateLimitActions) MarshalJSON() ([]byte, error) {
	str, err := RatelimitMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RateLimitActions
func (this *RateLimitActions) UnmarshalJSON(b []byte) error {
	return RatelimitUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RateLimit
func (this *RateLimit) MarshalJSON() ([]byte, error) {
	str, err := RatelimitMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RateLimit
func (this *RateLimit) UnmarshalJSON(b []byte) error {
	return RatelimitUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Action
func (this *Action) MarshalJSON() ([]byte, error) {
	str, err := RatelimitMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Action
func (this *Action) UnmarshalJSON(b []byte) error {
	return RatelimitUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Action_SourceCluster
func (this *Action_SourceCluster) MarshalJSON() ([]byte, error) {
	str, err := RatelimitMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Action_SourceCluster
func (this *Action_SourceCluster) UnmarshalJSON(b []byte) error {
	return RatelimitUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Action_DestinationCluster
func (this *Action_DestinationCluster) MarshalJSON() ([]byte, error) {
	str, err := RatelimitMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Action_DestinationCluster
func (this *Action_DestinationCluster) UnmarshalJSON(b []byte) error {
	return RatelimitUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Action_RequestHeaders
func (this *Action_RequestHeaders) MarshalJSON() ([]byte, error) {
	str, err := RatelimitMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Action_RequestHeaders
func (this *Action_RequestHeaders) UnmarshalJSON(b []byte) error {
	return RatelimitUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Action_RemoteAddress
func (this *Action_RemoteAddress) MarshalJSON() ([]byte, error) {
	str, err := RatelimitMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Action_RemoteAddress
func (this *Action_RemoteAddress) UnmarshalJSON(b []byte) error {
	return RatelimitUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Action_GenericKey
func (this *Action_GenericKey) MarshalJSON() ([]byte, error) {
	str, err := RatelimitMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Action_GenericKey
func (this *Action_GenericKey) UnmarshalJSON(b []byte) error {
	return RatelimitUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Action_HeaderValueMatch
func (this *Action_HeaderValueMatch) MarshalJSON() ([]byte, error) {
	str, err := RatelimitMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Action_HeaderValueMatch
func (this *Action_HeaderValueMatch) UnmarshalJSON(b []byte) error {
	return RatelimitUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Action_HeaderValueMatch_HeaderMatcher
func (this *Action_HeaderValueMatch_HeaderMatcher) MarshalJSON() ([]byte, error) {
	str, err := RatelimitMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Action_HeaderValueMatch_HeaderMatcher
func (this *Action_HeaderValueMatch_HeaderMatcher) UnmarshalJSON(b []byte) error {
	return RatelimitUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Action_HeaderValueMatch_HeaderMatcher_Int64Range
func (this *Action_HeaderValueMatch_HeaderMatcher_Int64Range) MarshalJSON() ([]byte, error) {
	str, err := RatelimitMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Action_HeaderValueMatch_HeaderMatcher_Int64Range
func (this *Action_HeaderValueMatch_HeaderMatcher_Int64Range) UnmarshalJSON(b []byte) error {
	return RatelimitUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	RatelimitMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	RatelimitUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
