#!/bin/bash

set -e

for file in $(find api/gloo -type f | grep ".proto")
do

# Gloo API changes
  sed "s|gloo/projects/gloo/api|solo-apis/api/gloo/gloo|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed "s|github.com/solo-io/gloo/projects/gloo/pkg/api|github.com/solo-io/solo-apis/pkg/api/gloo.solo.io|g" "$file" > "$file".tmp && mv "$file".tmp "$file"

# Gateway API changes
  sed "s|gloo/projects/gateway/api|solo-apis/api/gloo/gateway|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed "s|github.com/solo-io/gloo/projects/gateway/pkg/api|github.com/solo-io/solo-apis/pkg/api/gateway.solo.io|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
done

# for file in $(find pkg/api/gateway.solo.io/v1/types -type f | grep "_json.gen.go")
# do
#   sed "s|github_com_gogo_protobuf_jsonpb.Marshaler{}|github_com_gogo_protobuf_jsonpb.Marshaler{EnumsAsInts: true}|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
# done
