#!/usr/bin/env bash

export $(grep -v '^#' ../../golang-api-secrets/db.env | xargs)

docker run \
  --rm \
  --mount type=bind,src=$(pwd)/migrations,dst=/migrations \
  --mount type=bind,src=$(pwd)/../../golang-api-secrets/db_certs,dst=/certs \
  migrate/migrate -path /migrations \
  -database $DB_CONNECTIONSTRING \
  up
