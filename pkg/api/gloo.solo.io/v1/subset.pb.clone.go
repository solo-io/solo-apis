// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/subset.proto

package v1

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
func (m *Subset) Clone() proto.Message {
	var target *Subset
	if m == nil {
		return target
	}
	target = &Subset{}

	if m.GetValues() != nil {
		target.Values = make(map[string]string, len(m.GetValues()))
		for k, v := range m.GetValues() {

			target.Values[k] = v

		}
	}

	return target
}
