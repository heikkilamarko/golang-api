#!/bin/sh

KONG_URL=http://kong:8001

SERVICE_ID=36bab0c2-d5be-47ce-bacc-83351d75c2e5
ROUTE_ID=f6de4fce-876d-4f49-811b-f770808892ea
CONSUMER_ID=f047eb15-c4fa-4c02-b49c-f8dc4b0f2034
KEY_AUTH_PLUGIN_ID=3e451512-33ec-4cd7-b34a-e1c18446793d
RATE_LIMITING_PLUGIN_ID=e08471c7-a17e-47ba-9d44-a202282a18df

service_data()
{
  cat <<EOF
{
  "id": "$SERVICE_ID",
  "name": "products",
  "protocol": "http",
  "host": "api",
  "port": 8080
}
EOF
}

route_data()
{
  cat <<EOF
{
  "id": "$ROUTE_ID",
  "name": "products",
  "protocols": ["http"],
  "paths": ["/"],
  "service": {
    "id": "$SERVICE_ID"
  }
}
EOF
}

consumer_data()
{
  cat <<EOF
{
  "id": "$CONSUMER_ID",
  "username": "products",
  "custom_id": "products"
}
EOF
}

plugin_key_auth_data()
{
  cat <<EOF
{
  "id": "$KEY_AUTH_PLUGIN_ID",
  "name": "key-auth",
  "service": {
    "id": "$SERVICE_ID"
  },
  "config": {
    "key_in_query": false,
    "key_names": ["x-api-key"],
    "key_in_header": true,
    "run_on_preflight": true,
    "hide_credentials": false,
    "key_in_body": false
  },
  "protocols": ["http"]
}
EOF
}

plugin_rate_limiting_data()
{
  cat <<EOF
{
  "id": "$RATE_LIMITING_PLUGIN_ID",
  "name": "rate-limiting",
  "service": {
    "id": "$SERVICE_ID"
  },
  "consumer": {
    "id": "$CONSUMER_ID"
  },
  "config": {
    "minute": 60,
    "limit_by": "consumer",
    "policy": "local"
  },
  "protocols": ["http"]
}
EOF
}

key_auth_data()
{
  cat <<EOF
{
  "key": "$API_KEY"
}
EOF
}

curl \
  -i \
  -d "$(service_data)" \
  -H "Content-Type: application/json" \
  -X PUT $KONG_URL/services/products

curl \
  -i \
  -d "$(route_data)" \
  -H "Content-Type: application/json" \
  -X PUT $KONG_URL/routes/$ROUTE_ID

curl \
  -i \
  -d "$(consumer_data)" \
  -H "Content-Type: application/json" \
  -X PUT $KONG_URL/consumers/$CONSUMER_ID

curl \
  -i \
  -d "$(plugin_key_auth_data)" \
  -H "Content-Type: application/json" \
  -X PUT $KONG_URL/plugins/$KEY_AUTH_PLUGIN_ID

curl \
  -i \
  -d "$(plugin_rate_limiting_data)" \
  -H "Content-Type: application/json" \
  -X PUT $KONG_URL/plugins/$RATE_LIMITING_PLUGIN_ID

curl \
  -i \
  -d "$(key_auth_data)" \
  -H "Content-Type: application/json" \
  -X POST $KONG_URL/consumers/$CONSUMER_ID/key-auth
