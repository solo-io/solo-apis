// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/admin/v2/root_trust_policy.proto

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
func (m *RootTrustPolicySpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RootTrustPolicySpec)
	if !ok {
		that2, ok := that.(RootTrustPolicySpec)
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

	if len(m.GetApplyToMeshes()) != len(target.GetApplyToMeshes()) {
		return false
	}
	for idx, v := range m.GetApplyToMeshes() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetApplyToMeshes()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetApplyToMeshes()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetConfig(), target.GetConfig()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *RootTrustPolicyStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RootTrustPolicyStatus)
	if !ok {
		that2, ok := that.(RootTrustPolicyStatus)
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

	return true
}

// Equal function
func (m *RootTrustPolicySpec_Config) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RootTrustPolicySpec_Config)
	if !ok {
		that2, ok := that.(RootTrustPolicySpec_Config)
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

	if h, ok := interface{}(m.GetIntermediateCertOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIntermediateCertOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIntermediateCertOptions(), target.GetIntermediateCertOptions()) {
			return false
		}
	}

	if m.GetAutoRestartPods() != target.GetAutoRestartPods() {
		return false
	}

	if len(m.GetPassiveCertificateAuthorities()) != len(target.GetPassiveCertificateAuthorities()) {
		return false
	}
	for idx, v := range m.GetPassiveCertificateAuthorities() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPassiveCertificateAuthorities()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPassiveCertificateAuthorities()[idx]) {
				return false
			}
		}

	}

	switch m.CertificateAuthorityType.(type) {

	case *RootTrustPolicySpec_Config_MgmtServerCa:
		if _, ok := target.CertificateAuthorityType.(*RootTrustPolicySpec_Config_MgmtServerCa); !ok {
			return false
		}

		if h, ok := interface{}(m.GetMgmtServerCa()).(equality.Equalizer); ok {
			if !h.Equal(target.GetMgmtServerCa()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetMgmtServerCa(), target.GetMgmtServerCa()) {
				return false
			}
		}

	case *RootTrustPolicySpec_Config_AgentCa:
		if _, ok := target.CertificateAuthorityType.(*RootTrustPolicySpec_Config_AgentCa); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAgentCa()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAgentCa()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAgentCa(), target.GetAgentCa()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.CertificateAuthorityType != target.CertificateAuthorityType {
			return false
		}
	}

	return true
}

// Equal function
func (m *RootTrustPolicySpec_Config_MgmtServerCertificateAuthority) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RootTrustPolicySpec_Config_MgmtServerCertificateAuthority)
	if !ok {
		that2, ok := that.(RootTrustPolicySpec_Config_MgmtServerCertificateAuthority)
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

	switch m.CaSource.(type) {

	case *RootTrustPolicySpec_Config_MgmtServerCertificateAuthority_Generated:
		if _, ok := target.CaSource.(*RootTrustPolicySpec_Config_MgmtServerCertificateAuthority_Generated); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGenerated()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGenerated()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGenerated(), target.GetGenerated()) {
				return false
			}
		}

	case *RootTrustPolicySpec_Config_MgmtServerCertificateAuthority_SecretRef:
		if _, ok := target.CaSource.(*RootTrustPolicySpec_Config_MgmtServerCertificateAuthority_SecretRef); !ok {
			return false
		}

		if h, ok := interface{}(m.GetSecretRef()).(equality.Equalizer); ok {
			if !h.Equal(target.GetSecretRef()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetSecretRef(), target.GetSecretRef()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.CaSource != target.CaSource {
			return false
		}
	}

	return true
}
