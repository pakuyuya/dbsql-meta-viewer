FROM golang:1.11.5-alpine3.9
MAINTAINER pakuyuya

ENV REPOSITORY=https://github.com/pakuyuya/dbsql-meta-viewer.git \
    GODOMAIN=github.com/pakuyuya \
    GOPATH=/go \
    NODEPATH=/node_modules \
    FILEOWNER=root:root \
    APPDIR=/var/app \
    DATASSRC=/var/datas \
    DATASDIR=$APPDIR/datas

RUN apk update \
    && apk add --no-cache git dep nodejs=10.14.2-r0 npm=10.14.2-r0 \
    && mkdir -p "$GOPATH/src/$GODOMAIN" \
    \
    && git clone $REPOSITORY "$GOPATH/src/$GODOMAIN/dbsql-meta-viewer" \
    && cd "$GOPATH/src/$GODOMAIN/dbsql-meta-viewer" \
    \
    && echo building client.. \
    && /bin/sh `pwd`/client/build.sh \
    \
    && echo building server.. \
    && /bin/sh `pwd`/server/build.sh \
    \
    && echo copy dists... \
    && rm -rf $APPDIR \
    && mkdir  $APPDIR \
    && cp -rf ./server/dist/*        $APPDIR/ \
    && cp -f  ./client/dist/*.html   $APPDIR/templates \
    && cp -rf ./client/dist/static/* $APPDIR/static/ \
    && mkdir  $DATASDIR


