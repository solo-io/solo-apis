// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/api/v2/core/health_check.proto

package core

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"

	v3 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3"
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

	_ = v3.RequestMethod(0)
)

// Equal function
func (m *HealthCheck) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HealthCheck)
	if !ok {
		that2, ok := that.(HealthCheck)
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

	if h, ok := interface{}(m.GetTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTimeout(), target.GetTimeout()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetInterval()).(equality.Equalizer); ok {
		if !h.Equal(target.GetInterval()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetInterval(), target.GetInterval()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetInitialJitter()).(equality.Equalizer); ok {
		if !h.Equal(target.GetInitialJitter()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetInitialJitter(), target.GetInitialJitter()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetIntervalJitter()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIntervalJitter()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIntervalJitter(), target.GetIntervalJitter()) {
			return false
		}
	}

	if m.GetIntervalJitterPercent() != target.GetIntervalJitterPercent() {
		return false
	}

	if h, ok := interface{}(m.GetUnhealthyThreshold()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUnhealthyThreshold()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUnhealthyThreshold(), target.GetUnhealthyThreshold()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetHealthyThreshold()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHealthyThreshold()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHealthyThreshold(), target.GetHealthyThreshold()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetReuseConnection()).(equality.Equalizer); ok {
		if !h.Equal(target.GetReuseConnection()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetReuseConnection(), target.GetReuseConnection()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetNoTrafficInterval()).(equality.Equalizer); ok {
		if !h.Equal(target.GetNoTrafficInterval()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetNoTrafficInterval(), target.GetNoTrafficInterval()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetUnhealthyInterval()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUnhealthyInterval()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUnhealthyInterval(), target.GetUnhealthyInterval()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetUnhealthyEdgeInterval()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUnhealthyEdgeInterval()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUnhealthyEdgeInterval(), target.GetUnhealthyEdgeInterval()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetHealthyEdgeInterval()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHealthyEdgeInterval()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHealthyEdgeInterval(), target.GetHealthyEdgeInterval()) {
			return false
		}
	}

	if strings.Compare(m.GetEventLogPath(), target.GetEventLogPath()) != 0 {
		return false
	}

	if m.GetAlwaysLogHealthCheckFailures() != target.GetAlwaysLogHealthCheckFailures() {
		return false
	}

	switch m.HealthChecker.(type) {

	case *HealthCheck_HttpHealthCheck_:
		if _, ok := target.HealthChecker.(*HealthCheck_HttpHealthCheck_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetHttpHealthCheck()).(equality.Equalizer); ok {
			if !h.Equal(target.GetHttpHealthCheck()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetHttpHealthCheck(), target.GetHttpHealthCheck()) {
				return false
			}
		}

	case *HealthCheck_TcpHealthCheck_:
		if _, ok := target.HealthChecker.(*HealthCheck_TcpHealthCheck_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetTcpHealthCheck()).(equality.Equalizer); ok {
			if !h.Equal(target.GetTcpHealthCheck()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetTcpHealthCheck(), target.GetTcpHealthCheck()) {
				return false
			}
		}

	case *HealthCheck_GrpcHealthCheck_:
		if _, ok := target.HealthChecker.(*HealthCheck_GrpcHealthCheck_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGrpcHealthCheck()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGrpcHealthCheck()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGrpcHealthCheck(), target.GetGrpcHealthCheck()) {
				return false
			}
		}

	case *HealthCheck_CustomHealthCheck_:
		if _, ok := target.HealthChecker.(*HealthCheck_CustomHealthCheck_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetCustomHealthCheck()).(equality.Equalizer); ok {
			if !h.Equal(target.GetCustomHealthCheck()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetCustomHealthCheck(), target.GetCustomHealthCheck()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.HealthChecker != target.HealthChecker {
			return false
		}
	}

	return true
}

// Equal function
func (m *HealthCheck_Payload) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HealthCheck_Payload)
	if !ok {
		that2, ok := that.(HealthCheck_Payload)
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

	switch m.Payload.(type) {

	case *HealthCheck_Payload_Text:
		if _, ok := target.Payload.(*HealthCheck_Payload_Text); !ok {
			return false
		}

		if strings.Compare(m.GetText(), target.GetText()) != 0 {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.Payload != target.Payload {
			return false
		}
	}

	return true
}

// Equal function
func (m *HealthCheck_HttpHealthCheck) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HealthCheck_HttpHealthCheck)
	if !ok {
		that2, ok := that.(HealthCheck_HttpHealthCheck)
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

	if strings.Compare(m.GetHost(), target.GetHost()) != 0 {
		return false
	}

	if strings.Compare(m.GetPath(), target.GetPath()) != 0 {
		return false
	}

	if strings.Compare(m.GetServiceName(), target.GetServiceName()) != 0 {
		return false
	}

	if len(m.GetRequestHeadersToAdd()) != len(target.GetRequestHeadersToAdd()) {
		return false
	}
	for idx, v := range m.GetRequestHeadersToAdd() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetRequestHeadersToAdd()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetRequestHeadersToAdd()[idx]) {
				return false
			}
		}

	}

	if len(m.GetRequestHeadersToRemove()) != len(target.GetRequestHeadersToRemove()) {
		return false
	}
	for idx, v := range m.GetRequestHeadersToRemove() {

		if strings.Compare(v, target.GetRequestHeadersToRemove()[idx]) != 0 {
			return false
		}

	}

	if m.GetUseHttp2() != target.GetUseHttp2() {
		return false
	}

	if len(m.GetExpectedStatuses()) != len(target.GetExpectedStatuses()) {
		return false
	}
	for idx, v := range m.GetExpectedStatuses() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetExpectedStatuses()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetExpectedStatuses()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetResponseAssertions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetResponseAssertions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetResponseAssertions(), target.GetResponseAssertions()) {
			return false
		}
	}

	if m.GetMethod() != target.GetMethod() {
		return false
	}

	return true
}

