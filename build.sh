#!/bin/bash

set -e
set -x

export GOPATH=`pwd`

if [ ! -d ./src/github.com/LaynePeng/ ]; then
	mkdir -p ./src/github.com/LaynePeng
	ln -s `pwd` ./src/github.com/LaynePeng
fi

GOARCH=amd64 GOOS=linux go build -o cfd