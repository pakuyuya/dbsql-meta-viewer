#!/bin/bash
cd `dirname $0`

if [[ ! -z "${SHELL}" ]]; then
   echo using shell as "/bin/bash"
   SHELL=/bin/bash
fi

echo building client..
${SHELL} ./client/build.sh

echo building server..
${SHELL} ./server/build.sh

echo copy dists...
rm -rf ./dist

mkdir ./dist

cp -rf ./server/dist/*        ./dist/
cp -f  ./client/dist/*.html   ./dist/templates
cp -rf ./client/dist/static/* ./dist/static/

