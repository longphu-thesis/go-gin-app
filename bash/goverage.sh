#!/bin/bash

. ./bash/init_env.sh

goverage -v -covermode=count -coverprofile=coverage.cov "$GO_FILES_IGNORE_VENDOR"