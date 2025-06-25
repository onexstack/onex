# Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/onexstack/onex.
#

# build-allinone-image.sh 在 build/docker/onex-allinone 目录中以 build.sh 软连接
# 的形式存在，是为了防止误删 build/docker 目录
# Copy of scripts/build-allinone-image.sh
PROJ_ROOT_DIR=$(dirname "${BASH_SOURCE[0]}")/../../..
ONEX_ENV_FILE=${ONEX_ENV_FILE:-${PROJ_ROOT_DIR}/manifests/env.local}

source "${PROJ_ROOT_DIR}/scripts/lib/init.sh"
source ${ONEX_ENV_FILE}

cd ${PROJ_ROOT_DIR}

# 生成构建Dockerfile需要的构建产物
make build
make gen.systemd
make gen.appconfig
make gen.ca
make gen.kubeconfig

# 复制构建产物到指定目录
mkdir -p ${DST_DIR}/bin
cp ${OUTPUT_DIR}/platforms/${IMAGE_PLAT}/* ${DST_DIR}/bin/
cp -r ${OUTPUT_DIR}/appconfig ${DST_DIR}/
cp -r ${OUTPUT_DIR}/cert ${DST_DIR}/
cp -r ${OUTPUT_DIR}/config ${DST_DIR}/
cp -r ${OUTPUT_DIR}/systemd ${DST_DIR}/

