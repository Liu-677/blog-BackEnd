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

  go build -x -o $GOPATH/src/blog-BackEnd/build/docker
  popd
}

build_docker(){
  docker build -t blog-server:v1.0 $GOPATH/src/blog-BackEnd/build/docker/
  docker run blog-server
}

build_binary
build_docker