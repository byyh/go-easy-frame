#!/bin/sh

OS="linux"
ARCH="amd64"

if [ "$1" == "mac" -o "$1" == "windows" ]; then
    OS=$1
    if [ "$OS" == "mac" ]; then
        OS="darwin"
    fi
fi
echo "build for ${OS}, arch: ${ARCH}..."

export GOOS=$OS
export GOARCH=$ARCH
export CGO_ENABLED=0 
export GOPATH=/windows/go
export GO111MODULE=on
export GOPROXY=https://goproxy.io

make clean
make build_web
make build_cron
make build_consumer