#!/usr/bin/env bash
set -e

containerd --config config.toml &
ctr namespace create services.linuxkit
ctr -n services.linuxkit image pull docker.io/linuxkit/kubelet:496c57f6b7d7fbea671823aa17e5d5c80ab23501
ctr -n services.linuxkit run --net-host -d docker.io/linuxkit/kubelet:496c57f6b7d7fbea671823aa17e5d5c80ab23501 kubelet
ctr --namespace services.linuxkit tasks exec --tty --exec-id mkdir kubelet mkdir -p /var/lib/kubeadm


exec "$@"
