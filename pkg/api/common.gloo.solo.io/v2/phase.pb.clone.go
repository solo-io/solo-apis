// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/common/v2/phase.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"
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
func (m *PrioritizedPhase) Clone() proto.Message {
	var target *PrioritizedPhase
	if m == nil {
		return target
	}
	target = &PrioritizedPhase{}

	switch m.Phase.(type) {

	case *PrioritizedPhase_PreAuthz:

		if h, ok := interface{}(m.GetPreAuthz()).(clone.Cloner); ok {
			target.Phase = &PrioritizedPhase_PreAuthz{
				PreAuthz: h.Clone().(*PrioritizedPhase_Phase),
			}
		} else {
			target.Phase = &PrioritizedPhase_PreAuthz{
				PreAuthz: proto.Clone(m.GetPreAuthz()).(*PrioritizedPhase_Phase),
			}
		}

	case *PrioritizedPhase_PostAuthz:

		if h, ok := interface{}(m.GetPostAuthz()).(clone.Cloner); ok {
			target.Phase = &PrioritizedPhase_PostAuthz{
				PostAuthz: h.Clone().(*PrioritizedPhase_Phase),
			}
		} else {
			target.Phase = &PrioritizedPhase_PostAuthz{
				PostAuthz: proto.Clone(m.GetPostAuthz()).(*PrioritizedPhase_Phase),
			}
		}

	}

	return target
}

// Clone function
func (m *PrioritizedPhase_Phase) Clone() proto.Message {
	var target *PrioritizedPhase_Phase
	if m == nil {
		return target
	}
	target = &PrioritizedPhase_Phase{}

	if h, ok := interface{}(m.GetPriority()).(clone.Cloner); ok {
		target.Priority = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.Int32Value)
	} else {
		target.Priority = proto.Clone(m.GetPriority()).(*github_com_golang_protobuf_ptypes_wrappers.Int32Value)
	}

	return target
}