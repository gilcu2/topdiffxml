#!/bin/bash

VERSION=$1
go mod tidy
kacl-cli release $VERSION --modify --auto-link
git commit -a -m "updated CHANGELOG for release"
git tag $VERSION
git push



