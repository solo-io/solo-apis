// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/policy/v2/resilience/connection_policy.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_duration "github.com/golang/protobuf/ptypes/duration"

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
func (m *ConnectionPolicySpec) Clone() proto.Message {
	var target *ConnectionPolicySpec
	if m == nil {
		return target
	}
	target = &ConnectionPolicySpec{}

	if m.GetApplyToDestinations() != nil {
		target.ApplyToDestinations = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationSelector, len(m.GetApplyToDestinations()))
		for idx, v := range m.GetApplyToDestinations() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ApplyToDestinations[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationSelector)
			} else {
				target.ApplyToDestinations[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationSelector)
			}

		}
	}

	if h, ok := interface{}(m.GetConfig()).(clone.Cloner); ok {
		target.Config = h.Clone().(*ConnectionPolicySpec_Config)
	} else {
		target.Config = proto.Clone(m.GetConfig()).(*ConnectionPolicySpec_Config)
	}

	return target
}

// Clone function
func (m *ConnectionPolicyStatus) Clone() proto.Message {
	var target *ConnectionPolicyStatus
	if m == nil {
		return target
	}
	target = &ConnectionPolicyStatus{}

	if h, ok := interface{}(m.GetCommon()).(clone.Cloner); ok {
		target.Common = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	} else {
		target.Common = proto.Clone(m.GetCommon()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	}

	target.NumSelectedDestinationPorts = m.GetNumSelectedDestinationPorts()

	return target
}

// Clone function
func (m *ConnectionPolicyReport) Clone() proto.Message {
	var target *ConnectionPolicyReport
	if m == nil {
		return target
	}
	target = &ConnectionPolicyReport{}

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

	if m.GetSelectedDestinationPorts() != nil {
		target.SelectedDestinationPorts = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference, len(m.GetSelectedDestinationPorts()))
		for idx, v := range m.GetSelectedDestinationPorts() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SelectedDestinationPorts[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference)
			} else {
				target.SelectedDestinationPorts[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference)
			}

		}
	}

	return target
}

// Clone function
func (m *ConnectionPolicySpec_Config) Clone() proto.Message {
	var target *ConnectionPolicySpec_Config
	if m == nil {
		return target
	}
	target = &ConnectionPolicySpec_Config{}

	if h, ok := interface{}(m.GetTcp()).(clone.Cloner); ok {
		target.Tcp = h.Clone().(*ConnectionPolicySpec_Config_TCPConfig)
	} else {
		target.Tcp = proto.Clone(m.GetTcp()).(*ConnectionPolicySpec_Config_TCPConfig)
	}

	if h, ok := interface{}(m.GetHttp()).(clone.Cloner); ok {
		target.Http = h.Clone().(*ConnectionPolicySpec_Config_HTTPConfig)
	} else {
		target.Http = proto.Clone(m.GetHttp()).(*ConnectionPolicySpec_Config_HTTPConfig)
	}

	return target
}

// Clone function
func (m *ConnectionPolicySpec_Config_TCPConfig) Clone() proto.Message {
	var target *ConnectionPolicySpec_Config_TCPConfig
	if m == nil {
		return target
	}
	target = &ConnectionPolicySpec_Config_TCPConfig{}

	if h, ok := interface{}(m.GetTcpKeepalive()).(clone.Cloner); ok {
		target.TcpKeepalive = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.TCPKeepalive)
	} else {
		target.TcpKeepalive = proto.Clone(m.GetTcpKeepalive()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.TCPKeepalive)
	}

	target.MaxConnections = m.GetMaxConnections()

	if h, ok := interface{}(m.GetConnectTimeout()).(clone.Cloner); ok {
		target.ConnectTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.ConnectTimeout = proto.Clone(m.GetConnectTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	return target
}

// Clone function
func (m *ConnectionPolicySpec_Config_HTTPConfig) Clone() proto.Message {
	var target *ConnectionPolicySpec_Config_HTTPConfig
	if m == nil {
		return target
	}
	target = &ConnectionPolicySpec_Config_HTTPConfig{}

	target.MaxRequestsPerConnection = m.GetMaxRequestsPerConnection()

	target.MaxRetries = m.GetMaxRetries()

	if h, ok := interface{}(m.GetIdleTimeout()).(clone.Cloner); ok {
		target.IdleTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.IdleTimeout = proto.Clone(m.GetIdleTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	return target
}
