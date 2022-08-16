// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/policy/v2/security/jwt_policy.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_duration "github.com/golang/protobuf/ptypes/duration"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"

	github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2 "github.com/solo-io/solo-apis/pkg/api/common.gloo.solo.io/v2"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *JWTPolicySpec) Clone() proto.Message {
	var target *JWTPolicySpec
	if m == nil {
		return target
	}
	target = &JWTPolicySpec{}

	if m.GetApplyToRoutes() != nil {
		target.ApplyToRoutes = make([]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.RouteSelector, len(m.GetApplyToRoutes()))
		for idx, v := range m.GetApplyToRoutes() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ApplyToRoutes[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.RouteSelector)
			} else {
				target.ApplyToRoutes[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.RouteSelector)
			}

		}
	}

	if h, ok := interface{}(m.GetConfig()).(clone.Cloner); ok {
		target.Config = h.Clone().(*JWTPolicySpec_Config)
	} else {
		target.Config = proto.Clone(m.GetConfig()).(*JWTPolicySpec_Config)
	}

	return target
}

// Clone function
func (m *JWTPolicyStatus) Clone() proto.Message {
	var target *JWTPolicyStatus
	if m == nil {
		return target
	}
	target = &JWTPolicyStatus{}

	if h, ok := interface{}(m.GetGlobal()).(clone.Cloner); ok {
		target.Global = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.GenericGlobalStatus)
	} else {
		target.Global = proto.Clone(m.GetGlobal()).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.GenericGlobalStatus)
	}

	if m.GetWorkspaces() != nil {
		target.Workspaces = make(map[string]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.WorkspaceStatus, len(m.GetWorkspaces()))
		for k, v := range m.GetWorkspaces() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Workspaces[k] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.WorkspaceStatus)
			} else {
				target.Workspaces[k] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.WorkspaceStatus)
			}

		}
	}

	if m.GetSelectedRoutes() != nil {
		target.SelectedRoutes = make([]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.RouteReference, len(m.GetSelectedRoutes()))
		for idx, v := range m.GetSelectedRoutes() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SelectedRoutes[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.RouteReference)
			} else {
				target.SelectedRoutes[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.RouteReference)
			}

		}
	}

	return target
}

// Clone function
func (m *JWTPolicySpec_Config) Clone() proto.Message {
	var target *JWTPolicySpec_Config
	if m == nil {
		return target
	}
	target = &JWTPolicySpec_Config{}

	if m.GetProviders() != nil {
		target.Providers = make(map[string]*JWTPolicySpec_Config_Provider, len(m.GetProviders()))
		for k, v := range m.GetProviders() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Providers[k] = h.Clone().(*JWTPolicySpec_Config_Provider)
			} else {
				target.Providers[k] = proto.Clone(v).(*JWTPolicySpec_Config_Provider)
			}

		}
	}

	if h, ok := interface{}(m.GetPhase()).(clone.Cloner); ok {
		target.Phase = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.PrioritizedPhase)
	} else {
		target.Phase = proto.Clone(m.GetPhase()).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.PrioritizedPhase)
	}

	return target
}

// Clone function
func (m *JWTPolicySpec_Config_Provider) Clone() proto.Message {
	var target *JWTPolicySpec_Config_Provider
	if m == nil {
		return target
	}
	target = &JWTPolicySpec_Config_Provider{}

	target.Issuer = m.GetIssuer()

	if m.GetAudiences() != nil {
		target.Audiences = make([]string, len(m.GetAudiences()))
		for idx, v := range m.GetAudiences() {

			target.Audiences[idx] = v

		}
	}

	if h, ok := interface{}(m.GetTokenSource()).(clone.Cloner); ok {
		target.TokenSource = h.Clone().(*JWTPolicySpec_Config_Provider_TokenSource)
	} else {
		target.TokenSource = proto.Clone(m.GetTokenSource()).(*JWTPolicySpec_Config_Provider_TokenSource)
	}

	if m.GetClaimsToHeaders() != nil {
		target.ClaimsToHeaders = make([]*JWTPolicySpec_Config_Provider_ClaimsToHeader, len(m.GetClaimsToHeaders()))
		for idx, v := range m.GetClaimsToHeaders() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ClaimsToHeaders[idx] = h.Clone().(*JWTPolicySpec_Config_Provider_ClaimsToHeader)
			} else {
				target.ClaimsToHeaders[idx] = proto.Clone(v).(*JWTPolicySpec_Config_Provider_ClaimsToHeader)
			}

		}
	}

	target.KeepToken = m.GetKeepToken()

	if h, ok := interface{}(m.GetClockSkewSeconds()).(clone.Cloner); ok {
		target.ClockSkewSeconds = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.ClockSkewSeconds = proto.Clone(m.GetClockSkewSeconds()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	switch m.JwksSource.(type) {

	case *JWTPolicySpec_Config_Provider_Local:

		if h, ok := interface{}(m.GetLocal()).(clone.Cloner); ok {
			target.JwksSource = &JWTPolicySpec_Config_Provider_Local{
				Local: h.Clone().(*JWTPolicySpec_Config_Provider_LocalJWKS),
			}
		} else {
			target.JwksSource = &JWTPolicySpec_Config_Provider_Local{
				Local: proto.Clone(m.GetLocal()).(*JWTPolicySpec_Config_Provider_LocalJWKS),
			}
		}

	case *JWTPolicySpec_Config_Provider_Remote:

		if h, ok := interface{}(m.GetRemote()).(clone.Cloner); ok {
			target.JwksSource = &JWTPolicySpec_Config_Provider_Remote{
				Remote: h.Clone().(*JWTPolicySpec_Config_Provider_RemoteJWKS),
			}
		} else {
			target.JwksSource = &JWTPolicySpec_Config_Provider_Remote{
				Remote: proto.Clone(m.GetRemote()).(*JWTPolicySpec_Config_Provider_RemoteJWKS),
			}
		}

	}

	return target
}

// Clone function
func (m *JWTPolicySpec_Config_Provider_TokenSource) Clone() proto.Message {
	var target *JWTPolicySpec_Config_Provider_TokenSource
	if m == nil {
		return target
	}
	target = &JWTPolicySpec_Config_Provider_TokenSource{}

	if m.GetHeaders() != nil {
		target.Headers = make([]*JWTPolicySpec_Config_Provider_TokenSourceFromHeader, len(m.GetHeaders()))
		for idx, v := range m.GetHeaders() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Headers[idx] = h.Clone().(*JWTPolicySpec_Config_Provider_TokenSourceFromHeader)
			} else {
				target.Headers[idx] = proto.Clone(v).(*JWTPolicySpec_Config_Provider_TokenSourceFromHeader)
			}

		}
	}

	if m.GetQueryParams() != nil {
		target.QueryParams = make([]string, len(m.GetQueryParams()))
		for idx, v := range m.GetQueryParams() {

			target.QueryParams[idx] = v

		}
	}

	return target
}

