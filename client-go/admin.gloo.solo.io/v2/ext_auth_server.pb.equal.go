// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/admin/v2/ext_auth_server.proto

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
func (m *ExtAuthServerSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtAuthServerSpec)
	if !ok {
		that2, ok := that.(ExtAuthServerSpec)
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

	if h, ok := interface{}(m.GetDestinationServer()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDestinationServer()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDestinationServer(), target.GetDestinationServer()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetHttpService()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHttpService()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHttpService(), target.GetHttpService()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRequestTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRequestTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRequestTimeout(), target.GetRequestTimeout()) {
			return false
		}
	}

	if m.GetFailureModeAllow() != target.GetFailureModeAllow() {
		return false
	}

	if h, ok := interface{}(m.GetRequestBody()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRequestBody()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRequestBody(), target.GetRequestBody()) {
			return false
		}
	}

	if m.GetClearRouteCache() != target.GetClearRouteCache() {
		return false
	}

	if m.GetStatusOnError() != target.GetStatusOnError() {
		return false
	}

	if m.GetTransportApiVersion() != target.GetTransportApiVersion() {
		return false
	}

	if strings.Compare(m.GetStatPrefix(), target.GetStatPrefix()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *ExtAuthServerStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtAuthServerStatus)
	if !ok {
		that2, ok := that.(ExtAuthServerStatus)
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

	if m.GetObservedGeneration() != target.GetObservedGeneration() {
		return false
	}

	if m.GetState() != target.GetState() {
		return false
	}

	return true
}

// Equal function
func (m *ExtAuthServerSpec_HttpService) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtAuthServerSpec_HttpService)
	if !ok {
		that2, ok := that.(ExtAuthServerSpec_HttpService)
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

	if strings.Compare(m.GetPathPrefix(), target.GetPathPrefix()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetRequest()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRequest()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRequest(), target.GetRequest()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetResponse()).(equality.Equalizer); ok {
		if !h.Equal(target.GetResponse()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetResponse(), target.GetResponse()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ExtAuthServerSpec_BufferSettings) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtAuthServerSpec_BufferSettings)
	if !ok {
		that2, ok := that.(ExtAuthServerSpec_BufferSettings)
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

	if m.GetMaxRequestBytes() != target.GetMaxRequestBytes() {
		return false
	}

	if m.GetAllowPartialMessage() != target.GetAllowPartialMessage() {
		return false
	}

	if m.GetPackAsBytes() != target.GetPackAsBytes() {
		return false
	}

	return true
}

// Equal function
func (m *ExtAuthServerSpec_HttpService_Request) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtAuthServerSpec_HttpService_Request)
	if !ok {
		that2, ok := that.(ExtAuthServerSpec_HttpService_Request)
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

	if len(m.GetAllowedHeaders()) != len(target.GetAllowedHeaders()) {
		return false
	}
	for idx, v := range m.GetAllowedHeaders() {

		if strings.Compare(v, target.GetAllowedHeaders()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetHeadersToAdd()) != len(target.GetHeadersToAdd()) {
		return false
	}
	for k, v := range m.GetHeadersToAdd() {

		if strings.Compare(v, target.GetHeadersToAdd()[k]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *ExtAuthServerSpec_HttpService_Response) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtAuthServerSpec_HttpService_Response)
	if !ok {
		that2, ok := that.(ExtAuthServerSpec_HttpService_Response)
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

	if len(m.GetAllowedUpstreamHeaders()) != len(target.GetAllowedUpstreamHeaders()) {
		return false
	}
	for idx, v := range m.GetAllowedUpstreamHeaders() {

		if strings.Compare(v, target.GetAllowedUpstreamHeaders()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetAllowedClientHeaders()) != len(target.GetAllowedClientHeaders()) {
		return false
	}
	for idx, v := range m.GetAllowedClientHeaders() {

		if strings.Compare(v, target.GetAllowedClientHeaders()[idx]) != 0 {
			return false
		}

	}

	return true
}
