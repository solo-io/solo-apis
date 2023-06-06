// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/policy/v2/observability/access_log_policy.proto

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
func (m *AccessLogPolicySpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AccessLogPolicySpec)
	if !ok {
		that2, ok := that.(AccessLogPolicySpec)
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

	if len(m.GetApplyToWorkloads()) != len(target.GetApplyToWorkloads()) {
		return false
	}
	for idx, v := range m.GetApplyToWorkloads() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetApplyToWorkloads()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetApplyToWorkloads()[idx]) {
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
func (m *AccessLogPolicyStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AccessLogPolicyStatus)
	if !ok {
		that2, ok := that.(AccessLogPolicyStatus)
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

	if len(m.GetSelectedWorkloads()) != len(target.GetSelectedWorkloads()) {
		return false
	}
	for idx, v := range m.GetSelectedWorkloads() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSelectedWorkloads()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSelectedWorkloads()[idx]) {
				return false
			}
		}

	}

	if len(m.GetSelectedWorkloadRefs()) != len(target.GetSelectedWorkloadRefs()) {
		return false
	}
	for idx, v := range m.GetSelectedWorkloadRefs() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSelectedWorkloadRefs()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSelectedWorkloadRefs()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *AccessLogPolicySpec_Config) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AccessLogPolicySpec_Config)
	if !ok {
		that2, ok := that.(AccessLogPolicySpec_Config)
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

	if len(m.GetFilters()) != len(target.GetFilters()) {
		return false
	}
	for idx, v := range m.GetFilters() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetFilters()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetFilters()[idx]) {
				return false
			}
		}

	}

	if len(m.GetIncludedRequestHeaders()) != len(target.GetIncludedRequestHeaders()) {
		return false
	}
	for idx, v := range m.GetIncludedRequestHeaders() {

		if strings.Compare(v, target.GetIncludedRequestHeaders()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetIncludedResponseHeaders()) != len(target.GetIncludedResponseHeaders()) {
		return false
	}
	for idx, v := range m.GetIncludedResponseHeaders() {

		if strings.Compare(v, target.GetIncludedResponseHeaders()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetIncludedResponseTrailers()) != len(target.GetIncludedResponseTrailers()) {
		return false
	}
	for idx, v := range m.GetIncludedResponseTrailers() {

		if strings.Compare(v, target.GetIncludedResponseTrailers()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetIncludedFilterStateObjects()) != len(target.GetIncludedFilterStateObjects()) {
		return false
	}
	for idx, v := range m.GetIncludedFilterStateObjects() {

		if strings.Compare(v, target.GetIncludedFilterStateObjects()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *AccessLogPolicySpec_Config_Filter) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AccessLogPolicySpec_Config_Filter)
	if !ok {
		that2, ok := that.(AccessLogPolicySpec_Config_Filter)
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

	switch m.Type.(type) {

	case *AccessLogPolicySpec_Config_Filter_StatusCodeMatcher:
		if _, ok := target.Type.(*AccessLogPolicySpec_Config_Filter_StatusCodeMatcher); !ok {
			return false
		}

		if h, ok := interface{}(m.GetStatusCodeMatcher()).(equality.Equalizer); ok {
			if !h.Equal(target.GetStatusCodeMatcher()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetStatusCodeMatcher(), target.GetStatusCodeMatcher()) {
				return false
			}
		}

	case *AccessLogPolicySpec_Config_Filter_HeaderMatcher:
		if _, ok := target.Type.(*AccessLogPolicySpec_Config_Filter_HeaderMatcher); !ok {
			return false
		}

		if h, ok := interface{}(m.GetHeaderMatcher()).(equality.Equalizer); ok {
			if !h.Equal(target.GetHeaderMatcher()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetHeaderMatcher(), target.GetHeaderMatcher()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Type != target.Type {
			return false
		}
	}

	return true
}
