// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/admin/v2/dashboard.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	google_golang_org_protobuf_types_known_durationpb "google.golang.org/protobuf/types/known/durationpb"

	google_golang_org_protobuf_types_known_emptypb "google.golang.org/protobuf/types/known/emptypb"

	google_golang_org_protobuf_types_known_wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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
func (m *DashboardSpec) Clone() proto.Message {
	var target *DashboardSpec
	if m == nil {
		return target
	}
	target = &DashboardSpec{}

	if h, ok := interface{}(m.GetAuthn()).(clone.Cloner); ok {
		target.Authn = h.Clone().(*DashboardSpec_AuthnConfig)
	} else {
		target.Authn = proto.Clone(m.GetAuthn()).(*DashboardSpec_AuthnConfig)
	}

	if h, ok := interface{}(m.GetAuthz()).(clone.Cloner); ok {
		target.Authz = h.Clone().(*DashboardSpec_AuthzConfig)
	} else {
		target.Authz = proto.Clone(m.GetAuthz()).(*DashboardSpec_AuthzConfig)
	}

	return target
}

// Clone function
func (m *MultiClusterRbac) Clone() proto.Message {
	var target *MultiClusterRbac
	if m == nil {
		return target
	}
	target = &MultiClusterRbac{}

	return target
}

// Clone function
func (m *SessionConfig) Clone() proto.Message {
	var target *SessionConfig
	if m == nil {
		return target
	}
	target = &SessionConfig{}

	if h, ok := interface{}(m.GetCookieOptions()).(clone.Cloner); ok {
		target.CookieOptions = h.Clone().(*SessionConfig_CookieOptions)
	} else {
		target.CookieOptions = proto.Clone(m.GetCookieOptions()).(*SessionConfig_CookieOptions)
	}

	switch m.Backend.(type) {

	case *SessionConfig_Cookie:

		if h, ok := interface{}(m.GetCookie()).(clone.Cloner); ok {
			target.Backend = &SessionConfig_Cookie{
				Cookie: h.Clone().(*SessionConfig_CookieSession),
			}
		} else {
			target.Backend = &SessionConfig_Cookie{
				Cookie: proto.Clone(m.GetCookie()).(*SessionConfig_CookieSession),
			}
		}

	case *SessionConfig_Redis:

		if h, ok := interface{}(m.GetRedis()).(clone.Cloner); ok {
			target.Backend = &SessionConfig_Redis{
				Redis: h.Clone().(*SessionConfig_RedisSession),
			}
		} else {
			target.Backend = &SessionConfig_Redis{
				Redis: proto.Clone(m.GetRedis()).(*SessionConfig_RedisSession),
			}
		}

	}

	return target
}

