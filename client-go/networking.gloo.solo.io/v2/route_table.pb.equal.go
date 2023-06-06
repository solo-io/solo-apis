// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/networking/v2/route_table.proto

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
func (m *RouteTableSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteTableSpec)
	if !ok {
		that2, ok := that.(RouteTableSpec)
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

	if len(m.GetHosts()) != len(target.GetHosts()) {
		return false
	}
	for idx, v := range m.GetHosts() {

		if strings.Compare(v, target.GetHosts()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetVirtualGateways()) != len(target.GetVirtualGateways()) {
		return false
	}
	for idx, v := range m.GetVirtualGateways() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetVirtualGateways()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetVirtualGateways()[idx]) {
				return false
			}
		}

	}

	if len(m.GetWorkloadSelectors()) != len(target.GetWorkloadSelectors()) {
		return false
	}
	for idx, v := range m.GetWorkloadSelectors() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetWorkloadSelectors()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetWorkloadSelectors()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetDefaultDestination()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDefaultDestination()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDefaultDestination(), target.GetDefaultDestination()) {
			return false
		}
	}

	if len(m.GetHttp()) != len(target.GetHttp()) {
		return false
	}
	for idx, v := range m.GetHttp() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetHttp()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetHttp()[idx]) {
				return false
			}
		}

	}

	if m.GetWeight() != target.GetWeight() {
		return false
	}

	return true
}

// Equal function
func (m *HTTPRoute) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HTTPRoute)
	if !ok {
		that2, ok := that.(HTTPRoute)
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

	if len(m.GetLabels()) != len(target.GetLabels()) {
		return false
	}
	for k, v := range m.GetLabels() {

		if strings.Compare(v, target.GetLabels()[k]) != 0 {
			return false
		}

	}

	if len(m.GetMatchers()) != len(target.GetMatchers()) {
		return false
	}
	for idx, v := range m.GetMatchers() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetMatchers()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetMatchers()[idx]) {
				return false
			}
		}

	}

	switch m.ActionType.(type) {

	case *HTTPRoute_ForwardTo:
		if _, ok := target.ActionType.(*HTTPRoute_ForwardTo); !ok {
			return false
		}

		if h, ok := interface{}(m.GetForwardTo()).(equality.Equalizer); ok {
			if !h.Equal(target.GetForwardTo()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetForwardTo(), target.GetForwardTo()) {
				return false
			}
		}

	case *HTTPRoute_Delegate:
		if _, ok := target.ActionType.(*HTTPRoute_Delegate); !ok {
			return false
		}

		if h, ok := interface{}(m.GetDelegate()).(equality.Equalizer); ok {
			if !h.Equal(target.GetDelegate()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetDelegate(), target.GetDelegate()) {
				return false
			}
		}

	case *HTTPRoute_Redirect:
		if _, ok := target.ActionType.(*HTTPRoute_Redirect); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRedirect()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRedirect()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRedirect(), target.GetRedirect()) {
				return false
			}
		}

	case *HTTPRoute_DirectResponse:
		if _, ok := target.ActionType.(*HTTPRoute_DirectResponse); !ok {
			return false
		}

		if h, ok := interface{}(m.GetDirectResponse()).(equality.Equalizer); ok {
			if !h.Equal(target.GetDirectResponse()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetDirectResponse(), target.GetDirectResponse()) {
				return false
			}
		}

	case *HTTPRoute_Graphql:
		if _, ok := target.ActionType.(*HTTPRoute_Graphql); !ok {
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
		if m.ActionType != target.ActionType {
			return false
		}
	}

	return true
}

