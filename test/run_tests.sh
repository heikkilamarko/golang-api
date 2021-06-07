#!/bin/bash

if [[ -z "$1" ]]; then
  echo "error: You must pass an api key as the only argument."
  exit 0
fi

npx newman run postman_collection.json \
  -e postman_environment.json \
  --env-var "api_key=$1" \
  --reporters cli,junit \
  --reporter-junit-export results/junitReport.xml
