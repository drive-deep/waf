#!/bin/bash

# Variables
INFLUXDB_URL="http://influxdb:8086"  # URL of your InfluxDB container
INFLUXDB_ORG="my-org"                # Organization name
INFLUXDB_BUCKET="api_metrics"        # Bucket in InfluxDB
MEASUREMENT="api_requests"           # Measurement name

# Get the API endpoint and current timestamp
API_ENDPOINT=$1
TIMESTAMP=$(date +%s)

# Check if the API endpoint is provided
if [ -z "$API_ENDPOINT" ]; then
  echo "API endpoint is required"
  exit 1
fi

# Send data to InfluxDB v2 (uses org and bucket)
curl --request POST "$INFLUXDB_URL/api/v2/write?org=$INFLUXDB_ORG&bucket=$INFLUXDB_BUCKET&precision=s" \
  --header "Content-Type: text/plain; charset=utf-8" \
  --data-binary "$MEASUREMENT,endpoint=$API_ENDPOINT value=1 $TIMESTAMP"

echo "✅ API request data sent to InfluxDB for endpoint: $API_ENDPOINT"
