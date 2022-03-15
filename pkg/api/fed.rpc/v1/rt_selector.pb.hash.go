// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-fed/fed.rpc/v1/rt_selector.proto

package v1

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
func (m *SubRouteTableRow) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.rpc/v1.SubRouteTableRow")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetMatcher())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetMatchType())); err != nil {
		return 0, err
	}

	for _, v := range m.GetMethods() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	for _, v := range m.GetHeaders() {

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

	for _, v := range m.GetQueryParameters() {

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

	for _, v := range m.GetRtRoutes() {

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

	switch m.Action.(type) {

	case *SubRouteTableRow_RouteAction:

		if h, ok := interface{}(m.GetRouteAction()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("RouteAction")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRouteAction(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("RouteAction")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *SubRouteTableRow_RedirectAction:

		if h, ok := interface{}(m.GetRedirectAction()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("RedirectAction")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRedirectAction(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("RedirectAction")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *SubRouteTableRow_DirectResponseAction:

		if h, ok := interface{}(m.GetDirectResponseAction()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("DirectResponseAction")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetDirectResponseAction(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("DirectResponseAction")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *SubRouteTableRow_DelegateAction:

		if h, ok := interface{}(m.GetDelegateAction()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("DelegateAction")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetDelegateAction(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("DelegateAction")); err != nil {
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
func (m *GetVirtualServiceRoutesRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.rpc/v1.GetVirtualServiceRoutesRequest")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetVirtualServiceRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("VirtualServiceRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetVirtualServiceRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("VirtualServiceRef")); err != nil {
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
func (m *GetVirtualServiceRoutesResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.rpc/v1.GetVirtualServiceRoutesResponse")); err != nil {
		return 0, err
	}

	for _, v := range m.GetVsRoutes() {

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
