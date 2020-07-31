#!/bin/bash

set -e

for file in $(find pkg/api/gloo.solo.io/v1 -type f | grep "_json.gen.go")
do
  sed "s|github_com_gogo_protobuf_jsonpb.Marshaler{}|github_com_gogo_protobuf_jsonpb.Marshaler{EnumsAsInts: true}|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
done

for file in $(find pkg/api/gateway.solo.io/v1 -type f | grep "_json.gen.go")
do
  sed "s|github_com_gogo_protobuf_jsonpb.Marshaler{}|github_com_gogo_protobuf_jsonpb.Marshaler{EnumsAsInts: true}|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
done
