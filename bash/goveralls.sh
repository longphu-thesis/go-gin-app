#!/bin/bash

. ./bash/init_env.sh

# -repotoken $COVERALLS_TOKEN
goveralls -coverprofile=coverage.cov -service=travis-ci