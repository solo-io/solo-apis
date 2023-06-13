// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/config/matching/inputs/cipher_detection_input/v3/cipher_detection_input.proto

package v3

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
func (m *CipherDetectionInput) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CipherDetectionInput)
	if !ok {
		that2, ok := that.(CipherDetectionInput)
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

	if len(m.GetPassthroughCiphers()) != len(target.GetPassthroughCiphers()) {
		return false
	}
	for idx, v := range m.GetPassthroughCiphers() {

		if v != target.GetPassthroughCiphers()[idx] {
			return false
		}

	}

	if len(m.GetTerminatingCiphers()) != len(target.GetTerminatingCiphers()) {
		return false
	}
	for idx, v := range m.GetTerminatingCiphers() {

		if v != target.GetTerminatingCiphers()[idx] {
			return false
		}

	}

	return true
}
