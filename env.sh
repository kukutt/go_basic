#!/bin/bash
# 获取工作路径
homePath="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
workPath=$PWD
echo work dir [$homePath] [$workPath]

export GO111MODULE=on
export GOPATH=$homePath/
export GOROOT=$GOPATH/go
export GOPROXY=https://goproxy.cn

[ -d $GOROOT ] || {
    echo "no go env"
    pushd $GOPATH
    wget https://go.dev/dl/go1.16.linux-amd64.tar.gz
    tar -zxvf go1.16.linux-amd64.tar.gz
    popd
}
export PATH=$GOROOT/bin:$PATH
