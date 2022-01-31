// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/config/core/v3/grpc_service.proto

package v3

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
func (m *GrpcService) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GrpcService)
	if !ok {
		that2, ok := that.(GrpcService)
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

	if h, ok := interface{}(m.GetTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTimeout(), target.GetTimeout()) {
			return false
		}
	}

	if len(m.GetInitialMetadata()) != len(target.GetInitialMetadata()) {
		return false
	}
	for idx, v := range m.GetInitialMetadata() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetInitialMetadata()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetInitialMetadata()[idx]) {
				return false
			}
		}

	}

	switch m.TargetSpecifier.(type) {

	case *GrpcService_EnvoyGrpc_:
		if _, ok := target.TargetSpecifier.(*GrpcService_EnvoyGrpc_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetEnvoyGrpc()).(equality.Equalizer); ok {
			if !h.Equal(target.GetEnvoyGrpc()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetEnvoyGrpc(), target.GetEnvoyGrpc()) {
				return false
			}
		}

	case *GrpcService_GoogleGrpc_:
		if _, ok := target.TargetSpecifier.(*GrpcService_GoogleGrpc_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGoogleGrpc()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGoogleGrpc()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGoogleGrpc(), target.GetGoogleGrpc()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.TargetSpecifier != target.TargetSpecifier {
			return false
		}
	}

	return true
}

