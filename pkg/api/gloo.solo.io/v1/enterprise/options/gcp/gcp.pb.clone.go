// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/enterprise/options/gcp/gcp.proto

package gcp

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
func (m *UpstreamSpec) Clone() proto.Message {
	var target *UpstreamSpec
	if m == nil {
		return target
	}
	target = &UpstreamSpec{}

	target.Host = m.GetHost()

	target.Audience = m.GetAudience()

	return target
}