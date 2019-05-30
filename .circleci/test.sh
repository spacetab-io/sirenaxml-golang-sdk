#!/bin/bash  -eo pipefail

mkdir -p ./artifacts
go mod vendor
go test ./... -coverprofile=c.out
go tool cover -html=c.out -o coverage.html
mv coverage.html ./artifacts
mv c.out ./artifacts