// Equal function
func (m *GrpcService_EnvoyGrpc) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GrpcService_EnvoyGrpc)
	if !ok {
		that2, ok := that.(GrpcService_EnvoyGrpc)
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

	if strings.Compare(m.GetClusterName(), target.GetClusterName()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *GrpcService_GoogleGrpc) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GrpcService_GoogleGrpc)
	if !ok {
		that2, ok := that.(GrpcService_GoogleGrpc)
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

	if h, ok := interface{}(m.GetChannelCredentials()).(equality.Equalizer); ok {
		if !h.Equal(target.GetChannelCredentials()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetChannelCredentials(), target.GetChannelCredentials()) {
			return false
		}
	}

	if len(m.GetCallCredentials()) != len(target.GetCallCredentials()) {
		return false
	}
	for idx, v := range m.GetCallCredentials() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetCallCredentials()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetCallCredentials()[idx]) {
				return false
			}
		}

	}

	if strings.Compare(m.GetStatPrefix(), target.GetStatPrefix()) != 0 {
		return false
	}

	if strings.Compare(m.GetCredentialsFactoryName(), target.GetCredentialsFactoryName()) != 0 {
		return false
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

	if h, ok := interface{}(m.GetPerStreamBufferLimitBytes()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPerStreamBufferLimitBytes()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPerStreamBufferLimitBytes(), target.GetPerStreamBufferLimitBytes()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetChannelArgs()).(equality.Equalizer); ok {
		if !h.Equal(target.GetChannelArgs()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetChannelArgs(), target.GetChannelArgs()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GrpcService_GoogleGrpc_SslCredentials) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GrpcService_GoogleGrpc_SslCredentials)
	if !ok {
		that2, ok := that.(GrpcService_GoogleGrpc_SslCredentials)
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

	if h, ok := interface{}(m.GetRootCerts()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRootCerts()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRootCerts(), target.GetRootCerts()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPrivateKey()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPrivateKey()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPrivateKey(), target.GetPrivateKey()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetCertChain()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCertChain()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCertChain(), target.GetCertChain()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GrpcService_GoogleGrpc_GoogleLocalCredentials) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GrpcService_GoogleGrpc_GoogleLocalCredentials)
	if !ok {
		that2, ok := that.(GrpcService_GoogleGrpc_GoogleLocalCredentials)
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
func (m *GrpcService_GoogleGrpc_ChannelCredentials) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GrpcService_GoogleGrpc_ChannelCredentials)
	if !ok {
		that2, ok := that.(GrpcService_GoogleGrpc_ChannelCredentials)
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

	switch m.CredentialSpecifier.(type) {

	case *GrpcService_GoogleGrpc_ChannelCredentials_SslCredentials:
		if _, ok := target.CredentialSpecifier.(*GrpcService_GoogleGrpc_ChannelCredentials_SslCredentials); !ok {
			return false
		}

		if h, ok := interface{}(m.GetSslCredentials()).(equality.Equalizer); ok {
			if !h.Equal(target.GetSslCredentials()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetSslCredentials(), target.GetSslCredentials()) {
				return false
			}
		}

	case *GrpcService_GoogleGrpc_ChannelCredentials_GoogleDefault:
		if _, ok := target.CredentialSpecifier.(*GrpcService_GoogleGrpc_ChannelCredentials_GoogleDefault); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGoogleDefault()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGoogleDefault()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGoogleDefault(), target.GetGoogleDefault()) {
				return false
			}
		}

	case *GrpcService_GoogleGrpc_ChannelCredentials_LocalCredentials:
		if _, ok := target.CredentialSpecifier.(*GrpcService_GoogleGrpc_ChannelCredentials_LocalCredentials); !ok {
			return false
		}

		if h, ok := interface{}(m.GetLocalCredentials()).(equality.Equalizer); ok {
			if !h.Equal(target.GetLocalCredentials()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetLocalCredentials(), target.GetLocalCredentials()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.CredentialSpecifier != target.CredentialSpecifier {
			return false
		}
	}

	return true
}

// Equal function
func (m *GrpcService_GoogleGrpc_CallCredentials) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GrpcService_GoogleGrpc_CallCredentials)
	if !ok {
		that2, ok := that.(GrpcService_GoogleGrpc_CallCredentials)
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

	switch m.CredentialSpecifier.(type) {

	case *GrpcService_GoogleGrpc_CallCredentials_AccessToken:
		if _, ok := target.CredentialSpecifier.(*GrpcService_GoogleGrpc_CallCredentials_AccessToken); !ok {
			return false
		}

		if strings.Compare(m.GetAccessToken(), target.GetAccessToken()) != 0 {
			return false
		}

	case *GrpcService_GoogleGrpc_CallCredentials_GoogleComputeEngine:
		if _, ok := target.CredentialSpecifier.(*GrpcService_GoogleGrpc_CallCredentials_GoogleComputeEngine); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGoogleComputeEngine()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGoogleComputeEngine()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGoogleComputeEngine(), target.GetGoogleComputeEngine()) {
				return false
			}
		}

	case *GrpcService_GoogleGrpc_CallCredentials_GoogleRefreshToken:
		if _, ok := target.CredentialSpecifier.(*GrpcService_GoogleGrpc_CallCredentials_GoogleRefreshToken); !ok {
			return false
		}

		if strings.Compare(m.GetGoogleRefreshToken(), target.GetGoogleRefreshToken()) != 0 {
			return false
		}

	case *GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJwtAccess:
		if _, ok := target.CredentialSpecifier.(*GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJwtAccess); !ok {
			return false
		}

		if h, ok := interface{}(m.GetServiceAccountJwtAccess()).(equality.Equalizer); ok {
			if !h.Equal(target.GetServiceAccountJwtAccess()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetServiceAccountJwtAccess(), target.GetServiceAccountJwtAccess()) {
				return false
			}
		}

	case *GrpcService_GoogleGrpc_CallCredentials_GoogleIam:
		if _, ok := target.CredentialSpecifier.(*GrpcService_GoogleGrpc_CallCredentials_GoogleIam); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGoogleIam()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGoogleIam()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGoogleIam(), target.GetGoogleIam()) {
				return false
			}
		}

	case *GrpcService_GoogleGrpc_CallCredentials_FromPlugin:
		if _, ok := target.CredentialSpecifier.(*GrpcService_GoogleGrpc_CallCredentials_FromPlugin); !ok {
			return false
		}

		if h, ok := interface{}(m.GetFromPlugin()).(equality.Equalizer); ok {
			if !h.Equal(target.GetFromPlugin()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetFromPlugin(), target.GetFromPlugin()) {
				return false
			}
		}

	case *GrpcService_GoogleGrpc_CallCredentials_StsService_:
		if _, ok := target.CredentialSpecifier.(*GrpcService_GoogleGrpc_CallCredentials_StsService_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetStsService()).(equality.Equalizer); ok {
			if !h.Equal(target.GetStsService()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetStsService(), target.GetStsService()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.CredentialSpecifier != target.CredentialSpecifier {
			return false
		}
	}

	return true
}

