// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/internal/v2/certificate_request.proto

package v2

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
func (m *CertificateRequestSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CertificateRequestSpec)
	if !ok {
		that2, ok := that.(CertificateRequestSpec)
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

	if bytes.Compare(m.GetCertificateSigningRequest(), target.GetCertificateSigningRequest()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *CertificateRequestStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CertificateRequestStatus)
	if !ok {
		that2, ok := that.(CertificateRequestStatus)
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

	if m.GetObservedGeneration() != target.GetObservedGeneration() {
		return false
	}

	if strings.Compare(m.GetError(), target.GetError()) != 0 {
		return false
	}

	if m.GetState() != target.GetState() {
		return false
	}

	if bytes.Compare(m.GetSignedCertificate(), target.GetSignedCertificate()) != 0 {
		return false
	}

	if bytes.Compare(m.GetSigningRootCa(), target.GetSigningRootCa()) != 0 {
		return false
	}

	if bytes.Compare(m.GetCertChain(), target.GetCertChain()) != 0 {
		return false
	}

	return true
}
