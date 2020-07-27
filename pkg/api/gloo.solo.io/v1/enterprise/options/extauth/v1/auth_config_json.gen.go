// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/enterprise.gloo/v1/auth_config.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for AuthConfigSpec
func (this *AuthConfigSpec) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AuthConfigSpec
func (this *AuthConfigSpec) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for AuthConfigSpec_Config
func (this *AuthConfigSpec_Config) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AuthConfigSpec_Config
func (this *AuthConfigSpec_Config) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ExtAuthExtension
func (this *ExtAuthExtension) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ExtAuthExtension
func (this *ExtAuthExtension) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Settings
func (this *Settings) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Settings
func (this *Settings) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HttpService
func (this *HttpService) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HttpService
func (this *HttpService) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HttpService_Request
func (this *HttpService_Request) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HttpService_Request
func (this *HttpService_Request) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HttpService_Response
func (this *HttpService_Response) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HttpService_Response
func (this *HttpService_Response) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for BufferSettings
func (this *BufferSettings) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for BufferSettings
func (this *BufferSettings) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for CustomAuth
func (this *CustomAuth) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CustomAuth
func (this *CustomAuth) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for AuthPlugin
func (this *AuthPlugin) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AuthPlugin
func (this *AuthPlugin) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for BasicAuth
func (this *BasicAuth) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for BasicAuth
func (this *BasicAuth) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for BasicAuth_Apr
func (this *BasicAuth_Apr) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for BasicAuth_Apr
func (this *BasicAuth_Apr) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for BasicAuth_Apr_SaltedHashedPassword
func (this *BasicAuth_Apr_SaltedHashedPassword) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for BasicAuth_Apr_SaltedHashedPassword
func (this *BasicAuth_Apr_SaltedHashedPassword) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for OAuth
func (this *OAuth) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for OAuth
func (this *OAuth) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for OAuth2
func (this *OAuth2) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for OAuth2
func (this *OAuth2) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for OidcAuthorizationCode
func (this *OidcAuthorizationCode) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for OidcAuthorizationCode
func (this *OidcAuthorizationCode) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for AccessTokenValidation
func (this *AccessTokenValidation) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AccessTokenValidation
func (this *AccessTokenValidation) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for OauthSecret
func (this *OauthSecret) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for OauthSecret
func (this *OauthSecret) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ApiKeyAuth
func (this *ApiKeyAuth) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ApiKeyAuth
func (this *ApiKeyAuth) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ApiKeySecret
func (this *ApiKeySecret) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ApiKeySecret
func (this *ApiKeySecret) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for OpaAuth
func (this *OpaAuth) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for OpaAuth
func (this *OpaAuth) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Ldap
func (this *Ldap) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Ldap
func (this *Ldap) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Ldap_ConnectionPool
func (this *Ldap_ConnectionPool) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Ldap_ConnectionPool
func (this *Ldap_ConnectionPool) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ExtAuthConfig
func (this *ExtAuthConfig) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ExtAuthConfig
func (this *ExtAuthConfig) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ExtAuthConfig_OAuthConfig
func (this *ExtAuthConfig_OAuthConfig) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ExtAuthConfig_OAuthConfig
func (this *ExtAuthConfig_OAuthConfig) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ExtAuthConfig_OidcAuthorizationCodeConfig
func (this *ExtAuthConfig_OidcAuthorizationCodeConfig) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ExtAuthConfig_OidcAuthorizationCodeConfig
func (this *ExtAuthConfig_OidcAuthorizationCodeConfig) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ExtAuthConfig_OAuth2Config
func (this *ExtAuthConfig_OAuth2Config) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ExtAuthConfig_OAuth2Config
func (this *ExtAuthConfig_OAuth2Config) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ExtAuthConfig_ApiKeyAuthConfig
func (this *ExtAuthConfig_ApiKeyAuthConfig) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ExtAuthConfig_ApiKeyAuthConfig
func (this *ExtAuthConfig_ApiKeyAuthConfig) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ExtAuthConfig_OpaAuthConfig
func (this *ExtAuthConfig_OpaAuthConfig) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ExtAuthConfig_OpaAuthConfig
func (this *ExtAuthConfig_OpaAuthConfig) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ExtAuthConfig_Config
func (this *ExtAuthConfig_Config) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ExtAuthConfig_Config
func (this *ExtAuthConfig_Config) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for AuthConfigStatus
func (this *AuthConfigStatus) MarshalJSON() ([]byte, error) {
	str, err := AuthConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AuthConfigStatus
func (this *AuthConfigStatus) UnmarshalJSON(b []byte) error {
	return AuthConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	AuthConfigMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{EnumsAsInts: true}
	AuthConfigUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
