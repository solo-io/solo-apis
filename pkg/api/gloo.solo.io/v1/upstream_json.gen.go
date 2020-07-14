// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/upstream.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/api/v2/cluster"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/api/v2/core"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/aws"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/aws/ec2"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/azure"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/consul"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/kubernetes"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/pipe"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/static"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for Upstream
func (this *Upstream) MarshalJSON() ([]byte, error) {
	str, err := UpstreamMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Upstream
func (this *Upstream) UnmarshalJSON(b []byte) error {
	return UpstreamUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for DiscoveryMetadata
func (this *DiscoveryMetadata) MarshalJSON() ([]byte, error) {
	str, err := UpstreamMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for DiscoveryMetadata
func (this *DiscoveryMetadata) UnmarshalJSON(b []byte) error {
	return UpstreamUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	UpstreamMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	UpstreamUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
