// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/enterprise.gloo/v1/auth_config.proto

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
func (m *AuthConfigSpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.AuthConfigSpec")); err != nil {
		return 0, err
	}

	for _, v := range m.GetConfigs() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ExtAuthExtension) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.ExtAuthExtension")); err != nil {
		return 0, err
	}

	switch m.Spec.(type) {

	case *ExtAuthExtension_Disable:

		err = binary.Write(hasher, binary.LittleEndian, m.GetDisable())
		if err != nil {
			return 0, err
		}

	case *ExtAuthExtension_ConfigRef:

		if h, ok := interface{}(m.GetConfigRef()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetConfigRef(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	case *ExtAuthExtension_CustomAuth:

		if h, ok := interface{}(m.GetCustomAuth()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetCustomAuth(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *Settings) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.Settings")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetExtauthzServerRef()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetExtauthzServerRef(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetHttpService()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetHttpService(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetUserIdHeader())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetRequestTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetRequestTimeout(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetFailureModeAllow())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetRequestBody()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetRequestBody(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetClearRouteCache())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetStatusOnError())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *HttpService) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.HttpService")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetPathPrefix())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetRequest()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetRequest(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetResponse()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetResponse(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *BufferSettings) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.BufferSettings")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetMaxRequestBytes())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetAllowPartialMessage())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *CustomAuth) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.CustomAuth")); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetContextExtensions() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte(v)); err != nil {
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

	return hasher.Sum64(), nil
}

