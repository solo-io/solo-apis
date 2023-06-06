// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for proto-based Spec and Status fields

package v2alpha1

import (
	proto "github.com/golang/protobuf/proto"
	"github.com/solo-io/protoc-gen-ext/pkg/clone"
)

// DeepCopyInto for the WaypointLifecycleManager.Spec
func (in *WaypointLifecycleManagerSpec) DeepCopyInto(out *WaypointLifecycleManagerSpec) {
	var p *WaypointLifecycleManagerSpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*WaypointLifecycleManagerSpec)
	} else {
		p = proto.Clone(in).(*WaypointLifecycleManagerSpec)
	}
	*out = *p
}

// DeepCopyInto for the WaypointLifecycleManager.Status
func (in *WaypointLifecycleManagerStatus) DeepCopyInto(out *WaypointLifecycleManagerStatus) {
	var p *WaypointLifecycleManagerStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*WaypointLifecycleManagerStatus)
	} else {
		p = proto.Clone(in).(*WaypointLifecycleManagerStatus)
	}
	*out = *p
}
