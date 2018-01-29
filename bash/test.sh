#!/bin/bash

. ./bash/init_env.sh

# go test -v -race ./..

# Run all the tests with the race detector enabled
go test -v -race  $GO_FILES_IGNORE_VENDOR
# go test -v -race  $(go list ./... | grep -v /vendor/)
