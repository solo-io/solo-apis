// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/networking/v2/virtual_destination.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2 "github.com/solo-io/solo-apis/client-go/common.gloo.solo.io/v2"
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
func (m *VirtualDestinationSpec) Clone() proto.Message {
	var target *VirtualDestinationSpec
	if m == nil {
		return target
	}
	target = &VirtualDestinationSpec{}

	if m.GetHosts() != nil {
		target.Hosts = make([]string, len(m.GetHosts()))
		for idx, v := range m.GetHosts() {

			target.Hosts[idx] = v

		}
	}

	if m.GetServices() != nil {
		target.Services = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectSelector, len(m.GetServices()))
		for idx, v := range m.GetServices() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Services[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectSelector)
			} else {
				target.Services[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectSelector)
			}

		}
	}

	if m.GetExternalServices() != nil {
		target.ExternalServices = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectSelector, len(m.GetExternalServices()))
		for idx, v := range m.GetExternalServices() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ExternalServices[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectSelector)
			} else {
				target.ExternalServices[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectSelector)
			}

		}
	}

	if m.GetExternalWorkloads() != nil {
		target.ExternalWorkloads = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectSelector, len(m.GetExternalWorkloads()))
		for idx, v := range m.GetExternalWorkloads() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ExternalWorkloads[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectSelector)
			} else {
				target.ExternalWorkloads[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectSelector)
			}

		}
	}

	if m.GetPorts() != nil {
		target.Ports = make([]*VirtualDestinationSpec_PortMapping, len(m.GetPorts()))
		for idx, v := range m.GetPorts() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Ports[idx] = h.Clone().(*VirtualDestinationSpec_PortMapping)
			} else {
				target.Ports[idx] = proto.Clone(v).(*VirtualDestinationSpec_PortMapping)
			}

		}
	}

	if h, ok := interface{}(m.GetClientMode()).(clone.Cloner); ok {
		target.ClientMode = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ClientMode)
	} else {
		target.ClientMode = proto.Clone(m.GetClientMode()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ClientMode)
	}

	return target
}

// Clone function
func (m *VirtualDestinationStatus) Clone() proto.Message {
	var target *VirtualDestinationStatus
	if m == nil {
		return target
	}
	target = &VirtualDestinationStatus{}

	if h, ok := interface{}(m.GetCommon()).(clone.Cloner); ok {
		target.Common = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	} else {
		target.Common = proto.Clone(m.GetCommon()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	}

	if m.GetNumAppliedDestinationPolicies() != nil {
		target.NumAppliedDestinationPolicies = make(map[string]uint32, len(m.GetNumAppliedDestinationPolicies()))
		for k, v := range m.GetNumAppliedDestinationPolicies() {

			target.NumAppliedDestinationPolicies[k] = v

		}
	}

	target.NumSelectedBackingServices = m.GetNumSelectedBackingServices()

	target.OwnedByWorkspace = m.GetOwnedByWorkspace()

	return target
}

// Clone function
func (m *VirtualDestinationReport) Clone() proto.Message {
	var target *VirtualDestinationReport
	if m == nil {
		return target
	}
	target = &VirtualDestinationReport{}

	if m.GetWorkspaces() != nil {
		target.Workspaces = make(map[string]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Report, len(m.GetWorkspaces()))
		for k, v := range m.GetWorkspaces() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Workspaces[k] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Report)
			} else {
				target.Workspaces[k] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Report)
			}

		}
	}

	if m.GetAppliedDestinationPolicies() != nil {
		target.AppliedDestinationPolicies = make(map[string]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.AppliedDestinationPortPolicies, len(m.GetAppliedDestinationPolicies()))
		for k, v := range m.GetAppliedDestinationPolicies() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.AppliedDestinationPolicies[k] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.AppliedDestinationPortPolicies)
			} else {
				target.AppliedDestinationPolicies[k] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.AppliedDestinationPortPolicies)
			}

		}
	}

	if m.GetSelectedBackingServices() != nil {
		target.SelectedBackingServices = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference, len(m.GetSelectedBackingServices()))
		for idx, v := range m.GetSelectedBackingServices() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SelectedBackingServices[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference)
			} else {
				target.SelectedBackingServices[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference)
			}

		}
	}

	target.OwnerWorkspace = m.GetOwnerWorkspace()

	return target
}

// Clone function
func (m *VirtualDestinationSpec_PortMapping) Clone() proto.Message {
	var target *VirtualDestinationSpec_PortMapping
	if m == nil {
		return target
	}
	target = &VirtualDestinationSpec_PortMapping{}

	target.Number = m.GetNumber()

	target.Protocol = m.GetProtocol()

	if h, ok := interface{}(m.GetTargetPort()).(clone.Cloner); ok {
		target.TargetPort = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.PortSelector)
	} else {
		target.TargetPort = proto.Clone(m.GetTargetPort()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.PortSelector)
	}

	return target
}
