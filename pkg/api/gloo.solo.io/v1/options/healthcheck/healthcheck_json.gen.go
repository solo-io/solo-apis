// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/healthcheck/healthcheck.proto

package healthcheck

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for HealthCheck
func (this *HealthCheck) MarshalJSON() ([]byte, error) {
	str, err := HealthcheckMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HealthCheck
func (this *HealthCheck) UnmarshalJSON(b []byte) error {
	return HealthcheckUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	HealthcheckMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{EnumsAsInts: true}
	HealthcheckUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
