#!/usr/bin/env bash

CURRENT_DIR=`pwd`
OLD_GO_PATH="/opt/homebrew/bin/go"  #例如: /usr/local/go
OLD_GO_BIN="/opt/homebrew/bin/go/bin"    #例如: /usr/local/go/bin

export GOPATH="$CURRENT_DIR"
export GOBIN="$CURRENT_DIR/bin"

#指定并整理当前的源码路径
gofmt -w src

go install hello.go

export GOPATH="$OLD_GO_PATH"
export GOBIN="$OLD_GO_BIN"