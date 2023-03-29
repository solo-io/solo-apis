// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/policy/v2/security/dlp_policy.proto

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
func (m *DLPPolicySpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DLPPolicySpec)
	if !ok {
		that2, ok := that.(DLPPolicySpec)
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
func (m *DlpAction) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DlpAction)
	if !ok {
		that2, ok := that.(DlpAction)
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

	if h, ok := interface{}(m.GetShadow()).(equality.Equalizer); ok {
		if !h.Equal(target.GetShadow()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetShadow(), target.GetShadow()) {
			return false
		}
	}

	switch m.Action.(type) {

	case *DlpAction_PredefinedAction_:
		if _, ok := target.Action.(*DlpAction_PredefinedAction_); !ok {
			return false
		}

		if m.GetPredefinedAction() != target.GetPredefinedAction() {
			return false
		}

	case *DlpAction_CustomAction:
		if _, ok := target.Action.(*DlpAction_CustomAction); !ok {
			return false
		}

		if h, ok := interface{}(m.GetCustomAction()).(equality.Equalizer); ok {
			if !h.Equal(target.GetCustomAction()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetCustomAction(), target.GetCustomAction()) {
				return false
			}
		}

	case *DlpAction_KeyValueAction:
		if _, ok := target.Action.(*DlpAction_KeyValueAction); !ok {
			return false
		}

		if h, ok := interface{}(m.GetKeyValueAction()).(equality.Equalizer); ok {
			if !h.Equal(target.GetKeyValueAction()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetKeyValueAction(), target.GetKeyValueAction()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Action != target.Action {
			return false
		}
	}

	return true
}

// Equal function
func (m *DlpCustomAction) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DlpCustomAction)
	if !ok {
		that2, ok := that.(DlpCustomAction)
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

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetMaskChar()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMaskChar()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMaskChar(), target.GetMaskChar()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPercent()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPercent()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPercent(), target.GetPercent()) {
			return false
		}
	}

	if len(m.GetRegexActions()) != len(target.GetRegexActions()) {
		return false
	}
	for idx, v := range m.GetRegexActions() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetRegexActions()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetRegexActions()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *DlpKeyValueAction) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DlpKeyValueAction)
	if !ok {
		that2, ok := that.(DlpKeyValueAction)
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

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetMaskChar()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMaskChar()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMaskChar(), target.GetMaskChar()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPercent()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPercent()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPercent(), target.GetPercent()) {
			return false
		}
	}

	if len(m.GetKeysToMask()) != len(target.GetKeysToMask()) {
		return false
	}
	for idx, v := range m.GetKeysToMask() {

		if strings.Compare(v, target.GetKeysToMask()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *DLPPolicyStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DLPPolicyStatus)
	if !ok {
		that2, ok := that.(DLPPolicyStatus)
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
func (m *DLPPolicyReport) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DLPPolicyReport)
	if !ok {
		that2, ok := that.(DLPPolicyReport)
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
func (m *DLPPolicySpec_Config) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DLPPolicySpec_Config)
	if !ok {
		that2, ok := that.(DLPPolicySpec_Config)
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

	if len(m.GetActions()) != len(target.GetActions()) {
		return false
	}
	for idx, v := range m.GetActions() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetActions()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetActions()[idx]) {
				return false
			}
		}

	}

	if m.GetSanitize() != target.GetSanitize() {
		return false
	}

	return true
}
