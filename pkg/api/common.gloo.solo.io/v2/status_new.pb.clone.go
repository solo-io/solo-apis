// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/common/v2/status_new.proto

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
func (m *Status) Clone() proto.Message {
	var target *Status
	if m == nil {
		return target
	}
	target = &Status{}

	if h, ok := interface{}(m.GetState()).(clone.Cloner); ok {
		target.State = h.Clone().(*State)
	} else {
		target.State = proto.Clone(m.GetState()).(*State)
	}

	if m.GetWorkspaceConditions() != nil {
		target.WorkspaceConditions = make(map[string]uint32, len(m.GetWorkspaceConditions()))
		for k, v := range m.GetWorkspaceConditions() {

			target.WorkspaceConditions[k] = v

		}
	}

	return target
}

// Clone function
func (m *State) Clone() proto.Message {
	var target *State
	if m == nil {
		return target
	}
	target = &State{}

	target.ObservedGeneration = m.GetObservedGeneration()

	target.Approval = m.GetApproval()

	target.Message = m.GetMessage()

	return target
}

// Clone function
func (m *Report) Clone() proto.Message {
	var target *Report
	if m == nil {
		return target
	}
	target = &Report{}

	if m.GetClusters() != nil {
		target.Clusters = make(map[string]*State, len(m.GetClusters()))
		for k, v := range m.GetClusters() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Clusters[k] = h.Clone().(*State)
			} else {
				target.Clusters[k] = proto.Clone(v).(*State)
			}

		}
	}

	return target
}
