#!/bin/bash
set -e


build_binary(){
  echo $GOPATH
  mkdir -p $GOPATH/bin
  mkdir -p $GOPATH/pkg
  pushd $GOPATH/src/blog-BackEnd
  # 下载好需要的库
  go mod tidy

  # 编译
  go build  -o $GOPATH/bin/blog-backend
  popd
}

build_docker(){
  docker build -t blog-server:v1.0 ./docker
  docker run blog-server
}

build_binary
build_docker