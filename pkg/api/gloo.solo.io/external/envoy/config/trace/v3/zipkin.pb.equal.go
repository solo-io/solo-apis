// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/config/trace/v3/zipkin.proto

package v3

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
func (m *ZipkinConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ZipkinConfig)
	if !ok {
		that2, ok := that.(ZipkinConfig)
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

	if strings.Compare(m.GetCollectorEndpoint(), target.GetCollectorEndpoint()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetTraceId_128Bit()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTraceId_128Bit()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTraceId_128Bit(), target.GetTraceId_128Bit()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSharedSpanContext()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSharedSpanContext()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSharedSpanContext(), target.GetSharedSpanContext()) {
			return false
		}
	}

	if m.GetCollectorEndpointVersion() != target.GetCollectorEndpointVersion() {
		return false
	}

	switch m.CollectorCluster.(type) {

	case *ZipkinConfig_CollectorUpstreamRef:
		if _, ok := target.CollectorCluster.(*ZipkinConfig_CollectorUpstreamRef); !ok {
			return false
		}

		if h, ok := interface{}(m.GetCollectorUpstreamRef()).(equality.Equalizer); ok {
			if !h.Equal(target.GetCollectorUpstreamRef()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetCollectorUpstreamRef(), target.GetCollectorUpstreamRef()) {
				return false
			}
		}

	case *ZipkinConfig_ClusterName:
		if _, ok := target.CollectorCluster.(*ZipkinConfig_ClusterName); !ok {
			return false
		}

		if strings.Compare(m.GetClusterName(), target.GetClusterName()) != 0 {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.CollectorCluster != target.CollectorCluster {
			return false
		}
	}

	return true
}
