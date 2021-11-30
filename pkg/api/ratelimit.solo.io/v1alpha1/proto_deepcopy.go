// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for proto-based Spec and Status fields

package v1alpha1

import (
    proto "github.com/golang/protobuf/proto"
	"github.com/solo-io/protoc-gen-ext/pkg/clone"
)

// DeepCopyInto for the RateLimitConfig.Spec
func (in *RateLimitConfigSpec) DeepCopyInto(out *RateLimitConfigSpec) {
    var p *RateLimitConfigSpec
    if h, ok := interface{}(in).(clone.Cloner); ok {
        p = h.Clone().(*RateLimitConfigSpec)
    } else {
        p = proto.Clone(in).(*RateLimitConfigSpec)
    }
    *out = *p
}
// DeepCopyInto for the RateLimitConfig.Status
func (in *RateLimitConfigStatus) DeepCopyInto(out *RateLimitConfigStatus) {
    var p *RateLimitConfigStatus
    if h, ok := interface{}(in).(clone.Cloner); ok {
        p = h.Clone().(*RateLimitConfigStatus)
    } else {
        p = proto.Clone(in).(*RateLimitConfigStatus)
    }
    *out = *p
}
