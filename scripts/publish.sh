#!/bin/bash

VERSION=v$1

GOPROXY=proxy.golang.org go list -m github.com/gilcu2/topdiffxml@$VERSION
goreleaser release --clean



