#!/bin/bash
cd `dirname $0`
docker build --rm -t dbsql-meta-viewer_build-server:latest .
docker run --rm -it -v `pwd`/dist:/var/dist dbsql-meta-viewer_build-server

