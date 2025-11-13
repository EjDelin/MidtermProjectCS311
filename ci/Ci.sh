#!/bin/bash
set -e #it will stop if there is an error

# Generate enum, mocks, and static files
go generate ./...

# Install dependencies
go mod tidy

# Run tests (mock usage example)
go test ./... -v

go build ./...

