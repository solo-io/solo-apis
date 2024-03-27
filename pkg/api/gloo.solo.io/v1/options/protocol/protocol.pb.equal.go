// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/protocol/protocol.proto

package protocol

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
func (m *HttpProtocolOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HttpProtocolOptions)
	if !ok {
		that2, ok := that.(HttpProtocolOptions)
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

	if h, ok := interface{}(m.GetIdleTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIdleTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIdleTimeout(), target.GetIdleTimeout()) {
			return false
		}
	}

	if m.GetMaxHeadersCount() != target.GetMaxHeadersCount() {
		return false
	}

	if h, ok := interface{}(m.GetMaxStreamDuration()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMaxStreamDuration()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMaxStreamDuration(), target.GetMaxStreamDuration()) {
			return false
		}
	}

	if m.GetHeadersWithUnderscoresAction() != target.GetHeadersWithUnderscoresAction() {
		return false
	}

	return true
}

// Equal function
func (m *Http1ProtocolOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Http1ProtocolOptions)
	if !ok {
		that2, ok := that.(Http1ProtocolOptions)
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

	if m.GetEnableTrailers() != target.GetEnableTrailers() {
		return false
	}

	if h, ok := interface{}(m.GetOverrideStreamErrorOnInvalidHttpMessage()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOverrideStreamErrorOnInvalidHttpMessage()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOverrideStreamErrorOnInvalidHttpMessage(), target.GetOverrideStreamErrorOnInvalidHttpMessage()) {
			return false
		}
	}

	switch m.HeaderFormat.(type) {

	case *Http1ProtocolOptions_ProperCaseHeaderKeyFormat:

		if m.GetProperCaseHeaderKeyFormat() != target.GetProperCaseHeaderKeyFormat() {
			return false
		}

	case *Http1ProtocolOptions_PreserveCaseHeaderKeyFormat:

		if m.GetPreserveCaseHeaderKeyFormat() != target.GetPreserveCaseHeaderKeyFormat() {
			return false
		}

	}

	return true
}

// Equal function
func (m *Http2ProtocolOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Http2ProtocolOptions)
	if !ok {
		that2, ok := that.(Http2ProtocolOptions)
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

	if h, ok := interface{}(m.GetMaxConcurrentStreams()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMaxConcurrentStreams()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMaxConcurrentStreams(), target.GetMaxConcurrentStreams()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetInitialStreamWindowSize()).(equality.Equalizer); ok {
		if !h.Equal(target.GetInitialStreamWindowSize()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetInitialStreamWindowSize(), target.GetInitialStreamWindowSize()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetInitialConnectionWindowSize()).(equality.Equalizer); ok {
		if !h.Equal(target.GetInitialConnectionWindowSize()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetInitialConnectionWindowSize(), target.GetInitialConnectionWindowSize()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetOverrideStreamErrorOnInvalidHttpMessage()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOverrideStreamErrorOnInvalidHttpMessage()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOverrideStreamErrorOnInvalidHttpMessage(), target.GetOverrideStreamErrorOnInvalidHttpMessage()) {
			return false
		}
	}

	return true
}
