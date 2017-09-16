#!/bin/sh
set -e
go fmt ./taskset
make build_agent
make/release/ymir agent --debug --default-timeout=5s