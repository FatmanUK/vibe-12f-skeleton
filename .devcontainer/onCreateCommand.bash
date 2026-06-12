#!/bin/bash
set -euo pipefail

apt-get update
apt-get install -y podman fuse-overlayfs

CONTENT_SUBIDS='vscode:1001:1
vscode:100000:65536'

rm -f /etc/subuid /etc/subgid
<<<"${CONTENT_SUBIDS}" install -m0644 /dev/stdin /etc/subuid
<<<"${CONTENT_SUBIDS}" install -m0644 /dev/stdin /etc/subgid

DIR_SRC_ROOT='.devcontainer/files'
DIR_DEST='/etc/containers'

install -m0644 ${DIR_SRC_ROOT}${DIR_DEST}/registries.conf ${DIR_DEST}
