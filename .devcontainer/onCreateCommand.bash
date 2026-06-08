#!/bin/bash
set -euo pipefail

apt-get update
apt-get install -y podman fuse-overlayfs
rm -f /etc/subuid /etc/subgid
touch /etc/subuid /etc/subgid
pwd
install .devcontainer/files/etc/containers/registries.conf /etc/containers/
