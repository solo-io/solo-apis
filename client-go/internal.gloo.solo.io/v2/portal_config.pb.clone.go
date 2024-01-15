// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/internal/v2/portal_config.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_struct "github.com/golang/protobuf/ptypes/struct"

	github_com_solo_io_gloo_mesh_solo_apis_client_go_apimanagement_gloo_solo_io_v2 "github.com/solo-io/solo-apis/client-go/apimanagement.gloo.solo.io/v2"

	github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2 "github.com/solo-io/solo-apis/client-go/common.gloo.solo.io/v2"
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
func (m *PortalConfigSpec) Clone() proto.Message {
	var target *PortalConfigSpec
	if m == nil {
		return target
	}
	target = &PortalConfigSpec{}

	if h, ok := interface{}(m.GetPortalCustomMetadata()).(clone.Cloner); ok {
		target.PortalCustomMetadata = h.Clone().(*github_com_golang_protobuf_ptypes_struct.Value)
	} else {
		target.PortalCustomMetadata = proto.Clone(m.GetPortalCustomMetadata()).(*github_com_golang_protobuf_ptypes_struct.Value)
	}

	if m.GetApis() != nil {
		target.Apis = make([]*PortalConfigSpec_API, len(m.GetApis()))
		for idx, v := range m.GetApis() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Apis[idx] = h.Clone().(*PortalConfigSpec_API)
			} else {
				target.Apis[idx] = proto.Clone(v).(*PortalConfigSpec_API)
			}

		}
	}

	if m.GetUsagePlans() != nil {
		target.UsagePlans = make([]*PortalConfigSpec_UsagePlan, len(m.GetUsagePlans()))
		for idx, v := range m.GetUsagePlans() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.UsagePlans[idx] = h.Clone().(*PortalConfigSpec_UsagePlan)
			} else {
				target.UsagePlans[idx] = proto.Clone(v).(*PortalConfigSpec_UsagePlan)
			}

		}
	}

	if h, ok := interface{}(m.GetPortalRef()).(clone.Cloner); ok {
		target.PortalRef = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
	} else {
		target.PortalRef = proto.Clone(m.GetPortalRef()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
	}

	if m.GetDomains() != nil {
		target.Domains = make([]string, len(m.GetDomains()))
		for idx, v := range m.GetDomains() {

			target.Domains[idx] = v

		}
	}

	if m.GetGroups() != nil {
		target.Groups = make([]*PortalConfigSpec_Group, len(m.GetGroups()))
		for idx, v := range m.GetGroups() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Groups[idx] = h.Clone().(*PortalConfigSpec_Group)
			} else {
				target.Groups[idx] = proto.Clone(v).(*PortalConfigSpec_Group)
			}

		}
	}

	target.Public = m.GetPublic()

	return target
}

// Clone function
func (m *PortalConfigStatus) Clone() proto.Message {
	var target *PortalConfigStatus
	if m == nil {
		return target
	}
	target = &PortalConfigStatus{}

	if h, ok := interface{}(m.GetCommon()).(clone.Cloner); ok {
		target.Common = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	} else {
		target.Common = proto.Clone(m.GetCommon()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	}

	target.OwnedByWorkspace = m.GetOwnedByWorkspace()

	return target
}

// Clone function
func (m *PortalConfigReport) Clone() proto.Message {
	var target *PortalConfigReport
	if m == nil {
		return target
	}
	target = &PortalConfigReport{}

	if m.GetWorkspaces() != nil {
		target.Workspaces = make(map[string]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Report, len(m.GetWorkspaces()))
		for k, v := range m.GetWorkspaces() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Workspaces[k] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Report)
			} else {
				target.Workspaces[k] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Report)
			}

		}
	}

	target.OwnerWorkspace = m.GetOwnerWorkspace()

	return target
}