// Hash function
func (m *AuthPlugin) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.AuthPlugin")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetPluginFileName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetExportedSymbolName())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetConfig()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetConfig(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *BasicAuth) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.BasicAuth")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetRealm())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetApr()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetApr(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *OAuth) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.OAuth")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetClientId())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetClientSecretRef()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetClientSecretRef(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetIssuerUrl())); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetAuthEndpointQueryParams() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte(v)); err != nil {
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

	if _, err = hasher.Write([]byte(m.GetAppUrl())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetCallbackPath())); err != nil {
		return 0, err
	}

	for _, v := range m.GetScopes() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *OAuth2) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.OAuth2")); err != nil {
		return 0, err
	}

	switch m.OauthType.(type) {

	case *OAuth2_OidcAuthorizationCode:

		if h, ok := interface{}(m.GetOidcAuthorizationCode()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetOidcAuthorizationCode(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	case *OAuth2_AccessTokenValidation:

		if h, ok := interface{}(m.GetAccessTokenValidation()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetAccessTokenValidation(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *OidcAuthorizationCode) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.OidcAuthorizationCode")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetClientId())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetClientSecretRef()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetClientSecretRef(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetIssuerUrl())); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetAuthEndpointQueryParams() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte(v)); err != nil {
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

	if _, err = hasher.Write([]byte(m.GetAppUrl())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetCallbackPath())); err != nil {
		return 0, err
	}

	for _, v := range m.GetScopes() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *AccessTokenValidation) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.AccessTokenValidation")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetUserinfoUrl())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetCacheTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetCacheTimeout(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	switch m.ValidationType.(type) {

	case *AccessTokenValidation_IntrospectionUrl:

		if _, err = hasher.Write([]byte(m.GetIntrospectionUrl())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *OauthSecret) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.OauthSecret")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetClientSecret())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ApiKeyAuth) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.ApiKeyAuth")); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetLabelSelector() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte(v)); err != nil {
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

	for _, v := range m.GetApiKeySecretRefs() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ApiKeySecret) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.ApiKeySecret")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetGenerateApiKey())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetApiKey())); err != nil {
		return 0, err
	}

	for _, v := range m.GetLabels() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *OpaAuth) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.OpaAuth")); err != nil {
		return 0, err
	}

	for _, v := range m.GetModules() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	if _, err = hasher.Write([]byte(m.GetQuery())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *Ldap) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.Ldap")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetAddress())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetUserDnTemplate())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetMembershipAttributeName())); err != nil {
		return 0, err
	}

	for _, v := range m.GetAllowedGroups() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	if h, ok := interface{}(m.GetPool()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetPool(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ExtAuthConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.ExtAuthConfig")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetAuthConfigRefName())); err != nil {
		return 0, err
	}

	for _, v := range m.GetConfigs() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *AuthConfigStatus) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.AuthConfigStatus")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetState())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetReason())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetReportedBy())); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetSubresourceStatuses() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if val, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if err := binary.Write(innerHash, binary.LittleEndian, val); err != nil {
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

	if h, ok := interface{}(m.GetDetails()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetDetails(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *AuthConfigSpec_Config) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.AuthConfigSpec_Config")); err != nil {
		return 0, err
	}

	switch m.AuthConfig.(type) {

	case *AuthConfigSpec_Config_BasicAuth:

		if h, ok := interface{}(m.GetBasicAuth()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetBasicAuth(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	case *AuthConfigSpec_Config_Oauth:

		if h, ok := interface{}(m.GetOauth()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetOauth(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	case *AuthConfigSpec_Config_Oauth2:

		if h, ok := interface{}(m.GetOauth2()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetOauth2(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	case *AuthConfigSpec_Config_ApiKeyAuth:

		if h, ok := interface{}(m.GetApiKeyAuth()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetApiKeyAuth(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	case *AuthConfigSpec_Config_PluginAuth:

		if h, ok := interface{}(m.GetPluginAuth()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetPluginAuth(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	case *AuthConfigSpec_Config_OpaAuth:

		if h, ok := interface{}(m.GetOpaAuth()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetOpaAuth(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	case *AuthConfigSpec_Config_Ldap:

		if h, ok := interface{}(m.GetLdap()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetLdap(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *HttpService_Request) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.HttpService_Request")); err != nil {
		return 0, err
	}

	for _, v := range m.GetAllowedHeaders() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetHeadersToAdd() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte(v)); err != nil {
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

	return hasher.Sum64(), nil
}

// Hash function
func (m *HttpService_Response) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.HttpService_Response")); err != nil {
		return 0, err
	}

	for _, v := range m.GetAllowedUpstreamHeaders() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	for _, v := range m.GetAllowedClientHeaders() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *BasicAuth_Apr) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.BasicAuth_Apr")); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetUsers() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if val, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if err := binary.Write(innerHash, binary.LittleEndian, val); err != nil {
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
func (m *BasicAuth_Apr_SaltedHashedPassword) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.BasicAuth_Apr_SaltedHashedPassword")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetSalt())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetHashedPassword())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *Ldap_ConnectionPool) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.Ldap_ConnectionPool")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMaxSize()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetMaxSize(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetInitialSize()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetInitialSize(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ExtAuthConfig_OAuthConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.ExtAuthConfig_OAuthConfig")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetClientId())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetClientSecret())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetIssuerUrl())); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetAuthEndpointQueryParams() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte(v)); err != nil {
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

	if _, err = hasher.Write([]byte(m.GetAppUrl())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetCallbackPath())); err != nil {
		return 0, err
	}

	for _, v := range m.GetScopes() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ExtAuthConfig_OidcAuthorizationCodeConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.ExtAuthConfig_OidcAuthorizationCodeConfig")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetClientId())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetClientSecret())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetIssuerUrl())); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetAuthEndpointQueryParams() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte(v)); err != nil {
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

	if _, err = hasher.Write([]byte(m.GetAppUrl())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetCallbackPath())); err != nil {
		return 0, err
	}

	for _, v := range m.GetScopes() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ExtAuthConfig_OAuth2Config) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.ExtAuthConfig_OAuth2Config")); err != nil {
		return 0, err
	}

	switch m.OauthType.(type) {

	case *ExtAuthConfig_OAuth2Config_OidcAuthorizationCode:

		if h, ok := interface{}(m.GetOidcAuthorizationCode()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetOidcAuthorizationCode(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	case *ExtAuthConfig_OAuth2Config_AccessTokenValidation:

		if h, ok := interface{}(m.GetAccessTokenValidation()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetAccessTokenValidation(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ExtAuthConfig_ApiKeyAuthConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.ExtAuthConfig_ApiKeyAuthConfig")); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetValidApiKeyAndUser() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte(v)); err != nil {
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

	return hasher.Sum64(), nil
}

// Hash function
func (m *ExtAuthConfig_OpaAuthConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.ExtAuthConfig_OpaAuthConfig")); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetModules() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte(v)); err != nil {
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

	if _, err = hasher.Write([]byte(m.GetQuery())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ExtAuthConfig_Config) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("enterprise.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1.ExtAuthConfig_Config")); err != nil {
		return 0, err
	}

	switch m.AuthConfig.(type) {

	case *ExtAuthConfig_Config_Oauth:

		if h, ok := interface{}(m.GetOauth()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetOauth(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	case *ExtAuthConfig_Config_Oauth2:

		if h, ok := interface{}(m.GetOauth2()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetOauth2(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	case *ExtAuthConfig_Config_BasicAuth:

		if h, ok := interface{}(m.GetBasicAuth()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetBasicAuth(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	case *ExtAuthConfig_Config_ApiKeyAuth:

		if h, ok := interface{}(m.GetApiKeyAuth()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetApiKeyAuth(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	case *ExtAuthConfig_Config_PluginAuth:

		if h, ok := interface{}(m.GetPluginAuth()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetPluginAuth(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	case *ExtAuthConfig_Config_OpaAuth:

		if h, ok := interface{}(m.GetOpaAuth()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetOpaAuth(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	case *ExtAuthConfig_Config_Ldap:

		if h, ok := interface{}(m.GetLdap()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetLdap(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}
