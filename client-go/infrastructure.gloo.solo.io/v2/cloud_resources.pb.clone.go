// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/infrastructure/v2/cloud_resources.proto

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
func (m *CloudResourcesSpec) Clone() proto.Message {
	var target *CloudResourcesSpec
	if m == nil {
		return target
	}
	target = &CloudResourcesSpec{}

	target.Provider = m.GetProvider()

	switch m.ProviderType.(type) {

	case *CloudResourcesSpec_Aws:

		if h, ok := interface{}(m.GetAws()).(clone.Cloner); ok {
			target.ProviderType = &CloudResourcesSpec_Aws{
				Aws: h.Clone().(*CloudResourcesSpec_AWSResources),
			}
		} else {
			target.ProviderType = &CloudResourcesSpec_Aws{
				Aws: proto.Clone(m.GetAws()).(*CloudResourcesSpec_AWSResources),
			}
		}

	}

	return target
}

// Clone function
func (m *CloudResourcesSpec_AWSResources) Clone() proto.Message {
	var target *CloudResourcesSpec_AWSResources
	if m == nil {
		return target
	}
	target = &CloudResourcesSpec_AWSResources{}

	if m.GetLambda() != nil {
		target.Lambda = make([]*CloudResourcesSpec_AWSResources_Lambda, len(m.GetLambda()))
		for idx, v := range m.GetLambda() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Lambda[idx] = h.Clone().(*CloudResourcesSpec_AWSResources_Lambda)
			} else {
				target.Lambda[idx] = proto.Clone(v).(*CloudResourcesSpec_AWSResources_Lambda)
			}

		}
	}

	return target
}

// Clone function
func (m *CloudResourcesSpec_AWSResources_Lambda) Clone() proto.Message {
	var target *CloudResourcesSpec_AWSResources_Lambda
	if m == nil {
		return target
	}
	target = &CloudResourcesSpec_AWSResources_Lambda{}

	target.LambdaFunctionName = m.GetLambdaFunctionName()

	target.Qualifier = m.GetQualifier()

	return target
}
