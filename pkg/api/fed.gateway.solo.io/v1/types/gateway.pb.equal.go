// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-fed/fed.gateway/v1/gateway.proto

package types

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *FederatedGatewaySpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FederatedGatewaySpec)
	if !ok {
		that2, ok := that.(FederatedGatewaySpec)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetTemplate()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTemplate()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTemplate(), target.GetTemplate()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPlacement()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPlacement()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPlacement(), target.GetPlacement()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *FederatedGatewayStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FederatedGatewayStatus)
	if !ok {
		that2, ok := that.(FederatedGatewayStatus)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetPlacementStatus()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPlacementStatus()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPlacementStatus(), target.GetPlacementStatus()) {
			return false
		}
	}

	if len(m.GetNamespacedPlacementStatuses()) != len(target.GetNamespacedPlacementStatuses()) {
		return false
	}
	for k, v := range m.GetNamespacedPlacementStatuses() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetNamespacedPlacementStatuses()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetNamespacedPlacementStatuses()[k]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *FederatedGatewaySpec_Template) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FederatedGatewaySpec_Template)
	if !ok {
		that2, ok := that.(FederatedGatewaySpec_Template)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetSpec()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSpec()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSpec(), target.GetSpec()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetMetadata()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMetadata()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMetadata(), target.GetMetadata()) {
			return false
		}
	}

	return true
}
