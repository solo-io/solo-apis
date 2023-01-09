// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/als/als.proto

package als

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_struct "github.com/golang/protobuf/ptypes/struct"

	github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_core_v3 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3"

	github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_route_v3 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/route/v3"

	github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_type_v3 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/type/v3"
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
func (m *AccessLoggingService) Clone() proto.Message {
	var target *AccessLoggingService
	if m == nil {
		return target
	}
	target = &AccessLoggingService{}

	if m.GetAccessLog() != nil {
		target.AccessLog = make([]*AccessLog, len(m.GetAccessLog()))
		for idx, v := range m.GetAccessLog() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.AccessLog[idx] = h.Clone().(*AccessLog)
			} else {
				target.AccessLog[idx] = proto.Clone(v).(*AccessLog)
			}

		}
	}

	return target
}

// Clone function
func (m *AccessLog) Clone() proto.Message {
	var target *AccessLog
	if m == nil {
		return target
	}
	target = &AccessLog{}

	if h, ok := interface{}(m.GetFilter()).(clone.Cloner); ok {
		target.Filter = h.Clone().(*AccessLogFilter)
	} else {
		target.Filter = proto.Clone(m.GetFilter()).(*AccessLogFilter)
	}

	switch m.OutputDestination.(type) {

	case *AccessLog_FileSink:

		if h, ok := interface{}(m.GetFileSink()).(clone.Cloner); ok {
			target.OutputDestination = &AccessLog_FileSink{
				FileSink: h.Clone().(*FileSink),
			}
		} else {
			target.OutputDestination = &AccessLog_FileSink{
				FileSink: proto.Clone(m.GetFileSink()).(*FileSink),
			}
		}

	case *AccessLog_GrpcService:

		if h, ok := interface{}(m.GetGrpcService()).(clone.Cloner); ok {
			target.OutputDestination = &AccessLog_GrpcService{
				GrpcService: h.Clone().(*GrpcService),
			}
		} else {
			target.OutputDestination = &AccessLog_GrpcService{
				GrpcService: proto.Clone(m.GetGrpcService()).(*GrpcService),
			}
		}

	}

	return target
}

// Clone function
func (m *FileSink) Clone() proto.Message {
	var target *FileSink
	if m == nil {
		return target
	}
	target = &FileSink{}

	target.Path = m.GetPath()

	switch m.OutputFormat.(type) {

	case *FileSink_StringFormat:

		target.OutputFormat = &FileSink_StringFormat{
			StringFormat: m.GetStringFormat(),
		}

	case *FileSink_JsonFormat:

		if h, ok := interface{}(m.GetJsonFormat()).(clone.Cloner); ok {
			target.OutputFormat = &FileSink_JsonFormat{
				JsonFormat: h.Clone().(*github_com_golang_protobuf_ptypes_struct.Struct),
			}
		} else {
			target.OutputFormat = &FileSink_JsonFormat{
				JsonFormat: proto.Clone(m.GetJsonFormat()).(*github_com_golang_protobuf_ptypes_struct.Struct),
			}
		}

	}

	return target
}

// Clone function
func (m *GrpcService) Clone() proto.Message {
	var target *GrpcService
	if m == nil {
		return target
	}
	target = &GrpcService{}

	target.LogName = m.GetLogName()

	if m.GetAdditionalRequestHeadersToLog() != nil {
		target.AdditionalRequestHeadersToLog = make([]string, len(m.GetAdditionalRequestHeadersToLog()))
		for idx, v := range m.GetAdditionalRequestHeadersToLog() {

			target.AdditionalRequestHeadersToLog[idx] = v

		}
	}

	if m.GetAdditionalResponseHeadersToLog() != nil {
		target.AdditionalResponseHeadersToLog = make([]string, len(m.GetAdditionalResponseHeadersToLog()))
		for idx, v := range m.GetAdditionalResponseHeadersToLog() {

			target.AdditionalResponseHeadersToLog[idx] = v

		}
	}

	if m.GetAdditionalResponseTrailersToLog() != nil {
		target.AdditionalResponseTrailersToLog = make([]string, len(m.GetAdditionalResponseTrailersToLog()))
		for idx, v := range m.GetAdditionalResponseTrailersToLog() {

			target.AdditionalResponseTrailersToLog[idx] = v

		}
	}

	switch m.ServiceRef.(type) {

	case *GrpcService_StaticClusterName:

		target.ServiceRef = &GrpcService_StaticClusterName{
			StaticClusterName: m.GetStaticClusterName(),
		}

	}

	return target
}

