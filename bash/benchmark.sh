#!/bin/bash

GO_FILES_IGNORE_VENDOR=$(go list ./... | grep -v vendor)
#CURRENT_WORKING_DIRECTORY=$PWD
#PARENT_DIR="$(dirname "$CURRENT_WORKING_DIRECTORY")"
for d in $GO_FILES_IGNORE_VENDOR; do
    cd "$GOPATH/src/$d"
    go test -bench=.
done