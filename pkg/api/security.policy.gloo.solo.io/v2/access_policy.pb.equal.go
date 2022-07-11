// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/policy/v2/security/access_policy.proto

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
func (m *AccessPolicySpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AccessPolicySpec)
	if !ok {
		that2, ok := that.(AccessPolicySpec)
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
func (m *AccessPolicyStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AccessPolicyStatus)
	if !ok {
		that2, ok := that.(AccessPolicyStatus)
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
func (m *AccessPolicySpec_Config) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AccessPolicySpec_Config)
	if !ok {
		that2, ok := that.(AccessPolicySpec_Config)
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

	if h, ok := interface{}(m.GetAuthn()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAuthn()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAuthn(), target.GetAuthn()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetAuthz()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAuthz()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAuthz(), target.GetAuthz()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *AccessPolicySpec_Config_Authentication) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AccessPolicySpec_Config_Authentication)
	if !ok {
		that2, ok := that.(AccessPolicySpec_Config_Authentication)
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

	if m.GetTlsMode() != target.GetTlsMode() {
		return false
	}

	return true
}

// Equal function
func (m *AccessPolicySpec_Config_Authorization) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AccessPolicySpec_Config_Authorization)
	if !ok {
		that2, ok := that.(AccessPolicySpec_Config_Authorization)
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

	if len(m.GetAllowedClients()) != len(target.GetAllowedClients()) {
		return false
	}
	for idx, v := range m.GetAllowedClients() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAllowedClients()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAllowedClients()[idx]) {
				return false
			}
		}

	}

	if len(m.GetAllowedPaths()) != len(target.GetAllowedPaths()) {
		return false
	}
	for idx, v := range m.GetAllowedPaths() {

		if strings.Compare(v, target.GetAllowedPaths()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetAllowedMethods()) != len(target.GetAllowedMethods()) {
		return false
	}
	for idx, v := range m.GetAllowedMethods() {

		if strings.Compare(v, target.GetAllowedMethods()[idx]) != 0 {
			return false
		}

	}

	return true
}