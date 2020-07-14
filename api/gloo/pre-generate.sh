#!/bin/bash

set -e

rsync -ax --exclude 'solo-kit.json' --exclude 'grpc/v*'  ../gloo/projects/gloo/api/  ./api/gloo/gloo
rsync -ax --exclude 'solo-kit.json'  ../gloo/projects/gateway/api/  ./api/gloo/gateway

for file in $(find api/gloo -type f | grep ".proto")
do

# Gloo API changes
  sed "s|gloo/projects/gloo/api|solo-apis/api/gloo/gloo|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed "s|github.com/solo-io/gloo/projects/gloo/pkg/api|github.com/solo-io/solo-apis/pkg/api/gloo.solo.io|g" "$file" > "$file".tmp && mv "$file".tmp "$file"

# Gateway API changes
  sed "s|gloo/projects/gateway/api|solo-apis/api/gloo/gateway|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed "s|github.com/solo-io/gloo/projects/gateway/pkg/api|github.com/solo-io/solo-apis/pkg/api/gateway.solo.io|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
done

go run hack/convert.go