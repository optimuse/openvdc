#!/bin/bash

export ENV_ROOTDIR="$(cd "$(dirname $(readlink -f "$0"))/.." && pwd -P)"
export NODE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TMP_ROOT="${NODE_DIR}/tmp_root"

. "${ENV_ROOTDIR}/config.source"
. "${NODE_DIR}/vmspec.conf"
. "${ENV_ROOTDIR}/ind-steps/common.source"

zk_host=true

IND_STEPS=(
    "box"
    "ssh"
    "hosts"
    "disable-firewalld"
    "mesosphere-repo"
    "zookeeper"
)

build "${IND_STEPS[@]}"
