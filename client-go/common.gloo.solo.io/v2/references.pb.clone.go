// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/common/v2/references.proto

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
func (m *ObjectReference) Clone() proto.Message {
	var target *ObjectReference
	if m == nil {
		return target
	}
	target = &ObjectReference{}

	target.Name = m.GetName()

	target.Namespace = m.GetNamespace()

	target.Cluster = m.GetCluster()

	return target
}

// Clone function
func (m *DestinationReference) Clone() proto.Message {
	var target *DestinationReference
	if m == nil {
		return target
	}
	target = &DestinationReference{}

	target.Kind = m.GetKind()

	if h, ok := interface{}(m.GetPort()).(clone.Cloner); ok {
		target.Port = h.Clone().(*PortSelector)
	} else {
		target.Port = proto.Clone(m.GetPort()).(*PortSelector)
	}

	if m.GetSubset() != nil {
		target.Subset = make(map[string]string, len(m.GetSubset()))
		for k, v := range m.GetSubset() {

			target.Subset[k] = v

		}
	}

	target.Weight = m.GetWeight()

	switch m.RefKind.(type) {

	case *DestinationReference_Ref:

		if h, ok := interface{}(m.GetRef()).(clone.Cloner); ok {
			target.RefKind = &DestinationReference_Ref{
				Ref: h.Clone().(*ObjectReference),
			}
		} else {
			target.RefKind = &DestinationReference_Ref{
				Ref: proto.Clone(m.GetRef()).(*ObjectReference),
			}
		}

	case *DestinationReference_AwsLambda:

		if h, ok := interface{}(m.GetAwsLambda()).(clone.Cloner); ok {
			target.RefKind = &DestinationReference_AwsLambda{
				AwsLambda: h.Clone().(*AWSLambdaReference),
			}
		} else {
			target.RefKind = &DestinationReference_AwsLambda{
				AwsLambda: proto.Clone(m.GetAwsLambda()).(*AWSLambdaReference),
			}
		}

	}

	return target
}

// Clone function
func (m *AWSLambdaReference) Clone() proto.Message {
	var target *AWSLambdaReference
	if m == nil {
		return target
	}
	target = &AWSLambdaReference{}

	if h, ok := interface{}(m.GetCloudProvider()).(clone.Cloner); ok {
		target.CloudProvider = h.Clone().(*ObjectReference)
	} else {
		target.CloudProvider = proto.Clone(m.GetCloudProvider()).(*ObjectReference)
	}

	target.Function = m.GetFunction()

	target.Qualifier = m.GetQualifier()

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*LambdaOptions)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*LambdaOptions)
	}

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