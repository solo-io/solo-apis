#!/bin/bash

set -e

jsonGenFile="json.gen.go"

glooJsonGenFile="pkg/api/gloo.solo.io/v1/$jsonGenFile"
if [ -f "$glooJsonGenFile" ] ; then
    rm "$glooJsonGenFile"
fi

gatewayJsonGenFile="pkg/api/gateway.solo.io/v1/$jsonGenFile"
if [ -f "$gatewayJsonGenFile" ] ; then
    rm "$gatewayJsonGenFile"
fi

glooEnterpriseJsonGenFile="pkg/api/enterprise.gloo.solo.io/v1/$jsonGenFile"
if [ -f "$glooEnterpriseJsonGenFile" ] ; then
    rm "$glooEnterpriseJsonGenFile"
fi

glooGraphqlJsonGenFile="pkg/api/graphql.gloo.solo.io/v1beta1/$jsonGenFile"
if [ -f "$glooGraphqlJsonGenFile" ] ; then
    rm "$glooGraphqlJsonGenFile"
fi