// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/tracing/tracing.proto

package tracing

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"

	github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_trace_v3 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/trace/v3"
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
func (m *ListenerTracingSettings) Clone() proto.Message {
	var target *ListenerTracingSettings
	if m == nil {
		return target
	}
	target = &ListenerTracingSettings{}

	if m.GetRequestHeadersForTags() != nil {
		target.RequestHeadersForTags = make([]string, len(m.GetRequestHeadersForTags()))
		for idx, v := range m.GetRequestHeadersForTags() {

			target.RequestHeadersForTags[idx] = v

		}
	}

	target.Verbose = m.GetVerbose()

	if h, ok := interface{}(m.GetTracePercentages()).(clone.Cloner); ok {
		target.TracePercentages = h.Clone().(*TracePercentages)
	} else {
		target.TracePercentages = proto.Clone(m.GetTracePercentages()).(*TracePercentages)
	}

	if m.GetEnvironmentVariablesForTags() != nil {
		target.EnvironmentVariablesForTags = make([]*TracingTagEnvironmentVariable, len(m.GetEnvironmentVariablesForTags()))
		for idx, v := range m.GetEnvironmentVariablesForTags() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.EnvironmentVariablesForTags[idx] = h.Clone().(*TracingTagEnvironmentVariable)
			} else {
				target.EnvironmentVariablesForTags[idx] = proto.Clone(v).(*TracingTagEnvironmentVariable)
			}

		}
	}

	if m.GetLiteralsForTags() != nil {
		target.LiteralsForTags = make([]*TracingTagLiteral, len(m.GetLiteralsForTags()))
		for idx, v := range m.GetLiteralsForTags() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.LiteralsForTags[idx] = h.Clone().(*TracingTagLiteral)
			} else {
				target.LiteralsForTags[idx] = proto.Clone(v).(*TracingTagLiteral)
			}

		}
	}

	switch m.ProviderConfig.(type) {

	case *ListenerTracingSettings_ZipkinConfig:

		if h, ok := interface{}(m.GetZipkinConfig()).(clone.Cloner); ok {
			target.ProviderConfig = &ListenerTracingSettings_ZipkinConfig{
				ZipkinConfig: h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_trace_v3.ZipkinConfig),
			}
		} else {
			target.ProviderConfig = &ListenerTracingSettings_ZipkinConfig{
				ZipkinConfig: proto.Clone(m.GetZipkinConfig()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_trace_v3.ZipkinConfig),
			}
		}

	case *ListenerTracingSettings_DatadogConfig:

		if h, ok := interface{}(m.GetDatadogConfig()).(clone.Cloner); ok {
			target.ProviderConfig = &ListenerTracingSettings_DatadogConfig{
				DatadogConfig: h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_trace_v3.DatadogConfig),
			}
		} else {
			target.ProviderConfig = &ListenerTracingSettings_DatadogConfig{
				DatadogConfig: proto.Clone(m.GetDatadogConfig()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_trace_v3.DatadogConfig),
			}
		}

	}

	return target
}

// Clone function
func (m *RouteTracingSettings) Clone() proto.Message {
	var target *RouteTracingSettings
	if m == nil {
		return target
	}
	target = &RouteTracingSettings{}

	target.RouteDescriptor = m.GetRouteDescriptor()

	if h, ok := interface{}(m.GetTracePercentages()).(clone.Cloner); ok {
		target.TracePercentages = h.Clone().(*TracePercentages)
	} else {
		target.TracePercentages = proto.Clone(m.GetTracePercentages()).(*TracePercentages)
	}

	if h, ok := interface{}(m.GetPropagate()).(clone.Cloner); ok {
		target.Propagate = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.Propagate = proto.Clone(m.GetPropagate()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	return target
}

// Clone function
func (m *TracePercentages) Clone() proto.Message {
	var target *TracePercentages
	if m == nil {
		return target
	}
	target = &TracePercentages{}

	if h, ok := interface{}(m.GetClientSamplePercentage()).(clone.Cloner); ok {
		target.ClientSamplePercentage = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.FloatValue)
	} else {
		target.ClientSamplePercentage = proto.Clone(m.GetClientSamplePercentage()).(*github_com_golang_protobuf_ptypes_wrappers.FloatValue)
	}

	if h, ok := interface{}(m.GetRandomSamplePercentage()).(clone.Cloner); ok {
		target.RandomSamplePercentage = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.FloatValue)
	} else {
		target.RandomSamplePercentage = proto.Clone(m.GetRandomSamplePercentage()).(*github_com_golang_protobuf_ptypes_wrappers.FloatValue)
	}

	if h, ok := interface{}(m.GetOverallSamplePercentage()).(clone.Cloner); ok {
		target.OverallSamplePercentage = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.FloatValue)
	} else {
		target.OverallSamplePercentage = proto.Clone(m.GetOverallSamplePercentage()).(*github_com_golang_protobuf_ptypes_wrappers.FloatValue)
	}

	return target
}

// Clone function
func (m *TracingTagEnvironmentVariable) Clone() proto.Message {
	var target *TracingTagEnvironmentVariable
	if m == nil {
		return target
	}
	target = &TracingTagEnvironmentVariable{}

	target.Tag = m.GetTag()

	target.Name = m.GetName()

	target.DefaultValue = m.GetDefaultValue()

	return target
}

// Clone function
func (m *TracingTagLiteral) Clone() proto.Message {
	var target *TracingTagLiteral
	if m == nil {
		return target
	}
	target = &TracingTagLiteral{}

	target.Tag = m.GetTag()

	target.Value = m.GetValue()

	return target
}
