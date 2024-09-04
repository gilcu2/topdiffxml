#!/bin/bash

VERSION=${1:-patch}
go mod tidy
kacl-cli release $VERSION --modify --auto-link
NEWVERSION=$(kacl-cli get)
git commit -a -m "updated CHANGELOG for release $NEWVERSION"
git push
git tag v$NEWVERSION
git push origin v$NEWVERSION



