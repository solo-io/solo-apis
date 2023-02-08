// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/networking/v2/cloud_resources.proto

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

	if m.GetLambda() != nil {
		target.Lambda = make([]*LambdaSpec, len(m.GetLambda()))
		for idx, v := range m.GetLambda() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Lambda[idx] = h.Clone().(*LambdaSpec)
			} else {
				target.Lambda[idx] = proto.Clone(v).(*LambdaSpec)
			}

		}
	}

	return target
}

// Clone function
func (m *LambdaSpec) Clone() proto.Message {
	var target *LambdaSpec
	if m == nil {
		return target
	}
	target = &LambdaSpec{}

	target.LogicalName = m.GetLogicalName()

	target.LambdaFunctionName = m.GetLambdaFunctionName()

	target.Qualifier = m.GetQualifier()

	return target
}