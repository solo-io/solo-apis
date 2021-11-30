// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/ssl.proto

package v1

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
func (m *SslConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SslConfig)
	if !ok {
		that2, ok := that.(SslConfig)
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

	if len(m.GetSniDomains()) != len(target.GetSniDomains()) {
		return false
	}
	for idx, v := range m.GetSniDomains() {

		if strings.Compare(v, target.GetSniDomains()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetVerifySubjectAltName()) != len(target.GetVerifySubjectAltName()) {
		return false
	}
	for idx, v := range m.GetVerifySubjectAltName() {

		if strings.Compare(v, target.GetVerifySubjectAltName()[idx]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetParameters()).(equality.Equalizer); ok {
		if !h.Equal(target.GetParameters()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetParameters(), target.GetParameters()) {
			return false
		}
	}

	if len(m.GetAlpnProtocols()) != len(target.GetAlpnProtocols()) {
		return false
	}
	for idx, v := range m.GetAlpnProtocols() {

		if strings.Compare(v, target.GetAlpnProtocols()[idx]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetOneWayTls()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOneWayTls()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOneWayTls(), target.GetOneWayTls()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDisableTlsSessionResumption()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDisableTlsSessionResumption()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDisableTlsSessionResumption(), target.GetDisableTlsSessionResumption()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetTransportSocketConnectTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTransportSocketConnectTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTransportSocketConnectTimeout(), target.GetTransportSocketConnectTimeout()) {
			return false
		}
	}

	switch m.SslSecrets.(type) {

	case *SslConfig_SecretRef:
		if _, ok := target.SslSecrets.(*SslConfig_SecretRef); !ok {
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

	case *SslConfig_SslFiles:
		if _, ok := target.SslSecrets.(*SslConfig_SslFiles); !ok {
			return false
		}

		if h, ok := interface{}(m.GetSslFiles()).(equality.Equalizer); ok {
			if !h.Equal(target.GetSslFiles()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetSslFiles(), target.GetSslFiles()) {
				return false
			}
		}

	case *SslConfig_Sds:
		if _, ok := target.SslSecrets.(*SslConfig_Sds); !ok {
			return false
		}

		if h, ok := interface{}(m.GetSds()).(equality.Equalizer); ok {
			if !h.Equal(target.GetSds()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetSds(), target.GetSds()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.SslSecrets != target.SslSecrets {
			return false
		}
	}

	return true
}

// Equal function
func (m *SSLFiles) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SSLFiles)
	if !ok {
		that2, ok := that.(SSLFiles)
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

	if strings.Compare(m.GetTlsCert(), target.GetTlsCert()) != 0 {
		return false
	}

	if strings.Compare(m.GetTlsKey(), target.GetTlsKey()) != 0 {
		return false
	}

	if strings.Compare(m.GetRootCa(), target.GetRootCa()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *UpstreamSslConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpstreamSslConfig)
	if !ok {
		that2, ok := that.(UpstreamSslConfig)
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

	if len(m.GetVerifySubjectAltName()) != len(target.GetVerifySubjectAltName()) {
		return false
	}
	for idx, v := range m.GetVerifySubjectAltName() {

		if strings.Compare(v, target.GetVerifySubjectAltName()[idx]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetParameters()).(equality.Equalizer); ok {
		if !h.Equal(target.GetParameters()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetParameters(), target.GetParameters()) {
			return false
		}
	}

	if len(m.GetAlpnProtocols()) != len(target.GetAlpnProtocols()) {
		return false
	}
	for idx, v := range m.GetAlpnProtocols() {

		if strings.Compare(v, target.GetAlpnProtocols()[idx]) != 0 {
			return false
		}

	}

	switch m.SslSecrets.(type) {

	case *UpstreamSslConfig_SecretRef:
		if _, ok := target.SslSecrets.(*UpstreamSslConfig_SecretRef); !ok {
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

	case *UpstreamSslConfig_SslFiles:
		if _, ok := target.SslSecrets.(*UpstreamSslConfig_SslFiles); !ok {
			return false
		}

		if h, ok := interface{}(m.GetSslFiles()).(equality.Equalizer); ok {
			if !h.Equal(target.GetSslFiles()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetSslFiles(), target.GetSslFiles()) {
				return false
			}
		}

	case *UpstreamSslConfig_Sds:
		if _, ok := target.SslSecrets.(*UpstreamSslConfig_Sds); !ok {
			return false
		}

		if h, ok := interface{}(m.GetSds()).(equality.Equalizer); ok {
			if !h.Equal(target.GetSds()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetSds(), target.GetSds()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.SslSecrets != target.SslSecrets {
			return false
		}
	}

	return true
}

// Equal function
func (m *SDSConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SDSConfig)
	if !ok {
		that2, ok := that.(SDSConfig)
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

	if strings.Compare(m.GetTargetUri(), target.GetTargetUri()) != 0 {
		return false
	}

	if strings.Compare(m.GetCertificatesSecretName(), target.GetCertificatesSecretName()) != 0 {
		return false
	}

	if strings.Compare(m.GetValidationContextName(), target.GetValidationContextName()) != 0 {
		return false
	}

	switch m.SdsBuilder.(type) {

	case *SDSConfig_CallCredentials:
		if _, ok := target.SdsBuilder.(*SDSConfig_CallCredentials); !ok {
			return false
		}

		if h, ok := interface{}(m.GetCallCredentials()).(equality.Equalizer); ok {
			if !h.Equal(target.GetCallCredentials()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetCallCredentials(), target.GetCallCredentials()) {
				return false
			}
		}

	case *SDSConfig_ClusterName:
		if _, ok := target.SdsBuilder.(*SDSConfig_ClusterName); !ok {
			return false
		}

		if strings.Compare(m.GetClusterName(), target.GetClusterName()) != 0 {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.SdsBuilder != target.SdsBuilder {
			return false
		}
	}

	return true
}

// Equal function
func (m *CallCredentials) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CallCredentials)
	if !ok {
		that2, ok := that.(CallCredentials)
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

	if h, ok := interface{}(m.GetFileCredentialSource()).(equality.Equalizer); ok {
		if !h.Equal(target.GetFileCredentialSource()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetFileCredentialSource(), target.GetFileCredentialSource()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *SslParameters) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SslParameters)
	if !ok {
		that2, ok := that.(SslParameters)
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

	if m.GetMinimumProtocolVersion() != target.GetMinimumProtocolVersion() {
		return false
	}

	if m.GetMaximumProtocolVersion() != target.GetMaximumProtocolVersion() {
		return false
	}

	if len(m.GetCipherSuites()) != len(target.GetCipherSuites()) {
		return false
	}
	for idx, v := range m.GetCipherSuites() {

		if strings.Compare(v, target.GetCipherSuites()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetEcdhCurves()) != len(target.GetEcdhCurves()) {
		return false
	}
	for idx, v := range m.GetEcdhCurves() {

		if strings.Compare(v, target.GetEcdhCurves()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *CallCredentials_FileCredentialSource) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CallCredentials_FileCredentialSource)
	if !ok {
		that2, ok := that.(CallCredentials_FileCredentialSource)
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

	if strings.Compare(m.GetTokenFileName(), target.GetTokenFileName()) != 0 {
		return false
	}

	if strings.Compare(m.GetHeader(), target.GetHeader()) != 0 {
		return false
	}

	return true
}