// Clone function
func (m *PortalConfigSpec_Group) Clone() proto.Message {
	var target *PortalConfigSpec_Group
	if m == nil {
		return target
	}
	target = &PortalConfigSpec_Group{}

	target.Name = m.GetName()

	if m.GetApis() != nil {
		target.Apis = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference, len(m.GetApis()))
		for idx, v := range m.GetApis() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Apis[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
			} else {
				target.Apis[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
			}

		}
	}

	if m.GetUsagePlans() != nil {
		target.UsagePlans = make([]string, len(m.GetUsagePlans()))
		for idx, v := range m.GetUsagePlans() {

			target.UsagePlans[idx] = v

		}
	}

	if m.GetMembershipClaims() != nil {
		target.MembershipClaims = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_apimanagement_gloo_solo_io_v2.Membership, len(m.GetMembershipClaims()))
		for idx, v := range m.GetMembershipClaims() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.MembershipClaims[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_apimanagement_gloo_solo_io_v2.Membership)
			} else {
				target.MembershipClaims[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_apimanagement_gloo_solo_io_v2.Membership)
			}

		}
	}

	return target
}

// Clone function
func (m *PortalConfigSpec_API) Clone() proto.Message {
	var target *PortalConfigSpec_API
	if m == nil {
		return target
	}
	target = &PortalConfigSpec_API{}

	target.ApiProductId = m.GetApiProductId()

	target.ApiProductDisplayName = m.GetApiProductDisplayName()

	target.ApiId = m.GetApiId()

	target.ApiVersion = m.GetApiVersion()

	target.Title = m.GetTitle()

	target.Description = m.GetDescription()

	target.TermsOfService = m.GetTermsOfService()

	target.Contact = m.GetContact()

	target.License = m.GetLicense()

	target.Lifecycle = m.GetLifecycle()

	if h, ok := interface{}(m.GetApiSchema()).(clone.Cloner); ok {
		target.ApiSchema = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
	} else {
		target.ApiSchema = proto.Clone(m.GetApiSchema()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
	}

	if h, ok := interface{}(m.GetRouteTable()).(clone.Cloner); ok {
		target.RouteTable = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
	} else {
		target.RouteTable = proto.Clone(m.GetRouteTable()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
	}

	if m.GetUsagePlans() != nil {
		target.UsagePlans = make([]*PortalConfigSpec_UsagePlanRef, len(m.GetUsagePlans()))
		for idx, v := range m.GetUsagePlans() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.UsagePlans[idx] = h.Clone().(*PortalConfigSpec_UsagePlanRef)
			} else {
				target.UsagePlans[idx] = proto.Clone(v).(*PortalConfigSpec_UsagePlanRef)
			}

		}
	}

	target.IsPrivate = m.GetIsPrivate()

	if m.GetCustomMetadata() != nil {
		target.CustomMetadata = make(map[string]string, len(m.GetCustomMetadata()))
		for k, v := range m.GetCustomMetadata() {

			target.CustomMetadata[k] = v

		}
	}

	return target
}

// Clone function
func (m *PortalConfigSpec_UsagePlan) Clone() proto.Message {
	var target *PortalConfigSpec_UsagePlan
	if m == nil {
		return target
	}
	target = &PortalConfigSpec_UsagePlan{}

	target.Name = m.GetName()

	target.DisplayName = m.GetDisplayName()

	target.Description = m.GetDescription()

	if m.GetExtAuthPolicies() != nil {
		target.ExtAuthPolicies = make([]*PortalConfigSpec_ExtAuthPolicy, len(m.GetExtAuthPolicies()))
		for idx, v := range m.GetExtAuthPolicies() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ExtAuthPolicies[idx] = h.Clone().(*PortalConfigSpec_ExtAuthPolicy)
			} else {
				target.ExtAuthPolicies[idx] = proto.Clone(v).(*PortalConfigSpec_ExtAuthPolicy)
			}

		}
	}

	if h, ok := interface{}(m.GetRateLimitPolicy()).(clone.Cloner); ok {
		target.RateLimitPolicy = h.Clone().(*PortalConfigSpec_RateLimitPolicy)
	} else {
		target.RateLimitPolicy = proto.Clone(m.GetRateLimitPolicy()).(*PortalConfigSpec_RateLimitPolicy)
	}

	return target
}

// Clone function
func (m *PortalConfigSpec_UsagePlanRef) Clone() proto.Message {
	var target *PortalConfigSpec_UsagePlanRef
	if m == nil {
		return target
	}
	target = &PortalConfigSpec_UsagePlanRef{}

	target.Name = m.GetName()

	return target
}

