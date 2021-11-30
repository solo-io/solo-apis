// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/enterprise/options/graphql/v1alpha1/graphql.proto

package v1alpha1

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
func (m *PathSegment) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/graphql/v1alpha1.PathSegment")); err != nil {
		return 0, err
	}

	switch m.Segment.(type) {

	case *PathSegment_Key:

		if _, err = hasher.Write([]byte(m.GetKey())); err != nil {
			return 0, err
		}

	case *PathSegment_Index:

		err = binary.Write(hasher, binary.LittleEndian, m.GetIndex())
		if err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ValueProvider) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/graphql/v1alpha1.ValueProvider")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetProviderTemplate())); err != nil {
		return 0, err
	}

	switch m.Provider.(type) {

	case *ValueProvider_GraphqlArg:

		if h, ok := interface{}(m.GetGraphqlArg()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("GraphqlArg")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetGraphqlArg(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("GraphqlArg")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *ValueProvider_TypedProvider:

		if h, ok := interface{}(m.GetTypedProvider()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("TypedProvider")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetTypedProvider(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("TypedProvider")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *ValueProvider_GraphqlParent:

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

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *JsonKeyValue) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/graphql/v1alpha1.JsonKeyValue")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetKey())); err != nil {
		return 0, err
	}

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

	return hasher.Sum64(), nil
}

// Hash function
func (m *JsonNode) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/graphql/v1alpha1.JsonNode")); err != nil {
		return 0, err
	}

	for _, v := range m.GetKeyValues() {

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

	return hasher.Sum64(), nil
}

// Hash function
func (m *RequestTemplate) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/graphql/v1alpha1.RequestTemplate")); err != nil {
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

	switch m.OutgoingBody.(type) {

	case *RequestTemplate_Json:

		if h, ok := interface{}(m.GetJson()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Json")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetJson(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Json")); err != nil {
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
func (m *RESTResolver) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/graphql/v1alpha1.RESTResolver")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetUpstreamRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("UpstreamRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetUpstreamRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("UpstreamRef")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetRequestTransform()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RequestTransform")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRequestTransform(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RequestTransform")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetSpanName())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *QueryMatcher) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/graphql/v1alpha1.QueryMatcher")); err != nil {
		return 0, err
	}

	switch m.Match.(type) {

	case *QueryMatcher_FieldMatcher_:

		if h, ok := interface{}(m.GetFieldMatcher()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("FieldMatcher")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetFieldMatcher(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("FieldMatcher")); err != nil {
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
func (m *Resolution) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/graphql/v1alpha1.Resolution")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMatcher()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Matcher")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMatcher(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Matcher")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	switch m.Resolver.(type) {

	case *Resolution_RestResolver:

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
func (m *GraphQLSchema) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/graphql/v1alpha1.GraphQLSchema")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMetadata()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Metadata")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMetadata(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Metadata")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetSchema())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetEnableIntrospection())
	if err != nil {
		return 0, err
	}

	for _, v := range m.GetResolutions() {

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

	return hasher.Sum64(), nil
}

// Hash function
func (m *ValueProvider_GraphQLArgExtraction) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/graphql/v1alpha1.ValueProvider_GraphQLArgExtraction")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetArgName())); err != nil {
		return 0, err
	}

	for _, v := range m.GetPath() {

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

	return hasher.Sum64(), nil
}

// Hash function
func (m *ValueProvider_GraphQLParentExtraction) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/graphql/v1alpha1.ValueProvider_GraphQLParentExtraction")); err != nil {
		return 0, err
	}

	for _, v := range m.GetPath() {

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

	return hasher.Sum64(), nil
}

// Hash function
func (m *ValueProvider_TypedValueProvider) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/graphql/v1alpha1.ValueProvider_TypedValueProvider")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetType())
	if err != nil {
		return 0, err
	}

	switch m.ValProvider.(type) {

	case *ValueProvider_TypedValueProvider_Header:

		if _, err = hasher.Write([]byte(m.GetHeader())); err != nil {
			return 0, err
		}

	case *ValueProvider_TypedValueProvider_Value:

		if _, err = hasher.Write([]byte(m.GetValue())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *JsonKeyValue_JsonValueList) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/graphql/v1alpha1.JsonKeyValue_JsonValueList")); err != nil {
		return 0, err
	}

	for _, v := range m.GetValues() {

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

	return hasher.Sum64(), nil
}

// Hash function
func (m *JsonKeyValue_JsonValue) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/graphql/v1alpha1.JsonKeyValue_JsonValue")); err != nil {
		return 0, err
	}

	switch m.JsonVal.(type) {

	case *JsonKeyValue_JsonValue_Node:

		if h, ok := interface{}(m.GetNode()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Node")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetNode(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Node")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *JsonKeyValue_JsonValue_ValueProvider:

		if h, ok := interface{}(m.GetValueProvider()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ValueProvider")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetValueProvider(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ValueProvider")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *JsonKeyValue_JsonValue_List:

		if h, ok := interface{}(m.GetList()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("List")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetList(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("List")); err != nil {
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
func (m *QueryMatcher_FieldMatcher) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/graphql/v1alpha1.QueryMatcher_FieldMatcher")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetType())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetField())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
