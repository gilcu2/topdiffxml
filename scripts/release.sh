#!/bin/bash

VERSION=$1
go mod tidy
kacl-cli release $VERSION --modify --auto-link
git commit -a -m "updated CHANGELOG for release $VERSION"
git tag v$VERSION
git push origin v$VERSION



