// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/extensions/aws/filter.proto

package aws

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_duration "github.com/golang/protobuf/ptypes/duration"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *AWSLambdaPerRoute) Clone() proto.Message {
	var target *AWSLambdaPerRoute
	if m == nil {
		return target
	}
	target = &AWSLambdaPerRoute{}

	target.Name = m.GetName()

	target.Qualifier = m.GetQualifier()

	target.Async = m.GetAsync()

	if h, ok := interface{}(m.GetEmptyBodyOverride()).(clone.Cloner); ok {
		target.EmptyBodyOverride = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.StringValue)
	} else {
		target.EmptyBodyOverride = proto.Clone(m.GetEmptyBodyOverride()).(*github_com_golang_protobuf_ptypes_wrappers.StringValue)
	}

	target.UnwrapAsAlb = m.GetUnwrapAsAlb()

	return target
}

// Clone function
func (m *AWSLambdaProtocolExtension) Clone() proto.Message {
	var target *AWSLambdaProtocolExtension
	if m == nil {
		return target
	}
	target = &AWSLambdaProtocolExtension{}

	target.Host = m.GetHost()

	target.Region = m.GetRegion()

	target.AccessKey = m.GetAccessKey()

	target.SecretKey = m.GetSecretKey()

	target.SessionToken = m.GetSessionToken()

	target.RoleArn = m.GetRoleArn()

	return target
}

// Clone function
func (m *AWSLambdaConfig) Clone() proto.Message {
	var target *AWSLambdaConfig
	if m == nil {
		return target
	}
	target = &AWSLambdaConfig{}

	target.PropagateOriginalRouting = m.GetPropagateOriginalRouting()

	if h, ok := interface{}(m.GetCredentialRefreshDelay()).(clone.Cloner); ok {
		target.CredentialRefreshDelay = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.CredentialRefreshDelay = proto.Clone(m.GetCredentialRefreshDelay()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	switch m.CredentialsFetcher.(type) {

	case *AWSLambdaConfig_UseDefaultCredentials:

		if h, ok := interface{}(m.GetUseDefaultCredentials()).(clone.Cloner); ok {
			target.CredentialsFetcher = &AWSLambdaConfig_UseDefaultCredentials{
				UseDefaultCredentials: h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue),
			}
		} else {
			target.CredentialsFetcher = &AWSLambdaConfig_UseDefaultCredentials{
				UseDefaultCredentials: proto.Clone(m.GetUseDefaultCredentials()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue),
			}
		}

	case *AWSLambdaConfig_ServiceAccountCredentials_:

		if h, ok := interface{}(m.GetServiceAccountCredentials()).(clone.Cloner); ok {
			target.CredentialsFetcher = &AWSLambdaConfig_ServiceAccountCredentials_{
				ServiceAccountCredentials: h.Clone().(*AWSLambdaConfig_ServiceAccountCredentials),
			}
		} else {
			target.CredentialsFetcher = &AWSLambdaConfig_ServiceAccountCredentials_{
				ServiceAccountCredentials: proto.Clone(m.GetServiceAccountCredentials()).(*AWSLambdaConfig_ServiceAccountCredentials),
			}
		}

	}

	return target
}

// Clone function
func (m *AWSLambdaConfig_ServiceAccountCredentials) Clone() proto.Message {
	var target *AWSLambdaConfig_ServiceAccountCredentials
	if m == nil {
		return target
	}
	target = &AWSLambdaConfig_ServiceAccountCredentials{}

	target.Cluster = m.GetCluster()

	target.Uri = m.GetUri()

	if h, ok := interface{}(m.GetTimeout()).(clone.Cloner); ok {
		target.Timeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.Timeout = proto.Clone(m.GetTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	return target
}
