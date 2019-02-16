FROM node:8.15.0-alpine
MAINTAINER pakuyuya

ARG REPOSITORY=https://github.com/pakuyuya/dbsql-meta-viewer.git

ENV DIST=/var/dist 

RUN apk update \
    && apk add git

RUN git clone $REPOSITORY "~/dbsql-meta-viewer" \
    && cd "~/dbsql-meta-viewer/client" \
    && /bin/ash ./build.sh \
    && cp -rf ./dist /var/dist

