#!/bin/bash

# . ./export.sh go19
GO_VERSION="$1"
GO_PACKAGE="$HOME/package_$GO_VERSION"
BASENAME=$(basename $PWD)

if [ "$GO_VERSION" != "" ];
then
    export GOROOT="$HOME/$GO_VERSION"
    export GOPATH="$GO_PACKAGE:$GOPATH"
    export GOBIN="$GO_PACKAGE/bin"
    export PATH="$GOROOT/bin:$PATH:$GO_PACKAGE/bin"

    mkdir "$GO_PACKAGE/src"
    ln -s $PWD "$GO_PACKAGE/src/$BASENAME"

    echo "========================================================="
    echo "GOROOT : $GOROOT"
    echo "GOPATH : $GOPATH"
    echo "GOBIN : $GOBIN"
    echo "PATH : $PATH"
    echo "========================================================="
    ls -l "$GO_PACKAGE/src"
    echo "========================================================="

    cd "$GO_PACKAGE/src/$BASENAME"
else
    echo "run [. export.sh go19]"
fi