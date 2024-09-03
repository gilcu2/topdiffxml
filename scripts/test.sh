#!/bin/bash

VERSION=$1
kacl-cli release $VERSION --modify --auto-link
git commit -a -m "updated CHANGELOG for release"
git tag $VERSION

