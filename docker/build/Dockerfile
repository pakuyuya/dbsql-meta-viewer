FROM golang:1.11.5-alpine3.9
MAINTAINER pakuyuya

ENV REPOSITORY=https://github.com/pakuyuya/dbsql-meta-viewer.git \
    GODOMAIN=github.com/pakuyuya \
    DIST=/var/dist \
    GOPATH=/go \
    NODEPATH=/node_modules \
    FILEOWNER=root:root

ADD ./entrypoint.sh /

RUN chmod +x /entrypoint.sh \
    && apk update \
    && apk add --no-cache git dep nodejs=10.14.2-r0 npm=10.14.2-r0 \
    && mkdir -p "$GOPATH/src/$GODOMAIN"


ENTRYPOINT ["/entrypoint.sh"]
