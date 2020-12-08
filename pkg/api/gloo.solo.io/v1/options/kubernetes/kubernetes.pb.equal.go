// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/kubernetes/kubernetes.proto

package kubernetes

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
func (m *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
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

	if strings.Compare(m.GetServiceName(), target.GetServiceName()) != 0 {
		return false
	}

	if strings.Compare(m.GetServiceNamespace(), target.GetServiceNamespace()) != 0 {
		return false
	}

	if m.GetServicePort() != target.GetServicePort() {
		return false
	}

	if len(m.GetSelector()) != len(target.GetSelector()) {
		return false
	}
	for k, v := range m.GetSelector() {

		if strings.Compare(v, target.GetSelector()[k]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetServiceSpec()).(equality.Equalizer); ok {
		if !h.Equal(target.GetServiceSpec()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetServiceSpec(), target.GetServiceSpec()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSubsetSpec()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSubsetSpec()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSubsetSpec(), target.GetSubsetSpec()) {
			return false
		}
	}

	return true
}
