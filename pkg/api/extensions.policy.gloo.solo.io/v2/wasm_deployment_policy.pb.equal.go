// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/policy/v2/extensions/wasm_deployment_policy.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"

	v1alpha3 "istio.io/api/networking/v1alpha3"
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

	_ = v1alpha3.EnvoyFilter_PatchContext(0)
)

// Equal function
func (m *WasmDeploymentPolicySpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WasmDeploymentPolicySpec)
	if !ok {
		that2, ok := that.(WasmDeploymentPolicySpec)
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
func (m *WasmDeploymentPolicyStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WasmDeploymentPolicyStatus)
	if !ok {
		that2, ok := that.(WasmDeploymentPolicyStatus)
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
func (m *WasmDeploymentPolicyNewStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WasmDeploymentPolicyNewStatus)
	if !ok {
		that2, ok := that.(WasmDeploymentPolicyNewStatus)
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

	if h, ok := interface{}(m.GetStatus()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStatus()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStatus(), target.GetStatus()) {
			return false
		}
	}

	if m.GetSelectedWorkloads() != target.GetSelectedWorkloads() {
		return false
	}

	return true
}

// Equal function
func (m *WasmDeploymentPolicyReport) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WasmDeploymentPolicyReport)
	if !ok {
		that2, ok := that.(WasmDeploymentPolicyReport)
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

	return true
}

// Equal function
func (m *WasmDeploymentPolicySpec_Config) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WasmDeploymentPolicySpec_Config)
	if !ok {
		that2, ok := that.(WasmDeploymentPolicySpec_Config)
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

	if m.GetWeight() != target.GetWeight() {
		return false
	}

	return true
}

// Equal function
func (m *WasmDeploymentPolicySpec_Config_WasmFilter) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WasmDeploymentPolicySpec_Config_WasmFilter)
	if !ok {
		that2, ok := that.(WasmDeploymentPolicySpec_Config_WasmFilter)
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

	if strings.Compare(m.GetRootId(), target.GetRootId()) != 0 {
		return false
	}

	if strings.Compare(m.GetVmId(), target.GetVmId()) != 0 {
		return false
	}

	if m.GetFilterContext() != target.GetFilterContext() {
		return false
	}

	if strings.Compare(m.GetInsertBeforeFilter(), target.GetInsertBeforeFilter()) != 0 {
		return false
	}

	switch m.FilterSource.(type) {

	case *WasmDeploymentPolicySpec_Config_WasmFilter_LocalPathSource:
		if _, ok := target.FilterSource.(*WasmDeploymentPolicySpec_Config_WasmFilter_LocalPathSource); !ok {
			return false
		}

		if strings.Compare(m.GetLocalPathSource(), target.GetLocalPathSource()) != 0 {
			return false
		}

	case *WasmDeploymentPolicySpec_Config_WasmFilter_HttpUriSource:
		if _, ok := target.FilterSource.(*WasmDeploymentPolicySpec_Config_WasmFilter_HttpUriSource); !ok {
			return false
		}

		if h, ok := interface{}(m.GetHttpUriSource()).(equality.Equalizer); ok {
			if !h.Equal(target.GetHttpUriSource()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetHttpUriSource(), target.GetHttpUriSource()) {
				return false
			}
		}

	case *WasmDeploymentPolicySpec_Config_WasmFilter_WasmImageSource_:
		if _, ok := target.FilterSource.(*WasmDeploymentPolicySpec_Config_WasmFilter_WasmImageSource_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetWasmImageSource()).(equality.Equalizer); ok {
			if !h.Equal(target.GetWasmImageSource()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetWasmImageSource(), target.GetWasmImageSource()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.FilterSource != target.FilterSource {
			return false
		}
	}

	switch m.FilterConfigSource.(type) {

	case *WasmDeploymentPolicySpec_Config_WasmFilter_StaticFilterConfig:
		if _, ok := target.FilterConfigSource.(*WasmDeploymentPolicySpec_Config_WasmFilter_StaticFilterConfig); !ok {
			return false
		}

		if h, ok := interface{}(m.GetStaticFilterConfig()).(equality.Equalizer); ok {
			if !h.Equal(target.GetStaticFilterConfig()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetStaticFilterConfig(), target.GetStaticFilterConfig()) {
				return false
			}
		}

	case *WasmDeploymentPolicySpec_Config_WasmFilter_DynamicFilterConfig:
		if _, ok := target.FilterConfigSource.(*WasmDeploymentPolicySpec_Config_WasmFilter_DynamicFilterConfig); !ok {
			return false
		}

		if strings.Compare(m.GetDynamicFilterConfig(), target.GetDynamicFilterConfig()) != 0 {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.FilterConfigSource != target.FilterConfigSource {
			return false
		}
	}

	return true
}

// Equal function
func (m *WasmDeploymentPolicySpec_Config_WasmFilter_UriSource) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WasmDeploymentPolicySpec_Config_WasmFilter_UriSource)
	if !ok {
		that2, ok := that.(WasmDeploymentPolicySpec_Config_WasmFilter_UriSource)
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

	if strings.Compare(m.GetUri(), target.GetUri()) != 0 {
		return false
	}

	if strings.Compare(m.GetSha(), target.GetSha()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *WasmDeploymentPolicySpec_Config_WasmFilter_WasmImageSource) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WasmDeploymentPolicySpec_Config_WasmFilter_WasmImageSource)
	if !ok {
		that2, ok := that.(WasmDeploymentPolicySpec_Config_WasmFilter_WasmImageSource)
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

	if strings.Compare(m.GetWasmImageTag(), target.GetWasmImageTag()) != 0 {
		return false
	}

	return true
}
