#!/bin/bash

read -sp "apikey: " API_KEY

echo -e "\n\nRunning tests...\n\n"

npx newman run postman_collection.json \
  -e postman_environment.json \
  --env-var "api_key=$API_KEY" \
  --reporters cli,junit \
  --reporter-junit-export results/junitReport.xml
