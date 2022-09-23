// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-fed/fed.rpc/v1/rt_selector.proto

package v1

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
func (m *SubRouteTableRow) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SubRouteTableRow)
	if !ok {
		that2, ok := that.(SubRouteTableRow)
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

	if strings.Compare(m.GetMatcher(), target.GetMatcher()) != 0 {
		return false
	}

	if strings.Compare(m.GetMatchType(), target.GetMatchType()) != 0 {
		return false
	}

	if len(m.GetMethods()) != len(target.GetMethods()) {
		return false
	}
	for idx, v := range m.GetMethods() {

		if strings.Compare(v, target.GetMethods()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetHeaders()) != len(target.GetHeaders()) {
		return false
	}
	for idx, v := range m.GetHeaders() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetHeaders()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetHeaders()[idx]) {
				return false
			}
		}

	}

	if len(m.GetQueryParameters()) != len(target.GetQueryParameters()) {
		return false
	}
	for idx, v := range m.GetQueryParameters() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetQueryParameters()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetQueryParameters()[idx]) {
				return false
			}
		}

	}

	if len(m.GetRtRoutes()) != len(target.GetRtRoutes()) {
		return false
	}
	for idx, v := range m.GetRtRoutes() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetRtRoutes()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetRtRoutes()[idx]) {
				return false
			}
		}

	}

	switch m.Action.(type) {

	case *SubRouteTableRow_RouteAction:
		if _, ok := target.Action.(*SubRouteTableRow_RouteAction); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRouteAction()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRouteAction()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRouteAction(), target.GetRouteAction()) {
				return false
			}
		}

	case *SubRouteTableRow_RedirectAction:
		if _, ok := target.Action.(*SubRouteTableRow_RedirectAction); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRedirectAction()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRedirectAction()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRedirectAction(), target.GetRedirectAction()) {
				return false
			}
		}

	case *SubRouteTableRow_DirectResponseAction:
		if _, ok := target.Action.(*SubRouteTableRow_DirectResponseAction); !ok {
			return false
		}

		if h, ok := interface{}(m.GetDirectResponseAction()).(equality.Equalizer); ok {
			if !h.Equal(target.GetDirectResponseAction()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetDirectResponseAction(), target.GetDirectResponseAction()) {
				return false
			}
		}

	case *SubRouteTableRow_DelegateAction:
		if _, ok := target.Action.(*SubRouteTableRow_DelegateAction); !ok {
			return false
		}

		if h, ok := interface{}(m.GetDelegateAction()).(equality.Equalizer); ok {
			if !h.Equal(target.GetDelegateAction()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetDelegateAction(), target.GetDelegateAction()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Action != target.Action {
			return false
		}
	}

	return true
}

// Equal function
func (m *GetVirtualServiceRoutesRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetVirtualServiceRoutesRequest)
	if !ok {
		that2, ok := that.(GetVirtualServiceRoutesRequest)
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

	if h, ok := interface{}(m.GetVirtualServiceRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetVirtualServiceRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetVirtualServiceRef(), target.GetVirtualServiceRef()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GetVirtualServiceRoutesResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetVirtualServiceRoutesResponse)
	if !ok {
		that2, ok := that.(GetVirtualServiceRoutesResponse)
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

	if len(m.GetVsRoutes()) != len(target.GetVsRoutes()) {
		return false
	}
	for idx, v := range m.GetVsRoutes() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetVsRoutes()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetVsRoutes()[idx]) {
				return false
			}
		}

	}

	return true
}
