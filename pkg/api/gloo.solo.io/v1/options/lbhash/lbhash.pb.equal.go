// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/lbhash/lbhash.proto

package lbhash

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
func (m *RouteActionHashConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteActionHashConfig)
	if !ok {
		that2, ok := that.(RouteActionHashConfig)
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

	if len(m.GetHashPolicies()) != len(target.GetHashPolicies()) {
		return false
	}
	for idx, v := range m.GetHashPolicies() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetHashPolicies()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetHashPolicies()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *Cookie) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Cookie)
	if !ok {
		that2, ok := that.(Cookie)
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

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetTtl()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTtl()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTtl(), target.GetTtl()) {
			return false
		}
	}

	if strings.Compare(m.GetPath(), target.GetPath()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *HashPolicy) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HashPolicy)
	if !ok {
		that2, ok := that.(HashPolicy)
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

	if m.GetTerminal() != target.GetTerminal() {
		return false
	}

	switch m.KeyType.(type) {

	case *HashPolicy_Header:
		if _, ok := target.KeyType.(*HashPolicy_Header); !ok {
			return false
		}

		if strings.Compare(m.GetHeader(), target.GetHeader()) != 0 {
			return false
		}

	case *HashPolicy_Cookie:
		if _, ok := target.KeyType.(*HashPolicy_Cookie); !ok {
			return false
		}

		if h, ok := interface{}(m.GetCookie()).(equality.Equalizer); ok {
			if !h.Equal(target.GetCookie()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetCookie(), target.GetCookie()) {
				return false
			}
		}

	case *HashPolicy_SourceIp:
		if _, ok := target.KeyType.(*HashPolicy_SourceIp); !ok {
			return false
		}

		if m.GetSourceIp() != target.GetSourceIp() {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.KeyType != target.KeyType {
			return false
		}
	}

	return true
}
