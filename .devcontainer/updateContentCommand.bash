#!/bin/bash
set -euo pipefail

DIR_DEST='/home/vscode/.config/containers'
DIR_SRC='.devcontainer/files/home/vscode/.config/containers'
install -d ${DIR_DEST}
install -m0644 ${DIR_SRC}/storage.conf    ${DIR_DEST}/
install -m0644 ${DIR_SRC}/containers.conf ${DIR_DEST}/
