#!/usr/bin/env bash

APP=nft
OS=$1
ARCH=$2
VERSION=$3

DEST=build/${OS}/${ARCH}
TAR=${APP}.${VERSION}.${OS}-${ARCH}.tar.gz
BIN=${APP}

cd `dirname $0`

CGO_ENABLED=0 GOOS=$OS GOARCH=$ARCH go build -o ${DEST}/${BIN} ./cmd/${APP}
cd ${DEST}
tar -czf ${TAR} ${BIN}
rm ${BIN}
