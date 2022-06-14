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

// MarshalJSON is a custom marshaler for GatewaySpec
func (this *GatewaySpec) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GatewaySpec
func (this *GatewaySpec) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GatewayStatus
func (this *GatewayStatus) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GatewayStatus
func (this *GatewayStatus) UnmarshalJSON(b []byte) error {
	namespacedStatuses := GatewayNamespacedStatuses{}
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

// MarshalJSON is a custom marshaler for MatchableHttpGatewaySpec
func (this *MatchableHttpGatewaySpec) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MatchableHttpGatewaySpec
func (this *MatchableHttpGatewaySpec) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MatchableHttpGatewayStatus
func (this *MatchableHttpGatewayStatus) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MatchableHttpGatewayStatus
func (this *MatchableHttpGatewayStatus) UnmarshalJSON(b []byte) error {
	namespacedStatuses := MatchableHttpGatewayNamespacedStatuses{}
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

// MarshalJSON is a custom marshaler for RouteTableSpec
func (this *RouteTableSpec) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RouteTableSpec
func (this *RouteTableSpec) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RouteTableStatus
func (this *RouteTableStatus) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RouteTableStatus
func (this *RouteTableStatus) UnmarshalJSON(b []byte) error {
	namespacedStatuses := RouteTableNamespacedStatuses{}
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

// MarshalJSON is a custom marshaler for VirtualServiceSpec
func (this *VirtualServiceSpec) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualServiceSpec
func (this *VirtualServiceSpec) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for VirtualServiceStatus
func (this *VirtualServiceStatus) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualServiceStatus
func (this *VirtualServiceStatus) UnmarshalJSON(b []byte) error {
	namespacedStatuses := VirtualServiceNamespacedStatuses{}
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

// MarshalJSON is a custom marshaler for VirtualHostOptionSpec
func (this *VirtualHostOptionSpec) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualHostOptionSpec
func (this *VirtualHostOptionSpec) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for VirtualHostOptionStatus
func (this *VirtualHostOptionStatus) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualHostOptionStatus
func (this *VirtualHostOptionStatus) UnmarshalJSON(b []byte) error {
	namespacedStatuses := VirtualHostOptionNamespacedStatuses{}
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

// MarshalJSON is a custom marshaler for RouteOptionSpec
func (this *RouteOptionSpec) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RouteOptionSpec
func (this *RouteOptionSpec) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RouteOptionStatus
func (this *RouteOptionStatus) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RouteOptionStatus
func (this *RouteOptionStatus) UnmarshalJSON(b []byte) error {
	namespacedStatuses := RouteOptionNamespacedStatuses{}
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