// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/admin/v2/workspace_settings.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"

	github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2 "github.com/solo-io/solo-apis/pkg/api/common.gloo.solo.io/v2"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *WorkspaceSettingsSpec) Clone() proto.Message {
	var target *WorkspaceSettingsSpec
	if m == nil {
		return target
	}
	target = &WorkspaceSettingsSpec{}

	if m.GetImportFrom() != nil {
		target.ImportFrom = make([]*WorkspaceSettingsSpec_WorkspaceObjectSelector, len(m.GetImportFrom()))
		for idx, v := range m.GetImportFrom() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ImportFrom[idx] = h.Clone().(*WorkspaceSettingsSpec_WorkspaceObjectSelector)
			} else {
				target.ImportFrom[idx] = proto.Clone(v).(*WorkspaceSettingsSpec_WorkspaceObjectSelector)
			}

		}
	}

	if m.GetExportTo() != nil {
		target.ExportTo = make([]*WorkspaceSettingsSpec_WorkspaceObjectSelector, len(m.GetExportTo()))
		for idx, v := range m.GetExportTo() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ExportTo[idx] = h.Clone().(*WorkspaceSettingsSpec_WorkspaceObjectSelector)
			} else {
				target.ExportTo[idx] = proto.Clone(v).(*WorkspaceSettingsSpec_WorkspaceObjectSelector)
			}

		}
	}

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*WorkspaceSettingsSpec_Options)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*WorkspaceSettingsSpec_Options)
	}

	return target
}

// Clone function
func (m *WorkspaceSettingsStatus) Clone() proto.Message {
	var target *WorkspaceSettingsStatus
	if m == nil {
		return target
	}
	target = &WorkspaceSettingsStatus{}

	if h, ok := interface{}(m.GetGeneric()).(clone.Cloner); ok {
		target.Generic = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.GenericContextStatus)
	} else {
		target.Generic = proto.Clone(m.GetGeneric()).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.GenericContextStatus)
	}

	if h, ok := interface{}(m.GetWorkspace()).(clone.Cloner); ok {
		target.Workspace = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectReference)
	} else {
		target.Workspace = proto.Clone(m.GetWorkspace()).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectReference)
	}

	if m.GetSelectedEastWestGateways() != nil {
		target.SelectedEastWestGateways = make([]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectReference, len(m.GetSelectedEastWestGateways()))
		for idx, v := range m.GetSelectedEastWestGateways() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SelectedEastWestGateways[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectReference)
			} else {
				target.SelectedEastWestGateways[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectReference)
			}

		}
	}

	if m.GetFederatedServices() != nil {
		target.FederatedServices = make([]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectReference, len(m.GetFederatedServices()))
		for idx, v := range m.GetFederatedServices() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.FederatedServices[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectReference)
			} else {
				target.FederatedServices[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectReference)
			}

		}
	}

	return target
}

// Clone function
func (m *WorkspaceSettingsSpec_WorkspaceObjectSelector) Clone() proto.Message {
	var target *WorkspaceSettingsSpec_WorkspaceObjectSelector
	if m == nil {
		return target
	}
	target = &WorkspaceSettingsSpec_WorkspaceObjectSelector{}

	if m.GetWorkspaces() != nil {
		target.Workspaces = make([]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.WorkspaceSelector, len(m.GetWorkspaces()))
		for idx, v := range m.GetWorkspaces() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Workspaces[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.WorkspaceSelector)
			} else {
				target.Workspaces[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.WorkspaceSelector)
			}

		}
	}

	if m.GetResources() != nil {
		target.Resources = make([]*WorkspaceSettingsSpec_WorkspaceObjectSelector_TypedObjectSelector, len(m.GetResources()))
		for idx, v := range m.GetResources() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Resources[idx] = h.Clone().(*WorkspaceSettingsSpec_WorkspaceObjectSelector_TypedObjectSelector)
			} else {
				target.Resources[idx] = proto.Clone(v).(*WorkspaceSettingsSpec_WorkspaceObjectSelector_TypedObjectSelector)
			}

		}
	}

	return target
}

// Clone function
func (m *WorkspaceSettingsSpec_Options) Clone() proto.Message {
	var target *WorkspaceSettingsSpec_Options
	if m == nil {
		return target
	}
	target = &WorkspaceSettingsSpec_Options{}

	if h, ok := interface{}(m.GetServiceIsolation()).(clone.Cloner); ok {
		target.ServiceIsolation = h.Clone().(*WorkspaceSettingsSpec_Options_ServiceIsolation)
	} else {
		target.ServiceIsolation = proto.Clone(m.GetServiceIsolation()).(*WorkspaceSettingsSpec_Options_ServiceIsolation)
	}

	if h, ok := interface{}(m.GetFederation()).(clone.Cloner); ok {
		target.Federation = h.Clone().(*WorkspaceSettingsSpec_Options_Federation)
	} else {
		target.Federation = proto.Clone(m.GetFederation()).(*WorkspaceSettingsSpec_Options_Federation)
	}

	if m.GetEastWestGateways() != nil {
		target.EastWestGateways = make([]*WorkspaceSettingsSpec_Options_EastWestGatewaySelector, len(m.GetEastWestGateways()))
		for idx, v := range m.GetEastWestGateways() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.EastWestGateways[idx] = h.Clone().(*WorkspaceSettingsSpec_Options_EastWestGatewaySelector)
			} else {
				target.EastWestGateways[idx] = proto.Clone(v).(*WorkspaceSettingsSpec_Options_EastWestGatewaySelector)
			}

		}
	}

	return target
}

