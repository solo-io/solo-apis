// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/common/v2/status.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"
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
func (m *Status) Clone() proto.Message {
	var target *Status
	if m == nil {
		return target
	}
	target = &Status{}

	if h, ok := interface{}(m.GetState()).(clone.Cloner); ok {
		target.State = h.Clone().(*State)
	} else {
		target.State = proto.Clone(m.GetState()).(*State)
	}

	if m.GetWorkspaceConditions() != nil {
		target.WorkspaceConditions = make(map[string]uint32, len(m.GetWorkspaceConditions()))
		for k, v := range m.GetWorkspaceConditions() {

			target.WorkspaceConditions[k] = v

		}
	}

	return target
}

// Clone function
func (m *State) Clone() proto.Message {
	var target *State
	if m == nil {
		return target
	}
	target = &State{}

	target.ObservedGeneration = m.GetObservedGeneration()

	target.Approval = m.GetApproval()

	target.Message = m.GetMessage()

	return target
}

// Clone function
func (m *Report) Clone() proto.Message {
	var target *Report
	if m == nil {
		return target
	}
	target = &Report{}

	if m.GetClusters() != nil {
		target.Clusters = make(map[string]*State, len(m.GetClusters()))
		for k, v := range m.GetClusters() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Clusters[k] = h.Clone().(*State)
			} else {
				target.Clusters[k] = proto.Clone(v).(*State)
			}

		}
	}

	return target
}

// Clone function
func (m *RouteReference) Clone() proto.Message {
	var target *RouteReference
	if m == nil {
		return target
	}
	target = &RouteReference{}

	target.RouteName = m.GetRouteName()

	target.RouteIndex = m.GetRouteIndex()

	if h, ok := interface{}(m.GetRouteTable()).(clone.Cloner); ok {
		target.RouteTable = h.Clone().(*ObjectReference)
	} else {
		target.RouteTable = proto.Clone(m.GetRouteTable()).(*ObjectReference)
	}

	return target
}

// Clone function
func (m *AppliedDestinationPortPolicies) Clone() proto.Message {
	var target *AppliedDestinationPortPolicies
	if m == nil {
		return target
	}
	target = &AppliedDestinationPortPolicies{}

	if m.GetPolicies() != nil {
		target.Policies = make([]*AppliedDestinationPortPolicies_DestinationPolicyReference, len(m.GetPolicies()))
		for idx, v := range m.GetPolicies() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Policies[idx] = h.Clone().(*AppliedDestinationPortPolicies_DestinationPolicyReference)
			} else {
				target.Policies[idx] = proto.Clone(v).(*AppliedDestinationPortPolicies_DestinationPolicyReference)
			}

		}
	}

	return target
}

// Clone function
func (m *AppliedRoutePolicies) Clone() proto.Message {
	var target *AppliedRoutePolicies
	if m == nil {
		return target
	}
	target = &AppliedRoutePolicies{}

	if m.GetPolicies() != nil {
		target.Policies = make([]*AppliedRoutePolicies_RoutePolicyReference, len(m.GetPolicies()))
		for idx, v := range m.GetPolicies() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Policies[idx] = h.Clone().(*AppliedRoutePolicies_RoutePolicyReference)
			} else {
				target.Policies[idx] = proto.Clone(v).(*AppliedRoutePolicies_RoutePolicyReference)
			}

		}
	}

	return target
}

// Clone function
func (m *AppliedWorkloadPolicies) Clone() proto.Message {
	var target *AppliedWorkloadPolicies
	if m == nil {
		return target
	}
	target = &AppliedWorkloadPolicies{}

	if m.GetPolicies() != nil {
		target.Policies = make([]*ObjectReference, len(m.GetPolicies()))
		for idx, v := range m.GetPolicies() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Policies[idx] = h.Clone().(*ObjectReference)
			} else {
				target.Policies[idx] = proto.Clone(v).(*ObjectReference)
			}

		}
	}

	return target
}

// Clone function
func (m *AppliedDestinationPortPolicies_DestinationPolicyReference) Clone() proto.Message {
	var target *AppliedDestinationPortPolicies_DestinationPolicyReference
	if m == nil {
		return target
	}
	target = &AppliedDestinationPortPolicies_DestinationPolicyReference{}

	target.DestinationPort = m.GetDestinationPort()

	target.DestinationKind = m.GetDestinationKind()

	if h, ok := interface{}(m.GetPolicy()).(clone.Cloner); ok {
		target.Policy = h.Clone().(*ObjectReference)
	} else {
		target.Policy = proto.Clone(m.GetPolicy()).(*ObjectReference)
	}

	return target
}

// Clone function
func (m *AppliedRoutePolicies_RoutePolicyReference) Clone() proto.Message {
	var target *AppliedRoutePolicies_RoutePolicyReference
	if m == nil {
		return target
	}
	target = &AppliedRoutePolicies_RoutePolicyReference{}

	target.RouteName = m.GetRouteName()

	target.RouteIndex = m.GetRouteIndex()

	if h, ok := interface{}(m.GetPolicy()).(clone.Cloner); ok {
		target.Policy = h.Clone().(*ObjectReference)
	} else {
		target.Policy = proto.Clone(m.GetPolicy()).(*ObjectReference)
	}

	return target
}
