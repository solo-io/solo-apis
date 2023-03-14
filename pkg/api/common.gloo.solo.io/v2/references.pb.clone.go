// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/common/v2/references.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"
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
func (m *WorkloadReference) Clone() proto.Message {
	var target *WorkloadReference
	if m == nil {
		return target
	}
	target = &WorkloadReference{}

	if h, ok := interface{}(m.GetRef()).(clone.Cloner); ok {
		target.Ref = h.Clone().(*ObjectReference)
	} else {
		target.Ref = proto.Clone(m.GetRef()).(*ObjectReference)
	}

	target.Kind = m.GetKind()

	return target
}

// Clone function
func (m *ListenerPortReference) Clone() proto.Message {
	var target *ListenerPortReference
	if m == nil {
		return target
	}
	target = &ListenerPortReference{}

	if h, ok := interface{}(m.GetGatewayRef()).(clone.Cloner); ok {
		target.GatewayRef = h.Clone().(*ObjectReference)
	} else {
		target.GatewayRef = proto.Clone(m.GetGatewayRef()).(*ObjectReference)
	}

	target.Port = m.GetPort()

	return target
}
