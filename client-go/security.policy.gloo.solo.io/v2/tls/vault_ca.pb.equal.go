// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/policy/v2/security/tls/vault_ca.proto

package tls

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
func (m *VaultCA) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VaultCA)
	if !ok {
		that2, ok := that.(VaultCA)
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

	if strings.Compare(m.GetCaPath(), target.GetCaPath()) != 0 {
		return false
	}

	if strings.Compare(m.GetCsrPath(), target.GetCsrPath()) != 0 {
		return false
	}

	if strings.Compare(m.GetServer(), target.GetServer()) != 0 {
		return false
	}

	if strings.Compare(m.GetNamespace(), target.GetNamespace()) != 0 {
		return false
	}

	switch m.CaSource.(type) {

	case *VaultCA_CaBundle:
		if _, ok := target.CaSource.(*VaultCA_CaBundle); !ok {
			return false
		}

		if bytes.Compare(m.GetCaBundle(), target.GetCaBundle()) != 0 {
			return false
		}

	case *VaultCA_CaSecretRef:
		if _, ok := target.CaSource.(*VaultCA_CaSecretRef); !ok {
			return false
		}

		if h, ok := interface{}(m.GetCaSecretRef()).(equality.Equalizer); ok {
			if !h.Equal(target.GetCaSecretRef()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetCaSecretRef(), target.GetCaSecretRef()) {
				return false
			}
		}

	case *VaultCA_CaLocalPath:
		if _, ok := target.CaSource.(*VaultCA_CaLocalPath); !ok {
			return false
		}

		if strings.Compare(m.GetCaLocalPath(), target.GetCaLocalPath()) != 0 {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.CaSource != target.CaSource {
			return false
		}
	}

	switch m.AuthType.(type) {

	case *VaultCA_TokenSecretRef:
		if _, ok := target.AuthType.(*VaultCA_TokenSecretRef); !ok {
			return false
		}

		if h, ok := interface{}(m.GetTokenSecretRef()).(equality.Equalizer); ok {
			if !h.Equal(target.GetTokenSecretRef()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetTokenSecretRef(), target.GetTokenSecretRef()) {
				return false
			}
		}

	case *VaultCA_KubernetesAuth:
		if _, ok := target.AuthType.(*VaultCA_KubernetesAuth); !ok {
			return false
		}

		if h, ok := interface{}(m.GetKubernetesAuth()).(equality.Equalizer); ok {
			if !h.Equal(target.GetKubernetesAuth()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetKubernetesAuth(), target.GetKubernetesAuth()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.AuthType != target.AuthType {
			return false
		}
	}

	return true
}

// Equal function
func (m *VaultKubernetesAuth) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VaultKubernetesAuth)
	if !ok {
		that2, ok := that.(VaultKubernetesAuth)
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

	if strings.Compare(m.GetMountPath(), target.GetMountPath()) != 0 {
		return false
	}

	if strings.Compare(m.GetRole(), target.GetRole()) != 0 {
		return false
	}

	if strings.Compare(m.GetSecretTokenKey(), target.GetSecretTokenKey()) != 0 {
		return false
	}

	switch m.ServiceAccountLocation.(type) {

	case *VaultKubernetesAuth_ServiceAccountRef:
		if _, ok := target.ServiceAccountLocation.(*VaultKubernetesAuth_ServiceAccountRef); !ok {
			return false
		}

		if h, ok := interface{}(m.GetServiceAccountRef()).(equality.Equalizer); ok {
			if !h.Equal(target.GetServiceAccountRef()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetServiceAccountRef(), target.GetServiceAccountRef()) {
				return false
			}
		}

	case *VaultKubernetesAuth_MountedSaPath:
		if _, ok := target.ServiceAccountLocation.(*VaultKubernetesAuth_MountedSaPath); !ok {
			return false
		}

		if strings.Compare(m.GetMountedSaPath(), target.GetMountedSaPath()) != 0 {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.ServiceAccountLocation != target.ServiceAccountLocation {
			return false
		}
	}

	return true
}
