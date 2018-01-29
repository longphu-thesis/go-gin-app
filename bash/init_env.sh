#!/bin/bash

GO_FILES=$(find . -iname '*.go' -type f | grep -v /vendor/) # All the .go files, excluding vendor/
GO_FILES_IGNORE_VENDOR=$(go list ./... | grep -v /vendor/) # All the .go files, excluding vendor/

echo "$GO_FILES"
echo "$GO_FILES_IGNORE_VENDOR"