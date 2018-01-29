#!/bin/bash

# -repotoken $COVERALLS_TOKEN
goveralls -coverprofile=coverage.out -service=travis-ci