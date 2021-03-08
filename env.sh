#!/bin/bash
# 获取工作路径
homePath="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
workPath=$PWD
echo work dir [$homePath] [$workPath]

export GOPATH=$homePath/
export GOROOT=$GOPATH/go

[ -d $GOROOT ] || {
    echo "no go env"
    wget https://golang.org/dl/go1.16.linux-amd64.tar.gz
    tar -zxvf go1.16.linux-amd64.tar.gz
}
export PATH=$GOROOT/bin:$PATH
