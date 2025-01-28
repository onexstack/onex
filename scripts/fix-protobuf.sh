#!/bin/bash

PROJ_ROOT_DIR=$(dirname "${BASH_SOURCE[0]}")/..
source "${PROJ_ROOT_DIR}/scripts/lib/init.sh"

GOMODCACHE=$(go env GOMODCACHE)

sudo rm -rf ${GOMODCACHE}/k8s.io/apiextensions-apiserver*
sudo rm -rf ${GOMODCACHE}/k8s.io/api*
sudo rm -rf ${GOMODCACHE}/k8s.io/apimachinery*

go get k8s.io/apimachinery@$(egrep 'k8s.io/apimachinery .*=>' ${PROJ_ROOT_DIR}/go.mod |awk '{print $NF}')
go get k8s.io/api@$(egrep 'k8s.io/api .*=>' ${PROJ_ROOT_DIR}/go.mod |awk '{print $NF}')
go get k8s.io/apiextensions-apiserver@$(egrep 'k8s.io/apiextensions-apiserver .*=>' ${PROJ_ROOT_DIR}/go.mod |awk '{print $NF}')
sudo chmod 777 -R  ${GOMODCACHE}/k8s.io/*
