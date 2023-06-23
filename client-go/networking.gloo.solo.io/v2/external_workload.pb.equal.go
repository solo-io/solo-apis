// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/networking/v2/external_workload.proto

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
func (m *ExternalWorkloadSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExternalWorkloadSpec)
	if !ok {
		that2, ok := that.(ExternalWorkloadSpec)
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

	if h, ok := interface{}(m.GetIdentitySelector()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIdentitySelector()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIdentitySelector(), target.GetIdentitySelector()) {
			return false
		}
	}

	if len(m.GetConnectedClusters()) != len(target.GetConnectedClusters()) {
		return false
	}
	for k, v := range m.GetConnectedClusters() {

		if strings.Compare(v, target.GetConnectedClusters()[k]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetReadinessProbe()).(equality.Equalizer); ok {
		if !h.Equal(target.GetReadinessProbe()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetReadinessProbe(), target.GetReadinessProbe()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ExternalWorkloadStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExternalWorkloadStatus)
	if !ok {
		that2, ok := that.(ExternalWorkloadStatus)
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

	if len(m.GetNumAppliedPolicies()) != len(target.GetNumAppliedPolicies()) {
		return false
	}
	for k, v := range m.GetNumAppliedPolicies() {

		if v != target.GetNumAppliedPolicies()[k] {
			return false
		}

	}

	if strings.Compare(m.GetOwnedByWorkspace(), target.GetOwnedByWorkspace()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *ExternalWorkloadReport) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExternalWorkloadReport)
	if !ok {
		that2, ok := that.(ExternalWorkloadReport)
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
func (m *ExternalWorkloadSpec_Port) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExternalWorkloadSpec_Port)
	if !ok {
		that2, ok := that.(ExternalWorkloadSpec_Port)
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

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if strings.Compare(m.GetProtocol(), target.GetProtocol()) != 0 {
		return false
	}

	if m.GetNumber() != target.GetNumber() {
		return false
	}

	return true
}

// Equal function
func (m *ExternalWorkloadSpec_IdentitySelector) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExternalWorkloadSpec_IdentitySelector)
	if !ok {
		that2, ok := that.(ExternalWorkloadSpec_IdentitySelector)
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

	if len(m.GetAws()) != len(target.GetAws()) {
		return false
	}
	for idx, v := range m.GetAws() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAws()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAws()[idx]) {
				return false
			}
		}

	}

	if len(m.GetGcp()) != len(target.GetGcp()) {
		return false
	}
	for idx, v := range m.GetGcp() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetGcp()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetGcp()[idx]) {
				return false
			}
		}

	}

	if len(m.GetAzure()) != len(target.GetAzure()) {
		return false
	}
	for idx, v := range m.GetAzure() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAzure()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAzure()[idx]) {
				return false
			}
		}

	}

	if len(m.GetJoinTokenSpiffeId()) != len(target.GetJoinTokenSpiffeId()) {
		return false
	}
	for idx, v := range m.GetJoinTokenSpiffeId() {

		if strings.Compare(v, target.GetJoinTokenSpiffeId()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *ExternalWorkloadSpec_Probe) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExternalWorkloadSpec_Probe)
	if !ok {
		that2, ok := that.(ExternalWorkloadSpec_Probe)
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

	if h, ok := interface{}(m.GetInitialDelaySeconds()).(equality.Equalizer); ok {
		if !h.Equal(target.GetInitialDelaySeconds()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetInitialDelaySeconds(), target.GetInitialDelaySeconds()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetTimeoutSeconds()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTimeoutSeconds()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTimeoutSeconds(), target.GetTimeoutSeconds()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPeriodSeconds()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPeriodSeconds()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPeriodSeconds(), target.GetPeriodSeconds()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSuccessThreshold()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSuccessThreshold()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSuccessThreshold(), target.GetSuccessThreshold()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetFailureThreshold()).(equality.Equalizer); ok {
		if !h.Equal(target.GetFailureThreshold()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetFailureThreshold(), target.GetFailureThreshold()) {
			return false
		}
	}

	switch m.Handler.(type) {

	case *ExternalWorkloadSpec_Probe_HttpGet:
		if _, ok := target.Handler.(*ExternalWorkloadSpec_Probe_HttpGet); !ok {
			return false
		}

		if h, ok := interface{}(m.GetHttpGet()).(equality.Equalizer); ok {
			if !h.Equal(target.GetHttpGet()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetHttpGet(), target.GetHttpGet()) {
				return false
			}
		}

	case *ExternalWorkloadSpec_Probe_TcpSocket:
		if _, ok := target.Handler.(*ExternalWorkloadSpec_Probe_TcpSocket); !ok {
			return false
		}

		if h, ok := interface{}(m.GetTcpSocket()).(equality.Equalizer); ok {
			if !h.Equal(target.GetTcpSocket()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetTcpSocket(), target.GetTcpSocket()) {
				return false
			}
		}

	case *ExternalWorkloadSpec_Probe_Exec:
		if _, ok := target.Handler.(*ExternalWorkloadSpec_Probe_Exec); !ok {
			return false
		}

		if h, ok := interface{}(m.GetExec()).(equality.Equalizer); ok {
			if !h.Equal(target.GetExec()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetExec(), target.GetExec()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Handler != target.Handler {
			return false
		}
	}

	return true
}

// Equal function
func (m *ExternalWorkloadSpec_IdentitySelector_AWS) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExternalWorkloadSpec_IdentitySelector_AWS)
	if !ok {
		that2, ok := that.(ExternalWorkloadSpec_IdentitySelector_AWS)
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

	if strings.Compare(m.GetIamRole(), target.GetIamRole()) != 0 {
		return false
	}

	if strings.Compare(m.GetSecurityGroupName(), target.GetSecurityGroupName()) != 0 {
		return false
	}

	if strings.Compare(m.GetSecurityGroupId(), target.GetSecurityGroupId()) != 0 {
		return false
	}

	if strings.Compare(m.GetImageId(), target.GetImageId()) != 0 {
		return false
	}

	if strings.Compare(m.GetInstanceId(), target.GetInstanceId()) != 0 {
		return false
	}

	if strings.Compare(m.GetZone(), target.GetZone()) != 0 {
		return false
	}

	if strings.Compare(m.GetRegion(), target.GetRegion()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetTag()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTag()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTag(), target.GetTag()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ExternalWorkloadSpec_IdentitySelector_GCP) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExternalWorkloadSpec_IdentitySelector_GCP)
	if !ok {
		that2, ok := that.(ExternalWorkloadSpec_IdentitySelector_GCP)
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

	if strings.Compare(m.GetServiceAccount(), target.GetServiceAccount()) != 0 {
		return false
	}

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if strings.Compare(m.GetTag(), target.GetTag()) != 0 {
		return false
	}

	if strings.Compare(m.GetProjectId(), target.GetProjectId()) != 0 {
		return false
	}

	if strings.Compare(m.GetZone(), target.GetZone()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetLabel()).(equality.Equalizer); ok {
		if !h.Equal(target.GetLabel()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetLabel(), target.GetLabel()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ExternalWorkloadSpec_IdentitySelector_Azure) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExternalWorkloadSpec_IdentitySelector_Azure)
	if !ok {
		that2, ok := that.(ExternalWorkloadSpec_IdentitySelector_Azure)
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

	if strings.Compare(m.GetSubscriptionId(), target.GetSubscriptionId()) != 0 {
		return false
	}

	if strings.Compare(m.GetSecurityGroup(), target.GetSecurityGroup()) != 0 {
		return false
	}

	if strings.Compare(m.GetVirtualNetwork(), target.GetVirtualNetwork()) != 0 {
		return false
	}

	if strings.Compare(m.GetSubnet(), target.GetSubnet()) != 0 {
		return false
	}

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if strings.Compare(m.GetResourceGroup(), target.GetResourceGroup()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *ExternalWorkloadSpec_IdentitySelector_AWS_Tag) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExternalWorkloadSpec_IdentitySelector_AWS_Tag)
	if !ok {
		that2, ok := that.(ExternalWorkloadSpec_IdentitySelector_AWS_Tag)
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

	if strings.Compare(m.GetKey(), target.GetKey()) != 0 {
		return false
	}

	if strings.Compare(m.GetValue(), target.GetValue()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *ExternalWorkloadSpec_IdentitySelector_GCP_Label) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExternalWorkloadSpec_IdentitySelector_GCP_Label)
	if !ok {
		that2, ok := that.(ExternalWorkloadSpec_IdentitySelector_GCP_Label)
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

	if strings.Compare(m.GetKey(), target.GetKey()) != 0 {
		return false
	}

	if strings.Compare(m.GetValue(), target.GetValue()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *ExternalWorkloadSpec_Probe_HTTPGetConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExternalWorkloadSpec_Probe_HTTPGetConfig)
	if !ok {
		that2, ok := that.(ExternalWorkloadSpec_Probe_HTTPGetConfig)
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

	if m.GetPort() != target.GetPort() {
		return false
	}

	if strings.Compare(m.GetPath(), target.GetPath()) != 0 {
		return false
	}

	if m.GetScheme() != target.GetScheme() {
		return false
	}

	if len(m.GetHttpHeaders()) != len(target.GetHttpHeaders()) {
		return false
	}
	for idx, v := range m.GetHttpHeaders() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetHttpHeaders()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetHttpHeaders()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *ExternalWorkloadSpec_Probe_HTTPHeader) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExternalWorkloadSpec_Probe_HTTPHeader)
	if !ok {
		that2, ok := that.(ExternalWorkloadSpec_Probe_HTTPHeader)
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

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if strings.Compare(m.GetValue(), target.GetValue()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *ExternalWorkloadSpec_Probe_TCPSocketConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExternalWorkloadSpec_Probe_TCPSocketConfig)
	if !ok {
		that2, ok := that.(ExternalWorkloadSpec_Probe_TCPSocketConfig)
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

	if strings.Compare(m.GetHost(), target.GetHost()) != 0 {
		return false
	}

	if m.GetPort() != target.GetPort() {
		return false
	}

	return true
}

// Equal function
func (m *ExternalWorkloadSpec_Probe_ExecConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExternalWorkloadSpec_Probe_ExecConfig)
	if !ok {
		that2, ok := that.(ExternalWorkloadSpec_Probe_ExecConfig)
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

	if len(m.GetCommand()) != len(target.GetCommand()) {
		return false
	}
	for idx, v := range m.GetCommand() {

		if strings.Compare(v, target.GetCommand()[idx]) != 0 {
			return false
		}

	}

	return true
}