// Clone function
func (m *JWTPolicySpec_Config_Provider_LocalJWKS) Clone() proto.Message {
	var target *JWTPolicySpec_Config_Provider_LocalJWKS
	if m == nil {
		return target
	}
	target = &JWTPolicySpec_Config_Provider_LocalJWKS{}

	switch m.Specifier.(type) {

	case *JWTPolicySpec_Config_Provider_LocalJWKS_SecretRef_:

		if h, ok := interface{}(m.GetSecretRef()).(clone.Cloner); ok {
			target.Specifier = &JWTPolicySpec_Config_Provider_LocalJWKS_SecretRef_{
				SecretRef: h.Clone().(*JWTPolicySpec_Config_Provider_LocalJWKS_SecretRef),
			}
		} else {
			target.Specifier = &JWTPolicySpec_Config_Provider_LocalJWKS_SecretRef_{
				SecretRef: proto.Clone(m.GetSecretRef()).(*JWTPolicySpec_Config_Provider_LocalJWKS_SecretRef),
			}
		}

	case *JWTPolicySpec_Config_Provider_LocalJWKS_Inline:

		target.Specifier = &JWTPolicySpec_Config_Provider_LocalJWKS_Inline{
			Inline: m.GetInline(),
		}

	}

	return target
}

// Clone function
func (m *JWTPolicySpec_Config_Provider_RemoteJWKS) Clone() proto.Message {
	var target *JWTPolicySpec_Config_Provider_RemoteJWKS
	if m == nil {
		return target
	}
	target = &JWTPolicySpec_Config_Provider_RemoteJWKS{}

	target.Url = m.GetUrl()

	if h, ok := interface{}(m.GetDestinationRef()).(clone.Cloner); ok {
		target.DestinationRef = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.DestinationReference)
	} else {
		target.DestinationRef = proto.Clone(m.GetDestinationRef()).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.DestinationReference)
	}

	if h, ok := interface{}(m.GetCacheDuration()).(clone.Cloner); ok {
		target.CacheDuration = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.CacheDuration = proto.Clone(m.GetCacheDuration()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetTimeout()).(clone.Cloner); ok {
		target.Timeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.Timeout = proto.Clone(m.GetTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	target.EnableAsyncFetch = m.GetEnableAsyncFetch()

	return target
}

// Clone function
func (m *JWTPolicySpec_Config_Provider_ClaimsToHeader) Clone() proto.Message {
	var target *JWTPolicySpec_Config_Provider_ClaimsToHeader
	if m == nil {
		return target
	}
	target = &JWTPolicySpec_Config_Provider_ClaimsToHeader{}

	target.Claim = m.GetClaim()

	target.Header = m.GetHeader()

	target.Append = m.GetAppend()

	return target
}

// Clone function
func (m *JWTPolicySpec_Config_Provider_TokenSourceFromHeader) Clone() proto.Message {
	var target *JWTPolicySpec_Config_Provider_TokenSourceFromHeader
	if m == nil {
		return target
	}
	target = &JWTPolicySpec_Config_Provider_TokenSourceFromHeader{}

	target.Name = m.GetName()

	target.Prefix = m.GetPrefix()

	return target
}

// Clone function
func (m *JWTPolicySpec_Config_Provider_LocalJWKS_SecretRef) Clone() proto.Message {
	var target *JWTPolicySpec_Config_Provider_LocalJWKS_SecretRef
	if m == nil {
		return target
	}
	target = &JWTPolicySpec_Config_Provider_LocalJWKS_SecretRef{}

	if h, ok := interface{}(m.GetObjectRef()).(clone.Cloner); ok {
		target.ObjectRef = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectReference)
	} else {
		target.ObjectRef = proto.Clone(m.GetObjectRef()).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectReference)
	}

	target.Key = m.GetKey()

	return target
}
