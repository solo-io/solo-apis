// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/apimanagement/v2/graphql_resolver_map.proto

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
func (m *GraphQLResolverMapSpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("apimanagement.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/apimanagement.gloo.solo.io/v2.GraphQLResolverMapSpec")); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetTypes() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
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

	return hasher.Sum64(), nil
}

// Hash function
func (m *ExecutableSchema) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("apimanagement.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/apimanagement.gloo.solo.io/v2.ExecutableSchema")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetSchemaRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SchemaRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSchemaRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SchemaRef")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	switch m.ExecutableSchema.(type) {

	case *ExecutableSchema_Server:

		if h, ok := interface{}(m.GetServer()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Server")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetServer(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Server")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *ExecutableSchema_Local_:

		if h, ok := interface{}(m.GetLocal()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Local")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetLocal(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Local")); err != nil {
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
func (m *VariableTransformation) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("apimanagement.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/apimanagement.gloo.solo.io/v2.VariableTransformation")); err != nil {
		return 0, err
	}

	switch m.Transformation.(type) {

	case *VariableTransformation_Variable:

		if _, err = hasher.Write([]byte(m.GetVariable())); err != nil {
			return 0, err
		}

	case *VariableTransformation_Jq:

		if _, err = hasher.Write([]byte(m.GetJq())); err != nil {
			return 0, err
		}

	case *VariableTransformation_Value:

		if h, ok := interface{}(m.GetValue()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Value")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetValue(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Value")); err != nil {
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
func (m *Extraction) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("apimanagement.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/apimanagement.gloo.solo.io/v2.Extraction")); err != nil {
		return 0, err
	}

	switch m.ExtractionType.(type) {

	case *Extraction_RequestHeader:

		if _, err = hasher.Write([]byte(m.GetRequestHeader())); err != nil {
			return 0, err
		}

	case *Extraction_DynamicMetadata:

		if h, ok := interface{}(m.GetDynamicMetadata()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("DynamicMetadata")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetDynamicMetadata(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("DynamicMetadata")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Extraction_GraphqlParent:

		if h, ok := interface{}(m.GetGraphqlParent()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("GraphqlParent")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetGraphqlParent(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("GraphqlParent")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Extraction_GraphqlArg:

		if _, err = hasher.Write([]byte(m.GetGraphqlArg())); err != nil {
			return 0, err
		}

	case *Extraction_ResolverResult:

		if h, ok := interface{}(m.GetResolverResult()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ResolverResult")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetResolverResult(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ResolverResult")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Extraction_JsonValue:

		if h, ok := interface{}(m.GetJsonValue()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("JsonValue")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetJsonValue(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("JsonValue")); err != nil {
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
func (m *DynamicMetadataExtraction) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("apimanagement.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/apimanagement.gloo.solo.io/v2.DynamicMetadataExtraction")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetMetadataNamespace())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetKey())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GraphQLResolverMapStatus) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("apimanagement.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/apimanagement.gloo.solo.io/v2.GraphQLResolverMapStatus")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetGlobal()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Global")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetGlobal(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Global")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetWorkspaces() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
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

	if h, ok := interface{}(m.GetOwnerWorkspace()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("OwnerWorkspace")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetOwnerWorkspace(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("OwnerWorkspace")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GraphQLResolverMapSpec_Resolution) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("apimanagement.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/apimanagement.gloo.solo.io/v2.GraphQLResolverMapSpec_Resolution")); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetFields() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
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

	return hasher.Sum64(), nil
}

// Hash function
func (m *GraphQLResolverMapSpec_Resolution_Resolvers) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("apimanagement.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/apimanagement.gloo.solo.io/v2.GraphQLResolverMapSpec_Resolution_Resolvers")); err != nil {
		return 0, err
	}

	for _, v := range m.GetResolvers() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
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
		for k, v := range m.GetVariables() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
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

	if h, ok := interface{}(m.GetStatPrefix()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("StatPrefix")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetStatPrefix(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("StatPrefix")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GraphQLResolverMapSpec_Resolution_Resolvers_Resolver) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("apimanagement.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/apimanagement.gloo.solo.io/v2.GraphQLResolverMapSpec_Resolution_Resolvers_Resolver")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetResolverResultTransform()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ResolverResultTransform")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetResolverResultTransform(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ResolverResultTransform")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	switch m.Resolver.(type) {

	case *GraphQLResolverMapSpec_Resolution_Resolvers_Resolver_RestResolver_:

		if h, ok := interface{}(m.GetRestResolver()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("RestResolver")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRestResolver(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("RestResolver")); err != nil {
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
func (m *GraphQLResolverMapSpec_Resolution_Resolvers_Resolver_RestResolver) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("apimanagement.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/apimanagement.gloo.solo.io/v2.GraphQLResolverMapSpec_Resolution_Resolvers_Resolver_RestResolver")); err != nil {
		return 0, err
	}

	for _, v := range m.GetDestinations() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	if h, ok := interface{}(m.GetRequest()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Request")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRequest(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Request")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetVariables() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
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

	if h, ok := interface{}(m.GetSpanName()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SpanName")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSpanName(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SpanName")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Timeout")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTimeout(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Timeout")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GraphQLResolverMapSpec_Resolution_Resolvers_Resolver_RestResolver_RequestTemplate) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("apimanagement.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/apimanagement.gloo.solo.io/v2.GraphQLResolverMapSpec_Resolution_Resolvers_Resolver_RestResolver_RequestTemplate")); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetHeaders() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
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
		for k, v := range m.GetQueryParams() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
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

	if h, ok := interface{}(m.GetBody()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Body")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetBody(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Body")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GraphQLResolverMapSpec_Resolution_Resolvers_Resolver_RestResolver_RESTVariable) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("apimanagement.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/apimanagement.gloo.solo.io/v2.GraphQLResolverMapSpec_Resolution_Resolvers_Resolver_RestResolver_RESTVariable")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetResponseHeader())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ExecutableSchema_Local) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("apimanagement.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/apimanagement.gloo.solo.io/v2.ExecutableSchema_Local")); err != nil {
		return 0, err
	}

	for _, v := range m.GetResolverMapRefs() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	if h, ok := interface{}(m.GetOptions()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Options")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetOptions(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Options")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ExecutableSchema_GraphQLServer) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("apimanagement.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/apimanagement.gloo.solo.io/v2.ExecutableSchema_GraphQLServer")); err != nil {
		return 0, err
	}

	for _, v := range m.GetGraphqlServers() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
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
		for k, v := range m.GetVariables() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
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
		for k, v := range m.GetHeaders() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
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
		for k, v := range m.GetQueryParams() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
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

	if _, err = hasher.Write([]byte(m.GetSpanName())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ExecutableSchema_Local_Options) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("apimanagement.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/apimanagement.gloo.solo.io/v2.ExecutableSchema_Local_Options")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMaxDepth()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MaxDepth")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMaxDepth(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MaxDepth")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetEnableIntrospection())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ExecutableSchema_GraphQLServer_RequestExtraction) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("apimanagement.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/apimanagement.gloo.solo.io/v2.ExecutableSchema_GraphQLServer_RequestExtraction")); err != nil {
		return 0, err
	}

	switch m.Extraction.(type) {

	case *ExecutableSchema_GraphQLServer_RequestExtraction_DynamicMetadata:

		if h, ok := interface{}(m.GetDynamicMetadata()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("DynamicMetadata")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetDynamicMetadata(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("DynamicMetadata")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *ExecutableSchema_GraphQLServer_RequestExtraction_RequestHeader:

		if _, err = hasher.Write([]byte(m.GetRequestHeader())); err != nil {
			return 0, err
		}

	case *ExecutableSchema_GraphQLServer_RequestExtraction_JsonValue:

		if h, ok := interface{}(m.GetJsonValue()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("JsonValue")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetJsonValue(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("JsonValue")); err != nil {
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