// Equal function
func (m *GraphQLAction) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GraphQLAction)
	if !ok {
		that2, ok := that.(GraphQLAction)
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

	if h, ok := interface{}(m.GetOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOptions(), target.GetOptions()) {
			return false
		}
	}

	switch m.GraphqlSchema.(type) {

	case *GraphQLAction_Schema:
		if _, ok := target.GraphqlSchema.(*GraphQLAction_Schema); !ok {
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

	case *GraphQLAction_StitchedSchema:
		if _, ok := target.GraphqlSchema.(*GraphQLAction_StitchedSchema); !ok {
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
func (m *ForwardToAction) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ForwardToAction)
	if !ok {
		that2, ok := that.(ForwardToAction)
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

	if len(m.GetDestinations()) != len(target.GetDestinations()) {
		return false
	}
	for idx, v := range m.GetDestinations() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetDestinations()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetDestinations()[idx]) {
				return false
			}
		}

	}

	if strings.Compare(m.GetPathRewrite(), target.GetPathRewrite()) != 0 {
		return false
	}

	if strings.Compare(m.GetHostRewrite(), target.GetHostRewrite()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *RedirectAction) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RedirectAction)
	if !ok {
		that2, ok := that.(RedirectAction)
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

	if strings.Compare(m.GetHostRedirect(), target.GetHostRedirect()) != 0 {
		return false
	}

	if m.GetResponseCode() != target.GetResponseCode() {
		return false
	}

	switch m.PathRewriteSpecifier.(type) {

	case *RedirectAction_PathRedirect:
		if _, ok := target.PathRewriteSpecifier.(*RedirectAction_PathRedirect); !ok {
			return false
		}

		if strings.Compare(m.GetPathRedirect(), target.GetPathRedirect()) != 0 {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.PathRewriteSpecifier != target.PathRewriteSpecifier {
			return false
		}
	}

	return true
}

// Equal function
func (m *DirectResponseAction) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DirectResponseAction)
	if !ok {
		that2, ok := that.(DirectResponseAction)
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

	if m.GetStatus() != target.GetStatus() {
		return false
	}

	if strings.Compare(m.GetBody(), target.GetBody()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *DelegateAction) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DelegateAction)
	if !ok {
		that2, ok := that.(DelegateAction)
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

	if len(m.GetRouteTables()) != len(target.GetRouteTables()) {
		return false
	}
	for idx, v := range m.GetRouteTables() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetRouteTables()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetRouteTables()[idx]) {
				return false
			}
		}

	}

	if len(m.GetAllowedRoutes()) != len(target.GetAllowedRoutes()) {
		return false
	}
	for idx, v := range m.GetAllowedRoutes() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAllowedRoutes()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAllowedRoutes()[idx]) {
				return false
			}
		}

	}

	if m.GetSortMethod() != target.GetSortMethod() {
		return false
	}

	return true
}

// Equal function
func (m *RouteTableStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteTableStatus)
	if !ok {
		that2, ok := that.(RouteTableStatus)
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

	if len(m.GetAppliedRoutePolicies()) != len(target.GetAppliedRoutePolicies()) {
		return false
	}
	for k, v := range m.GetAppliedRoutePolicies() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAppliedRoutePolicies()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAppliedRoutePolicies()[k]) {
				return false
			}
		}

	}

	if len(m.GetParentRouteTables()) != len(target.GetParentRouteTables()) {
		return false
	}
	for idx, v := range m.GetParentRouteTables() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetParentRouteTables()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetParentRouteTables()[idx]) {
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

	if len(m.GetAllowedVirtualGateways()) != len(target.GetAllowedVirtualGateways()) {
		return false
	}
	for idx, v := range m.GetAllowedVirtualGateways() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAllowedVirtualGateways()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAllowedVirtualGateways()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *GraphQLAction_Options) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GraphQLAction_Options)
	if !ok {
		that2, ok := that.(GraphQLAction_Options)
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

	if h, ok := interface{}(m.GetLogSensitiveInfo()).(equality.Equalizer); ok {
		if !h.Equal(target.GetLogSensitiveInfo()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetLogSensitiveInfo(), target.GetLogSensitiveInfo()) {
			return false
		}
	}

	return true
}
