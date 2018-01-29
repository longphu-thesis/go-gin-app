#!/bin/bash

# $0 is the script name, $1 id the first ARG, $2 is second...
# . ./init.sh go19 go1.9.2.linux-amd64.tar.gz

GO_VERSION="$1"
NAME="$2"
GO_PACKAGE="$HOME/package_$GO_VERSION"
GO_ROOT="$HOME/$GO_VERSION"

setup_go () {
   if [ -f "$NAME" ]; then
        rm "$NAME"
    fi
    # Download file with [name] file copy in page https://golang.org/dl/
    # wget https://redirector.gvt1.com/edgedl/go/go1.9.2.linux-amd64.tar.gz
    wget https://redirector.gvt1.com/edgedl/go/$NAME

    # tar file to folder
    tar -C ./ -xzf ./$NAME

    # move go to folder goversion
    mv ./go "$GO_ROOT"

    if [ ! -d "$GO_PACKAGE" ]; then
        mkdir "$GO_PACKAGE"
    else
        echo "Please check $GO_PACKAGE"
    fi

    rm "$NAME"
}

if [ "$GO_VERSION" != "" ] && [ "$NAME" != "" ];
then
    # check if a directory doesn't exist:
    if [ ! -d "$GO_ROOT" ];
    then
        setup_go
        . ./export.sh "$GO_VERSION"
        go version
        echo "Init Done"

        bash ./bash/go_get_dependency.sh
        bash ./bash/go_get_dependency_govendor.sh
    else
        echo "folder go exist please check and remove $GO_ROOT"
    fi
else
    echo "run [. ./init.sh go19 go1.9.2.linux-amd64.tar.gz]"
fi