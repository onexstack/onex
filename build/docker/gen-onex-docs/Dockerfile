# syntax=docker/dockerfile:1.4

# Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/onexstack/onex.

# Dockerfile generated by scripts/gen-dockerfile.sh. DO NOT EDIT.

# Build the gen-onex-docs binary
# Run this with docker build --build-arg prod_image=<golang:x.y.z>
# Default <prod_image> is debian:trixie
ARG prod_image=debian:trixie

FROM ${prod_image}
LABEL maintainer="<colin404@foxmail.com>"

WORKDIR /opt/onex

# Note: the <prod_image> is required to support
# setting timezone otherwise the build will fail
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
      echo "Asia/Shanghai" > /etc/timezone

COPY gen-onex-docs /opt/onex/bin/

ENTRYPOINT ["/opt/onex/bin/gen-onex-docs"]
