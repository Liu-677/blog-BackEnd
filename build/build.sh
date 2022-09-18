#!/bin/bash
set -e


build_binary(){
  echo $GOPATH
  mkdir $GOPATH/bin
  mkdir $GOPATH/pkg
  pushd $GOPATH/src/blog-BackEnd
  # 下载好需要的库
  go mod tidy

  # 编译
  go build main.go -o $GOPATH/bin/blog-backend
  popd
}

build_docker(){
  docker build -t blog-server
  docker run blog-server
}

build_binary
build_docker