// Clone function
func (m *OidcConfig) Clone() proto.Message {
	var target *OidcConfig
	if m == nil {
		return target
	}
	target = &OidcConfig{}

	target.ClientId = m.GetClientId()

	target.ClientSecretName = m.GetClientSecretName()

	target.IssuerUrl = m.GetIssuerUrl()

	if m.GetAuthEndpointQueryParams() != nil {
		target.AuthEndpointQueryParams = make(map[string]string, len(m.GetAuthEndpointQueryParams()))
		for k, v := range m.GetAuthEndpointQueryParams() {

			target.AuthEndpointQueryParams[k] = v

		}
	}

	if m.GetTokenEndpointQueryParams() != nil {
		target.TokenEndpointQueryParams = make(map[string]string, len(m.GetTokenEndpointQueryParams()))
		for k, v := range m.GetTokenEndpointQueryParams() {

			target.TokenEndpointQueryParams[k] = v

		}
	}

	target.AppUrl = m.GetAppUrl()

	target.CallbackPath = m.GetCallbackPath()

	target.LogoutPath = m.GetLogoutPath()

	if m.GetScopes() != nil {
		target.Scopes = make([]string, len(m.GetScopes()))
		for idx, v := range m.GetScopes() {

			target.Scopes[idx] = v

		}
	}

	if h, ok := interface{}(m.GetSession()).(clone.Cloner); ok {
		target.Session = h.Clone().(*SessionConfig)
	} else {
		target.Session = proto.Clone(m.GetSession()).(*SessionConfig)
	}

	if h, ok := interface{}(m.GetDiscoveryOverride()).(clone.Cloner); ok {
		target.DiscoveryOverride = h.Clone().(*OidcConfig_DiscoveryOverride)
	} else {
		target.DiscoveryOverride = proto.Clone(m.GetDiscoveryOverride()).(*OidcConfig_DiscoveryOverride)
	}

	if h, ok := interface{}(m.GetDiscoveryPollInterval()).(clone.Cloner); ok {
		target.DiscoveryPollInterval = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.DiscoveryPollInterval = proto.Clone(m.GetDiscoveryPollInterval()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetJwksCacheRefreshPolicy()).(clone.Cloner); ok {
		target.JwksCacheRefreshPolicy = h.Clone().(*JwksOnDemandCacheRefreshPolicy)
	} else {
		target.JwksCacheRefreshPolicy = proto.Clone(m.GetJwksCacheRefreshPolicy()).(*JwksOnDemandCacheRefreshPolicy)
	}

	if h, ok := interface{}(m.GetUserMapping()).(clone.Cloner); ok {
		target.UserMapping = h.Clone().(*UserMapping)
	} else {
		target.UserMapping = proto.Clone(m.GetUserMapping()).(*UserMapping)
	}

	target.CaCertConfigmapName = m.GetCaCertConfigmapName()

	return target
}

// Clone function
func (m *JwksOnDemandCacheRefreshPolicy) Clone() proto.Message {
	var target *JwksOnDemandCacheRefreshPolicy
	if m == nil {
		return target
	}
	target = &JwksOnDemandCacheRefreshPolicy{}

	switch m.Policy.(type) {

	case *JwksOnDemandCacheRefreshPolicy_Never:

		if h, ok := interface{}(m.GetNever()).(clone.Cloner); ok {
			target.Policy = &JwksOnDemandCacheRefreshPolicy_Never{
				Never: h.Clone().(*google_golang_org_protobuf_types_known_emptypb.Empty),
			}
		} else {
			target.Policy = &JwksOnDemandCacheRefreshPolicy_Never{
				Never: proto.Clone(m.GetNever()).(*google_golang_org_protobuf_types_known_emptypb.Empty),
			}
		}

	case *JwksOnDemandCacheRefreshPolicy_Always:

		if h, ok := interface{}(m.GetAlways()).(clone.Cloner); ok {
			target.Policy = &JwksOnDemandCacheRefreshPolicy_Always{
				Always: h.Clone().(*google_golang_org_protobuf_types_known_emptypb.Empty),
			}
		} else {
			target.Policy = &JwksOnDemandCacheRefreshPolicy_Always{
				Always: proto.Clone(m.GetAlways()).(*google_golang_org_protobuf_types_known_emptypb.Empty),
			}
		}

	case *JwksOnDemandCacheRefreshPolicy_MaxIdpReqPerPollingInterval:

		target.Policy = &JwksOnDemandCacheRefreshPolicy_MaxIdpReqPerPollingInterval{
			MaxIdpReqPerPollingInterval: m.GetMaxIdpReqPerPollingInterval(),
		}

	}

	return target
}

// Clone function
func (m *UserMapping) Clone() proto.Message {
	var target *UserMapping
	if m == nil {
		return target
	}
	target = &UserMapping{}

	target.UsernameClaim = m.GetUsernameClaim()

	target.UsernamePrefix = m.GetUsernamePrefix()

	target.GroupsClaim = m.GetGroupsClaim()

	target.GroupsPrefix = m.GetGroupsPrefix()

	return target
}

// Clone function
func (m *DashboardStatus) Clone() proto.Message {
	var target *DashboardStatus
	if m == nil {
		return target
	}
	target = &DashboardStatus{}

	target.ObservedGeneration = m.GetObservedGeneration()

	target.State = m.GetState()

	if m.GetErrors() != nil {
		target.Errors = make([]string, len(m.GetErrors()))
		for idx, v := range m.GetErrors() {

			target.Errors[idx] = v

		}
	}

	return target
}

// Clone function
func (m *DashboardSpec_AuthnConfig) Clone() proto.Message {
	var target *DashboardSpec_AuthnConfig
	if m == nil {
		return target
	}
	target = &DashboardSpec_AuthnConfig{}

	switch m.Backend.(type) {

	case *DashboardSpec_AuthnConfig_Oidc:

		if h, ok := interface{}(m.GetOidc()).(clone.Cloner); ok {
			target.Backend = &DashboardSpec_AuthnConfig_Oidc{
				Oidc: h.Clone().(*OidcConfig),
			}
		} else {
			target.Backend = &DashboardSpec_AuthnConfig_Oidc{
				Oidc: proto.Clone(m.GetOidc()).(*OidcConfig),
			}
		}

	}

	return target
}

// Clone function
func (m *DashboardSpec_AuthzConfig) Clone() proto.Message {
	var target *DashboardSpec_AuthzConfig
	if m == nil {
		return target
	}
	target = &DashboardSpec_AuthzConfig{}

	switch m.Backend.(type) {

	case *DashboardSpec_AuthzConfig_MultiClusterRbac:

		if h, ok := interface{}(m.GetMultiClusterRbac()).(clone.Cloner); ok {
			target.Backend = &DashboardSpec_AuthzConfig_MultiClusterRbac{
				MultiClusterRbac: h.Clone().(*MultiClusterRbac),
			}
		} else {
			target.Backend = &DashboardSpec_AuthzConfig_MultiClusterRbac{
				MultiClusterRbac: proto.Clone(m.GetMultiClusterRbac()).(*MultiClusterRbac),
			}
		}

	}

	return target
}

// Clone function
func (m *SessionConfig_CookieSession) Clone() proto.Message {
	var target *SessionConfig_CookieSession
	if m == nil {
		return target
	}
	target = &SessionConfig_CookieSession{}

	return target
}

// Clone function
func (m *SessionConfig_RedisSession) Clone() proto.Message {
	var target *SessionConfig_RedisSession
	if m == nil {
		return target
	}
	target = &SessionConfig_RedisSession{}

	target.Host = m.GetHost()

	target.Db = m.GetDb()

	target.PoolSize = m.GetPoolSize()

	target.KeyPrefix = m.GetKeyPrefix()

	target.CookieName = m.GetCookieName()

	if h, ok := interface{}(m.GetAllowRefreshing()).(clone.Cloner); ok {
		target.AllowRefreshing = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.AllowRefreshing = proto.Clone(m.GetAllowRefreshing()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	return target
}

// Clone function
func (m *SessionConfig_CookieOptions) Clone() proto.Message {
	var target *SessionConfig_CookieOptions
	if m == nil {
		return target
	}
	target = &SessionConfig_CookieOptions{}

	if h, ok := interface{}(m.GetMaxAge()).(clone.Cloner); ok {
		target.MaxAge = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.MaxAge = proto.Clone(m.GetMaxAge()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	target.NotSecure = m.GetNotSecure()

	if h, ok := interface{}(m.GetPath()).(clone.Cloner); ok {
		target.Path = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.Path = proto.Clone(m.GetPath()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	target.Domain = m.GetDomain()

	return target
}

// Clone function
func (m *OidcConfig_DiscoveryOverride) Clone() proto.Message {
	var target *OidcConfig_DiscoveryOverride
	if m == nil {
		return target
	}
	target = &OidcConfig_DiscoveryOverride{}

	target.AuthEndpoint = m.GetAuthEndpoint()

	target.TokenEndpoint = m.GetTokenEndpoint()

	target.JwksUri = m.GetJwksUri()

	if m.GetScopes() != nil {
		target.Scopes = make([]string, len(m.GetScopes()))
		for idx, v := range m.GetScopes() {

			target.Scopes[idx] = v

		}
	}

	if m.GetResponseTypes() != nil {
		target.ResponseTypes = make([]string, len(m.GetResponseTypes()))
		for idx, v := range m.GetResponseTypes() {

			target.ResponseTypes[idx] = v

		}
	}

	if m.GetSubjects() != nil {
		target.Subjects = make([]string, len(m.GetSubjects()))
		for idx, v := range m.GetSubjects() {

			target.Subjects[idx] = v

		}
	}

	if m.GetIdTokenAlgs() != nil {
		target.IdTokenAlgs = make([]string, len(m.GetIdTokenAlgs()))
		for idx, v := range m.GetIdTokenAlgs() {

			target.IdTokenAlgs[idx] = v

		}
	}

	if m.GetAuthMethods() != nil {
		target.AuthMethods = make([]string, len(m.GetAuthMethods()))
		for idx, v := range m.GetAuthMethods() {

			target.AuthMethods[idx] = v

		}
	}

	if m.GetClaims() != nil {
		target.Claims = make([]string, len(m.GetClaims()))
		for idx, v := range m.GetClaims() {

			target.Claims[idx] = v

		}
	}

	return target
}
