// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for proto-based Spec and Status fields

package v2alpha1

import (
	proto "github.com/golang/protobuf/proto"
	"github.com/solo-io/protoc-gen-ext/pkg/clone"
)

// DeepCopyInto for the ExternalWorkload.Spec
func (in *ExternalWorkloadSpec) DeepCopyInto(out *ExternalWorkloadSpec) {
	var p *ExternalWorkloadSpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*ExternalWorkloadSpec)
	} else {
		p = proto.Clone(in).(*ExternalWorkloadSpec)
	}
	*out = *p
}

// DeepCopyInto for the ExternalWorkload.Status
func (in *ExternalWorkloadStatus) DeepCopyInto(out *ExternalWorkloadStatus) {
	var p *ExternalWorkloadStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*ExternalWorkloadStatus)
	} else {
		p = proto.Clone(in).(*ExternalWorkloadStatus)
	}
	*out = *p
}