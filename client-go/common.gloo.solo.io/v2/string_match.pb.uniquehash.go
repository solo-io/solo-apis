// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/common/v2/string_match.proto

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

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *StringMatch) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("common.gloo.solo.io.github.com/solo-io/solo-apis/client-go/common.gloo.solo.io/v2.StringMatch")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("IgnoreCase")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetIgnoreCase())
	if err != nil {
		return 0, err
	}

	switch m.MatchType.(type) {

	case *StringMatch_Exact:

		if _, err = hasher.Write([]byte("Exact")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(m.GetExact())); err != nil {
			return 0, err
		}

	case *StringMatch_Prefix:

		if _, err = hasher.Write([]byte("Prefix")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(m.GetPrefix())); err != nil {
			return 0, err
		}

	case *StringMatch_Regex:

		if _, err = hasher.Write([]byte("Regex")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(m.GetRegex())); err != nil {
			return 0, err
		}

	case *StringMatch_Suffix:

		if _, err = hasher.Write([]byte("Suffix")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(m.GetSuffix())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}