// Clone function
func (m *AccessLogFilter) Clone() proto.Message {
	var target *AccessLogFilter
	if m == nil {
		return target
	}
	target = &AccessLogFilter{}

	switch m.FilterSpecifier.(type) {

	case *AccessLogFilter_StatusCodeFilter:

		if h, ok := interface{}(m.GetStatusCodeFilter()).(clone.Cloner); ok {
			target.FilterSpecifier = &AccessLogFilter_StatusCodeFilter{
				StatusCodeFilter: h.Clone().(*StatusCodeFilter),
			}
		} else {
			target.FilterSpecifier = &AccessLogFilter_StatusCodeFilter{
				StatusCodeFilter: proto.Clone(m.GetStatusCodeFilter()).(*StatusCodeFilter),
			}
		}

	case *AccessLogFilter_DurationFilter:

		if h, ok := interface{}(m.GetDurationFilter()).(clone.Cloner); ok {
			target.FilterSpecifier = &AccessLogFilter_DurationFilter{
				DurationFilter: h.Clone().(*DurationFilter),
			}
		} else {
			target.FilterSpecifier = &AccessLogFilter_DurationFilter{
				DurationFilter: proto.Clone(m.GetDurationFilter()).(*DurationFilter),
			}
		}

	case *AccessLogFilter_NotHealthCheckFilter:

		if h, ok := interface{}(m.GetNotHealthCheckFilter()).(clone.Cloner); ok {
			target.FilterSpecifier = &AccessLogFilter_NotHealthCheckFilter{
				NotHealthCheckFilter: h.Clone().(*NotHealthCheckFilter),
			}
		} else {
			target.FilterSpecifier = &AccessLogFilter_NotHealthCheckFilter{
				NotHealthCheckFilter: proto.Clone(m.GetNotHealthCheckFilter()).(*NotHealthCheckFilter),
			}
		}

	case *AccessLogFilter_TraceableFilter:

		if h, ok := interface{}(m.GetTraceableFilter()).(clone.Cloner); ok {
			target.FilterSpecifier = &AccessLogFilter_TraceableFilter{
				TraceableFilter: h.Clone().(*TraceableFilter),
			}
		} else {
			target.FilterSpecifier = &AccessLogFilter_TraceableFilter{
				TraceableFilter: proto.Clone(m.GetTraceableFilter()).(*TraceableFilter),
			}
		}

	case *AccessLogFilter_RuntimeFilter:

		if h, ok := interface{}(m.GetRuntimeFilter()).(clone.Cloner); ok {
			target.FilterSpecifier = &AccessLogFilter_RuntimeFilter{
				RuntimeFilter: h.Clone().(*RuntimeFilter),
			}
		} else {
			target.FilterSpecifier = &AccessLogFilter_RuntimeFilter{
				RuntimeFilter: proto.Clone(m.GetRuntimeFilter()).(*RuntimeFilter),
			}
		}

	case *AccessLogFilter_AndFilter:

		if h, ok := interface{}(m.GetAndFilter()).(clone.Cloner); ok {
			target.FilterSpecifier = &AccessLogFilter_AndFilter{
				AndFilter: h.Clone().(*AndFilter),
			}
		} else {
			target.FilterSpecifier = &AccessLogFilter_AndFilter{
				AndFilter: proto.Clone(m.GetAndFilter()).(*AndFilter),
			}
		}

	case *AccessLogFilter_OrFilter:

		if h, ok := interface{}(m.GetOrFilter()).(clone.Cloner); ok {
			target.FilterSpecifier = &AccessLogFilter_OrFilter{
				OrFilter: h.Clone().(*OrFilter),
			}
		} else {
			target.FilterSpecifier = &AccessLogFilter_OrFilter{
				OrFilter: proto.Clone(m.GetOrFilter()).(*OrFilter),
			}
		}

	case *AccessLogFilter_HeaderFilter:

		if h, ok := interface{}(m.GetHeaderFilter()).(clone.Cloner); ok {
			target.FilterSpecifier = &AccessLogFilter_HeaderFilter{
				HeaderFilter: h.Clone().(*HeaderFilter),
			}
		} else {
			target.FilterSpecifier = &AccessLogFilter_HeaderFilter{
				HeaderFilter: proto.Clone(m.GetHeaderFilter()).(*HeaderFilter),
			}
		}

	case *AccessLogFilter_ResponseFlagFilter:

		if h, ok := interface{}(m.GetResponseFlagFilter()).(clone.Cloner); ok {
			target.FilterSpecifier = &AccessLogFilter_ResponseFlagFilter{
				ResponseFlagFilter: h.Clone().(*ResponseFlagFilter),
			}
		} else {
			target.FilterSpecifier = &AccessLogFilter_ResponseFlagFilter{
				ResponseFlagFilter: proto.Clone(m.GetResponseFlagFilter()).(*ResponseFlagFilter),
			}
		}

	case *AccessLogFilter_GrpcStatusFilter:

		if h, ok := interface{}(m.GetGrpcStatusFilter()).(clone.Cloner); ok {
			target.FilterSpecifier = &AccessLogFilter_GrpcStatusFilter{
				GrpcStatusFilter: h.Clone().(*GrpcStatusFilter),
			}
		} else {
			target.FilterSpecifier = &AccessLogFilter_GrpcStatusFilter{
				GrpcStatusFilter: proto.Clone(m.GetGrpcStatusFilter()).(*GrpcStatusFilter),
			}
		}

	}

	return target
}

