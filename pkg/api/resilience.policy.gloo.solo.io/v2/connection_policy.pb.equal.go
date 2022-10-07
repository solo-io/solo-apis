// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/policy/v2/resilience/connection_policy.proto

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
func (m *ConnectionPolicySpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ConnectionPolicySpec)
	if !ok {
		that2, ok := that.(ConnectionPolicySpec)
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
func (m *ConnectionPolicyStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ConnectionPolicyStatus)
	if !ok {
		that2, ok := that.(ConnectionPolicyStatus)
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

	if len(m.GetSelectedDestiantionPorts()) != len(target.GetSelectedDestiantionPorts()) {
		return false
	}
	for idx, v := range m.GetSelectedDestiantionPorts() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSelectedDestiantionPorts()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSelectedDestiantionPorts()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *ConnectionPolicySpec_Config) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ConnectionPolicySpec_Config)
	if !ok {
		that2, ok := that.(ConnectionPolicySpec_Config)
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

	if h, ok := interface{}(m.GetTcp()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTcp()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTcp(), target.GetTcp()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ConnectionPolicySpec_Config_TCPConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ConnectionPolicySpec_Config_TCPConfig)
	if !ok {
		that2, ok := that.(ConnectionPolicySpec_Config_TCPConfig)
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

	if h, ok := interface{}(m.GetTcpKeepalive()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTcpKeepalive()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTcpKeepalive(), target.GetTcpKeepalive()) {
			return false
		}
	}

	if m.GetMaxConnections() != target.GetMaxConnections() {
		return false
	}

	if h, ok := interface{}(m.GetConnectTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetConnectTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetConnectTimeout(), target.GetConnectTimeout()) {
			return false
		}
	}

	return true
}