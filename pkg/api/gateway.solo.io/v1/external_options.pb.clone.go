// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gateway/v1/external_options.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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
func (m *VirtualHostOption) Clone() proto.Message {
	var target *VirtualHostOption
	if m == nil {
		return target
	}
	target = &VirtualHostOption{}

	if h, ok := interface{}(m.GetStatus()).(clone.Cloner); ok {
		target.Status = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Status)
	} else {
		target.Status = proto.Clone(m.GetStatus()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Status)
	}

	if h, ok := interface{}(m.GetMetadata()).(clone.Cloner); ok {
		target.Metadata = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	} else {
		target.Metadata = proto.Clone(m.GetMetadata()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	}

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.VirtualHostOptions)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.VirtualHostOptions)
	}

	return target
}

// Clone function
func (m *RouteOption) Clone() proto.Message {
	var target *RouteOption
	if m == nil {
		return target
	}
	target = &RouteOption{}

	if h, ok := interface{}(m.GetStatus()).(clone.Cloner); ok {
		target.Status = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Status)
	} else {
		target.Status = proto.Clone(m.GetStatus()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Status)
	}

	if h, ok := interface{}(m.GetMetadata()).(clone.Cloner); ok {
		target.Metadata = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	} else {
		target.Metadata = proto.Clone(m.GetMetadata()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	}

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.RouteOptions)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.RouteOptions)
	}

	return target
}
