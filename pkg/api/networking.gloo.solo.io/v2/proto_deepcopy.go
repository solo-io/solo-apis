// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for proto-based Spec and Status fields

package v2

import (
	proto "github.com/golang/protobuf/proto"
	"github.com/solo-io/protoc-gen-ext/pkg/clone"
)

// DeepCopyInto for the ExternalService.Spec
func (in *ExternalServiceSpec) DeepCopyInto(out *ExternalServiceSpec) {
	var p *ExternalServiceSpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*ExternalServiceSpec)
	} else {
		p = proto.Clone(in).(*ExternalServiceSpec)
	}
	*out = *p
}

// DeepCopyInto for the ExternalService.Status
func (in *ExternalServiceStatus) DeepCopyInto(out *ExternalServiceStatus) {
	var p *ExternalServiceStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*ExternalServiceStatus)
	} else {
		p = proto.Clone(in).(*ExternalServiceStatus)
	}
	*out = *p
}

// DeepCopyInto for the ExternalEndpoint.Spec
func (in *ExternalEndpointSpec) DeepCopyInto(out *ExternalEndpointSpec) {
	var p *ExternalEndpointSpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*ExternalEndpointSpec)
	} else {
		p = proto.Clone(in).(*ExternalEndpointSpec)
	}
	*out = *p
}

// DeepCopyInto for the ExternalEndpoint.Status
func (in *ExternalEndpointStatus) DeepCopyInto(out *ExternalEndpointStatus) {
	var p *ExternalEndpointStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*ExternalEndpointStatus)
	} else {
		p = proto.Clone(in).(*ExternalEndpointStatus)
	}
	*out = *p
}

// DeepCopyInto for the RouteTable.Spec
func (in *RouteTableSpec) DeepCopyInto(out *RouteTableSpec) {
	var p *RouteTableSpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*RouteTableSpec)
	} else {
		p = proto.Clone(in).(*RouteTableSpec)
	}
	*out = *p
}

// DeepCopyInto for the RouteTable.Status
func (in *RouteTableStatus) DeepCopyInto(out *RouteTableStatus) {
	var p *RouteTableStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*RouteTableStatus)
	} else {
		p = proto.Clone(in).(*RouteTableStatus)
	}
	*out = *p
}

// DeepCopyInto for the VirtualDestination.Spec
func (in *VirtualDestinationSpec) DeepCopyInto(out *VirtualDestinationSpec) {
	var p *VirtualDestinationSpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*VirtualDestinationSpec)
	} else {
		p = proto.Clone(in).(*VirtualDestinationSpec)
	}
	*out = *p
}

// DeepCopyInto for the VirtualDestination.Status
func (in *VirtualDestinationStatus) DeepCopyInto(out *VirtualDestinationStatus) {
	var p *VirtualDestinationStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*VirtualDestinationStatus)
	} else {
		p = proto.Clone(in).(*VirtualDestinationStatus)
	}
	*out = *p
}

// DeepCopyInto for the VirtualGateway.Spec
func (in *VirtualGatewaySpec) DeepCopyInto(out *VirtualGatewaySpec) {
	var p *VirtualGatewaySpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*VirtualGatewaySpec)
	} else {
		p = proto.Clone(in).(*VirtualGatewaySpec)
	}
	*out = *p
}

// DeepCopyInto for the VirtualGateway.Status
func (in *VirtualGatewayStatus) DeepCopyInto(out *VirtualGatewayStatus) {
	var p *VirtualGatewayStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*VirtualGatewayStatus)
	} else {
		p = proto.Clone(in).(*VirtualGatewayStatus)
	}
	*out = *p
}
