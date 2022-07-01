// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/connection.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_duration "github.com/golang/protobuf/ptypes/duration"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"

	github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1_options_protocol "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/protocol"
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
func (m *ConnectionConfig) Clone() proto.Message {
	var target *ConnectionConfig
	if m == nil {
		return target
	}
	target = &ConnectionConfig{}

	target.MaxRequestsPerConnection = m.GetMaxRequestsPerConnection()

	if h, ok := interface{}(m.GetConnectTimeout()).(clone.Cloner); ok {
		target.ConnectTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.ConnectTimeout = proto.Clone(m.GetConnectTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetTcpKeepalive()).(clone.Cloner); ok {
		target.TcpKeepalive = h.Clone().(*ConnectionConfig_TcpKeepAlive)
	} else {
		target.TcpKeepalive = proto.Clone(m.GetTcpKeepalive()).(*ConnectionConfig_TcpKeepAlive)
	}

	if h, ok := interface{}(m.GetPerConnectionBufferLimitBytes()).(clone.Cloner); ok {
		target.PerConnectionBufferLimitBytes = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.PerConnectionBufferLimitBytes = proto.Clone(m.GetPerConnectionBufferLimitBytes()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if h, ok := interface{}(m.GetCommonHttpProtocolOptions()).(clone.Cloner); ok {
		target.CommonHttpProtocolOptions = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1_options_protocol.HttpProtocolOptions)
	} else {
		target.CommonHttpProtocolOptions = proto.Clone(m.GetCommonHttpProtocolOptions()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1_options_protocol.HttpProtocolOptions)
	}

	if h, ok := interface{}(m.GetHttp1ProtocolOptions()).(clone.Cloner); ok {
		target.Http1ProtocolOptions = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1_options_protocol.Http1ProtocolOptions)
	} else {
		target.Http1ProtocolOptions = proto.Clone(m.GetHttp1ProtocolOptions()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1_options_protocol.Http1ProtocolOptions)
	}

	return target
}

// Clone function
func (m *ConnectionConfig_TcpKeepAlive) Clone() proto.Message {
	var target *ConnectionConfig_TcpKeepAlive
	if m == nil {
		return target
	}
	target = &ConnectionConfig_TcpKeepAlive{}

	target.KeepaliveProbes = m.GetKeepaliveProbes()

	if h, ok := interface{}(m.GetKeepaliveTime()).(clone.Cloner); ok {
		target.KeepaliveTime = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.KeepaliveTime = proto.Clone(m.GetKeepaliveTime()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetKeepaliveInterval()).(clone.Cloner); ok {
		target.KeepaliveInterval = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.KeepaliveInterval = proto.Clone(m.GetKeepaliveInterval()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	return target
}
