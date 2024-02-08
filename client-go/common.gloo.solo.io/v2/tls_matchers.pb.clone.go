// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/common/v2/tls_matchers.proto

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
func (m *TLSRequestMatcher) Clone() proto.Message {
	var target *TLSRequestMatcher
	if m == nil {
		return target
	}
	target = &TLSRequestMatcher{}

	if m.GetSniHosts() != nil {
		target.SniHosts = make([]string, len(m.GetSniHosts()))
		for idx, v := range m.GetSniHosts() {

			target.SniHosts[idx] = v

		}
	}

	target.Port = m.GetPort()

	return target
}
