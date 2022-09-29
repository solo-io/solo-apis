// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/admin/v2/istio_lifecycle_manager.proto

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
func (m *IstioLifecycleManagerSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*IstioLifecycleManagerSpec)
	if !ok {
		that2, ok := that.(IstioLifecycleManagerSpec)
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

	if strings.Compare(m.GetRevision(), target.GetRevision()) != 0 {
		return false
	}

	if strings.Compare(m.GetDefaultRevision(), target.GetDefaultRevision()) != 0 {
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

	if len(m.GetUninstallRevisions()) != len(target.GetUninstallRevisions()) {
		return false
	}
	for idx, v := range m.GetUninstallRevisions() {

		if strings.Compare(v, target.GetUninstallRevisions()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *IstioLifecycleManagerStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*IstioLifecycleManagerStatus)
	if !ok {
		that2, ok := that.(IstioLifecycleManagerStatus)
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
func (m *IstioInstallation) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*IstioInstallation)
	if !ok {
		that2, ok := that.(IstioInstallation)
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

	if h, ok := interface{}(m.GetIstioOperatorSpec()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIstioOperatorSpec()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIstioOperatorSpec(), target.GetIstioOperatorSpec()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *IstioLifecycleManagerSpec_Cluster) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*IstioLifecycleManagerSpec_Cluster)
	if !ok {
		that2, ok := that.(IstioLifecycleManagerSpec_Cluster)
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

	if strings.Compare(m.GetTrustDomain(), target.GetTrustDomain()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *IstioLifecycleManagerStatus_ClusterStatuses) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*IstioLifecycleManagerStatus_ClusterStatuses)
	if !ok {
		that2, ok := that.(IstioLifecycleManagerStatus_ClusterStatuses)
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
func (m *IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus)
	if !ok {
		that2, ok := that.(IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus)
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

	if strings.Compare(m.GetRevision(), target.GetRevision()) != 0 {
		return false
	}

	if len(m.GetObservedInstallations()) != len(target.GetObservedInstallations()) {
		return false
	}
	for idx, v := range m.GetObservedInstallations() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetObservedInstallations()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetObservedInstallations()[idx]) {
				return false
			}
		}

	}

	return true
}
