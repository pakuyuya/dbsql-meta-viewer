#!/bin/bash

BASEDIR=`dirname "$0"`

npm install --no-audit --prefix $BASEDIR/
npm run build --prefix $BASEDIR/

echo build client is successful

