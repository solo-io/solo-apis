#!/bin/bash

set -e

for file in $(find api/gloo-mesh -type f | grep ".proto")
do
  # Re-map imports within oss-imported and enterprise-networking
  sed 's|github.com/solo-io/gloo-mesh-enterprise/oss-imported/api/|github.com/solo-io/solo-apis/api/gloo-mesh/|g' "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed 's|github.com/solo-io/gloo-mesh-enterprise/enterprise-networking/api/|github.com/solo-io/solo-apis/api/gloo-mesh/|g' "$file" > "$file".tmp && mv "$file".tmp "$file"

  sed 's|github.com/solo-io/skv2|github.com/solo-io/solo-apis/api/gloo-mesh/external/skv2|g' "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed 's|"networking/|"github.com/solo-io/solo-apis/api/gloo-mesh/external/istio.io/api/networking/|g' "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed 's|"udpa/|"github.com/solo-io/solo-apis/api/gloo-mesh/external/cncf/udpa/|g' "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed 's|"envoy/|"github.com/solo-io/solo-apis/api/gloo-mesh/external/envoyproxy/data-plane-api/envoy/|g' "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed 's|"encoding/protobuf/|"github.com/solo-io/solo-apis/api/gloo-mesh/external/cuelang.org/go/encoding/protobuf/|g' "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed 's|"xds/|"github.com/solo-io/solo-apis/api/gloo-mesh/external/cncf/udpa/xds/|g' "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed 's|"gogoproto/|"github.com/solo-io/solo-apis/api/gloo-mesh/external/gogo/protobuf/gogoproto/|g' "$file" > "$file".tmp && mv "$file".tmp "$file"
  

done
