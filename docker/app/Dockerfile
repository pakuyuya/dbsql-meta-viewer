FROM dbsql-meta-viewer-base:latest
MAINTAINER pakuyuya

ARG REVISION=HEAD
ARG REPOSITORY=https://github.com/pakuyuya/dbsql-meta-viewer.git
ARG GODOMAIN=github.com/pakuyuya
ARG FILEOWNER=root:root

ENV GOPATH=/go \
    NODEPATH=/node_modules \
    APPDIR=/var/app \
    DATASSRC=/var/datas \
    DATASDIR=$APPDIR/datas

RUN cd "$GOPATH/src/$GODOMAIN/dbsql-meta-viewer" \
    && git pull origin \
    && echo copy dists... \
    && rm -rf $APPDIR \
    && mkdir  $APPDIR \
    && cp -rf ./server/dist/*        $APPDIR/ \
    && cp -f  ./client/dist/*.html   $APPDIR/templates \
    && cp -rf ./client/dist/static/* $APPDIR/static/ \
    && mkdir  $DATASDIR