// Clone function
func (m *PortalConfigSpec_ExtAuthPolicy) Clone() proto.Message {
	var target *PortalConfigSpec_ExtAuthPolicy
	if m == nil {
		return target
	}
	target = &PortalConfigSpec_ExtAuthPolicy{}

	if h, ok := interface{}(m.GetExtAuthPolicyRef()).(clone.Cloner); ok {
		target.ExtAuthPolicyRef = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
	} else {
		target.ExtAuthPolicyRef = proto.Clone(m.GetExtAuthPolicyRef()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
	}

	target.AuthConfigId = m.GetAuthConfigId()

	switch m.AuthCfg.(type) {

	case *PortalConfigSpec_ExtAuthPolicy_ApiKeyAuth:

		if h, ok := interface{}(m.GetApiKeyAuth()).(clone.Cloner); ok {
			target.AuthCfg = &PortalConfigSpec_ExtAuthPolicy_ApiKeyAuth{
				ApiKeyAuth: h.Clone().(*PortalConfigSpec_ApiKeyAuth),
			}
		} else {
			target.AuthCfg = &PortalConfigSpec_ExtAuthPolicy_ApiKeyAuth{
				ApiKeyAuth: proto.Clone(m.GetApiKeyAuth()).(*PortalConfigSpec_ApiKeyAuth),
			}
		}

	case *PortalConfigSpec_ExtAuthPolicy_OidcAuth:

		if h, ok := interface{}(m.GetOidcAuth()).(clone.Cloner); ok {
			target.AuthCfg = &PortalConfigSpec_ExtAuthPolicy_OidcAuth{
				OidcAuth: h.Clone().(*PortalConfigSpec_OidcAuth),
			}
		} else {
			target.AuthCfg = &PortalConfigSpec_ExtAuthPolicy_OidcAuth{
				OidcAuth: proto.Clone(m.GetOidcAuth()).(*PortalConfigSpec_OidcAuth),
			}
		}

	case *PortalConfigSpec_ExtAuthPolicy_AccessTokenValidation:

		if h, ok := interface{}(m.GetAccessTokenValidation()).(clone.Cloner); ok {
			target.AuthCfg = &PortalConfigSpec_ExtAuthPolicy_AccessTokenValidation{
				AccessTokenValidation: h.Clone().(*PortalConfigSpec_AccessTokenValidation),
			}
		} else {
			target.AuthCfg = &PortalConfigSpec_ExtAuthPolicy_AccessTokenValidation{
				AccessTokenValidation: proto.Clone(m.GetAccessTokenValidation()).(*PortalConfigSpec_AccessTokenValidation),
			}
		}

	}

	return target
}

// Clone function
func (m *PortalConfigSpec_ApiKeyAuth) Clone() proto.Message {
	var target *PortalConfigSpec_ApiKeyAuth
	if m == nil {
		return target
	}
	target = &PortalConfigSpec_ApiKeyAuth{}

	if m.GetExtAuthLabelSelector() != nil {
		target.ExtAuthLabelSelector = make(map[string]string, len(m.GetExtAuthLabelSelector()))
		for k, v := range m.GetExtAuthLabelSelector() {

			target.ExtAuthLabelSelector[k] = v

		}
	}

	return target
}

// Clone function
func (m *PortalConfigSpec_OidcAuth) Clone() proto.Message {
	var target *PortalConfigSpec_OidcAuth
	if m == nil {
		return target
	}
	target = &PortalConfigSpec_OidcAuth{}

	target.WellKnownOpenidConfig = m.GetWellKnownOpenidConfig()

	return target
}

// Clone function
func (m *PortalConfigSpec_AccessTokenValidation) Clone() proto.Message {
	var target *PortalConfigSpec_AccessTokenValidation
	if m == nil {
		return target
	}
	target = &PortalConfigSpec_AccessTokenValidation{}

	target.BearerFormat = m.GetBearerFormat()

	return target
}

// Clone function
func (m *PortalConfigSpec_RateLimitPolicy) Clone() proto.Message {
	var target *PortalConfigSpec_RateLimitPolicy
	if m == nil {
		return target
	}
	target = &PortalConfigSpec_RateLimitPolicy{}

	target.Unit = m.GetUnit()

	target.RequestsPerUnit = m.GetRequestsPerUnit()

	if h, ok := interface{}(m.GetRateLimitPolicyRef()).(clone.Cloner); ok {
		target.RateLimitPolicyRef = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
	} else {
		target.RateLimitPolicyRef = proto.Clone(m.GetRateLimitPolicyRef()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
	}

	return target
}
