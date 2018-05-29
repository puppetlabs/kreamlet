#!/usr/bin/env bash
set -e

containerd --config config.toml &
exec "$@"
