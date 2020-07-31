// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/config/filter/http/gzip/v2/gzip.proto

package v2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for Gzip
func (this *Gzip) MarshalJSON() ([]byte, error) {
	str, err := GzipMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Gzip
func (this *Gzip) UnmarshalJSON(b []byte) error {
	return GzipUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Gzip_CompressionLevel
func (this *Gzip_CompressionLevel) MarshalJSON() ([]byte, error) {
	str, err := GzipMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Gzip_CompressionLevel
func (this *Gzip_CompressionLevel) UnmarshalJSON(b []byte) error {
	return GzipUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	GzipMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	GzipUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
