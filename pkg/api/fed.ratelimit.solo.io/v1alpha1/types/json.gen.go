// Code generated by skv2. DO NOT EDIT.

// Generated json marshal and unmarshal functions

package types

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	jsonpb "github.com/golang/protobuf/jsonpb"
	proto "github.com/golang/protobuf/proto"
	skv2jsonpb "github.com/solo-io/skv2/pkg/kube_jsonpb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var (
	marshaller   = &skv2jsonpb.Marshaler{}
	unmarshaller = &jsonpb.Unmarshaler{}
)

// MarshalJSON is a custom marshaler for FederatedRateLimitConfigSpec
func (this *FederatedRateLimitConfigSpec) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for FederatedRateLimitConfigSpec
func (this *FederatedRateLimitConfigSpec) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for FederatedRateLimitConfigStatus
func (this *FederatedRateLimitConfigStatus) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for FederatedRateLimitConfigStatus
func (this *FederatedRateLimitConfigStatus) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}
