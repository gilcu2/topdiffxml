#!/bin/bash

#go test -coverpkg=./... -coverprofile=coverage.out ./...
go test -coverprofile=coverage.out ./...
go tool cover -func coverage.out
#go tool cover -html coverage.out


