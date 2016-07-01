#!/bin/bash

set -e
set -x

export GOPATH=`pwd`

GOARCH=amd64 GOOS=linux go build -o cfd