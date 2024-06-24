// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/common/v2/istio_helm.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	k8s_io_api_core_v1 "k8s.io/api/core/v1"
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
func (m *IstioLifecycleHelmGlobals) Clone() proto.Message {
	var target *IstioLifecycleHelmGlobals
	if m == nil {
		return target
	}
	target = &IstioLifecycleHelmGlobals{}

	target.Repo = m.GetRepo()

	if m.GetRepoSecrets() != nil {
		target.RepoSecrets = make([]*k8s_io_api_core_v1.ObjectReference, len(m.GetRepoSecrets()))
		for idx, v := range m.GetRepoSecrets() {

			target.RepoSecrets[idx] = v.DeepCopy()

		}
	}

	target.RepoSkipTlsVerify = m.GetRepoSkipTlsVerify()

	return target
}
