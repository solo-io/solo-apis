#!/bin/bash

set -e

rsync --delete -ax ../solo-projects/projects/gloo-fed/api/  ./api/gloo-fed

for file in $(find api/gloo-fed -type f | grep ".proto")
do

# Gloo Fed API changes
  sed "s|github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed|github.com/solo-io/solo-apis/pkg/api/fed|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed "s|github.com/solo-io/solo-projects/projects/gloo-fed/api/fed|github.com/solo-io/solo-apis/api/gloo-fed/fed|g" "$file" > "$file".tmp && mv "$file".tmp "$file"

# Skv2 Enterprise Api changes
  sed "s|github.com/solo-io/skv2-enterprise/multicluster-admission-webhook/api/multicluster|github.com/solo-io/solo-apis/api/skv2-enterprise|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
  
done
