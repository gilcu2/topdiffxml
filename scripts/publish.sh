#!/bin/bash

VERSION=v${1:-$(kacl-cli current)}
echo Publising $VERSION
GOPROXY=proxy.golang.org go list -m github.com/gilcu2/topdiffxml@$VERSION
goreleaser release --clean