// Clone function
func (m *WorkspaceSettingsSpec_WorkspaceObjectSelector_TypedObjectSelector) Clone() proto.Message {
	var target *WorkspaceSettingsSpec_WorkspaceObjectSelector_TypedObjectSelector
	if m == nil {
		return target
	}
	target = &WorkspaceSettingsSpec_WorkspaceObjectSelector_TypedObjectSelector{}

	target.Kind = m.GetKind()

	if m.GetLabels() != nil {
		target.Labels = make(map[string]string, len(m.GetLabels()))
		for k, v := range m.GetLabels() {

			target.Labels[k] = v

		}
	}

	target.Name = m.GetName()

	target.Namespace = m.GetNamespace()

	target.Cluster = m.GetCluster()

	return target
}

// Clone function
func (m *WorkspaceSettingsSpec_Options_ServiceIsolation) Clone() proto.Message {
	var target *WorkspaceSettingsSpec_Options_ServiceIsolation
	if m == nil {
		return target
	}
	target = &WorkspaceSettingsSpec_Options_ServiceIsolation{}

	target.Enabled = m.GetEnabled()

	if h, ok := interface{}(m.GetTrimProxyConfig()).(clone.Cloner); ok {
		target.TrimProxyConfig = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.TrimProxyConfig = proto.Clone(m.GetTrimProxyConfig()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	return target
}

// Clone function
func (m *WorkspaceSettingsSpec_Options_Federation) Clone() proto.Message {
	var target *WorkspaceSettingsSpec_Options_Federation
	if m == nil {
		return target
	}
	target = &WorkspaceSettingsSpec_Options_Federation{}

	target.Enabled = m.GetEnabled()

	target.HostSuffix = m.GetHostSuffix()

	if m.GetServiceSelector() != nil {
		target.ServiceSelector = make([]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectSelector, len(m.GetServiceSelector()))
		for idx, v := range m.GetServiceSelector() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ServiceSelector[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectSelector)
			} else {
				target.ServiceSelector[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectSelector)
			}

		}
	}

	if m.GetPorts() != nil {
		target.Ports = make([]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.PortSelector, len(m.GetPorts()))
		for idx, v := range m.GetPorts() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Ports[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.PortSelector)
			} else {
				target.Ports[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.PortSelector)
			}

		}
	}

	return target
}

// Clone function
func (m *WorkspaceSettingsSpec_Options_EastWestGatewaySelector) Clone() proto.Message {
	var target *WorkspaceSettingsSpec_Options_EastWestGatewaySelector
	if m == nil {
		return target
	}
	target = &WorkspaceSettingsSpec_Options_EastWestGatewaySelector{}

	if h, ok := interface{}(m.GetSelector()).(clone.Cloner); ok {
		target.Selector = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectSelector)
	} else {
		target.Selector = proto.Clone(m.GetSelector()).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectSelector)
	}

	if h, ok := interface{}(m.GetPort()).(clone.Cloner); ok {
		target.Port = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.PortSelector)
	} else {
		target.Port = proto.Clone(m.GetPort()).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.PortSelector)
	}

	if m.GetHostInfoOverrides() != nil {
		target.HostInfoOverrides = make([]*WorkspaceSettingsSpec_Options_EastWestGatewaySelector_HostInfo, len(m.GetHostInfoOverrides()))
		for idx, v := range m.GetHostInfoOverrides() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.HostInfoOverrides[idx] = h.Clone().(*WorkspaceSettingsSpec_Options_EastWestGatewaySelector_HostInfo)
			} else {
				target.HostInfoOverrides[idx] = proto.Clone(v).(*WorkspaceSettingsSpec_Options_EastWestGatewaySelector_HostInfo)
			}

		}
	}

	return target
}

// Clone function
func (m *WorkspaceSettingsSpec_Options_EastWestGatewaySelector_HostInfo) Clone() proto.Message {
	var target *WorkspaceSettingsSpec_Options_EastWestGatewaySelector_HostInfo
	if m == nil {
		return target
	}
	target = &WorkspaceSettingsSpec_Options_EastWestGatewaySelector_HostInfo{}

	target.Addr = m.GetAddr()

	target.Port = m.GetPort()

	return target
}
