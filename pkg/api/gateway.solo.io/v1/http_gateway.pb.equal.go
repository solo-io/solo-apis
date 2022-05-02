// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gateway/v1/http_gateway.proto

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
func (m *HttpGatewaySpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HttpGatewaySpec)
	if !ok {
		that2, ok := that.(HttpGatewaySpec)
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

	if h, ok := interface{}(m.GetMatcher()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMatcher()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMatcher(), target.GetMatcher()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetHttpGateway()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHttpGateway()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHttpGateway(), target.GetHttpGateway()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *HttpGateway) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HttpGateway)
	if !ok {
		that2, ok := that.(HttpGateway)
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

	if len(m.GetVirtualServices()) != len(target.GetVirtualServices()) {
		return false
	}
	for idx, v := range m.GetVirtualServices() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetVirtualServices()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetVirtualServices()[idx]) {
				return false
			}
		}

	}

	if len(m.GetVirtualServiceSelector()) != len(target.GetVirtualServiceSelector()) {
		return false
	}
	for k, v := range m.GetVirtualServiceSelector() {

		if strings.Compare(v, target.GetVirtualServiceSelector()[k]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetVirtualServiceExpressions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetVirtualServiceExpressions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetVirtualServiceExpressions(), target.GetVirtualServiceExpressions()) {
			return false
		}
	}

	if len(m.GetVirtualServiceNamespaces()) != len(target.GetVirtualServiceNamespaces()) {
		return false
	}
	for idx, v := range m.GetVirtualServiceNamespaces() {

		if strings.Compare(v, target.GetVirtualServiceNamespaces()[idx]) != 0 {
			return false
		}

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

	return true
}

// Equal function
func (m *VirtualServiceSelectorExpressions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualServiceSelectorExpressions)
	if !ok {
		that2, ok := that.(VirtualServiceSelectorExpressions)
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

	if len(m.GetExpressions()) != len(target.GetExpressions()) {
		return false
	}
	for idx, v := range m.GetExpressions() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetExpressions()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetExpressions()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *HttpGatewayStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HttpGatewayStatus)
	if !ok {
		that2, ok := that.(HttpGatewayStatus)
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

	if m.GetState() != target.GetState() {
		return false
	}

	if strings.Compare(m.GetReason(), target.GetReason()) != 0 {
		return false
	}

	if strings.Compare(m.GetReportedBy(), target.GetReportedBy()) != 0 {
		return false
	}

	if len(m.GetSubresourceStatuses()) != len(target.GetSubresourceStatuses()) {
		return false
	}
	for k, v := range m.GetSubresourceStatuses() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSubresourceStatuses()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSubresourceStatuses()[k]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetDetails()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDetails()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDetails(), target.GetDetails()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *HttpGatewayNamespacedStatuses) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HttpGatewayNamespacedStatuses)
	if !ok {
		that2, ok := that.(HttpGatewayNamespacedStatuses)
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

	if len(m.GetStatuses()) != len(target.GetStatuses()) {
		return false
	}
	for k, v := range m.GetStatuses() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetStatuses()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetStatuses()[k]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *HttpGatewaySpec_Matcher) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HttpGatewaySpec_Matcher)
	if !ok {
		that2, ok := that.(HttpGatewaySpec_Matcher)
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

	if len(m.GetSourcePrefixRanges()) != len(target.GetSourcePrefixRanges()) {
		return false
	}
	for idx, v := range m.GetSourcePrefixRanges() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSourcePrefixRanges()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSourcePrefixRanges()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetSslConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSslConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSslConfig(), target.GetSslConfig()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualServiceSelectorExpressions_Expression) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualServiceSelectorExpressions_Expression)
	if !ok {
		that2, ok := that.(VirtualServiceSelectorExpressions_Expression)
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

	if strings.Compare(m.GetKey(), target.GetKey()) != 0 {
		return false
	}

	if m.GetOperator() != target.GetOperator() {
		return false
	}

	if len(m.GetValues()) != len(target.GetValues()) {
		return false
	}
	for idx, v := range m.GetValues() {

		if strings.Compare(v, target.GetValues()[idx]) != 0 {
			return false
		}

	}

	return true
}
