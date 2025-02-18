package influxdb

import (
	"context"
	"fmt"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var client influxdb2.Client
var org, bucket string

func InitInfluxDB(url, token, orgName, bucketName string) {
	client = influxdb2.NewClient(url, token)
	org = orgName
	bucket = bucketName
}

func GetInfluxDBClient() influxdb2.Client {
	return client
}

func QueryHits(endpoint string) (int, int, error) {
	queryAPI := client.QueryAPI(org)

	// Flux query to filter by endpoint and measurement
	query := fmt.Sprintf(`
		from(bucket: "%s")
		|> range(start: -7d)
		|> filter(fn: (r) => r["_measurement"] == "api_requests" and r["endpoint"] == "%s")
		|> group(columns: ["_time"])
		|> keep(columns: ["_value", "_time"])
	`, bucket, endpoint)

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Printf("InfluxDB query error: %v", err)
		return 0, 0, err
	}

	// Time-based filtering
	hitsToday := 0
	hitsWeek := 0
	now := time.Now()
	startOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	for result.Next() {
		record := result.Record()
		timestamp := record.Time()

		hitsWeek++ // Count all hits in the last 7 days
		if timestamp.After(startOfToday) {
			hitsToday++ // Count hits from today only
		}
	}

	if result.Err() != nil {
		log.Printf("Query iteration error: %v", result.Err())
		return 0, 0, result.Err()
	}

	return hitsToday, hitsWeek, nil
}
