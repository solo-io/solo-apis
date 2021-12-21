#!/bin/bash

set -e

jsonGenFile="json.gen.go"
# There are two templates from skv2 that are modified for gloo-fed. We generate both versions and delete the skv2 versions here.
# See codegen/gloo.go for more details
echo "Cleaning up duplicated generated code"
glooJsonGenFile="pkg/api/gloo.apis.solo.io/v1/$jsonGenFile"
if [ -f "$glooJsonGenFile" ] ; then
    rm "$glooJsonGenFile"
fi

gatewayJsonGenFile="pkg/api/gateway.apis.solo.io/v1/$jsonGenFile"
if [ -f "$gatewayJsonGenFile" ] ; then
    rm "$gatewayJsonGenFile"
fi

glooEnterpriseJsonGenFile="pkg/api/enterprise.gloo.apis.solo.io/v1/$jsonGenFile"
if [ -f "$glooEnterpriseJsonGenFile" ] ; then
    rm "$glooEnterpriseJsonGenFile"
fi

registerFile="register.go"
glooRegisterFile="pkg/api/gloo.apis.solo.io/v1/$registerFile"
if [ -f "$glooRegisterFile" ] ; then
    rm "$glooRegisterFile"
fi

gatewayRegisterFile="pkg/api/gateway.apis.solo.io/v1/$registerFile"
if [ -f "$gatewayRegisterFile" ] ; then
    rm "$gatewayRegisterFile"
fi

enterpriseRegisterFile="pkg/api/enterprise.gloo.apis.solo.io/v1/$registerFile"
if [ -f "$enterpriseRegisterFile" ] ; then
    rm "$enterpriseRegisterFile"
fi
