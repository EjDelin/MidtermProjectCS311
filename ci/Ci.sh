#!/bin/bash
set -e

go generate ./...

go mod tidy

go test ./... -v

go build ./...

