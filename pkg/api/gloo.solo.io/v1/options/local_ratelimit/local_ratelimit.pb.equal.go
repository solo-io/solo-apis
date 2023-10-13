// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/local_ratelimit/local_ratelimit.proto

package local_ratelimit

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
func (m *TokenBucket) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TokenBucket)
	if !ok {
		that2, ok := that.(TokenBucket)
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

	if m.GetMaxTokens() != target.GetMaxTokens() {
		return false
	}

	if h, ok := interface{}(m.GetTokensPerFill()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTokensPerFill()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTokensPerFill(), target.GetTokensPerFill()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetFillInterval()).(equality.Equalizer); ok {
		if !h.Equal(target.GetFillInterval()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetFillInterval(), target.GetFillInterval()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *Settings) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Settings)
	if !ok {
		that2, ok := that.(Settings)
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

	if h, ok := interface{}(m.GetDefaultLimit()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDefaultLimit()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDefaultLimit(), target.GetDefaultLimit()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetLocalRateLimitPerDownstreamConnection()).(equality.Equalizer); ok {
		if !h.Equal(target.GetLocalRateLimitPerDownstreamConnection()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetLocalRateLimitPerDownstreamConnection(), target.GetLocalRateLimitPerDownstreamConnection()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetEnableXRatelimitHeaders()).(equality.Equalizer); ok {
		if !h.Equal(target.GetEnableXRatelimitHeaders()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetEnableXRatelimitHeaders(), target.GetEnableXRatelimitHeaders()) {
			return false
		}
	}

	return true
}
