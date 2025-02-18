package influxdb

import (
	"context"
	"fmt"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var (
	client influxdb2.Client
	org    string
	bucket string
)

func InitInfluxDB(url, orgName, bucketName string) {
	// Retry parameters
	maxRetries := 10
	retryInterval := 5 * time.Second

	// Connect to InfluxDB without authentication
	client = influxdb2.NewClient(url, "mytoken") 

	// Set global org and bucket
	org = orgName
	bucket = bucketName

	// Retry logic
	for i := 0; i < maxRetries; i++ {
		// Check if InfluxDB is ready
		ready, err := client.Ready(context.Background())
		if err != nil {
			log.Printf("Attempt %d: InfluxDB connection failed: %v", i+1, err)
		} else if ready != nil && ready.Status != nil && *ready.Status == "ready" {
			log.Println("✅ InfluxDB is connected and ready")
			return
		} else {
			log.Printf("Attempt %d: InfluxDB not ready, status: %v", i+1, ready)
		}

		// If not ready, wait and retry
		log.Printf("Retrying in %v...", retryInterval)
		time.Sleep(retryInterval)
	}

	// If InfluxDB is not ready after maxRetries
	log.Fatalf("InfluxDB is not ready after %d attempts", maxRetries)
}

func GetInfluxDBClient() influxdb2.Client {
	return client
}

type APIHits struct {
	APIEndpoint string `json:"api_endpoint"`
	Hits        int    `json:"hits"`
}

func QueryAllEndpointHitsWeek() ([]APIHits, error) {
	return queryAllEndpointHits("-7d")
}

func QueryAllEndpointHitsDay() ([]APIHits, error) {
	return queryAllEndpointHits("-1d")
}

func queryAllEndpointHits(duration string) ([]APIHits, error) {
	log.Printf("client : %v", client)
	log.Printf("org : %v", org)
	log.Printf("bucket : %v", bucket)
	queryAPI := client.QueryAPI(org)

	// Flux query to get hits for all endpoints for the specified duration
	query := fmt.Sprintf(`
	from(bucket: "%s")
		|> range(start: %s)
		|> filter(fn: (r) => r["_measurement"] == "api_requests")
		|> group(columns: ["endpoint"])
		|> count()`,
		bucket, duration)

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Printf("InfluxDB query error: %v", err)
		return nil, err
	}
	defer result.Close()

	var apiHits []APIHits
	for result.Next() {
		record := result.Record()
		endpoint, ok := record.ValueByKey("endpoint").(string)
		if !ok {
			log.Printf("Invalid endpoint data type")
			continue
		}

		hits, ok := record.Value().(int64)
		if !ok {
			log.Printf("Invalid hits data type")
			continue
		}

		apiHits = append(apiHits, APIHits{
			APIEndpoint: endpoint,
			Hits:        int(hits),
		})
	}

	if result.Err() != nil {
		log.Printf("Query iteration error: %v", result.Err())
		return nil, result.Err()
	}

	return apiHits, nil
}
