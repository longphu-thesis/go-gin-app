#!/bin/bash

GO_FILES=$(find . -iname '*.go' -type f | grep -v /vendor/) # All the .go files, excluding vendor/
GO_FILES_IGNORE_VENDOR=$(go list ./... | grep -v /vendor/) # All the .go files, excluding vendor/

go vet $GO_FILES_IGNORE_VENDOR                              # go vet is the official Go static analyzer
megacheck $GO_FILES_IGNORE_VENDOR                           # "go vet on steroids" + linter
gocyclo -over 19 $GO_FILES                                  # forbid code with huge functions
golint -set_exit_status $GO_FILES_IGNORE_VENDOR             # one last linter