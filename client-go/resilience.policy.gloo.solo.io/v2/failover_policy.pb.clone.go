// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/policy/v2/resilience/failover_policy.proto

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

	google_golang_org_protobuf_types_known_wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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
func (m *FailoverPolicySpec) Clone() proto.Message {
	var target *FailoverPolicySpec
	if m == nil {
		return target
	}
	target = &FailoverPolicySpec{}

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
		target.Config = h.Clone().(*FailoverPolicySpec_Config)
	} else {
		target.Config = proto.Clone(m.GetConfig()).(*FailoverPolicySpec_Config)
	}

	return target
}

// Clone function
func (m *FailoverPolicyStatus) Clone() proto.Message {
	var target *FailoverPolicyStatus
	if m == nil {
		return target
	}
	target = &FailoverPolicyStatus{}

	if h, ok := interface{}(m.GetCommon()).(clone.Cloner); ok {
		target.Common = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	} else {
		target.Common = proto.Clone(m.GetCommon()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	}

	target.NumSelectedDestinationPorts = m.GetNumSelectedDestinationPorts()

	return target
}

// Clone function
func (m *FailoverPolicyReport) Clone() proto.Message {
	var target *FailoverPolicyReport
	if m == nil {
		return target
	}
	target = &FailoverPolicyReport{}

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
func (m *FailoverPolicySpec_Config) Clone() proto.Message {
	var target *FailoverPolicySpec_Config
	if m == nil {
		return target
	}
	target = &FailoverPolicySpec_Config{}

	if m.GetLocalityMappings() != nil {
		target.LocalityMappings = make([]*FailoverPolicySpec_Config_LocalityMappings, len(m.GetLocalityMappings()))
		for idx, v := range m.GetLocalityMappings() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.LocalityMappings[idx] = h.Clone().(*FailoverPolicySpec_Config_LocalityMappings)
			} else {
				target.LocalityMappings[idx] = proto.Clone(v).(*FailoverPolicySpec_Config_LocalityMappings)
			}

		}
	}

	return target
}

// Clone function
func (m *FailoverPolicySpec_Config_LocalityMappings) Clone() proto.Message {
	var target *FailoverPolicySpec_Config_LocalityMappings
	if m == nil {
		return target
	}
	target = &FailoverPolicySpec_Config_LocalityMappings{}

	if h, ok := interface{}(m.GetFrom()).(clone.Cloner); ok {
		target.From = h.Clone().(*FailoverPolicySpec_Config_LocalityMappings_OriginatingLocality)
	} else {
		target.From = proto.Clone(m.GetFrom()).(*FailoverPolicySpec_Config_LocalityMappings_OriginatingLocality)
	}

	if m.GetTo() != nil {
		target.To = make([]*FailoverPolicySpec_Config_LocalityMappings_DestinationLocality, len(m.GetTo()))
		for idx, v := range m.GetTo() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.To[idx] = h.Clone().(*FailoverPolicySpec_Config_LocalityMappings_DestinationLocality)
			} else {
				target.To[idx] = proto.Clone(v).(*FailoverPolicySpec_Config_LocalityMappings_DestinationLocality)
			}

		}
	}

	return target
}

// Clone function
func (m *FailoverPolicySpec_Config_LocalityMappings_OriginatingLocality) Clone() proto.Message {
	var target *FailoverPolicySpec_Config_LocalityMappings_OriginatingLocality
	if m == nil {
		return target
	}
	target = &FailoverPolicySpec_Config_LocalityMappings_OriginatingLocality{}

	target.Region = m.GetRegion()

	target.Zone = m.GetZone()

	target.SubZone = m.GetSubZone()

	return target
}

// Clone function
func (m *FailoverPolicySpec_Config_LocalityMappings_DestinationLocality) Clone() proto.Message {
	var target *FailoverPolicySpec_Config_LocalityMappings_DestinationLocality
	if m == nil {
		return target
	}
	target = &FailoverPolicySpec_Config_LocalityMappings_DestinationLocality{}

	target.Region = m.GetRegion()

	target.Zone = m.GetZone()

	target.SubZone = m.GetSubZone()

	if h, ok := interface{}(m.GetWeight()).(clone.Cloner); ok {
		target.Weight = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.Weight = proto.Clone(m.GetWeight()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	return target
}
