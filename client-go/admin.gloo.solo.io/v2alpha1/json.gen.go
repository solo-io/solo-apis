// Code generated by skv2. DO NOT EDIT.

// Generated json marshal and unmarshal functions

package v2alpha1

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

// MarshalJSON is a custom marshaler for WaypointLifecycleManagerSpec
func (this *WaypointLifecycleManagerSpec) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for WaypointLifecycleManagerSpec
func (this *WaypointLifecycleManagerSpec) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for WaypointLifecycleManagerStatus
func (this *WaypointLifecycleManagerStatus) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for WaypointLifecycleManagerStatus
func (this *WaypointLifecycleManagerStatus) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for InsightsConfigSpec
func (this *InsightsConfigSpec) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for InsightsConfigSpec
func (this *InsightsConfigSpec) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for InsightsConfigStatus
func (this *InsightsConfigStatus) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for InsightsConfigStatus
func (this *InsightsConfigStatus) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}
