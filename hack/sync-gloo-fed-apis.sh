#!/bin/bash

set -e

rsync -ax ../solo-projects/projects/gloo-fed/api/  ./api/gloo-fed

# Change the paths in the Gloo Fed protos from solo-projects to solo-apis
for file in $(find api/gloo-fed -type f | grep ".proto")
do
  # rename the go_package from fed protos to have a solo-apis path
  sed "s|github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/|github.com/solo-io/solo-apis/pkg/api/|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
  # rename the proto imports to refer to the solo-apis path
  sed "s|github.com/solo-io/solo-projects/projects/gloo-fed/api/|github.com/solo-io/solo-apis/api/gloo-fed/|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
done
