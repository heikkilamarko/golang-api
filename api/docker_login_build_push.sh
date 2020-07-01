#!/usr/bin/env bash

export $(grep -v '^#' ../../golang-api-secrets/acr.env | xargs)

IMG=$ACR_SERVER/goapi:$1

docker login -u $ACR_USER -p $ACR_PASSWORD $ACR_SERVER
docker build -t $IMG .
docker push $IMG
