#!/bin/bash

# Variables (using values from docker-compose.yml)
INFLUXDB_URL="http://influxdb:8086"  # Service name 'influxdb' resolves to the container
INFLUXDB_ORG="myorg"                  # Organization name from docker-compose.yml
INFLUXDB_BUCKET="mybucket"                # Bucket name from docker-compose.yml
MEASUREMENT="api_requests"           # Measurement name (you can customize this)

# Get the API endpoint and current timestamp
API_ENDPOINT="$1"
TIMESTAMP=$(date +%s)

# Check if the API endpoint is provided
if [ -z "$API_ENDPOINT" ]; then
  echo "API endpoint is required"
  exit 1
fi

# Sanitize the API endpoint for InfluxDB tags
SANITIZED_ENDPOINT=$(echo "$API_ENDPOINT" | sed 's/[^a-zA-Z0-9_]/_/g')

# Send data to InfluxDB v2 (without authentication)
curl --request POST "$INFLUXDB_URL/api/v2/write?org=$INFLUXDB_ORG&bucket=$INFLUXDB_BUCKET&precision=s" \
  --header "Content-Type: text/plain; charset=utf-8" \
  --data-binary "$MEASUREMENT,endpoint=$SANITIZED_ENDPOINT value=1 $TIMESTAMP"

# Check the response from InfluxDB
RESPONSE=$(curl --request POST -s "$INFLUXDB_URL/api/v2/write?org=$INFLUXDB_ORG&bucket=$INFLUXDB_BUCKET&precision=s" \
  --header "Content-Type: text/plain; charset=utf-8" \
  --data-binary "$MEASUREMENT,endpoint=$SANITIZED_ENDPOINT value=1 $TIMESTAMP")

if [[ "$RESPONSE" == "" ]]; then
  echo "✅ API request data sent to InfluxDB for endpoint: $API_ENDPOINT"
else
  echo "❌ Error sending data to InfluxDB:"
  echo "$RESPONSE"
  exit 1
fi