// Clone function
func (m *ComparisonFilter) Clone() proto.Message {
	var target *ComparisonFilter
	if m == nil {
		return target
	}
	target = &ComparisonFilter{}

	target.Op = m.GetOp()

	if h, ok := interface{}(m.GetValue()).(clone.Cloner); ok {
		target.Value = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_core_v3.RuntimeUInt32)
	} else {
		target.Value = proto.Clone(m.GetValue()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_core_v3.RuntimeUInt32)
	}

	return target
}

// Clone function
func (m *StatusCodeFilter) Clone() proto.Message {
	var target *StatusCodeFilter
	if m == nil {
		return target
	}
	target = &StatusCodeFilter{}

	if h, ok := interface{}(m.GetComparison()).(clone.Cloner); ok {
		target.Comparison = h.Clone().(*ComparisonFilter)
	} else {
		target.Comparison = proto.Clone(m.GetComparison()).(*ComparisonFilter)
	}

	return target
}

// Clone function
func (m *DurationFilter) Clone() proto.Message {
	var target *DurationFilter
	if m == nil {
		return target
	}
	target = &DurationFilter{}

	if h, ok := interface{}(m.GetComparison()).(clone.Cloner); ok {
		target.Comparison = h.Clone().(*ComparisonFilter)
	} else {
		target.Comparison = proto.Clone(m.GetComparison()).(*ComparisonFilter)
	}

	return target
}

// Clone function
func (m *NotHealthCheckFilter) Clone() proto.Message {
	var target *NotHealthCheckFilter
	if m == nil {
		return target
	}
	target = &NotHealthCheckFilter{}

	return target
}

// Clone function
func (m *TraceableFilter) Clone() proto.Message {
	var target *TraceableFilter
	if m == nil {
		return target
	}
	target = &TraceableFilter{}

	return target
}

// Clone function
func (m *RuntimeFilter) Clone() proto.Message {
	var target *RuntimeFilter
	if m == nil {
		return target
	}
	target = &RuntimeFilter{}

	target.RuntimeKey = m.GetRuntimeKey()

	if h, ok := interface{}(m.GetPercentSampled()).(clone.Cloner); ok {
		target.PercentSampled = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_type_v3.FractionalPercent)
	} else {
		target.PercentSampled = proto.Clone(m.GetPercentSampled()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_type_v3.FractionalPercent)
	}

	target.UseIndependentRandomness = m.GetUseIndependentRandomness()

	return target
}

// Clone function
func (m *AndFilter) Clone() proto.Message {
	var target *AndFilter
	if m == nil {
		return target
	}
	target = &AndFilter{}

	if m.GetFilters() != nil {
		target.Filters = make([]*AccessLogFilter, len(m.GetFilters()))
		for idx, v := range m.GetFilters() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Filters[idx] = h.Clone().(*AccessLogFilter)
			} else {
				target.Filters[idx] = proto.Clone(v).(*AccessLogFilter)
			}

		}
	}

	return target
}

// Clone function
func (m *OrFilter) Clone() proto.Message {
	var target *OrFilter
	if m == nil {
		return target
	}
	target = &OrFilter{}

	if m.GetFilters() != nil {
		target.Filters = make([]*AccessLogFilter, len(m.GetFilters()))
		for idx, v := range m.GetFilters() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Filters[idx] = h.Clone().(*AccessLogFilter)
			} else {
				target.Filters[idx] = proto.Clone(v).(*AccessLogFilter)
			}

		}
	}

	return target
}

// Clone function
func (m *HeaderFilter) Clone() proto.Message {
	var target *HeaderFilter
	if m == nil {
		return target
	}
	target = &HeaderFilter{}

	if h, ok := interface{}(m.GetHeader()).(clone.Cloner); ok {
		target.Header = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_route_v3.HeaderMatcher)
	} else {
		target.Header = proto.Clone(m.GetHeader()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_route_v3.HeaderMatcher)
	}

	return target
}

// Clone function
func (m *ResponseFlagFilter) Clone() proto.Message {
	var target *ResponseFlagFilter
	if m == nil {
		return target
	}
	target = &ResponseFlagFilter{}

	if m.GetFlags() != nil {
		target.Flags = make([]string, len(m.GetFlags()))
		for idx, v := range m.GetFlags() {

			target.Flags[idx] = v

		}
	}

	return target
}

// Clone function
func (m *GrpcStatusFilter) Clone() proto.Message {
	var target *GrpcStatusFilter
	if m == nil {
		return target
	}
	target = &GrpcStatusFilter{}

	if m.GetStatuses() != nil {
		target.Statuses = make([]GrpcStatusFilter_Status, len(m.GetStatuses()))
		for idx, v := range m.GetStatuses() {

			target.Statuses[idx] = v

		}
	}

	target.Exclude = m.GetExclude()

	return target
}
