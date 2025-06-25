#!/usr/bin/env bash

# Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/onexstack/onex.
#


# This file is not intended to be run automatically. It is meant to be run
# immediately before exporting docs. We do not want to check these documents in
# by default.

set -o errexit
set -o nounset
set -o pipefail

PROJ_ROOT_DIR=$(dirname "${BASH_SOURCE[0]}")/..
source "${PROJ_ROOT_DIR}/scripts/lib/init.sh"

KINDS=(deployment service)

for component in $(ls ${PROJ_ROOT_DIR}/cmd)
do
  truncate -s 0 ${PROJ_ROOT_DIR}/deployments/${component}.yaml

  for kind in ${KINDS[@]}
  do
    echo -e "---\n# Source: deployments/${component}-${kind}.yaml" >> ${PROJ_ROOT_DIR}/deployments/${component}.yaml
    sed '/^#\|^$/d' ${PROJ_ROOT_DIR}/deployments/${component}-${kind}.yaml >> ${PROJ_ROOT_DIR}/deployments/${component}.yaml
  done

  onex::log::info "generate ${PROJ_ROOT_DIR}/deployments/${component}.yaml success"
done
