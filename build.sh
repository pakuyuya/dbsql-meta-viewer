#!/bin/bash
cd `dirname $0`

pwd

echo building client..
./client/build.sh

echo building server..
./server/build.sh

echo copy dists...
rm -rf ./dist

mkdir ./dist

cp -rf ./server/dist/*        ./dist/
cp -f  ./client/dist/*.html   ./dist/templates
cp -rf ./client/dist/static/* ./dist/static/

