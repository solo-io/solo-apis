// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/admin/v2/ext_auth_server.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2 "github.com/solo-io/solo-apis/client-go/common.gloo.solo.io/v2"

	google_golang_org_protobuf_types_known_durationpb "google.golang.org/protobuf/types/known/durationpb"
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
func (m *ExtAuthServerSpec) Clone() proto.Message {
	var target *ExtAuthServerSpec
	if m == nil {
		return target
	}
	target = &ExtAuthServerSpec{}

	if h, ok := interface{}(m.GetDestinationServer()).(clone.Cloner); ok {
		target.DestinationServer = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference)
	} else {
		target.DestinationServer = proto.Clone(m.GetDestinationServer()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference)
	}

	if h, ok := interface{}(m.GetHttpService()).(clone.Cloner); ok {
		target.HttpService = h.Clone().(*ExtAuthServerSpec_HttpService)
	} else {
		target.HttpService = proto.Clone(m.GetHttpService()).(*ExtAuthServerSpec_HttpService)
	}

	if h, ok := interface{}(m.GetRequestTimeout()).(clone.Cloner); ok {
		target.RequestTimeout = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.RequestTimeout = proto.Clone(m.GetRequestTimeout()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	target.FailureModeAllow = m.GetFailureModeAllow()

	if h, ok := interface{}(m.GetRequestBody()).(clone.Cloner); ok {
		target.RequestBody = h.Clone().(*ExtAuthServerSpec_BufferSettings)
	} else {
		target.RequestBody = proto.Clone(m.GetRequestBody()).(*ExtAuthServerSpec_BufferSettings)
	}

	target.ClearRouteCache = m.GetClearRouteCache()

	target.StatusOnError = m.GetStatusOnError()

	target.TransportApiVersion = m.GetTransportApiVersion()

	target.StatPrefix = m.GetStatPrefix()

	return target
}

// Clone function
func (m *ExtAuthServerReport) Clone() proto.Message {
	var target *ExtAuthServerReport
	if m == nil {
		return target
	}
	target = &ExtAuthServerReport{}

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

	if m.GetAppliedPolicies() != nil {
		target.AppliedPolicies = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference, len(m.GetAppliedPolicies()))
		for idx, v := range m.GetAppliedPolicies() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.AppliedPolicies[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
			} else {
				target.AppliedPolicies[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
			}

		}
	}

	if m.GetSelectedBackingServices() != nil {
		target.SelectedBackingServices = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference, len(m.GetSelectedBackingServices()))
		for idx, v := range m.GetSelectedBackingServices() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SelectedBackingServices[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
			} else {
				target.SelectedBackingServices[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
			}

		}
	}

	return target
}

// Clone function
func (m *ExtAuthServerStatus) Clone() proto.Message {
	var target *ExtAuthServerStatus
	if m == nil {
		return target
	}
	target = &ExtAuthServerStatus{}

	if h, ok := interface{}(m.GetCommon()).(clone.Cloner); ok {
		target.Common = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	} else {
		target.Common = proto.Clone(m.GetCommon()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	}

	target.NumAppliedPolicies = m.GetNumAppliedPolicies()

	target.NumSelectedBackingServices = m.GetNumSelectedBackingServices()

	return target
}

// Clone function
func (m *ExtAuthServerSpec_HttpService) Clone() proto.Message {
	var target *ExtAuthServerSpec_HttpService
	if m == nil {
		return target
	}
	target = &ExtAuthServerSpec_HttpService{}

	target.PathPrefix = m.GetPathPrefix()

	if h, ok := interface{}(m.GetRequest()).(clone.Cloner); ok {
		target.Request = h.Clone().(*ExtAuthServerSpec_HttpService_Request)
	} else {
		target.Request = proto.Clone(m.GetRequest()).(*ExtAuthServerSpec_HttpService_Request)
	}

	if h, ok := interface{}(m.GetResponse()).(clone.Cloner); ok {
		target.Response = h.Clone().(*ExtAuthServerSpec_HttpService_Response)
	} else {
		target.Response = proto.Clone(m.GetResponse()).(*ExtAuthServerSpec_HttpService_Response)
	}

	return target
}

// Clone function
func (m *ExtAuthServerSpec_BufferSettings) Clone() proto.Message {
	var target *ExtAuthServerSpec_BufferSettings
	if m == nil {
		return target
	}
	target = &ExtAuthServerSpec_BufferSettings{}

	target.MaxRequestBytes = m.GetMaxRequestBytes()

	target.AllowPartialMessage = m.GetAllowPartialMessage()

	target.PackAsBytes = m.GetPackAsBytes()

	return target
}

// Clone function
func (m *ExtAuthServerSpec_HttpService_Request) Clone() proto.Message {
	var target *ExtAuthServerSpec_HttpService_Request
	if m == nil {
		return target
	}
	target = &ExtAuthServerSpec_HttpService_Request{}

	if m.GetAllowedHeaders() != nil {
		target.AllowedHeaders = make([]string, len(m.GetAllowedHeaders()))
		for idx, v := range m.GetAllowedHeaders() {

			target.AllowedHeaders[idx] = v

		}
	}

	if m.GetHeadersToAdd() != nil {
		target.HeadersToAdd = make(map[string]string, len(m.GetHeadersToAdd()))
		for k, v := range m.GetHeadersToAdd() {

			target.HeadersToAdd[k] = v

		}
	}

	return target
}

// Clone function
func (m *ExtAuthServerSpec_HttpService_Response) Clone() proto.Message {
	var target *ExtAuthServerSpec_HttpService_Response
	if m == nil {
		return target
	}
	target = &ExtAuthServerSpec_HttpService_Response{}

	if m.GetAllowedUpstreamHeaders() != nil {
		target.AllowedUpstreamHeaders = make([]string, len(m.GetAllowedUpstreamHeaders()))
		for idx, v := range m.GetAllowedUpstreamHeaders() {

			target.AllowedUpstreamHeaders[idx] = v

		}
	}

	if m.GetAllowedClientHeaders() != nil {
		target.AllowedClientHeaders = make([]string, len(m.GetAllowedClientHeaders()))
		for idx, v := range m.GetAllowedClientHeaders() {

			target.AllowedClientHeaders[idx] = v

		}
	}

	return target
}
