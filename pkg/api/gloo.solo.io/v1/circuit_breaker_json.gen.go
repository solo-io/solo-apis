// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/circuit_breaker.proto

package v1

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

// MarshalJSON is a custom marshaler for CircuitBreakerConfig
func (this *CircuitBreakerConfig) MarshalJSON() ([]byte, error) {
	str, err := CircuitBreakerMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CircuitBreakerConfig
func (this *CircuitBreakerConfig) UnmarshalJSON(b []byte) error {
	return CircuitBreakerUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	CircuitBreakerMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	CircuitBreakerUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
