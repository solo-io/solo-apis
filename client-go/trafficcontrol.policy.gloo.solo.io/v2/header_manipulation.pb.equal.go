// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/policy/v2/trafficcontrol/header_manipulation.proto

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
func (m *HeaderManipulationPolicySpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HeaderManipulationPolicySpec)
	if !ok {
		that2, ok := that.(HeaderManipulationPolicySpec)
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

	if len(m.GetApplyToRouteDestinations()) != len(target.GetApplyToRouteDestinations()) {
		return false
	}
	for idx, v := range m.GetApplyToRouteDestinations() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetApplyToRouteDestinations()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetApplyToRouteDestinations()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *HeaderManipulationPolicyStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HeaderManipulationPolicyStatus)
	if !ok {
		that2, ok := that.(HeaderManipulationPolicyStatus)
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
func (m *HeaderManipulationPolicyReport) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HeaderManipulationPolicyReport)
	if !ok {
		that2, ok := that.(HeaderManipulationPolicyReport)
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
func (m *HeaderManipulationPolicySpec_Config) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HeaderManipulationPolicySpec_Config)
	if !ok {
		that2, ok := that.(HeaderManipulationPolicySpec_Config)
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

	if len(m.GetRemoveResponseHeaders()) != len(target.GetRemoveResponseHeaders()) {
		return false
	}
	for idx, v := range m.GetRemoveResponseHeaders() {

		if strings.Compare(v, target.GetRemoveResponseHeaders()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetAppendResponseHeaders()) != len(target.GetAppendResponseHeaders()) {
		return false
	}
	for k, v := range m.GetAppendResponseHeaders() {

		if strings.Compare(v, target.GetAppendResponseHeaders()[k]) != 0 {
			return false
		}

	}

	if len(m.GetRemoveRequestHeaders()) != len(target.GetRemoveRequestHeaders()) {
		return false
	}
	for idx, v := range m.GetRemoveRequestHeaders() {

		if strings.Compare(v, target.GetRemoveRequestHeaders()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetAppendRequestHeaders()) != len(target.GetAppendRequestHeaders()) {
		return false
	}
	for k, v := range m.GetAppendRequestHeaders() {

		if strings.Compare(v, target.GetAppendRequestHeaders()[k]) != 0 {
			return false
		}

	}

	return true
}
