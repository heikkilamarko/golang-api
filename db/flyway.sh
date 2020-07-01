#!/usr/bin/env bash

docker run \
  --rm \
  --mount type=bind,src=$(pwd)/migrations,dst=/flyway/sql \
  --env-file ../../golang-api-secrets/flyway.env \
  flyway/flyway $1
