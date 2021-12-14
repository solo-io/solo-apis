// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/static/static.proto

package static

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"

	github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1_options "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options"
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
func (m *UpstreamSpec) Clone() proto.Message {
	var target *UpstreamSpec
	if m == nil {
		return target
	}
	target = &UpstreamSpec{}

	if m.GetHosts() != nil {
		target.Hosts = make([]*Host, len(m.GetHosts()))
		for idx, v := range m.GetHosts() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Hosts[idx] = h.Clone().(*Host)
			} else {
				target.Hosts[idx] = proto.Clone(v).(*Host)
			}

		}
	}

	target.UseTls = m.GetUseTls()

	if h, ok := interface{}(m.GetServiceSpec()).(clone.Cloner); ok {
		target.ServiceSpec = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1_options.ServiceSpec)
	} else {
		target.ServiceSpec = proto.Clone(m.GetServiceSpec()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1_options.ServiceSpec)
	}

	if h, ok := interface{}(m.GetAutoSniRewrite()).(clone.Cloner); ok {
		target.AutoSniRewrite = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.AutoSniRewrite = proto.Clone(m.GetAutoSniRewrite()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	return target
}

// Clone function
func (m *Host) Clone() proto.Message {
	var target *Host
	if m == nil {
		return target
	}
	target = &Host{}

	target.Addr = m.GetAddr()

	target.Port = m.GetPort()

	target.SniAddr = m.GetSniAddr()

	if h, ok := interface{}(m.GetLoadBalancingWeight()).(clone.Cloner); ok {
		target.LoadBalancingWeight = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.LoadBalancingWeight = proto.Clone(m.GetLoadBalancingWeight()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if h, ok := interface{}(m.GetHealthCheckConfig()).(clone.Cloner); ok {
		target.HealthCheckConfig = h.Clone().(*Host_HealthCheckConfig)
	} else {
		target.HealthCheckConfig = proto.Clone(m.GetHealthCheckConfig()).(*Host_HealthCheckConfig)
	}

	return target
}

// Clone function
func (m *Host_HealthCheckConfig) Clone() proto.Message {
	var target *Host_HealthCheckConfig
	if m == nil {
		return target
	}
	target = &Host_HealthCheckConfig{}

	target.Path = m.GetPath()

	target.Method = m.GetMethod()

	return target
}
