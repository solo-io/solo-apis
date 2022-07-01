#!/bin/bash

set -e

EXTERNAL_API_PKG=github.com/solo-io/solo-apis/api/gloo-mesh/external

# Emits a sed statement that will add an "option go_package" statement to a .proto file
# cat my.proto | $(add_go_package "my.proto.package", "github.com/solo-io/my-repo/pkg/my/proto/package") > dest.proto
add_go_package() {
  subst="s|package ${1};|package ${1};\noption go_package = \"${2}\";|g"
  echo "eval sed '$subst'"
}

for file in $(find api/gloo-mesh -type f | grep ".proto")
do
  cat $file | \
  # Re-map imports
  sed 's|github.com/solo-io/gloo-mesh-enterprise/pkg/api/|github.com/solo-io/solo-apis/pkg/api/|g' | \
  sed 's|github.com/solo-io/gloo-mesh-enterprise/api/|github.com/solo-io/solo-apis/api/gloo-mesh/|g' | \
  sed 's|"networking/|"github.com/solo-io/solo-apis/api/gloo-mesh/external/istio.io/api/networking/|g' | \
  sed 's|"udpa/|"github.com/solo-io/solo-apis/api/gloo-mesh/external/cncf/udpa/udpa/|g' | \
  sed 's|"envoy/|"github.com/solo-io/solo-apis/api/gloo-mesh/external/envoyproxy/data-plane-api/envoy/|g' | \
  sed 's|"validate/|"github.com/envoyproxy/protoc-gen-validate/validate/|g' | \
  sed 's|"encoding/protobuf/|"github.com/solo-io/solo-apis/api/gloo-mesh/external/cue/encoding/protobuf/|g' | \
  sed 's|"xds/|"github.com/solo-io/solo-apis/api/gloo-mesh/external/cncf/udpa/xds/|g' | \
  sed 's|"gogoproto/|"github.com/solo-io/solo-apis/api/gloo-mesh/external/gogo/protobuf/gogoproto/|g' | \
  sed 's|github.com/solo-io/envoy-gloo|github.com/solo-io/solo-apis/api/gloo-mesh/external/envoy-gloo|g' | \
  sed 's|"transformation";|"github.com/solo-io/solo-apis/api/gloo-mesh/external/envoy-gloo/api/envoy/config/filter/http/transformation/v2/";|g' | \
  # Re-map imports and go_packages
  sed 's|github.com/envoyproxy/|github.com/solo-io/solo-apis/api/gloo-mesh/external/envoyproxy/|g' | \
  $(add_go_package "udpa.annotations" "${EXTERNAL_API_PKG}/cncf/udpa/udpa/annotations") | \
  $(add_go_package "xds.core.v3" "${EXTERNAL_API_PKG}/cncf/udpa/xds.core/v3") > "$file".tmp && mv "$file".tmp "$file"
done

