// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/policy/v2/security/csrf_policy.proto

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
func (m *CSRFPolicySpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CSRFPolicySpec)
	if !ok {
		that2, ok := that.(CSRFPolicySpec)
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
func (m *CSRFPolicyStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CSRFPolicyStatus)
	if !ok {
		that2, ok := that.(CSRFPolicyStatus)
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
func (m *CSRFPolicySpec_Config) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CSRFPolicySpec_Config)
	if !ok {
		that2, ok := that.(CSRFPolicySpec_Config)
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

	if m.GetFilterEnabled() != target.GetFilterEnabled() {
		return false
	}

	if m.GetShadowEnabled() != target.GetShadowEnabled() {
		return false
	}

	if h, ok := interface{}(m.GetPercentage()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPercentage()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPercentage(), target.GetPercentage()) {
			return false
		}
	}

	if len(m.GetAdditionalOrigins()) != len(target.GetAdditionalOrigins()) {
		return false
	}
	for idx, v := range m.GetAdditionalOrigins() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAdditionalOrigins()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAdditionalOrigins()[idx]) {
				return false
			}
		}

	}

	return true
}
