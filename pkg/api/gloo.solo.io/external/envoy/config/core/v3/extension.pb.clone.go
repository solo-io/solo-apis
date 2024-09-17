// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/config/core/v3/extension.proto

package v3

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	google_golang_org_protobuf_types_known_anypb "google.golang.org/protobuf/types/known/anypb"
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
func (m *TypedExtensionConfig) Clone() proto.Message {
	var target *TypedExtensionConfig
	if m == nil {
		return target
	}
	target = &TypedExtensionConfig{}

	target.Name = m.GetName()

	if h, ok := interface{}(m.GetTypedConfig()).(clone.Cloner); ok {
		target.TypedConfig = h.Clone().(*google_golang_org_protobuf_types_known_anypb.Any)
	} else {
		target.TypedConfig = proto.Clone(m.GetTypedConfig()).(*google_golang_org_protobuf_types_known_anypb.Any)
	}

	return target
}
