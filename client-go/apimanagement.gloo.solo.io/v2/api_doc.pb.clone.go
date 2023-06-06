// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/apimanagement/v2/api_doc.proto

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
func (m *ApiDocSpec) Clone() proto.Message {
	var target *ApiDocSpec
	if m == nil {
		return target
	}
	target = &ApiDocSpec{}

	if m.GetServedBy() != nil {
		target.ServedBy = make([]*ApiDocSpec_ServedBy, len(m.GetServedBy()))
		for idx, v := range m.GetServedBy() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ServedBy[idx] = h.Clone().(*ApiDocSpec_ServedBy)
			} else {
				target.ServedBy[idx] = proto.Clone(v).(*ApiDocSpec_ServedBy)
			}

		}
	}

	switch m.SchemaType.(type) {

	case *ApiDocSpec_Openapi:

		if h, ok := interface{}(m.GetOpenapi()).(clone.Cloner); ok {
			target.SchemaType = &ApiDocSpec_Openapi{
				Openapi: h.Clone().(*ApiDocSpec_OpenAPISchema),
			}
		} else {
			target.SchemaType = &ApiDocSpec_Openapi{
				Openapi: proto.Clone(m.GetOpenapi()).(*ApiDocSpec_OpenAPISchema),
			}
		}

	case *ApiDocSpec_Grpc:

		if h, ok := interface{}(m.GetGrpc()).(clone.Cloner); ok {
			target.SchemaType = &ApiDocSpec_Grpc{
				Grpc: h.Clone().(*ApiDocSpec_GrpcSchema),
			}
		} else {
			target.SchemaType = &ApiDocSpec_Grpc{
				Grpc: proto.Clone(m.GetGrpc()).(*ApiDocSpec_GrpcSchema),
			}
		}

	case *ApiDocSpec_Graphql:

		if h, ok := interface{}(m.GetGraphql()).(clone.Cloner); ok {
			target.SchemaType = &ApiDocSpec_Graphql{
				Graphql: h.Clone().(*ApiDocSpec_GraphQLSchema),
			}
		} else {
			target.SchemaType = &ApiDocSpec_Graphql{
				Graphql: proto.Clone(m.GetGraphql()).(*ApiDocSpec_GraphQLSchema),
			}
		}

	}

	return target
}

// Clone function
func (m *ApiDocStatus) Clone() proto.Message {
	var target *ApiDocStatus
	if m == nil {
		return target
	}
	target = &ApiDocStatus{}

	if h, ok := interface{}(m.GetCommon()).(clone.Cloner); ok {
		target.Common = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	} else {
		target.Common = proto.Clone(m.GetCommon()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	}

	target.OwnerWorkspace = m.GetOwnerWorkspace()

	target.SelectedServingDestinations = m.GetSelectedServingDestinations()

	return target
}

// Clone function
func (m *ApiDocReport) Clone() proto.Message {
	var target *ApiDocReport
	if m == nil {
		return target
	}
	target = &ApiDocReport{}

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

	if m.GetServingDestinations() != nil {
		target.ServingDestinations = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference, len(m.GetServingDestinations()))
		for idx, v := range m.GetServingDestinations() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ServingDestinations[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference)
			} else {
				target.ServingDestinations[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference)
			}

		}
	}

	return target
}

// Clone function
func (m *ApiDocSpec_ServedBy) Clone() proto.Message {
	var target *ApiDocSpec_ServedBy
	if m == nil {
		return target
	}
	target = &ApiDocSpec_ServedBy{}

	switch m.ServedByType.(type) {

	case *ApiDocSpec_ServedBy_DestinationSelector:

		if h, ok := interface{}(m.GetDestinationSelector()).(clone.Cloner); ok {
			target.ServedByType = &ApiDocSpec_ServedBy_DestinationSelector{
				DestinationSelector: h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationSelector),
			}
		} else {
			target.ServedByType = &ApiDocSpec_ServedBy_DestinationSelector{
				DestinationSelector: proto.Clone(m.GetDestinationSelector()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationSelector),
			}
		}

	case *ApiDocSpec_ServedBy_RouteTable:

		if h, ok := interface{}(m.GetRouteTable()).(clone.Cloner); ok {
			target.ServedByType = &ApiDocSpec_ServedBy_RouteTable{
				RouteTable: h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference),
			}
		} else {
			target.ServedByType = &ApiDocSpec_ServedBy_RouteTable{
				RouteTable: proto.Clone(m.GetRouteTable()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference),
			}
		}

	}

	return target
}

// Clone function
func (m *ApiDocSpec_OpenAPISchema) Clone() proto.Message {
	var target *ApiDocSpec_OpenAPISchema
	if m == nil {
		return target
	}
	target = &ApiDocSpec_OpenAPISchema{}

	target.InlineString = m.GetInlineString()

	return target
}

// Clone function
func (m *ApiDocSpec_GrpcSchema) Clone() proto.Message {
	var target *ApiDocSpec_GrpcSchema
	if m == nil {
		return target
	}
	target = &ApiDocSpec_GrpcSchema{}

	if m.GetDescriptors() != nil {
		target.Descriptors = make([]byte, len(m.GetDescriptors()))
		copy(target.Descriptors, m.GetDescriptors())
	}

	return target
}

// Clone function
func (m *ApiDocSpec_GraphQLSchema) Clone() proto.Message {
	var target *ApiDocSpec_GraphQLSchema
	if m == nil {
		return target
	}
	target = &ApiDocSpec_GraphQLSchema{}

	target.SchemaDefinition = m.GetSchemaDefinition()

	return target
}
