#!/bin/sh
set -e 
mkdir taskset
base64 -d tasksetencode/testtask.go > taskset/testtask.go 
cat taskset/testtask.go 
go fmt ./taskset
make build_agent
make/release/ymir agent --debug --default-timeout=5s