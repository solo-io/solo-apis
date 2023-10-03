// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/admin/v2/gateway_lifecycle_manager.proto

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
func (m *GatewayLifecycleManagerSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GatewayLifecycleManagerSpec)
	if !ok {
		that2, ok := that.(GatewayLifecycleManagerSpec)
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

	if len(m.GetInstallations()) != len(target.GetInstallations()) {
		return false
	}
	for idx, v := range m.GetInstallations() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetInstallations()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetInstallations()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *GatewayClusterSelector) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GatewayClusterSelector)
	if !ok {
		that2, ok := that.(GatewayClusterSelector)
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

	if m.GetActiveGateway() != target.GetActiveGateway() {
		return false
	}

	if strings.Compare(m.GetTrustDomain(), target.GetTrustDomain()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *GatewayInstallation) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GatewayInstallation)
	if !ok {
		that2, ok := that.(GatewayInstallation)
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

	if strings.Compare(m.GetControlPlaneRevision(), target.GetControlPlaneRevision()) != 0 {
		return false
	}

	if strings.Compare(m.GetGatewayRevision(), target.GetGatewayRevision()) != 0 {
		return false
	}

	if len(m.GetClusters()) != len(target.GetClusters()) {
		return false
	}
	for idx, v := range m.GetClusters() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetClusters()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetClusters()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetIstioOperatorSpec()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIstioOperatorSpec()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIstioOperatorSpec(), target.GetIstioOperatorSpec()) {
			return false
		}
	}

	if m.GetSkipUpgradeValidation() != target.GetSkipUpgradeValidation() {
		return false
	}

	return true
}

// Equal function
func (m *GatewayLifecycleManagerStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GatewayLifecycleManagerStatus)
	if !ok {
		that2, ok := that.(GatewayLifecycleManagerStatus)
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
func (m *GatewayLifecycleManagerNewStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GatewayLifecycleManagerNewStatus)
	if !ok {
		that2, ok := that.(GatewayLifecycleManagerNewStatus)
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

	return true
}

// Equal function
func (m *GatewayLifecycleManagerReport) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GatewayLifecycleManagerReport)
	if !ok {
		that2, ok := that.(GatewayLifecycleManagerReport)
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

	return true
}

// Equal function
func (m *GatewayLifecycleManagerStatus_ClusterStatuses) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GatewayLifecycleManagerStatus_ClusterStatuses)
	if !ok {
		that2, ok := that.(GatewayLifecycleManagerStatus_ClusterStatuses)
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

	if len(m.GetInstallations()) != len(target.GetInstallations()) {
		return false
	}
	for k, v := range m.GetInstallations() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetInstallations()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetInstallations()[k]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus)
	if !ok {
		that2, ok := that.(GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus)
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

	if strings.Compare(m.GetObservedRevision(), target.GetObservedRevision()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetObservedOperator()).(equality.Equalizer); ok {
		if !h.Equal(target.GetObservedOperator()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetObservedOperator(), target.GetObservedOperator()) {
			return false
		}
	}

	return true
}
