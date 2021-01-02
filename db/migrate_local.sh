#!/usr/bin/env bash

export $(grep -v '^#' ../../golang-api-secrets/db_local.env | xargs)

docker run \
  --rm \
  --network golang-api_golang-api-net \
  --mount type=bind,src=$(pwd)/migrations,dst=/migrations \
  migrate/migrate -path /migrations \
  -database $DB_CONNECTIONSTRING \
  up