// Equal function
func (m *GrpcService_GoogleGrpc_ChannelArgs) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GrpcService_GoogleGrpc_ChannelArgs)
	if !ok {
		that2, ok := that.(GrpcService_GoogleGrpc_ChannelArgs)
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

	if len(m.GetArgs()) != len(target.GetArgs()) {
		return false
	}
	for k, v := range m.GetArgs() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetArgs()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetArgs()[k]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentials) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentials)
	if !ok {
		that2, ok := that.(GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentials)
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

	if strings.Compare(m.GetJsonKey(), target.GetJsonKey()) != 0 {
		return false
	}

	if m.GetTokenLifetimeSeconds() != target.GetTokenLifetimeSeconds() {
		return false
	}

	return true
}

// Equal function
func (m *GrpcService_GoogleGrpc_CallCredentials_GoogleIAMCredentials) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GrpcService_GoogleGrpc_CallCredentials_GoogleIAMCredentials)
	if !ok {
		that2, ok := that.(GrpcService_GoogleGrpc_CallCredentials_GoogleIAMCredentials)
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

	if strings.Compare(m.GetAuthorizationToken(), target.GetAuthorizationToken()) != 0 {
		return false
	}

	if strings.Compare(m.GetAuthoritySelector(), target.GetAuthoritySelector()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin)
	if !ok {
		that2, ok := that.(GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin)
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

	switch m.ConfigType.(type) {

	case *GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin_TypedConfig:
		if _, ok := target.ConfigType.(*GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin_TypedConfig); !ok {
			return false
		}

		if h, ok := interface{}(m.GetTypedConfig()).(equality.Equalizer); ok {
			if !h.Equal(target.GetTypedConfig()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetTypedConfig(), target.GetTypedConfig()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.ConfigType != target.ConfigType {
			return false
		}
	}

	return true
}

// Equal function
func (m *GrpcService_GoogleGrpc_CallCredentials_StsService) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GrpcService_GoogleGrpc_CallCredentials_StsService)
	if !ok {
		that2, ok := that.(GrpcService_GoogleGrpc_CallCredentials_StsService)
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

	if strings.Compare(m.GetTokenExchangeServiceUri(), target.GetTokenExchangeServiceUri()) != 0 {
		return false
	}

	if strings.Compare(m.GetResource(), target.GetResource()) != 0 {
		return false
	}

	if strings.Compare(m.GetAudience(), target.GetAudience()) != 0 {
		return false
	}

	if strings.Compare(m.GetScope(), target.GetScope()) != 0 {
		return false
	}

	if strings.Compare(m.GetRequestedTokenType(), target.GetRequestedTokenType()) != 0 {
		return false
	}

	if strings.Compare(m.GetSubjectTokenPath(), target.GetSubjectTokenPath()) != 0 {
		return false
	}

	if strings.Compare(m.GetSubjectTokenType(), target.GetSubjectTokenType()) != 0 {
		return false
	}

	if strings.Compare(m.GetActorTokenPath(), target.GetActorTokenPath()) != 0 {
		return false
	}

	if strings.Compare(m.GetActorTokenType(), target.GetActorTokenType()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *GrpcService_GoogleGrpc_ChannelArgs_Value) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GrpcService_GoogleGrpc_ChannelArgs_Value)
	if !ok {
		that2, ok := that.(GrpcService_GoogleGrpc_ChannelArgs_Value)
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

	switch m.ValueSpecifier.(type) {

	case *GrpcService_GoogleGrpc_ChannelArgs_Value_StringValue:
		if _, ok := target.ValueSpecifier.(*GrpcService_GoogleGrpc_ChannelArgs_Value_StringValue); !ok {
			return false
		}

		if strings.Compare(m.GetStringValue(), target.GetStringValue()) != 0 {
			return false
		}

	case *GrpcService_GoogleGrpc_ChannelArgs_Value_IntValue:
		if _, ok := target.ValueSpecifier.(*GrpcService_GoogleGrpc_ChannelArgs_Value_IntValue); !ok {
			return false
		}

		if m.GetIntValue() != target.GetIntValue() {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.ValueSpecifier != target.ValueSpecifier {
			return false
		}
	}

	return true
}
