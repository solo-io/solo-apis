// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/apimanagement/v2/api_doc.proto

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
func (m *ApiDocSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ApiDocSpec)
	if !ok {
		that2, ok := that.(ApiDocSpec)
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

	if len(m.GetServedBy()) != len(target.GetServedBy()) {
		return false
	}
	for idx, v := range m.GetServedBy() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetServedBy()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetServedBy()[idx]) {
				return false
			}
		}

	}

	switch m.SchemaType.(type) {

	case *ApiDocSpec_Openapi:
		if _, ok := target.SchemaType.(*ApiDocSpec_Openapi); !ok {
			return false
		}

		if h, ok := interface{}(m.GetOpenapi()).(equality.Equalizer); ok {
			if !h.Equal(target.GetOpenapi()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetOpenapi(), target.GetOpenapi()) {
				return false
			}
		}

	case *ApiDocSpec_Grpc:
		if _, ok := target.SchemaType.(*ApiDocSpec_Grpc); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGrpc()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGrpc()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGrpc(), target.GetGrpc()) {
				return false
			}
		}

	case *ApiDocSpec_Graphql:
		if _, ok := target.SchemaType.(*ApiDocSpec_Graphql); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGraphql()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGraphql()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGraphql(), target.GetGraphql()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.SchemaType != target.SchemaType {
			return false
		}
	}

	return true
}

// Equal function
func (m *ApiDocStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ApiDocStatus)
	if !ok {
		that2, ok := that.(ApiDocStatus)
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

	if len(m.GetServingDestinations()) != len(target.GetServingDestinations()) {
		return false
	}
	for idx, v := range m.GetServingDestinations() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetServingDestinations()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetServingDestinations()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *ApiDocSpec_ServedBy) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ApiDocSpec_ServedBy)
	if !ok {
		that2, ok := that.(ApiDocSpec_ServedBy)
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

	switch m.ServedByType.(type) {

	case *ApiDocSpec_ServedBy_DestinationSelector:
		if _, ok := target.ServedByType.(*ApiDocSpec_ServedBy_DestinationSelector); !ok {
			return false
		}

		if h, ok := interface{}(m.GetDestinationSelector()).(equality.Equalizer); ok {
			if !h.Equal(target.GetDestinationSelector()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetDestinationSelector(), target.GetDestinationSelector()) {
				return false
			}
		}

	case *ApiDocSpec_ServedBy_RouteTable:
		if _, ok := target.ServedByType.(*ApiDocSpec_ServedBy_RouteTable); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRouteTable()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRouteTable()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRouteTable(), target.GetRouteTable()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.ServedByType != target.ServedByType {
			return false
		}
	}

	return true
}

// Equal function
func (m *ApiDocSpec_OpenAPISchema) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ApiDocSpec_OpenAPISchema)
	if !ok {
		that2, ok := that.(ApiDocSpec_OpenAPISchema)
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

	if strings.Compare(m.GetInlineString(), target.GetInlineString()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *ApiDocSpec_GrpcSchema) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ApiDocSpec_GrpcSchema)
	if !ok {
		that2, ok := that.(ApiDocSpec_GrpcSchema)
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

	if bytes.Compare(m.GetDescriptors(), target.GetDescriptors()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *ApiDocSpec_GraphQLSchema) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ApiDocSpec_GraphQLSchema)
	if !ok {
		that2, ok := that.(ApiDocSpec_GraphQLSchema)
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

	if strings.Compare(m.GetSchemaDefinition(), target.GetSchemaDefinition()) != 0 {
		return false
	}

	return true
}
