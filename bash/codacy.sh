#!/bin/bash

. ./bash/init_env.sh

goverage -v -coverprofile=coverage.out "$GO_FILES_IGNORE_VENDOR"