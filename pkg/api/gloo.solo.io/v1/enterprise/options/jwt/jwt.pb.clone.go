// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/enterprise/options/jwt/jwt.proto

package jwt

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_extensions_filters_http_jwt_authn_v3 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/extensions/filters/http/jwt_authn/v3"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	google_golang_org_protobuf_types_known_durationpb "google.golang.org/protobuf/types/known/durationpb"

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
func (m *JwtStagedVhostExtension) Clone() proto.Message {
	var target *JwtStagedVhostExtension
	if m == nil {
		return target
	}
	target = &JwtStagedVhostExtension{}

	if h, ok := interface{}(m.GetBeforeExtAuth()).(clone.Cloner); ok {
		target.BeforeExtAuth = h.Clone().(*VhostExtension)
	} else {
		target.BeforeExtAuth = proto.Clone(m.GetBeforeExtAuth()).(*VhostExtension)
	}

	if h, ok := interface{}(m.GetAfterExtAuth()).(clone.Cloner); ok {
		target.AfterExtAuth = h.Clone().(*VhostExtension)
	} else {
		target.AfterExtAuth = proto.Clone(m.GetAfterExtAuth()).(*VhostExtension)
	}

	return target
}

// Clone function
func (m *JwtStagedRouteProvidersExtension) Clone() proto.Message {
	var target *JwtStagedRouteProvidersExtension
	if m == nil {
		return target
	}
	target = &JwtStagedRouteProvidersExtension{}

	if h, ok := interface{}(m.GetBeforeExtAuth()).(clone.Cloner); ok {
		target.BeforeExtAuth = h.Clone().(*VhostExtension)
	} else {
		target.BeforeExtAuth = proto.Clone(m.GetBeforeExtAuth()).(*VhostExtension)
	}

	if h, ok := interface{}(m.GetAfterExtAuth()).(clone.Cloner); ok {
		target.AfterExtAuth = h.Clone().(*VhostExtension)
	} else {
		target.AfterExtAuth = proto.Clone(m.GetAfterExtAuth()).(*VhostExtension)
	}

	return target
}

// Clone function
func (m *JwtStagedRouteExtension) Clone() proto.Message {
	var target *JwtStagedRouteExtension
	if m == nil {
		return target
	}
	target = &JwtStagedRouteExtension{}

	if h, ok := interface{}(m.GetBeforeExtAuth()).(clone.Cloner); ok {
		target.BeforeExtAuth = h.Clone().(*RouteExtension)
	} else {
		target.BeforeExtAuth = proto.Clone(m.GetBeforeExtAuth()).(*RouteExtension)
	}

	if h, ok := interface{}(m.GetAfterExtAuth()).(clone.Cloner); ok {
		target.AfterExtAuth = h.Clone().(*RouteExtension)
	} else {
		target.AfterExtAuth = proto.Clone(m.GetAfterExtAuth()).(*RouteExtension)
	}

	return target
}

// Clone function
func (m *VhostExtension) Clone() proto.Message {
	var target *VhostExtension
	if m == nil {
		return target
	}
	target = &VhostExtension{}

	if m.GetProviders() != nil {
		target.Providers = make(map[string]*Provider, len(m.GetProviders()))
		for k, v := range m.GetProviders() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Providers[k] = h.Clone().(*Provider)
			} else {
				target.Providers[k] = proto.Clone(v).(*Provider)
			}

		}
	}

	target.AllowMissingOrFailedJwt = m.GetAllowMissingOrFailedJwt()

	target.ValidationPolicy = m.GetValidationPolicy()

	return target
}

// Clone function
func (m *RouteExtension) Clone() proto.Message {
	var target *RouteExtension
	if m == nil {
		return target
	}
	target = &RouteExtension{}

	target.Disable = m.GetDisable()

	return target
}

