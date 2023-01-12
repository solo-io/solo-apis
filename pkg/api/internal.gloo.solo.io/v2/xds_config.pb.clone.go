// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/internal/v2/xds_config.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
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
func (m *XdsConfigSpec) Clone() proto.Message {
	var target *XdsConfigSpec
	if m == nil {
		return target
	}
	target = &XdsConfigSpec{}

	if m.GetWorkloads() != nil {
		target.Workloads = make([]*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.ObjectRef, len(m.GetWorkloads()))
		for idx, v := range m.GetWorkloads() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Workloads[idx] = h.Clone().(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.ObjectRef)
			} else {
				target.Workloads[idx] = proto.Clone(v).(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.ObjectRef)
			}

		}
	}

	if m.GetTypes() != nil {
		target.Types = make([]*XdsConfigSpec_TypedResources, len(m.GetTypes()))
		for idx, v := range m.GetTypes() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Types[idx] = h.Clone().(*XdsConfigSpec_TypedResources)
			} else {
				target.Types[idx] = proto.Clone(v).(*XdsConfigSpec_TypedResources)
			}

		}
	}

	return target
}

// Clone function
func (m *XdsConfigStatus) Clone() proto.Message {
	var target *XdsConfigStatus
	if m == nil {
		return target
	}
	target = &XdsConfigStatus{}

	target.ObservedGeneration = m.GetObservedGeneration()

	target.Error = m.GetError()

	return target
}

// Clone function
func (m *XdsConfigSpec_TypedResources) Clone() proto.Message {
	var target *XdsConfigSpec_TypedResources
	if m == nil {
		return target
	}
	target = &XdsConfigSpec_TypedResources{}

	target.TypeUrl = m.GetTypeUrl()

	if m.GetResources() != nil {
		target.Resources = make([]*XdsConfigSpec_Resource, len(m.GetResources()))
		for idx, v := range m.GetResources() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Resources[idx] = h.Clone().(*XdsConfigSpec_Resource)
			} else {
				target.Resources[idx] = proto.Clone(v).(*XdsConfigSpec_Resource)
			}

		}
	}

	return target
}

// Clone function
func (m *XdsConfigSpec_Resource) Clone() proto.Message {
	var target *XdsConfigSpec_Resource
	if m == nil {
		return target
	}
	target = &XdsConfigSpec_Resource{}

	target.Name = m.GetName()

	if m.GetCompressedData() != nil {
		target.CompressedData = make([]byte, len(m.GetCompressedData()))
		copy(target.CompressedData, m.GetCompressedData())
	}

	return target
}
