// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-fed/fed.rpc/v1/glooinstance.proto

package v1

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
func (m *GlooInstance) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlooInstance)
	if !ok {
		that2, ok := that.(GlooInstance)
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

	if h, ok := interface{}(m.GetMetadata()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMetadata()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMetadata(), target.GetMetadata()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSpec()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSpec()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSpec(), target.GetSpec()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetStatus()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStatus()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStatus(), target.GetStatus()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ListGlooInstancesRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListGlooInstancesRequest)
	if !ok {
		that2, ok := that.(ListGlooInstancesRequest)
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

	return true
}

// Equal function
func (m *ListGlooInstancesResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListGlooInstancesResponse)
	if !ok {
		that2, ok := that.(ListGlooInstancesResponse)
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

	if len(m.GetGlooInstances()) != len(target.GetGlooInstances()) {
		return false
	}
	for idx, v := range m.GetGlooInstances() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetGlooInstances()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetGlooInstances()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *ClusterDetails) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ClusterDetails)
	if !ok {
		that2, ok := that.(ClusterDetails)
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

	if strings.Compare(m.GetCluster(), target.GetCluster()) != 0 {
		return false
	}

	if len(m.GetGlooInstances()) != len(target.GetGlooInstances()) {
		return false
	}
	for idx, v := range m.GetGlooInstances() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetGlooInstances()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetGlooInstances()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *ListClusterDetailsRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListClusterDetailsRequest)
	if !ok {
		that2, ok := that.(ListClusterDetailsRequest)
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

	return true
}

// Equal function
func (m *ListClusterDetailsResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListClusterDetailsResponse)
	if !ok {
		that2, ok := that.(ListClusterDetailsResponse)
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

	if len(m.GetClusterDetails()) != len(target.GetClusterDetails()) {
		return false
	}
	for idx, v := range m.GetClusterDetails() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetClusterDetails()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetClusterDetails()[idx]) {
				return false
			}
		}

	}

	return true
}
