#!/bin/bash

# load environments from .env file
PRIVATE_DOCKER_REPOS=localhost:5000
PRIVATE_DOCKER_ORG=pakuyuya

docker build -t dbsql-meta-viewer-base:latest ./base

# in use private repository
#docker tag  dbsql-meta-viewer-base:latest $PRIVATE_DOCKER_REPOS/$PRIVATE_DOCKER_ORG/dbsql-meta-viewer-base:latest
#docker push $PRIVATE_DOCKER_REPOS/$PRIVATE_DOCKER_ORG/dbsql-meta-viewer-base:latest

