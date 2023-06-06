// Code generated by skv2. DO NOT EDIT.

// Generated json marshal and unmarshal functions

package v2

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
	unmarshaller = &jsonpb.Unmarshaler{
		AllowUnknownFields: true,
	}
	strictUnmarshaller = &jsonpb.Unmarshaler{}
)

// MarshalJSON is a custom marshaler for GraphQLPersistedQueryCachePolicySpec
func (this *GraphQLPersistedQueryCachePolicySpec) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GraphQLPersistedQueryCachePolicySpec
func (this *GraphQLPersistedQueryCachePolicySpec) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GraphQLPersistedQueryCachePolicyStatus
func (this *GraphQLPersistedQueryCachePolicyStatus) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GraphQLPersistedQueryCachePolicyStatus
func (this *GraphQLPersistedQueryCachePolicyStatus) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for FailoverPolicySpec
func (this *FailoverPolicySpec) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for FailoverPolicySpec
func (this *FailoverPolicySpec) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for FailoverPolicyStatus
func (this *FailoverPolicyStatus) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for FailoverPolicyStatus
func (this *FailoverPolicyStatus) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for OutlierDetectionPolicySpec
func (this *OutlierDetectionPolicySpec) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for OutlierDetectionPolicySpec
func (this *OutlierDetectionPolicySpec) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for OutlierDetectionPolicyStatus
func (this *OutlierDetectionPolicyStatus) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for OutlierDetectionPolicyStatus
func (this *OutlierDetectionPolicyStatus) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for FaultInjectionPolicySpec
func (this *FaultInjectionPolicySpec) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for FaultInjectionPolicySpec
func (this *FaultInjectionPolicySpec) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for FaultInjectionPolicyStatus
func (this *FaultInjectionPolicyStatus) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for FaultInjectionPolicyStatus
func (this *FaultInjectionPolicyStatus) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RetryTimeoutPolicySpec
func (this *RetryTimeoutPolicySpec) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RetryTimeoutPolicySpec
func (this *RetryTimeoutPolicySpec) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RetryTimeoutPolicyStatus
func (this *RetryTimeoutPolicyStatus) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RetryTimeoutPolicyStatus
func (this *RetryTimeoutPolicyStatus) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ConnectionPolicySpec
func (this *ConnectionPolicySpec) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ConnectionPolicySpec
func (this *ConnectionPolicySpec) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ConnectionPolicyStatus
func (this *ConnectionPolicyStatus) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ConnectionPolicyStatus
func (this *ConnectionPolicyStatus) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrimProxyConfigPolicySpec
func (this *TrimProxyConfigPolicySpec) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrimProxyConfigPolicySpec
func (this *TrimProxyConfigPolicySpec) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrimProxyConfigPolicyStatus
func (this *TrimProxyConfigPolicyStatus) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrimProxyConfigPolicyStatus
func (this *TrimProxyConfigPolicyStatus) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}
