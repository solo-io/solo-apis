// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/networking/v2/external_service.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

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
func (m *ExternalServiceSpec) Clone() proto.Message {
	var target *ExternalServiceSpec
	if m == nil {
		return target
	}
	target = &ExternalServiceSpec{}

	if m.GetHosts() != nil {
		target.Hosts = make([]string, len(m.GetHosts()))
		for idx, v := range m.GetHosts() {

			target.Hosts[idx] = v

		}
	}

	if m.GetAddresses() != nil {
		target.Addresses = make([]string, len(m.GetAddresses()))
		for idx, v := range m.GetAddresses() {

			target.Addresses[idx] = v

		}
	}

	if m.GetPorts() != nil {
		target.Ports = make([]*ExternalServiceSpec_Port, len(m.GetPorts()))
		for idx, v := range m.GetPorts() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Ports[idx] = h.Clone().(*ExternalServiceSpec_Port)
			} else {
				target.Ports[idx] = proto.Clone(v).(*ExternalServiceSpec_Port)
			}

		}
	}

	if m.GetSelector() != nil {
		target.Selector = make(map[string]string, len(m.GetSelector()))
		for k, v := range m.GetSelector() {

			target.Selector[k] = v

		}
	}

	if m.GetSubjectAltNames() != nil {
		target.SubjectAltNames = make([]string, len(m.GetSubjectAltNames()))
		for idx, v := range m.GetSubjectAltNames() {

			target.SubjectAltNames[idx] = v

		}
	}

	return target
}

// Clone function
func (m *ExternalServiceStatus) Clone() proto.Message {
	var target *ExternalServiceStatus
	if m == nil {
		return target
	}
	target = &ExternalServiceStatus{}

	if h, ok := interface{}(m.GetGlobal()).(clone.Cloner); ok {
		target.Global = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.GenericGlobalStatus)
	} else {
		target.Global = proto.Clone(m.GetGlobal()).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.GenericGlobalStatus)
	}

	if m.GetWorkspaces() != nil {
		target.Workspaces = make(map[string]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.WorkspaceStatus, len(m.GetWorkspaces()))
		for k, v := range m.GetWorkspaces() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Workspaces[k] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.WorkspaceStatus)
			} else {
				target.Workspaces[k] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.WorkspaceStatus)
			}

		}
	}

	if m.GetAppliedDestinationPolicies() != nil {
		target.AppliedDestinationPolicies = make(map[string]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.AppliedDestinationPortPolicies, len(m.GetAppliedDestinationPolicies()))
		for k, v := range m.GetAppliedDestinationPolicies() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.AppliedDestinationPolicies[k] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.AppliedDestinationPortPolicies)
			} else {
				target.AppliedDestinationPolicies[k] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.AppliedDestinationPortPolicies)
			}

		}
	}

	if m.GetSelectedExternalEndpoints() != nil {
		target.SelectedExternalEndpoints = make([]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectReference, len(m.GetSelectedExternalEndpoints()))
		for idx, v := range m.GetSelectedExternalEndpoints() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SelectedExternalEndpoints[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectReference)
			} else {
				target.SelectedExternalEndpoints[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectReference)
			}

		}
	}

	if h, ok := interface{}(m.GetOwnerWorkspace()).(clone.Cloner); ok {
		target.OwnerWorkspace = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.OwnerWorkspace)
	} else {
		target.OwnerWorkspace = proto.Clone(m.GetOwnerWorkspace()).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.OwnerWorkspace)
	}

	return target
}

// Clone function
func (m *ExternalServiceSpec_Port) Clone() proto.Message {
	var target *ExternalServiceSpec_Port
	if m == nil {
		return target
	}
	target = &ExternalServiceSpec_Port{}

	target.Number = m.GetNumber()

	if h, ok := interface{}(m.GetTargetPort()).(clone.Cloner); ok {
		target.TargetPort = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.PortSelector)
	} else {
		target.TargetPort = proto.Clone(m.GetTargetPort()).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.PortSelector)
	}

	target.Name = m.GetName()

	target.Protocol = m.GetProtocol()

	if h, ok := interface{}(m.GetClientsideTls()).(clone.Cloner); ok {
		target.ClientsideTls = h.Clone().(*ExternalServiceSpec_Port_TlsConfig)
	} else {
		target.ClientsideTls = proto.Clone(m.GetClientsideTls()).(*ExternalServiceSpec_Port_TlsConfig)
	}

	return target
}

// Clone function
func (m *ExternalServiceSpec_Port_TlsConfig) Clone() proto.Message {
	var target *ExternalServiceSpec_Port_TlsConfig
	if m == nil {
		return target
	}
	target = &ExternalServiceSpec_Port_TlsConfig{}

	target.Sni = m.GetSni()

	return target
}
