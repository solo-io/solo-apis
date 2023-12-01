// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/apimanagement/v2/graphql_stitched_schema.proto

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

	github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
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
func (m *GraphQLStitchedSchemaSpec) Clone() proto.Message {
	var target *GraphQLStitchedSchemaSpec
	if m == nil {
		return target
	}
	target = &GraphQLStitchedSchemaSpec{}

	if m.GetSubschemas() != nil {
		target.Subschemas = make([]*GraphQLStitchedSchemaSpec_Subschema, len(m.GetSubschemas()))
		for idx, v := range m.GetSubschemas() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Subschemas[idx] = h.Clone().(*GraphQLStitchedSchemaSpec_Subschema)
			} else {
				target.Subschemas[idx] = proto.Clone(v).(*GraphQLStitchedSchemaSpec_Subschema)
			}

		}
	}

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*GraphQLStitchedSchemaSpec_Options)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*GraphQLStitchedSchemaSpec_Options)
	}

	return target
}

// Clone function
func (m *GraphQLStitchedSchemaStatus) Clone() proto.Message {
	var target *GraphQLStitchedSchemaStatus
	if m == nil {
		return target
	}
	target = &GraphQLStitchedSchemaStatus{}

	if h, ok := interface{}(m.GetCommon()).(clone.Cloner); ok {
		target.Common = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	} else {
		target.Common = proto.Clone(m.GetCommon()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	}

	target.OwnedByWorkspace = m.GetOwnedByWorkspace()

	return target
}

// Clone function
func (m *GraphQLStitchedSchemaReport) Clone() proto.Message {
	var target *GraphQLStitchedSchemaReport
	if m == nil {
		return target
	}
	target = &GraphQLStitchedSchemaReport{}

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
func (m *GraphQLStitchedSchemaSpec_Subschema) Clone() proto.Message {
	var target *GraphQLStitchedSchemaSpec_Subschema
	if m == nil {
		return target
	}
	target = &GraphQLStitchedSchemaSpec_Subschema{}

	if m.GetTypeMerge() != nil {
		target.TypeMerge = make(map[string]*GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig, len(m.GetTypeMerge()))
		for k, v := range m.GetTypeMerge() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.TypeMerge[k] = h.Clone().(*GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig)
			} else {
				target.TypeMerge[k] = proto.Clone(v).(*GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig)
			}

		}
	}

	switch m.GraphqlSchema.(type) {

	case *GraphQLStitchedSchemaSpec_Subschema_Schema:

		if h, ok := interface{}(m.GetSchema()).(clone.Cloner); ok {
			target.GraphqlSchema = &GraphQLStitchedSchemaSpec_Subschema_Schema{
				Schema: h.Clone().(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.ClusterObjectRef),
			}
		} else {
			target.GraphqlSchema = &GraphQLStitchedSchemaSpec_Subschema_Schema{
				Schema: proto.Clone(m.GetSchema()).(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.ClusterObjectRef),
			}
		}

	case *GraphQLStitchedSchemaSpec_Subschema_StitchedSchema:

		if h, ok := interface{}(m.GetStitchedSchema()).(clone.Cloner); ok {
			target.GraphqlSchema = &GraphQLStitchedSchemaSpec_Subschema_StitchedSchema{
				StitchedSchema: h.Clone().(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.ClusterObjectRef),
			}
		} else {
			target.GraphqlSchema = &GraphQLStitchedSchemaSpec_Subschema_StitchedSchema{
				StitchedSchema: proto.Clone(m.GetStitchedSchema()).(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.ClusterObjectRef),
			}
		}

	}

	return target
}

// Clone function
func (m *GraphQLStitchedSchemaSpec_Options) Clone() proto.Message {
	var target *GraphQLStitchedSchemaSpec_Options
	if m == nil {
		return target
	}
	target = &GraphQLStitchedSchemaSpec_Options{}

	target.EnableIntrospection = m.GetEnableIntrospection()

	return target
}

// Clone function
func (m *GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig) Clone() proto.Message {
	var target *GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig
	if m == nil {
		return target
	}
	target = &GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig{}

	target.SelectionSet = m.GetSelectionSet()

	target.QueryName = m.GetQueryName()

	if m.GetArgs() != nil {
		target.Args = make(map[string]string, len(m.GetArgs()))
		for k, v := range m.GetArgs() {

			target.Args[k] = v

		}
	}

	return target
}