#!/bin/ash
git clone $REPOSITORY "$GOPATH/src/$GODOMAIN/dbsql-meta-viewer"

cd "$GOPATH/src/$GODOMAIN/dbsql-meta-viewer"

SHELL=/bin/ash
$SHELL ./build.sh

cp -prf ./dist/* $DIST
chmod -R 777 $DIST

