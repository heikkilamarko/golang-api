#!/usr/bin/env bash

npx newman run kong.postman_collection.json \
  -e kong.postman_environment.json \
  --reporters cli,junit \
  --reporter-junit-export results/junitReport.xml
