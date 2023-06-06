// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/common/v2/phase.proto

package v2

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
	"github.com/solo-io/protoc-gen-ext/pkg/hasher/hashstructure"
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
func (m *PrioritizedPhase) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("common.gloo.solo.io.github.com/solo-io/solo-apis/client-go/common.gloo.solo.io/v2.PrioritizedPhase")); err != nil {
		return 0, err
	}

	switch m.Phase.(type) {

	case *PrioritizedPhase_PreAuthz:

		if h, ok := interface{}(m.GetPreAuthz()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("PreAuthz")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetPreAuthz(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("PreAuthz")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *PrioritizedPhase_PostAuthz:

		if h, ok := interface{}(m.GetPostAuthz()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("PostAuthz")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetPostAuthz(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("PostAuthz")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *PrioritizedPhase_Phase) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("common.gloo.solo.io.github.com/solo-io/solo-apis/client-go/common.gloo.solo.io/v2.PrioritizedPhase_Phase")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetPriority()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Priority")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetPriority(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Priority")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
