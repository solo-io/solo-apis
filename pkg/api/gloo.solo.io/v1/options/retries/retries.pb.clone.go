// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/retries/retries.proto

package retries

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	google_golang_org_protobuf_types_known_durationpb "google.golang.org/protobuf/types/known/durationpb"

	google_golang_org_protobuf_types_known_wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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
func (m *RetryBackOff) Clone() proto.Message {
	var target *RetryBackOff
	if m == nil {
		return target
	}
	target = &RetryBackOff{}

	if h, ok := interface{}(m.GetBaseInterval()).(clone.Cloner); ok {
		target.BaseInterval = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.BaseInterval = proto.Clone(m.GetBaseInterval()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetMaxInterval()).(clone.Cloner); ok {
		target.MaxInterval = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.MaxInterval = proto.Clone(m.GetMaxInterval()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	return target
}

// Clone function
func (m *RetryPolicy) Clone() proto.Message {
	var target *RetryPolicy
	if m == nil {
		return target
	}
	target = &RetryPolicy{}

	target.RetryOn = m.GetRetryOn()

	target.NumRetries = m.GetNumRetries()

	if h, ok := interface{}(m.GetPerTryTimeout()).(clone.Cloner); ok {
		target.PerTryTimeout = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.PerTryTimeout = proto.Clone(m.GetPerTryTimeout()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetRetryBackOff()).(clone.Cloner); ok {
		target.RetryBackOff = h.Clone().(*RetryBackOff)
	} else {
		target.RetryBackOff = proto.Clone(m.GetRetryBackOff()).(*RetryBackOff)
	}

	switch m.PriorityPredicate.(type) {

	case *RetryPolicy_PreviousPriorities_:

		if h, ok := interface{}(m.GetPreviousPriorities()).(clone.Cloner); ok {
			target.PriorityPredicate = &RetryPolicy_PreviousPriorities_{
				PreviousPriorities: h.Clone().(*RetryPolicy_PreviousPriorities),
			}
		} else {
			target.PriorityPredicate = &RetryPolicy_PreviousPriorities_{
				PreviousPriorities: proto.Clone(m.GetPreviousPriorities()).(*RetryPolicy_PreviousPriorities),
			}
		}

	}

	return target
}

// Clone function
func (m *RetryPolicy_PreviousPriorities) Clone() proto.Message {
	var target *RetryPolicy_PreviousPriorities
	if m == nil {
		return target
	}
	target = &RetryPolicy_PreviousPriorities{}

	if h, ok := interface{}(m.GetUpdateFrequency()).(clone.Cloner); ok {
		target.UpdateFrequency = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.UpdateFrequency = proto.Clone(m.GetUpdateFrequency()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	return target
}
