// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/networking/v2/virtual_destination.proto

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
func (m *VirtualDestinationSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualDestinationSpec)
	if !ok {
		that2, ok := that.(VirtualDestinationSpec)
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

	if len(m.GetHosts()) != len(target.GetHosts()) {
		return false
	}
	for idx, v := range m.GetHosts() {

		if strings.Compare(v, target.GetHosts()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetServices()) != len(target.GetServices()) {
		return false
	}
	for idx, v := range m.GetServices() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetServices()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetServices()[idx]) {
				return false
			}
		}

	}

	if len(m.GetExternalServices()) != len(target.GetExternalServices()) {
		return false
	}
	for idx, v := range m.GetExternalServices() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetExternalServices()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetExternalServices()[idx]) {
				return false
			}
		}

	}

	if len(m.GetPorts()) != len(target.GetPorts()) {
		return false
	}
	for idx, v := range m.GetPorts() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPorts()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPorts()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetClientMode()).(equality.Equalizer); ok {
		if !h.Equal(target.GetClientMode()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetClientMode(), target.GetClientMode()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualDestinationStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualDestinationStatus)
	if !ok {
		that2, ok := that.(VirtualDestinationStatus)
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

	if len(m.GetAppliedDestinationPolicies()) != len(target.GetAppliedDestinationPolicies()) {
		return false
	}
	for k, v := range m.GetAppliedDestinationPolicies() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAppliedDestinationPolicies()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAppliedDestinationPolicies()[k]) {
				return false
			}
		}

	}

	if len(m.GetSelectedBackingServices()) != len(target.GetSelectedBackingServices()) {
		return false
	}
	for idx, v := range m.GetSelectedBackingServices() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSelectedBackingServices()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSelectedBackingServices()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetOwnerWorkspace()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOwnerWorkspace()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOwnerWorkspace(), target.GetOwnerWorkspace()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualDestinationSpec_PortMapping) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualDestinationSpec_PortMapping)
	if !ok {
		that2, ok := that.(VirtualDestinationSpec_PortMapping)
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

	if m.GetNumber() != target.GetNumber() {
		return false
	}

	if strings.Compare(m.GetProtocol(), target.GetProtocol()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetTargetPort()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTargetPort()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTargetPort(), target.GetTargetPort()) {
			return false
		}
	}

	return true
}
