// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/admin/v2/workspace.proto

package v2

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"
	"strconv"

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
func (m *WorkspaceSpec) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("admin.gloo.solo.io.github.com/solo-io/solo-apis/client-go/admin.gloo.solo.io/v2.WorkspaceSpec")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("WorkloadClusters")); err != nil {
		return 0, err
	}
	for i, v := range m.GetWorkloadClusters() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
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

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *WorkspaceStatus) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("admin.gloo.solo.io.github.com/solo-io/solo-apis/client-go/admin.gloo.solo.io/v2.WorkspaceStatus")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetCommon()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Common")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetCommon(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Common")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetWorkspaceSettings()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("WorkspaceSettings")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetWorkspaceSettings(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("WorkspaceSettings")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("NumSelectedClusters")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetNumSelectedClusters())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("NumSelectedNamespaces")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetNumSelectedNamespaces())
	if err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetNumResources() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("v")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("v")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
			}

			if _, err = innerHash.Write([]byte("k")); err != nil {
				return 0, err
			}
			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetPolicyCounts() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte("v")); err != nil {
				return 0, err
			}
			err = binary.Write(innerHash, binary.LittleEndian, v)
			if err != nil {
				return 0, err
			}

			if _, err = innerHash.Write([]byte("k")); err != nil {
				return 0, err
			}
			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetDestinationCounts() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte("v")); err != nil {
				return 0, err
			}
			err = binary.Write(innerHash, binary.LittleEndian, v)
			if err != nil {
				return 0, err
			}

			if _, err = innerHash.Write([]byte("k")); err != nil {
				return 0, err
			}
			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	if _, err = hasher.Write([]byte("NumImportedWorkspaces")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetNumImportedWorkspaces())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *WorkspaceReport) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("admin.gloo.solo.io.github.com/solo-io/solo-apis/client-go/admin.gloo.solo.io/v2.WorkspaceReport")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetReport()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Report")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetReport(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Report")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetWorkspaceSettings()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("WorkspaceSettings")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetWorkspaceSettings(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("WorkspaceSettings")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("SelectedClusters")); err != nil {
		return 0, err
	}
	for i, v := range m.GetSelectedClusters() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetResources() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("v")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("v")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
			}

			if _, err = innerHash.Write([]byte("k")); err != nil {
				return 0, err
			}
			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetPolicyCounts() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("v")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("v")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
			}

			if _, err = innerHash.Write([]byte("k")); err != nil {
				return 0, err
			}
			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetDestinationCounts() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("v")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("v")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
			}

			if _, err = innerHash.Write([]byte("k")); err != nil {
				return 0, err
			}
			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	if _, err = hasher.Write([]byte("ImportedWorkspaces")); err != nil {
		return 0, err
	}
	for i, v := range m.GetImportedWorkspaces() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if _, err = hasher.Write([]byte("v")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *WorkspaceStatus_SourceType) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("admin.gloo.solo.io.github.com/solo-io/solo-apis/client-go/admin.gloo.solo.io/v2.WorkspaceStatus_SourceType")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Imported")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetImported())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Exported")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetExported())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Private")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetPrivate())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *WorkspaceReport_SelectedCluster) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("admin.gloo.solo.io.github.com/solo-io/solo-apis/client-go/admin.gloo.solo.io/v2.WorkspaceReport_SelectedCluster")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Name")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Namespaces")); err != nil {
		return 0, err
	}
	for i, v := range m.GetNamespaces() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if _, err = hasher.Write([]byte("v")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *WorkspaceReport_SourceType) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("admin.gloo.solo.io.github.com/solo-io/solo-apis/client-go/admin.gloo.solo.io/v2.WorkspaceReport_SourceType")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetImported()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Imported")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetImported(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Imported")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetExported()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Exported")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetExported(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Exported")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetPrivate()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Private")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetPrivate(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Private")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *WorkspaceReport_ObjectReferences) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("admin.gloo.solo.io.github.com/solo-io/solo-apis/client-go/admin.gloo.solo.io/v2.WorkspaceReport_ObjectReferences")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("ObjectReferences")); err != nil {
		return 0, err
	}
	for i, v := range m.GetObjectReferences() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
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