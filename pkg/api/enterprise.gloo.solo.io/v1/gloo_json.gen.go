// Code generated by skv2. DO NOT EDIT.

// Generated json marshal and unmarshal functions

package v1

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
	marshaller   = &skv2jsonpb.Marshaler{EnumsAsInts: true}
	unmarshaller = &jsonpb.Unmarshaler{}
)

// MarshalJSON is a custom marshaler for AuthConfigSpec
func (this *AuthConfigSpec) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AuthConfigSpec
func (this *AuthConfigSpec) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for AuthConfigStatus
func (this *AuthConfigStatus) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AuthConfigStatus
func (this *AuthConfigStatus) UnmarshalJSON(b []byte) error {
	namespacedStatuses := AuthConfigNamespacedStatuses{}
	if err := unmarshaller.Unmarshal(bytes.NewReader(b), &namespacedStatuses); err != nil {
		return unmarshaller.Unmarshal(bytes.NewReader(b), this)
	}

	for _, status := range namespacedStatuses.GetStatuses() {
		// take the first status
		if status != nil {
			status.DeepCopyInto(this)
			return nil
		}
	}
	return nil
}
