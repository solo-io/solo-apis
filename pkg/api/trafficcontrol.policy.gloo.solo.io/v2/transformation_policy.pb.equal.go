// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/policy/v2/trafficcontrol/transformation_policy.proto

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
func (m *TransformationPolicySpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TransformationPolicySpec)
	if !ok {
		that2, ok := that.(TransformationPolicySpec)
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

	if len(m.GetApplyToRoutes()) != len(target.GetApplyToRoutes()) {
		return false
	}
	for idx, v := range m.GetApplyToRoutes() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetApplyToRoutes()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetApplyToRoutes()[idx]) {
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
func (m *TransformationPolicyStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TransformationPolicyStatus)
	if !ok {
		that2, ok := that.(TransformationPolicyStatus)
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

	if m.GetNumSelectedRoutes() != target.GetNumSelectedRoutes() {
		return false
	}

	return true
}

// Equal function
func (m *TransformationPolicyReport) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TransformationPolicyReport)
	if !ok {
		that2, ok := that.(TransformationPolicyReport)
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

	if len(m.GetSelectedRoutes()) != len(target.GetSelectedRoutes()) {
		return false
	}
	for idx, v := range m.GetSelectedRoutes() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSelectedRoutes()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSelectedRoutes()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *TransformationPolicySpec_Config) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TransformationPolicySpec_Config)
	if !ok {
		that2, ok := that.(TransformationPolicySpec_Config)
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

	if h, ok := interface{}(m.GetPhase()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPhase()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPhase(), target.GetPhase()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRequest()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRequest()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRequest(), target.GetRequest()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetResponse()).(equality.Equalizer); ok {
		if !h.Equal(target.GetResponse()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetResponse(), target.GetResponse()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *TransformationPolicySpec_Config_RequestTransformation) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TransformationPolicySpec_Config_RequestTransformation)
	if !ok {
		that2, ok := that.(TransformationPolicySpec_Config_RequestTransformation)
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

	if m.GetRecalculateRoutingDestination() != target.GetRecalculateRoutingDestination() {
		return false
	}

	if h, ok := interface{}(m.GetInjaTemplate()).(equality.Equalizer); ok {
		if !h.Equal(target.GetInjaTemplate()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetInjaTemplate(), target.GetInjaTemplate()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *TransformationPolicySpec_Config_ResponseTransformation) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TransformationPolicySpec_Config_ResponseTransformation)
	if !ok {
		that2, ok := that.(TransformationPolicySpec_Config_ResponseTransformation)
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

	if h, ok := interface{}(m.GetInjaTemplate()).(equality.Equalizer); ok {
		if !h.Equal(target.GetInjaTemplate()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetInjaTemplate(), target.GetInjaTemplate()) {
			return false
		}
	}

	return true
}
