// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/networking/v2/external_service.proto

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
func (m *ExternalServiceSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExternalServiceSpec)
	if !ok {
		that2, ok := that.(ExternalServiceSpec)
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

	if len(m.GetHosts()) != len(target.GetHosts()) {
		return false
	}
	for idx, v := range m.GetHosts() {

		if strings.Compare(v, target.GetHosts()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetAddresses()) != len(target.GetAddresses()) {
		return false
	}
	for idx, v := range m.GetAddresses() {

		if strings.Compare(v, target.GetAddresses()[idx]) != 0 {
			return false
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

	if len(m.GetSelector()) != len(target.GetSelector()) {
		return false
	}
	for k, v := range m.GetSelector() {

		if strings.Compare(v, target.GetSelector()[k]) != 0 {
			return false
		}

	}

	if len(m.GetSubjectAltNames()) != len(target.GetSubjectAltNames()) {
		return false
	}
	for idx, v := range m.GetSubjectAltNames() {

		if strings.Compare(v, target.GetSubjectAltNames()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *ExternalServiceStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExternalServiceStatus)
	if !ok {
		that2, ok := that.(ExternalServiceStatus)
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

	if len(m.GetSelectedExternalEndpoints()) != len(target.GetSelectedExternalEndpoints()) {
		return false
	}
	for idx, v := range m.GetSelectedExternalEndpoints() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSelectedExternalEndpoints()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSelectedExternalEndpoints()[idx]) {
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
func (m *ExternalServiceNewStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExternalServiceNewStatus)
	if !ok {
		that2, ok := that.(ExternalServiceNewStatus)
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

	if len(m.GetAppliedDestinationPolicies()) != len(target.GetAppliedDestinationPolicies()) {
		return false
	}
	for k, v := range m.GetAppliedDestinationPolicies() {

		if v != target.GetAppliedDestinationPolicies()[k] {
			return false
		}

	}

	if m.GetSelectedExternalEndpoints() != target.GetSelectedExternalEndpoints() {
		return false
	}

	if strings.Compare(m.GetOwnerWorkspace(), target.GetOwnerWorkspace()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *ExternalServiceReport) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExternalServiceReport)
	if !ok {
		that2, ok := that.(ExternalServiceReport)
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

	if len(m.GetSelectedExternalEndpoints()) != len(target.GetSelectedExternalEndpoints()) {
		return false
	}
	for idx, v := range m.GetSelectedExternalEndpoints() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSelectedExternalEndpoints()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSelectedExternalEndpoints()[idx]) {
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
func (m *ExternalServiceSpec_Port) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExternalServiceSpec_Port)
	if !ok {
		that2, ok := that.(ExternalServiceSpec_Port)
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

	if m.GetNumber() != target.GetNumber() {
		return false
	}

	if h, ok := interface{}(m.GetTargetPort()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTargetPort()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTargetPort(), target.GetTargetPort()) {
			return false
		}
	}

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if strings.Compare(m.GetProtocol(), target.GetProtocol()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetClientsideTls()).(equality.Equalizer); ok {
		if !h.Equal(target.GetClientsideTls()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetClientsideTls(), target.GetClientsideTls()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ExternalServiceSpec_Port_TlsConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExternalServiceSpec_Port_TlsConfig)
	if !ok {
		that2, ok := that.(ExternalServiceSpec_Port_TlsConfig)
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

	if strings.Compare(m.GetSni(), target.GetSni()) != 0 {
		return false
	}

	if m.GetMode() != target.GetMode() {
		return false
	}

	if strings.Compare(m.GetClientCertificate(), target.GetClientCertificate()) != 0 {
		return false
	}

	if strings.Compare(m.GetPrivateKey(), target.GetPrivateKey()) != 0 {
		return false
	}

	if strings.Compare(m.GetCaCertificates(), target.GetCaCertificates()) != 0 {
		return false
	}

	return true
}
