// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/internal/v2/certificate_request.proto

package v2

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	"github.com/mitchellh/hashstructure"
	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = new(hash.Hash64)
	_ = fnv.New64
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// Hash function
func (m *CertificateRequestSpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("internal.gloo.solo.io.github.com/solo-io/solo-apis/client-go/internal.gloo.solo.io/v2.CertificateRequestSpec")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write(m.GetCertificateSigningRequest()); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *CertificateRequestStatus) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("internal.gloo.solo.io.github.com/solo-io/solo-apis/client-go/internal.gloo.solo.io/v2.CertificateRequestStatus")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetObservedGeneration())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetError())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetState())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write(m.GetSignedCertificate()); err != nil {
		return 0, err
	}

	if _, err = hasher.Write(m.GetSigningRootCa()); err != nil {
		return 0, err
	}

	if _, err = hasher.Write(m.GetCertChain()); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
