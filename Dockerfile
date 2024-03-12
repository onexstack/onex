# syntax=docker/dockerfile:1.4

# Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/superproj/onex.

# Dockerfile generated by scripts/gen-dockerfile.sh. DO NOT EDIT.

# Build the onex-usercenter binary
# Run this with docker build --build-arg prod_image=<golang:x.y.z>
# Default <prod_image> is debian:jessie
ARG prod_image=debian:jessie

FROM ${prod_image}
LABEL maintainer="<colin404@foxmail.com>"

WORKDIR /opt/onex

# Note: the <prod_image> is required to support
# setting timezone otherwise the build will fail
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
      echo "Asia/Shanghai" > /etc/timezone

COPY onex-usercenter /opt/onex/bin/

ENTRYPOINT ["/opt/onex/bin/onex-usercenter"]
