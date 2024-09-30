#!/bin/bash

set -e

# Sync the proto files from the Gloo Repository checked out at ../gloo
# Exclude gloo/api/grpc because we do not currently want to expose any of our grpc apis
rsync -ax --delete --exclude 'solo-kit.json' --exclude 'grpc'  ../gloo/projects/gloo/api/  ./api/gloo/gloo
rsync -ax --delete --exclude 'solo-kit.json'  ../gloo/projects/gateway/api/  ./api/gloo/gateway

# Clean up the grpc directory (remove this line if we decide to expose a grpc api
if [ -d 'api/gloo/gloo/grpc' ]; then rmdir 'api/gloo/gloo/grpc'; fi

# Create Enterprise Gloo directory
mkdir -p ./api/gloo/enterprise.gloo/v1
mv ./api/gloo/gloo/v1/enterprise/options/extauth/v1/extauth.proto ./api/gloo/enterprise.gloo/v1/auth_config.proto

# Create GraphQL directory
mkdir -p ./api/gloo/graphql.gloo/v1beta1
mv ./api/gloo/gloo/v1/enterprise/options/graphql/v1beta1/graphql.proto ./api/gloo/graphql.gloo/v1beta1/graphql.proto

# Fix paths
for file in $(find api/gloo -type f | grep ".proto")
do

# Gloo API changes
  sed "s|gloo/projects/gloo/api|solo-apis/api/gloo/gloo|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed "s|github.com/solo-io/gloo/projects/gloo/pkg/api|github.com/solo-io/solo-apis/pkg/api/gloo.solo.io|g" "$file" > "$file".tmp && mv "$file".tmp "$file"

# Gateway API changes
  sed "s|gloo/projects/gateway/api|solo-apis/api/gloo/gateway|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed "s|github.com/solo-io/gloo/projects/gateway/pkg/api|github.com/solo-io/solo-apis/pkg/api/gateway.solo.io|g" "$file" > "$file".tmp && mv "$file".tmp "$file"

# Gloo Enterprise API changes
  sed "s|gloo/v1/enterprise/options/extauth/v1/extauth.proto|enterprise.gloo/v1/auth_config.proto|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed "s|gloo.solo.io/v1/enterprise/options/extauth|enterprise.gloo.solo.io|g" "$file" > "$file".tmp && mv "$file".tmp "$file"

# GraphQL API changes
  sed "s|gloo/v1/enterprise/options/graphql/v1beta1/graphql.proto|graphql.gloo/v1beta1/graphql.proto|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed "s|gloo.solo.io/v1/enterprise/options/graphql|graphql.gloo.solo.io|g" "$file" > "$file".tmp && mv "$file".tmp "$file"

done

# convert protos used by skv1 into protos that can be used by skv2.
go run hack/convert.go

# Gateway2 changes (Go only, no protos)
mkdir -p ./pkg/api/gateway.gloo.solo.io/
rsync -ax --delete  ../gloo/projects/gateway2/api/  ./pkg/api/gateway.gloo.solo.io/
