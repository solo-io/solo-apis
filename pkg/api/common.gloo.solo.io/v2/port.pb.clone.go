// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/common/v2/port.proto

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
func (m *PortSelector) Clone() proto.Message {
	var target *PortSelector
	if m == nil {
		return target
	}
	target = &PortSelector{}

	switch m.Specifier.(type) {

	case *PortSelector_Number:

		target.Specifier = &PortSelector_Number{
			Number: m.GetNumber(),
		}

	case *PortSelector_Name:

		target.Specifier = &PortSelector_Name{
			Name: m.GetName(),
		}

	}

	return target
}
