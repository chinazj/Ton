#!/bin/bash

#"client,lister,informer"


set -e

GV="ton:v1alpha1"

rm -rf ./pkg/client
./hack/generate-groups.sh all ton/pkg/client ton/pkg/apis "$GV" --output-base=./  -h "$PWD/hack/boilerplate.go.txt"
mv ton/pkg/client ./pkg/
mv ton/pkg/apis/ton/v1alpha1/. ./pkg/apis/ton/v1alpha1/.
rm -rf ./ton