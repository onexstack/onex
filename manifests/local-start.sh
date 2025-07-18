#!/bin/bash

set -euo pipefail  # 安全设置以捕获错误

PROJ_ROOT_DIR=$(dirname "${BASH_SOURCE[0]}")/..
BIN_DIR="${PROJ_ROOT_DIR}/_output/platforms/linux/amd64"
CONFIG_DIR="${HOME}/.onex"

# 检查二进制文件是否存在
check_binary() {
  local binary="$1"
  if [ ! -x "${binary}" ]; then
    echo "Error: ${binary} not found or not executable"
    exit 1
  fi
}

# 启动 API 服务器
start_api_server() {
  check_binary "${BIN_DIR}/onex-apiserver"

  ${BIN_DIR}/onex-apiserver \
    --secure-port=52443 \
    --bind-address=0.0.0.0 \
    --etcd-servers=127.0.0.1:2379 \
    --client-ca-file="${CONFIG_DIR}/cert/ca.pem" \
    --tls-cert-file="${CONFIG_DIR}/cert/onex-apiserver.pem" \
    --tls-private-key-file="${CONFIG_DIR}/cert/onex-apiserver-key.pem" \
    --enable-admission-plugins=NamespaceAutoProvision,NamespaceExists,NamespaceLifecycle \
    --v=10
}

# 启动网关
start_gateway() {
  check_binary "${BIN_DIR}/onex-gateway"

  ${BIN_DIR}/onex-gateway --config="${CONFIG_DIR}/onex-gateway.yaml"
}

# 启动 onex-controller-manager
start_onex_controller_manager() {
  check_binary "${BIN_DIR}/onex-controller-manager"

  ${BIN_DIR}/onex-controller-manager --kubeconfig="${CONFIG_DIR}/config" --config="${CONFIG_DIR}/onex-controller-manager.yaml" --v 10 --controllers='*'
}

# 启动 onex-blockchain-controller
start_onex_blockchain_controller() {
  check_binary "${BIN_DIR}/onex-blockchain-controller"

  ${BIN_DIR}/onex-blockchain-controller --v 10 --kubeconfig="${CONFIG_DIR}/config" --config ${HOME}/.onex/onex-blockchain-controller.yaml
}

# 启动 onex-job-controller
start_onex_job_controller() {
  check_binary "${BIN_DIR}/onex-job-controller"

  ${BIN_DIR}/onex-job-controller --v 10 --kubeconfig="${CONFIG_DIR}/config" --config ${HOME}/.onex/onex-job-controller.yaml
}

# 主逻辑
case "$1" in
  api)
    start_api_server
    ;;
  cm)
    start_onex_controller_manager
    ;;
  gw)
    start_gateway
    ;;
  blc-ctrl)
    start_onex_blockchain_controller
    ;;
  job)
    start_onex_job_controller
    ;;
  *)
    echo "Usage: $0 {api|gw}"
    exit 1
    ;;
esac
