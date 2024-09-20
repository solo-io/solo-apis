// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/tcp/tcp.proto

package tcp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	google_golang_org_protobuf_types_known_durationpb "google.golang.org/protobuf/types/known/durationpb"

	google_golang_org_protobuf_types_known_wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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
func (m *TcpProxySettings) Clone() proto.Message {
	var target *TcpProxySettings
	if m == nil {
		return target
	}
	target = &TcpProxySettings{}

	if h, ok := interface{}(m.GetMaxConnectAttempts()).(clone.Cloner); ok {
		target.MaxConnectAttempts = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.MaxConnectAttempts = proto.Clone(m.GetMaxConnectAttempts()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	if h, ok := interface{}(m.GetIdleTimeout()).(clone.Cloner); ok {
		target.IdleTimeout = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.IdleTimeout = proto.Clone(m.GetIdleTimeout()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetTunnelingConfig()).(clone.Cloner); ok {
		target.TunnelingConfig = h.Clone().(*TcpProxySettings_TunnelingConfig)
	} else {
		target.TunnelingConfig = proto.Clone(m.GetTunnelingConfig()).(*TcpProxySettings_TunnelingConfig)
	}

	if h, ok := interface{}(m.GetAccessLogFlushInterval()).(clone.Cloner); ok {
		target.AccessLogFlushInterval = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.AccessLogFlushInterval = proto.Clone(m.GetAccessLogFlushInterval()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	return target
}

// Clone function
func (m *HeaderValueOption) Clone() proto.Message {
	var target *HeaderValueOption
	if m == nil {
		return target
	}
	target = &HeaderValueOption{}

	if h, ok := interface{}(m.GetHeader()).(clone.Cloner); ok {
		target.Header = h.Clone().(*HeaderValue)
	} else {
		target.Header = proto.Clone(m.GetHeader()).(*HeaderValue)
	}

	if h, ok := interface{}(m.GetAppend()).(clone.Cloner); ok {
		target.Append = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.Append = proto.Clone(m.GetAppend()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	return target
}

// Clone function
func (m *HeaderValue) Clone() proto.Message {
	var target *HeaderValue
	if m == nil {
		return target
	}
	target = &HeaderValue{}

	target.Key = m.GetKey()

	target.Value = m.GetValue()

	return target
}

// Clone function
func (m *TcpProxySettings_TunnelingConfig) Clone() proto.Message {
	var target *TcpProxySettings_TunnelingConfig
	if m == nil {
		return target
	}
	target = &TcpProxySettings_TunnelingConfig{}

	target.Hostname = m.GetHostname()

	if m.GetHeadersToAdd() != nil {
		target.HeadersToAdd = make([]*HeaderValueOption, len(m.GetHeadersToAdd()))
		for idx, v := range m.GetHeadersToAdd() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.HeadersToAdd[idx] = h.Clone().(*HeaderValueOption)
			} else {
				target.HeadersToAdd[idx] = proto.Clone(v).(*HeaderValueOption)
			}

		}
	}

	return target
}
