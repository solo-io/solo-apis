// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/common/v2/string_match.proto

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
func (m *StringMatch) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("common.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/common.gloo.solo.io/v2.StringMatch")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetIgnoreCase())
	if err != nil {
		return 0, err
	}

	switch m.MatchType.(type) {

	case *StringMatch_Exact:

		if _, err = hasher.Write([]byte(m.GetExact())); err != nil {
			return 0, err
		}

	case *StringMatch_Prefix:

		if _, err = hasher.Write([]byte(m.GetPrefix())); err != nil {
			return 0, err
		}

	case *StringMatch_Regex:

		if _, err = hasher.Write([]byte(m.GetRegex())); err != nil {
			return 0, err
		}

	case *StringMatch_Suffix:

		if _, err = hasher.Write([]byte(m.GetSuffix())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}