// Equal function
func (m *HealthCheck_TcpHealthCheck) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HealthCheck_TcpHealthCheck)
	if !ok {
		that2, ok := that.(HealthCheck_TcpHealthCheck)
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

	if h, ok := interface{}(m.GetSend()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSend()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSend(), target.GetSend()) {
			return false
		}
	}

	if len(m.GetReceive()) != len(target.GetReceive()) {
		return false
	}
	for idx, v := range m.GetReceive() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetReceive()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetReceive()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *HealthCheck_RedisHealthCheck) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HealthCheck_RedisHealthCheck)
	if !ok {
		that2, ok := that.(HealthCheck_RedisHealthCheck)
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

	return true
}

// Equal function
func (m *HealthCheck_GrpcHealthCheck) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HealthCheck_GrpcHealthCheck)
	if !ok {
		that2, ok := that.(HealthCheck_GrpcHealthCheck)
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

	if strings.Compare(m.GetServiceName(), target.GetServiceName()) != 0 {
		return false
	}

	if strings.Compare(m.GetAuthority(), target.GetAuthority()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *HealthCheck_CustomHealthCheck) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HealthCheck_CustomHealthCheck)
	if !ok {
		that2, ok := that.(HealthCheck_CustomHealthCheck)
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

	switch m.ConfigType.(type) {

	case *HealthCheck_CustomHealthCheck_Config:
		if _, ok := target.ConfigType.(*HealthCheck_CustomHealthCheck_Config); !ok {
			return false
		}

		if h, ok := interface{}(m.GetConfig()).(equality.Equalizer); ok {
			if !h.Equal(target.GetConfig()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetConfig(), target.GetConfig()) {
				return false
			}
		}

	case *HealthCheck_CustomHealthCheck_TypedConfig:
		if _, ok := target.ConfigType.(*HealthCheck_CustomHealthCheck_TypedConfig); !ok {
			return false
		}

		if h, ok := interface{}(m.GetTypedConfig()).(equality.Equalizer); ok {
			if !h.Equal(target.GetTypedConfig()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetTypedConfig(), target.GetTypedConfig()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.ConfigType != target.ConfigType {
			return false
		}
	}

	return true
}
