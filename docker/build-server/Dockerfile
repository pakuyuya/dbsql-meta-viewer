FROM golang:1.11.5-alpine3.9
MAINTAINER pakuyuya

ARG REPOSITORY=https://github.com/pakuyuya/dbsql-meta-viewer.git \
    GODOMAIN=github.com/pakuyuya

ENV DIST=/var/dist \
    GOPATH=/go

RUN apk update \
    && apk add git dep \
    && mkdir -p "$GOPATH/src/$GODOMAIN"

RUN git clone $REPOSITORY "$GOPATH/src/$GODOMAIN/dbsql-meta-viewer" \
    && cd "$GOPATH/src/$GODOMAIN/dbsql-meta-viewer/server" \
    && /bin/ash ./build.sh \
    && cp -rf ./dist /var/dist

