#!/usr/bin/env bash

npx newman run postman_collection.json \
  -e ../env/postman_environment.json \
  --reporters cli,junit \
  --reporter-junit-export results/junitReport.xml
