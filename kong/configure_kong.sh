#!/bin/sh

curl \
  -i \
  -d "$(envsubst < resources/service.json)" \
  -H "Content-Type: application/json" \
  -X PUT $KONG_URL/services/products

curl \
  -i \
  -d "$(envsubst < resources/route.json)" \
  -H "Content-Type: application/json" \
  -X PUT $KONG_URL/routes/$ROUTE_ID

curl \
  -i \
  -d "$(envsubst < resources/consumer.json)" \
  -H "Content-Type: application/json" \
  -X PUT $KONG_URL/consumers/$CONSUMER_ID

curl \
  -i \
  -d "$(envsubst < resources/plugin_key_auth.json)" \
  -H "Content-Type: application/json" \
  -X PUT $KONG_URL/plugins/$KEY_AUTH_PLUGIN_ID

curl \
  -i \
  -d "$(envsubst < resources/plugin_rate_limiting.json)" \
  -H "Content-Type: application/json" \
  -X PUT $KONG_URL/plugins/$RATE_LIMITING_PLUGIN_ID

curl \
  -i \
  -d "$(envsubst < resources/key_auth.json)" \
  -H "Content-Type: application/json" \
  -X POST $KONG_URL/consumers/$CONSUMER_ID/key-auth
