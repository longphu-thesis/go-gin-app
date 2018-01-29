#!/bin/bash

GO_FILES=$(find . -iname '*.go' -type f | grep -v /vendor/)

# Fail if a .go file hasn't been formatted with gofmt
test -z $(gofmt -s -l $GO_FILES)