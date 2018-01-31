#!/bin/bash

# -repotoken $COVERALLS_TOKEN
goveralls -coverprofile=coverage.txt -service=travis-ci