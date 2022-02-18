#!/bin/bash

set -e

EXTERNAL_API_PKG=github.com/solo-io/solo-apis/api/gloo-mesh/external
ENVOY_API_PKG=${EXTERNAL_API_PKG}/envoyproxy/data-plane-api/envoy

# Emits a sed statement that will add an "option go_package" statement to a .proto file
# cat my.proto | $(add_go_package "my.proto.package", "github.com/solo-io/my-repo/pkg/my/proto/package") > dest.proto
add_go_package() {
  subst="s|package ${1};|package ${1};\noption go_package = \"${2}\";|g"
  echo "eval sed '$subst'"
}

for file in $(find api/gloo-mesh -type f | grep ".proto")
do
  # Re-map imports within oss-imported and enterprise-networking
  cat $file | \
  sed 's|github.com/solo-io/gloo-mesh-enterprise/oss-imported/api/|github.com/solo-io/solo-apis/api/gloo-mesh/|g' | \
  sed 's|github.com/solo-io/gloo-mesh-enterprise/enterprise-networking/api/|github.com/solo-io/solo-apis/api/gloo-mesh/|g' | \
  sed 's|"networking/|"github.com/solo-io/solo-apis/api/gloo-mesh/external/istio.io/api/networking/|g' | \
  sed 's|"udpa/|"github.com/solo-io/solo-apis/api/gloo-mesh/external/cncf/udpa/udpa/|g' | \
  sed 's|"envoy/|"github.com/solo-io/solo-apis/api/gloo-mesh/external/envoyproxy/data-plane-api/envoy/|g' | \
  sed 's|"validate/|"github.com/envoyproxy/protoc-gen-validate/validate/|g' | \
  sed 's|"encoding/protobuf/|"github.com/solo-io/solo-apis/api/gloo-mesh/external/cue/encoding/protobuf/|g' | \
  sed 's|"xds/|"github.com/solo-io/solo-apis/api/gloo-mesh/external/cncf/udpa/xds/|g' | \
  sed 's|"gogoproto/|"github.com/solo-io/solo-apis/api/gloo-mesh/external/gogo/protobuf/gogoproto/|g' | \
  sed 's|github.com/solo-io/envoy-gloo|github.com/solo-io/solo-apis/api/gloo-mesh/external/envoy-gloo|g' | \
  sed 's|"transformation";|"github.com/solo-io/solo-apis/api/gloo-mesh/external/envoy-gloo/api/envoy/config/filter/http/transformation/v2/";|g' | \
  $(add_go_package "envoy.annotations" "${ENVOY_API_PKG}/annotations") | \
  $(add_go_package "envoy.api.v2.core" "${ENVOY_API_PKG}/api/v2/core") | \
  $(add_go_package "envoy.api.v2.route" "${ENVOY_API_PKG}/api/v2/route") | \
  $(add_go_package "envoy.config.core.v3" "${ENVOY_API_PKG}/config/core/v3") | \
  $(add_go_package "envoy.config.route.v3" "${ENVOY_API_PKG}/config/route/v3") | \
  $(add_go_package "envoy.data.accesslog.v3" "${ENVOY_API_PKG}/data/accesslog/v3") | \
  $(add_go_package "envoy.extensions.filters.http.jwt_authn.v3" "${ENVOY_API_PKG}/extensions/filters/http/jwt_authn/v3") | \
  $(add_go_package "envoy.type" "${ENVOY_API_PKG}/type") | \
  $(add_go_package "envoy.type.v3" "${ENVOY_API_PKG}/type/v3") | \
  $(add_go_package "envoy.type.matcher" "${ENVOY_API_PKG}/type/matcher") | \
  $(add_go_package "envoy.type.matcher.v3" "${ENVOY_API_PKG}/type/matcher/v3") | \
  $(add_go_package "envoy.type.metadata.v2" "${ENVOY_API_PKG}/type/metadata/v2") | \
  $(add_go_package "envoy.type.metadata.v3" "${ENVOY_API_PKG}/type/metadata/v3") | \
  $(add_go_package "envoy.type.tracing.v2" "${ENVOY_API_PKG}/type/tracing/v2") | \
  $(add_go_package "envoy.type.tracing.v3" "${ENVOY_API_PKG}/type/tracing/v3") | \
  $(add_go_package "udpa.annotations" "${EXTERNAL_API_PKG}/cncf/udpa/udpa/annotations") | \
  $(add_go_package "xds.core.v3" "${EXTERNAL_API_PKG}/cncf/udpa/xds.core/v3") > "$file".tmp && mv "$file".tmp "$file"
done

