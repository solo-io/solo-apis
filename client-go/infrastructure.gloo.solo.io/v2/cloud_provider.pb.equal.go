// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/infrastructure/v2/cloud_provider.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *CloudProviderSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CloudProviderSpec)
	if !ok {
		that2, ok := that.(CloudProviderSpec)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.Provider.(type) {

	case *CloudProviderSpec_Aws:
		if _, ok := target.Provider.(*CloudProviderSpec_Aws); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAws()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAws()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAws(), target.GetAws()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Provider != target.Provider {
			return false
		}
	}

	return true
}

// Equal function
func (m *AWSProvider) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AWSProvider)
	if !ok {
		that2, ok := that.(AWSProvider)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetAccountId(), target.GetAccountId()) != 0 {
		return false
	}

	if strings.Compare(m.GetRegion(), target.GetRegion()) != 0 {
		return false
	}

	if strings.Compare(m.GetStsEndpoint(), target.GetStsEndpoint()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetLambda()).(equality.Equalizer); ok {
		if !h.Equal(target.GetLambda()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetLambda(), target.GetLambda()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *CloudProviderStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CloudProviderStatus)
	if !ok {
		that2, ok := that.(CloudProviderStatus)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetCommon()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCommon()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCommon(), target.GetCommon()) {
			return false
		}
	}

	if len(m.GetNumAppliedRoutableFunctions()) != len(target.GetNumAppliedRoutableFunctions()) {
		return false
	}
	for k, v := range m.GetNumAppliedRoutableFunctions() {

		if v != target.GetNumAppliedRoutableFunctions()[k] {
			return false
		}

	}

	if m.GetNumChildCloudResources() != target.GetNumChildCloudResources() {
		return false
	}

	return true
}

// Equal function
func (m *CloudProviderReport) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CloudProviderReport)
	if !ok {
		that2, ok := that.(CloudProviderReport)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetReport()).(equality.Equalizer); ok {
		if !h.Equal(target.GetReport()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetReport(), target.GetReport()) {
			return false
		}
	}

	if len(m.GetAppliedRoutableFunctions()) != len(target.GetAppliedRoutableFunctions()) {
		return false
	}
	for k, v := range m.GetAppliedRoutableFunctions() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAppliedRoutableFunctions()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAppliedRoutableFunctions()[k]) {
				return false
			}
		}

	}

	if len(m.GetChildCloudResources()) != len(target.GetChildCloudResources()) {
		return false
	}
	for idx, v := range m.GetChildCloudResources() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetChildCloudResources()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetChildCloudResources()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *AWSProvider_LambdaOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AWSProvider_LambdaOptions)
	if !ok {
		that2, ok := that.(AWSProvider_LambdaOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetInvokeRoleName(), target.GetInvokeRoleName()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetDiscovery()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDiscovery()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDiscovery(), target.GetDiscovery()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *AWSProvider_LambdaOptions_LambdaDiscovery) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AWSProvider_LambdaOptions_LambdaDiscovery)
	if !ok {
		that2, ok := that.(AWSProvider_LambdaOptions_LambdaDiscovery)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetEnabled() != target.GetEnabled() {
		return false
	}

	if strings.Compare(m.GetRoleName(), target.GetRoleName()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetFilter()).(equality.Equalizer); ok {
		if !h.Equal(target.GetFilter()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetFilter(), target.GetFilter()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *AWSProvider_LambdaOptions_LambdaDiscovery_LambdaFilter) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AWSProvider_LambdaOptions_LambdaDiscovery_LambdaFilter)
	if !ok {
		that2, ok := that.(AWSProvider_LambdaOptions_LambdaDiscovery_LambdaFilter)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetNames()) != len(target.GetNames()) {
		return false
	}
	for idx, v := range m.GetNames() {

		if strings.Compare(v, target.GetNames()[idx]) != 0 {
			return false
		}

	}

	if m.GetLatestOnly() != target.GetLatestOnly() {
		return false
	}

	return true
}

// Equal function
func (m *CloudProviderReport_LambdaFunctions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CloudProviderReport_LambdaFunctions)
	if !ok {
		that2, ok := that.(CloudProviderReport_LambdaFunctions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetFunctions()) != len(target.GetFunctions()) {
		return false
	}
	for idx, v := range m.GetFunctions() {

		if strings.Compare(v, target.GetFunctions()[idx]) != 0 {
			return false
		}

	}

	return true
}
