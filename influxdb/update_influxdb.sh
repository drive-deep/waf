#!/bin/sh

# Variables (using values from docker-compose.yml)
INFLUXDB_URL="http://influxdb:8086"  # Service name 'influxdb' resolves to the container
INFLUXDB_ORG="myorg"                  # Organization name from docker-compose.yml
INFLUXDB_BUCKET="mybucket"            # Bucket name from docker-compose.yml
MEASUREMENT="api_requests"            # Measurement name (you can customize this)
INFLUXDB_TOKEN="1un_2_x6qWT8k-g2vSD3UDzH2qENtlK7kUuydbJwJkiMzq5lI0I7-mFERZ7HKqeDqK1UObpr_mHqZV0Uchw2_Q=="  # Your InfluxDB authentication token

# Get the API endpoint and current timestamp
API_ENDPOINT="$1"
TIMESTAMP=$(date +%s)

# Check if the API endpoint is provided
if [ -z "$API_ENDPOINT" ]; then
  echo "API endpoint is required"
  exit 1
fi

# Send data to InfluxDB v2 using the authentication token, storing the endpoint as a tag
RESPONSE=$(curl --request POST -s "$INFLUXDB_URL/api/v2/write?org=$INFLUXDB_ORG&bucket=$INFLUXDB_BUCKET&precision=s" \
  --header "Authorization: Token $INFLUXDB_TOKEN" \
  --header "Content-Type: text/plain; charset=utf-8" \
  --data-binary "$MEASUREMENT,endpoint=$API_ENDPOINT value=1 $TIMESTAMP")

# Check the response from InfluxDB
if [ -z "$RESPONSE" ]; then
  echo "✅ API request data sent to InfluxDB for endpoint: $API_ENDPOINT"
else
  echo "Error sending data to InfluxDB:"
  echo "$RESPONSE"
  exit 1
fi
