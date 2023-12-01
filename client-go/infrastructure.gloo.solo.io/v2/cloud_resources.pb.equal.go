// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/infrastructure/v2/cloud_resources.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *CloudResourcesSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CloudResourcesSpec)
	if !ok {
		that2, ok := that.(CloudResourcesSpec)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetProvider(), target.GetProvider()) != 0 {
		return false
	}

	switch m.ProviderType.(type) {

	case *CloudResourcesSpec_Aws:
		if _, ok := target.ProviderType.(*CloudResourcesSpec_Aws); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAws()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAws()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAws(), target.GetAws()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.ProviderType != target.ProviderType {
			return false
		}
	}

	return true
}

// Equal function
func (m *CloudResourcesSpec_AWSResources) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CloudResourcesSpec_AWSResources)
	if !ok {
		that2, ok := that.(CloudResourcesSpec_AWSResources)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetLambda()) != len(target.GetLambda()) {
		return false
	}
	for idx, v := range m.GetLambda() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetLambda()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetLambda()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *CloudResourcesSpec_AWSResources_Lambda) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CloudResourcesSpec_AWSResources_Lambda)
	if !ok {
		that2, ok := that.(CloudResourcesSpec_AWSResources_Lambda)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetLambdaFunctionName(), target.GetLambdaFunctionName()) != 0 {
		return false
	}

	if strings.Compare(m.GetQualifier(), target.GetQualifier()) != 0 {
		return false
	}

	return true
}