// Clone function
func (m *Provider) Clone() proto.Message {
	var target *Provider
	if m == nil {
		return target
	}
	target = &Provider{}

	if h, ok := interface{}(m.GetJwks()).(clone.Cloner); ok {
		target.Jwks = h.Clone().(*Jwks)
	} else {
		target.Jwks = proto.Clone(m.GetJwks()).(*Jwks)
	}

	if m.GetAudiences() != nil {
		target.Audiences = make([]string, len(m.GetAudiences()))
		for idx, v := range m.GetAudiences() {

			target.Audiences[idx] = v

		}
	}

	target.Issuer = m.GetIssuer()

	if h, ok := interface{}(m.GetTokenSource()).(clone.Cloner); ok {
		target.TokenSource = h.Clone().(*TokenSource)
	} else {
		target.TokenSource = proto.Clone(m.GetTokenSource()).(*TokenSource)
	}

	target.KeepToken = m.GetKeepToken()

	if m.GetClaimsToHeaders() != nil {
		target.ClaimsToHeaders = make([]*ClaimToHeader, len(m.GetClaimsToHeaders()))
		for idx, v := range m.GetClaimsToHeaders() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ClaimsToHeaders[idx] = h.Clone().(*ClaimToHeader)
			} else {
				target.ClaimsToHeaders[idx] = proto.Clone(v).(*ClaimToHeader)
			}

		}
	}

	if h, ok := interface{}(m.GetClockSkewSeconds()).(clone.Cloner); ok {
		target.ClockSkewSeconds = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.ClockSkewSeconds = proto.Clone(m.GetClockSkewSeconds()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	return target
}

// Clone function
func (m *Jwks) Clone() proto.Message {
	var target *Jwks
	if m == nil {
		return target
	}
	target = &Jwks{}

	switch m.Jwks.(type) {

	case *Jwks_Remote:

		if h, ok := interface{}(m.GetRemote()).(clone.Cloner); ok {
			target.Jwks = &Jwks_Remote{
				Remote: h.Clone().(*RemoteJwks),
			}
		} else {
			target.Jwks = &Jwks_Remote{
				Remote: proto.Clone(m.GetRemote()).(*RemoteJwks),
			}
		}

	case *Jwks_Local:

		if h, ok := interface{}(m.GetLocal()).(clone.Cloner); ok {
			target.Jwks = &Jwks_Local{
				Local: h.Clone().(*LocalJwks),
			}
		} else {
			target.Jwks = &Jwks_Local{
				Local: proto.Clone(m.GetLocal()).(*LocalJwks),
			}
		}

	}

	return target
}

// Clone function
func (m *RemoteJwks) Clone() proto.Message {
	var target *RemoteJwks
	if m == nil {
		return target
	}
	target = &RemoteJwks{}

	target.Url = m.GetUrl()

	if h, ok := interface{}(m.GetUpstreamRef()).(clone.Cloner); ok {
		target.UpstreamRef = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	} else {
		target.UpstreamRef = proto.Clone(m.GetUpstreamRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	}

	if h, ok := interface{}(m.GetCacheDuration()).(clone.Cloner); ok {
		target.CacheDuration = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.CacheDuration = proto.Clone(m.GetCacheDuration()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetAsyncFetch()).(clone.Cloner); ok {
		target.AsyncFetch = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_extensions_filters_http_jwt_authn_v3.JwksAsyncFetch)
	} else {
		target.AsyncFetch = proto.Clone(m.GetAsyncFetch()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_extensions_filters_http_jwt_authn_v3.JwksAsyncFetch)
	}

	return target
}

// Clone function
func (m *LocalJwks) Clone() proto.Message {
	var target *LocalJwks
	if m == nil {
		return target
	}
	target = &LocalJwks{}

	target.Key = m.GetKey()

	return target
}

// Clone function
func (m *TokenSource) Clone() proto.Message {
	var target *TokenSource
	if m == nil {
		return target
	}
	target = &TokenSource{}

	if m.GetHeaders() != nil {
		target.Headers = make([]*TokenSource_HeaderSource, len(m.GetHeaders()))
		for idx, v := range m.GetHeaders() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Headers[idx] = h.Clone().(*TokenSource_HeaderSource)
			} else {
				target.Headers[idx] = proto.Clone(v).(*TokenSource_HeaderSource)
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
func (m *ClaimToHeader) Clone() proto.Message {
	var target *ClaimToHeader
	if m == nil {
		return target
	}
	target = &ClaimToHeader{}

	target.Claim = m.GetClaim()

	target.Header = m.GetHeader()

	target.Append = m.GetAppend()

	return target
}

// Clone function
func (m *TokenSource_HeaderSource) Clone() proto.Message {
	var target *TokenSource_HeaderSource
	if m == nil {
		return target
	}
	target = &TokenSource_HeaderSource{}

	target.Header = m.GetHeader()

	target.Prefix = m.GetPrefix()

	return target
}
