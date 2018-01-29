#!/bin/bash

# . ./export.sh go19
GO_VERSION="$1"
GO_PACKAGE="$HOME/package_$GO_VERSION"
BASENAME=$(basename $PWD)
REPO="github.com/longphu-thesis"
WORK_DIR="$GO_PACKAGE/src/$REPO"

if [ "$GO_VERSION" != "" ];
then
    export GOROOT="$HOME/$GO_VERSION"
    export GOPATH="$GO_PACKAGE:$GOPATH"
    export GOBIN="$GO_PACKAGE/bin"
    export PATH="$GOROOT/bin:$PATH:$GO_PACKAGE/bin"

    mkdir -p "$WORK_DIR"
    ln -s $PWD "$WORK_DIR/$BASENAME"

    echo "========================================================="
    echo "GOROOT : $GOROOT"
    echo "GOPATH : $GOPATH"
    echo "GOBIN : $GOBIN"
    echo "PATH : $PATH"
    echo "========================================================="
    ls -l "$WORK_DIR/$BASENAME"
    echo "========================================================="

    cd "$WORK_DIR/$BASENAME"
else
    echo "run [. export.sh go19]"
fi