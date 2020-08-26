#!/bin/bash

set -e

GV="ton:v1alpha1"

rm -rf ./pkg/client
./hack/generate-groups.sh "client,lister,informer" ton/pkg/client ton/pkg/apis "$GV" --output-base=./  -h "$PWD/hack/boilerplate.go.txt"
mv ton/pkg/client ./pkg/
rm -rf ./ton