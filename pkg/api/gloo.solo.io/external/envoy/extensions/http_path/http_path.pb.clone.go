// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/extensions/http_path/http_path.proto

package http_path

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_core_v3 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3"
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
func (m *HttpPath) Clone() proto.Message {
	var target *HttpPath
	if m == nil {
		return target
	}
	target = &HttpPath{}

	if h, ok := interface{}(m.GetHttpHealthCheck()).(clone.Cloner); ok {
		target.HttpHealthCheck = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_core_v3.HealthCheck_HttpHealthCheck)
	} else {
		target.HttpHealthCheck = proto.Clone(m.GetHttpHealthCheck()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_core_v3.HealthCheck_HttpHealthCheck)
	}

	return target
}
