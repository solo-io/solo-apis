// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/apimanagement/v2/graphql_stitched_schema.proto

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
func (m *GraphQLStitchedSchemaSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GraphQLStitchedSchemaSpec)
	if !ok {
		that2, ok := that.(GraphQLStitchedSchemaSpec)
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

	if len(m.GetSubschemas()) != len(target.GetSubschemas()) {
		return false
	}
	for idx, v := range m.GetSubschemas() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSubschemas()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSubschemas()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *GraphQLStitchedSchemaStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GraphQLStitchedSchemaStatus)
	if !ok {
		that2, ok := that.(GraphQLStitchedSchemaStatus)
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

	if h, ok := interface{}(m.GetOwnerWorkspace()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOwnerWorkspace()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOwnerWorkspace(), target.GetOwnerWorkspace()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GraphQLStitchedSchemaSpec_Subschema) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GraphQLStitchedSchemaSpec_Subschema)
	if !ok {
		that2, ok := that.(GraphQLStitchedSchemaSpec_Subschema)
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

	if len(m.GetTypeMerge()) != len(target.GetTypeMerge()) {
		return false
	}
	for k, v := range m.GetTypeMerge() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetTypeMerge()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetTypeMerge()[k]) {
				return false
			}
		}

	}

	switch m.GraphqlSchema.(type) {

	case *GraphQLStitchedSchemaSpec_Subschema_Schema:
		if _, ok := target.GraphqlSchema.(*GraphQLStitchedSchemaSpec_Subschema_Schema); !ok {
			return false
		}

		if h, ok := interface{}(m.GetSchema()).(equality.Equalizer); ok {
			if !h.Equal(target.GetSchema()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetSchema(), target.GetSchema()) {
				return false
			}
		}

	case *GraphQLStitchedSchemaSpec_Subschema_StitchedSchema:
		if _, ok := target.GraphqlSchema.(*GraphQLStitchedSchemaSpec_Subschema_StitchedSchema); !ok {
			return false
		}

		if h, ok := interface{}(m.GetStitchedSchema()).(equality.Equalizer); ok {
			if !h.Equal(target.GetStitchedSchema()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetStitchedSchema(), target.GetStitchedSchema()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.GraphqlSchema != target.GraphqlSchema {
			return false
		}
	}

	return true
}

// Equal function
func (m *GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig)
	if !ok {
		that2, ok := that.(GraphQLStitchedSchemaSpec_Subschema_TypeMergeConfig)
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

	if strings.Compare(m.GetSelectionSet(), target.GetSelectionSet()) != 0 {
		return false
	}

	if strings.Compare(m.GetQueryName(), target.GetQueryName()) != 0 {
		return false
	}

	if len(m.GetArgs()) != len(target.GetArgs()) {
		return false
	}
	for k, v := range m.GetArgs() {

		if strings.Compare(v, target.GetArgs()[k]) != 0 {
			return false
		}

	}

	return true
}