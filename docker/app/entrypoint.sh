#!/bin/sh
SHELL=/bin/sh

git clone $REPOSITORY "$GOPATH/src/$GODOMAIN/dbsql-meta-viewer"


cd "$GOPATH/src/$GODOMAIN/dbsql-meta-viewer"

echo building client..
${SHELL} `pwd`/client/build.sh

echo building server..
${SHELL} `pwd`/server/build.sh

echo copy dists...
rm -rf ./dist

mkdir ./dist

cp -rf ./server/dist/*        ./dist/
cp -f  ./client/dist/*.html   ./dist/templates
cp -rf ./client/dist/static/* ./dist/static/

chown -R $FILEOWNER $DIST

cp -prf ./dist/* $DIST

