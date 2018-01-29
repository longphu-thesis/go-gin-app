#!/bin/bash

#GO_FILES_IGNORE_VENDOR=$(go list ./... | grep -v /vendor/)

#go test -cover -covermode=count -coverprofile=.coverprofile -coverpkg=$(go list ./... | grep -v /vendor/)
#goveralls -coverprofile=.coverprofile -service=travis-ci -repotoken $COVERALLS_TOKEN
goveralls -service=travis-ci