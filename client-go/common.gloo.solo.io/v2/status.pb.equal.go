// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/common/v2/status.proto

package v2

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
func (m *Status) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Status)
	if !ok {
		that2, ok := that.(Status)
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

	if h, ok := interface{}(m.GetState()).(equality.Equalizer); ok {
		if !h.Equal(target.GetState()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetState(), target.GetState()) {
			return false
		}
	}

	if len(m.GetWorkspaceConditions()) != len(target.GetWorkspaceConditions()) {
		return false
	}
	for k, v := range m.GetWorkspaceConditions() {

		if v != target.GetWorkspaceConditions()[k] {
			return false
		}

	}

	return true
}

// Equal function
func (m *State) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*State)
	if !ok {
		that2, ok := that.(State)
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

	if m.GetObservedGeneration() != target.GetObservedGeneration() {
		return false
	}

	if m.GetApproval() != target.GetApproval() {
		return false
	}

	if strings.Compare(m.GetMessage(), target.GetMessage()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *Report) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Report)
	if !ok {
		that2, ok := that.(Report)
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

	if len(m.GetClusters()) != len(target.GetClusters()) {
		return false
	}
	for k, v := range m.GetClusters() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetClusters()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetClusters()[k]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *RouteReference) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteReference)
	if !ok {
		that2, ok := that.(RouteReference)
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

	if strings.Compare(m.GetRouteName(), target.GetRouteName()) != 0 {
		return false
	}

	if m.GetRouteIndex() != target.GetRouteIndex() {
		return false
	}

	if h, ok := interface{}(m.GetRouteTable()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRouteTable()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRouteTable(), target.GetRouteTable()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *AppliedDestinationPortPolicies) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AppliedDestinationPortPolicies)
	if !ok {
		that2, ok := that.(AppliedDestinationPortPolicies)
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

	if len(m.GetPolicies()) != len(target.GetPolicies()) {
		return false
	}
	for idx, v := range m.GetPolicies() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPolicies()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPolicies()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *AppliedRoutePolicies) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AppliedRoutePolicies)
	if !ok {
		that2, ok := that.(AppliedRoutePolicies)
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

	if len(m.GetPolicies()) != len(target.GetPolicies()) {
		return false
	}
	for idx, v := range m.GetPolicies() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPolicies()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPolicies()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *AppliedWorkloadPolicies) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AppliedWorkloadPolicies)
	if !ok {
		that2, ok := that.(AppliedWorkloadPolicies)
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

	if len(m.GetPolicies()) != len(target.GetPolicies()) {
		return false
	}
	for idx, v := range m.GetPolicies() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPolicies()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPolicies()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *AppliedDestinationPortPolicies_DestinationPolicyReference) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AppliedDestinationPortPolicies_DestinationPolicyReference)
	if !ok {
		that2, ok := that.(AppliedDestinationPortPolicies_DestinationPolicyReference)
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

	if m.GetDestinationPort() != target.GetDestinationPort() {
		return false
	}

	if m.GetDestinationKind() != target.GetDestinationKind() {
		return false
	}

	if h, ok := interface{}(m.GetPolicy()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPolicy()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPolicy(), target.GetPolicy()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *AppliedRoutePolicies_RoutePolicyReference) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AppliedRoutePolicies_RoutePolicyReference)
	if !ok {
		that2, ok := that.(AppliedRoutePolicies_RoutePolicyReference)
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

	if strings.Compare(m.GetRouteName(), target.GetRouteName()) != 0 {
		return false
	}

	if m.GetRouteIndex() != target.GetRouteIndex() {
		return false
	}

	if h, ok := interface{}(m.GetPolicy()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPolicy()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPolicy(), target.GetPolicy()) {
			return false
		}
	}

	return true
}