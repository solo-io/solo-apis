// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/policy/v2/security/jwt_policy.proto

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
func (m *JWTPolicySpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*JWTPolicySpec)
	if !ok {
		that2, ok := that.(JWTPolicySpec)
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

	if len(m.GetApplyToRoutes()) != len(target.GetApplyToRoutes()) {
		return false
	}
	for idx, v := range m.GetApplyToRoutes() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetApplyToRoutes()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetApplyToRoutes()[idx]) {
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
func (m *JWTPolicyStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*JWTPolicyStatus)
	if !ok {
		that2, ok := that.(JWTPolicyStatus)
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

	if len(m.GetSelectedRoutes()) != len(target.GetSelectedRoutes()) {
		return false
	}
	for idx, v := range m.GetSelectedRoutes() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSelectedRoutes()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSelectedRoutes()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *JWTPolicySpec_Config) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*JWTPolicySpec_Config)
	if !ok {
		that2, ok := that.(JWTPolicySpec_Config)
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

	if len(m.GetProviders()) != len(target.GetProviders()) {
		return false
	}
	for k, v := range m.GetProviders() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetProviders()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetProviders()[k]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetPhase()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPhase()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPhase(), target.GetPhase()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *JWTPolicySpec_Config_Provider) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*JWTPolicySpec_Config_Provider)
	if !ok {
		that2, ok := that.(JWTPolicySpec_Config_Provider)
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

	if strings.Compare(m.GetIssuer(), target.GetIssuer()) != 0 {
		return false
	}

	if len(m.GetAudiences()) != len(target.GetAudiences()) {
		return false
	}
	for idx, v := range m.GetAudiences() {

		if strings.Compare(v, target.GetAudiences()[idx]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetTokenSource()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTokenSource()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTokenSource(), target.GetTokenSource()) {
			return false
		}
	}

	if len(m.GetClaimsToHeaders()) != len(target.GetClaimsToHeaders()) {
		return false
	}
	for idx, v := range m.GetClaimsToHeaders() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetClaimsToHeaders()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetClaimsToHeaders()[idx]) {
				return false
			}
		}

	}

	if m.GetKeepToken() != target.GetKeepToken() {
		return false
	}

	switch m.JwksSource.(type) {

	case *JWTPolicySpec_Config_Provider_Local:
		if _, ok := target.JwksSource.(*JWTPolicySpec_Config_Provider_Local); !ok {
			return false
		}

		if h, ok := interface{}(m.GetLocal()).(equality.Equalizer); ok {
			if !h.Equal(target.GetLocal()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetLocal(), target.GetLocal()) {
				return false
			}
		}

	case *JWTPolicySpec_Config_Provider_Remote:
		if _, ok := target.JwksSource.(*JWTPolicySpec_Config_Provider_Remote); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRemote()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRemote()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRemote(), target.GetRemote()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.JwksSource != target.JwksSource {
			return false
		}
	}

	return true
}

// Equal function
func (m *JWTPolicySpec_Config_Provider_TokenSource) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*JWTPolicySpec_Config_Provider_TokenSource)
	if !ok {
		that2, ok := that.(JWTPolicySpec_Config_Provider_TokenSource)
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

	if len(m.GetHeaders()) != len(target.GetHeaders()) {
		return false
	}
	for idx, v := range m.GetHeaders() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetHeaders()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetHeaders()[idx]) {
				return false
			}
		}

	}

	if len(m.GetQueryParams()) != len(target.GetQueryParams()) {
		return false
	}
	for idx, v := range m.GetQueryParams() {

		if strings.Compare(v, target.GetQueryParams()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *JWTPolicySpec_Config_Provider_LocalJWKS) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*JWTPolicySpec_Config_Provider_LocalJWKS)
	if !ok {
		that2, ok := that.(JWTPolicySpec_Config_Provider_LocalJWKS)
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

	switch m.Specifier.(type) {

	case *JWTPolicySpec_Config_Provider_LocalJWKS_SecretRef_:
		if _, ok := target.Specifier.(*JWTPolicySpec_Config_Provider_LocalJWKS_SecretRef_); !ok {
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

	case *JWTPolicySpec_Config_Provider_LocalJWKS_Inline:
		if _, ok := target.Specifier.(*JWTPolicySpec_Config_Provider_LocalJWKS_Inline); !ok {
			return false
		}

		if strings.Compare(m.GetInline(), target.GetInline()) != 0 {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.Specifier != target.Specifier {
			return false
		}
	}

	return true
}

// Equal function
func (m *JWTPolicySpec_Config_Provider_RemoteJWKS) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*JWTPolicySpec_Config_Provider_RemoteJWKS)
	if !ok {
		that2, ok := that.(JWTPolicySpec_Config_Provider_RemoteJWKS)
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

	if strings.Compare(m.GetUrl(), target.GetUrl()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetDestinationRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDestinationRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDestinationRef(), target.GetDestinationRef()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetCacheDuration()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCacheDuration()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCacheDuration(), target.GetCacheDuration()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTimeout(), target.GetTimeout()) {
			return false
		}
	}

	if m.GetEnableAsyncFetch() != target.GetEnableAsyncFetch() {
		return false
	}

	return true
}

// Equal function
func (m *JWTPolicySpec_Config_Provider_ClaimsToHeader) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*JWTPolicySpec_Config_Provider_ClaimsToHeader)
	if !ok {
		that2, ok := that.(JWTPolicySpec_Config_Provider_ClaimsToHeader)
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

	if strings.Compare(m.GetClaim(), target.GetClaim()) != 0 {
		return false
	}

	if strings.Compare(m.GetHeader(), target.GetHeader()) != 0 {
		return false
	}

	if m.GetAppend() != target.GetAppend() {
		return false
	}

	return true
}

// Equal function
func (m *JWTPolicySpec_Config_Provider_TokenSourceFromHeader) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*JWTPolicySpec_Config_Provider_TokenSourceFromHeader)
	if !ok {
		that2, ok := that.(JWTPolicySpec_Config_Provider_TokenSourceFromHeader)
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

	if strings.Compare(m.GetPrefix(), target.GetPrefix()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *JWTPolicySpec_Config_Provider_LocalJWKS_SecretRef) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*JWTPolicySpec_Config_Provider_LocalJWKS_SecretRef)
	if !ok {
		that2, ok := that.(JWTPolicySpec_Config_Provider_LocalJWKS_SecretRef)
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

	if h, ok := interface{}(m.GetObjectRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetObjectRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetObjectRef(), target.GetObjectRef()) {
			return false
		}
	}

	if strings.Compare(m.GetKey(), target.GetKey()) != 0 {
		return false
	}

	return true
}
