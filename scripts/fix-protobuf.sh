#!/bin/bash

GOMODCACHE=$(go env GOMODCACHE)

sudo rm -rf ${GOMODCACHE}/k8s.io/apiextensions-apiserver*
sudo rm -rf ${GOMODCACHE}/k8s.io/api*
sudo rm -rf ${GOMODCACHE}/k8s.io/apimachinery*

go get k8s.io/apimachinery@v0.31.2
go get k8s.io/api@v0.31.2
go get k8s.io/apiextensions-apiserver@v0.31.2
sudo chmod 777 -R  ${GOMODCACHE}/k8s.io/*
