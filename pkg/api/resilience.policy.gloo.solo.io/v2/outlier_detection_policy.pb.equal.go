// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/policy/v2/resilience/outlier_detection_policy.proto

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
func (m *OutlierDetectionPolicySpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*OutlierDetectionPolicySpec)
	if !ok {
		that2, ok := that.(OutlierDetectionPolicySpec)
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

	if len(m.GetApplyToDestinations()) != len(target.GetApplyToDestinations()) {
		return false
	}
	for idx, v := range m.GetApplyToDestinations() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetApplyToDestinations()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetApplyToDestinations()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetConfig(), target.GetConfig()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *OutlierDetectionPolicyStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*OutlierDetectionPolicyStatus)
	if !ok {
		that2, ok := that.(OutlierDetectionPolicyStatus)
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

	if h, ok := interface{}(m.GetGlobal()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGlobal()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGlobal(), target.GetGlobal()) {
			return false
		}
	}

	if len(m.GetWorkspaces()) != len(target.GetWorkspaces()) {
		return false
	}
	for k, v := range m.GetWorkspaces() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetWorkspaces()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetWorkspaces()[k]) {
				return false
			}
		}

	}

	if len(m.GetSelectedDestinationPorts()) != len(target.GetSelectedDestinationPorts()) {
		return false
	}
	for idx, v := range m.GetSelectedDestinationPorts() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSelectedDestinationPorts()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSelectedDestinationPorts()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *OutlierDetectionPolicyNewStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*OutlierDetectionPolicyNewStatus)
	if !ok {
		that2, ok := that.(OutlierDetectionPolicyNewStatus)
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

	if h, ok := interface{}(m.GetCommon()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCommon()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCommon(), target.GetCommon()) {
			return false
		}
	}

	if m.GetSelectedDestinationPorts() != target.GetSelectedDestinationPorts() {
		return false
	}

	return true
}

// Equal function
func (m *OutlierDetectionPolicyReport) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*OutlierDetectionPolicyReport)
	if !ok {
		that2, ok := that.(OutlierDetectionPolicyReport)
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

	if len(m.GetWorkspaces()) != len(target.GetWorkspaces()) {
		return false
	}
	for k, v := range m.GetWorkspaces() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetWorkspaces()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetWorkspaces()[k]) {
				return false
			}
		}

	}

	if len(m.GetSelectedDestinationPorts()) != len(target.GetSelectedDestinationPorts()) {
		return false
	}
	for idx, v := range m.GetSelectedDestinationPorts() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSelectedDestinationPorts()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSelectedDestinationPorts()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *OutlierDetectionPolicySpec_Config) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*OutlierDetectionPolicySpec_Config)
	if !ok {
		that2, ok := that.(OutlierDetectionPolicySpec_Config)
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

	if m.GetConsecutiveErrors() != target.GetConsecutiveErrors() {
		return false
	}

	if h, ok := interface{}(m.GetInterval()).(equality.Equalizer); ok {
		if !h.Equal(target.GetInterval()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetInterval(), target.GetInterval()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetBaseEjectionTime()).(equality.Equalizer); ok {
		if !h.Equal(target.GetBaseEjectionTime()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetBaseEjectionTime(), target.GetBaseEjectionTime()) {
			return false
		}
	}

	if m.GetMaxEjectionPercent() != target.GetMaxEjectionPercent() {
		return false
	}

	return true
}
