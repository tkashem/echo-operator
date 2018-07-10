#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

vendor/k8s.io/code-generator/generate-groups.sh \
deepcopy \
github.com/tkashem/echo-operator/pkg/generated \
github.com/tkashem/echo-operator/pkg/apis \
echo:v1 \
--go-header-file "./tmp/codegen/boilerplate.go.txt"
