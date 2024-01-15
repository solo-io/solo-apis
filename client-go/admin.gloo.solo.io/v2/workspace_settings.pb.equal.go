// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/admin/v2/workspace_settings.proto

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
func (m *WorkspaceSettingsSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WorkspaceSettingsSpec)
	if !ok {
		that2, ok := that.(WorkspaceSettingsSpec)
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

	if len(m.GetImportFrom()) != len(target.GetImportFrom()) {
		return false
	}
	for idx, v := range m.GetImportFrom() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetImportFrom()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetImportFrom()[idx]) {
				return false
			}
		}

	}

	if len(m.GetExportTo()) != len(target.GetExportTo()) {
		return false
	}
	for idx, v := range m.GetExportTo() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetExportTo()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetExportTo()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOptions(), target.GetOptions()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *WorkspaceSettingsStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WorkspaceSettingsStatus)
	if !ok {
		that2, ok := that.(WorkspaceSettingsStatus)
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

	if h, ok := interface{}(m.GetWorkspace()).(equality.Equalizer); ok {
		if !h.Equal(target.GetWorkspace()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetWorkspace(), target.GetWorkspace()) {
			return false
		}
	}

	if m.GetNumSelectedEastWestGateways() != target.GetNumSelectedEastWestGateways() {
		return false
	}

	if m.GetNumFederatedServices() != target.GetNumFederatedServices() {
		return false
	}

	return true
}

// Equal function
func (m *WorkspaceSettingsReport) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WorkspaceSettingsReport)
	if !ok {
		that2, ok := that.(WorkspaceSettingsReport)
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

	if h, ok := interface{}(m.GetWorkspace()).(equality.Equalizer); ok {
		if !h.Equal(target.GetWorkspace()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetWorkspace(), target.GetWorkspace()) {
			return false
		}
	}

	if len(m.GetSelectedEastWestGateways()) != len(target.GetSelectedEastWestGateways()) {
		return false
	}
	for idx, v := range m.GetSelectedEastWestGateways() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSelectedEastWestGateways()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSelectedEastWestGateways()[idx]) {
				return false
			}
		}

	}

	if len(m.GetFederatedServices()) != len(target.GetFederatedServices()) {
		return false
	}
	for idx, v := range m.GetFederatedServices() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetFederatedServices()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetFederatedServices()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *WorkspaceSettingsSpec_WorkspaceObjectSelector) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WorkspaceSettingsSpec_WorkspaceObjectSelector)
	if !ok {
		that2, ok := that.(WorkspaceSettingsSpec_WorkspaceObjectSelector)
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
	for idx, v := range m.GetWorkspaces() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetWorkspaces()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetWorkspaces()[idx]) {
				return false
			}
		}

	}

	if len(m.GetResources()) != len(target.GetResources()) {
		return false
	}
	for idx, v := range m.GetResources() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetResources()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetResources()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *WorkspaceSettingsSpec_Options) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WorkspaceSettingsSpec_Options)
	if !ok {
		that2, ok := that.(WorkspaceSettingsSpec_Options)
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

	if h, ok := interface{}(m.GetServiceIsolation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetServiceIsolation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetServiceIsolation(), target.GetServiceIsolation()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetFederation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetFederation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetFederation(), target.GetFederation()) {
			return false
		}
	}

	if len(m.GetEastWestGateways()) != len(target.GetEastWestGateways()) {
		return false
	}
	for idx, v := range m.GetEastWestGateways() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetEastWestGateways()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetEastWestGateways()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetVirtualDestClientMode()).(equality.Equalizer); ok {
		if !h.Equal(target.GetVirtualDestClientMode()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetVirtualDestClientMode(), target.GetVirtualDestClientMode()) {
			return false
		}
	}

	if m.GetTrimAllProxyConfig() != target.GetTrimAllProxyConfig() {
		return false
	}

	return true
}

// Equal function
func (m *WorkspaceSettingsSpec_WorkspaceObjectSelector_TypedObjectSelector) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WorkspaceSettingsSpec_WorkspaceObjectSelector_TypedObjectSelector)
	if !ok {
		that2, ok := that.(WorkspaceSettingsSpec_WorkspaceObjectSelector_TypedObjectSelector)
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

	if m.GetKind() != target.GetKind() {
		return false
	}

	if len(m.GetLabels()) != len(target.GetLabels()) {
		return false
	}
	for k, v := range m.GetLabels() {

		if strings.Compare(v, target.GetLabels()[k]) != 0 {
			return false
		}

	}

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if strings.Compare(m.GetNamespace(), target.GetNamespace()) != 0 {
		return false
	}

	if strings.Compare(m.GetCluster(), target.GetCluster()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *WorkspaceSettingsSpec_Options_ServiceIsolation) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WorkspaceSettingsSpec_Options_ServiceIsolation)
	if !ok {
		that2, ok := that.(WorkspaceSettingsSpec_Options_ServiceIsolation)
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

	if m.GetEnabled() != target.GetEnabled() {
		return false
	}

	if h, ok := interface{}(m.GetTrimProxyConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTrimProxyConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTrimProxyConfig(), target.GetTrimProxyConfig()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetEnforcementLayers()).(equality.Equalizer); ok {
		if !h.Equal(target.GetEnforcementLayers()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetEnforcementLayers(), target.GetEnforcementLayers()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *WorkspaceSettingsSpec_Options_Federation) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WorkspaceSettingsSpec_Options_Federation)
	if !ok {
		that2, ok := that.(WorkspaceSettingsSpec_Options_Federation)
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

	if m.GetEnabled() != target.GetEnabled() {
		return false
	}

	if strings.Compare(m.GetHostSuffix(), target.GetHostSuffix()) != 0 {
		return false
	}

	if len(m.GetServiceSelector()) != len(target.GetServiceSelector()) {
		return false
	}
	for idx, v := range m.GetServiceSelector() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetServiceSelector()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetServiceSelector()[idx]) {
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

	return true
}

// Equal function
func (m *WorkspaceSettingsSpec_Options_EastWestGatewaySelector) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WorkspaceSettingsSpec_Options_EastWestGatewaySelector)
	if !ok {
		that2, ok := that.(WorkspaceSettingsSpec_Options_EastWestGatewaySelector)
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

	if h, ok := interface{}(m.GetSelector()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSelector()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSelector(), target.GetSelector()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPort()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPort()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPort(), target.GetPort()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetTlsTerminationPort()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTlsTerminationPort()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTlsTerminationPort(), target.GetTlsTerminationPort()) {
			return false
		}
	}

	if len(m.GetHostInfoOverrides()) != len(target.GetHostInfoOverrides()) {
		return false
	}
	for idx, v := range m.GetHostInfoOverrides() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetHostInfoOverrides()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetHostInfoOverrides()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *WorkspaceSettingsSpec_Options_EastWestGatewaySelector_HostInfo) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WorkspaceSettingsSpec_Options_EastWestGatewaySelector_HostInfo)
	if !ok {
		that2, ok := that.(WorkspaceSettingsSpec_Options_EastWestGatewaySelector_HostInfo)
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

	if strings.Compare(m.GetAddr(), target.GetAddr()) != 0 {
		return false
	}

	if m.GetPort() != target.GetPort() {
		return false
	}

	return true
}
