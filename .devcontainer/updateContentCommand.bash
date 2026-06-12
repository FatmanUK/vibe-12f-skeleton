#!/bin/bash
set -euo pipefail

DIR_SRC_ROOT='.devcontainer/files'
DIR_DEST_ROOT='/home/vscode'

DIR_DEST="${DIR_DEST_ROOT}/.config/containers"
DIR_SRC="${DIR_SRC_ROOT}/.config/containers"

install -m0755 -d ${DIR_DEST}
install -m0644    ${DIR_DEST} ${DIR_SRC}/storage.conf
install -m0644    ${DIR_DEST} ${DIR_SRC}/containers.conf
