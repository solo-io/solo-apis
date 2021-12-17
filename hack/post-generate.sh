#!/bin/bash

set -e

jsonGenFile="json.gen.go"
echo "Cleaning up extra json.gen.go files"
glooJsonGenFile="pkg/api/gloo.apis.solo.io/v1/$jsonGenFile"
if [ -f "$glooJsonGenFile" ] ; then
    rm "$glooJsonGenFile"
fi

gatewayJsonGenFile="pkg/api/gateway.solo.io/v1/$jsonGenFile"
if [ -f "$gatewayJsonGenFile" ] ; then
    rm "$gatewayJsonGenFile"
fi

glooEnterpriseJsonGenFile="pkg/api/enterprise.gloo.apis.solo.io/v1/$jsonGenFile"
if [ -f "$glooEnterpriseJsonGenFile" ] ; then
    rm "$glooEnterpriseJsonGenFile"
fi
#TODO: finish cleaning up extra files instead of deleting manually
glooRegisterFile=""
if [ -f "$glooRegisterFile" ] ; then
    rm "$glooRegisterFile"
fi