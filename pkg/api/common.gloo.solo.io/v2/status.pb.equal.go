// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/common/v2/status.proto

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
func (m *OwnerWorkspace) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*OwnerWorkspace)
	if !ok {
		that2, ok := that.(OwnerWorkspace)
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

	if strings.Compare(m.GetWorkspace(), target.GetWorkspace()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *WorkspaceStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WorkspaceStatus)
	if !ok {
		that2, ok := that.(WorkspaceStatus)
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

	if len(m.GetClusters()) != len(target.GetClusters()) {
		return false
	}
	for k, v := range m.GetClusters() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetClusters()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetClusters()[k]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *K8SWorkloadStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*K8SWorkloadStatus)
	if !ok {
		that2, ok := that.(K8SWorkloadStatus)
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

	if len(m.GetAppliedWorkloadPolicies()) != len(target.GetAppliedWorkloadPolicies()) {
		return false
	}
	for k, v := range m.GetAppliedWorkloadPolicies() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAppliedWorkloadPolicies()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAppliedWorkloadPolicies()[k]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetOwnerWorkspace()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOwnerWorkspace()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOwnerWorkspace(), target.GetOwnerWorkspace()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *K8SServiceStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*K8SServiceStatus)
	if !ok {
		that2, ok := that.(K8SServiceStatus)
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

	if len(m.GetAppliedDestinationPolicies()) != len(target.GetAppliedDestinationPolicies()) {
		return false
	}
	for k, v := range m.GetAppliedDestinationPolicies() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAppliedDestinationPolicies()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAppliedDestinationPolicies()[k]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetOwnerWorkspace()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOwnerWorkspace()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOwnerWorkspace(), target.GetOwnerWorkspace()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *K8SServiceReport) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*K8SServiceReport)
	if !ok {
		that2, ok := that.(K8SServiceReport)
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

	if len(m.GetAppliedDestinationPolicies()) != len(target.GetAppliedDestinationPolicies()) {
		return false
	}
	for k, v := range m.GetAppliedDestinationPolicies() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAppliedDestinationPolicies()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAppliedDestinationPolicies()[k]) {
				return false
			}
		}

	}

	if strings.Compare(m.GetOwnerWorkspace(), target.GetOwnerWorkspace()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *GenericGlobalStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GenericGlobalStatus)
	if !ok {
		that2, ok := that.(GenericGlobalStatus)
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

	if m.GetState() != target.GetState() {
		return false
	}

	if strings.Compare(m.GetMessage(), target.GetMessage()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *GenericContextStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GenericContextStatus)
	if !ok {
		that2, ok := that.(GenericContextStatus)
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

	if m.GetObservedGeneration() != target.GetObservedGeneration() {
		return false
	}

	if m.GetState() != target.GetState() {
		return false
	}

	if strings.Compare(m.GetMessage(), target.GetMessage()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *RouteReference) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteReference)
	if !ok {
		that2, ok := that.(RouteReference)
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

	if strings.Compare(m.GetRouteName(), target.GetRouteName()) != 0 {
		return false
	}

	if m.GetRouteIndex() != target.GetRouteIndex() {
		return false
	}

	if h, ok := interface{}(m.GetRouteTable()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRouteTable()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRouteTable(), target.GetRouteTable()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *AppliedDestinationPortPolicies) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AppliedDestinationPortPolicies)
	if !ok {
		that2, ok := that.(AppliedDestinationPortPolicies)
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

	if len(m.GetPolicies()) != len(target.GetPolicies()) {
		return false
	}
	for idx, v := range m.GetPolicies() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPolicies()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPolicies()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *AppliedRoutePolicies) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AppliedRoutePolicies)
	if !ok {
		that2, ok := that.(AppliedRoutePolicies)
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

	if len(m.GetPolicies()) != len(target.GetPolicies()) {
		return false
	}
	for idx, v := range m.GetPolicies() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPolicies()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPolicies()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *AppliedWorkloadPolicies) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AppliedWorkloadPolicies)
	if !ok {
		that2, ok := that.(AppliedWorkloadPolicies)
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

	if len(m.GetPolicies()) != len(target.GetPolicies()) {
		return false
	}
	for idx, v := range m.GetPolicies() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPolicies()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPolicies()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *WorkspaceStatus_ClusterStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WorkspaceStatus_ClusterStatus)
	if !ok {
		that2, ok := that.(WorkspaceStatus_ClusterStatus)
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

	if h, ok := interface{}(m.GetGeneric()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGeneric()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGeneric(), target.GetGeneric()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *AppliedDestinationPortPolicies_DestinationPolicyReference) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AppliedDestinationPortPolicies_DestinationPolicyReference)
	if !ok {
		that2, ok := that.(AppliedDestinationPortPolicies_DestinationPolicyReference)
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

	if m.GetDestinationPort() != target.GetDestinationPort() {
		return false
	}

	if m.GetDestinationKind() != target.GetDestinationKind() {
		return false
	}

	if h, ok := interface{}(m.GetPolicy()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPolicy()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPolicy(), target.GetPolicy()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *AppliedRoutePolicies_RoutePolicyReference) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AppliedRoutePolicies_RoutePolicyReference)
	if !ok {
		that2, ok := that.(AppliedRoutePolicies_RoutePolicyReference)
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

	if strings.Compare(m.GetRouteName(), target.GetRouteName()) != 0 {
		return false
	}

	if m.GetRouteIndex() != target.GetRouteIndex() {
		return false
	}

	if h, ok := interface{}(m.GetPolicy()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPolicy()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPolicy(), target.GetPolicy()) {
			return false
		}
	}

	return true
}
