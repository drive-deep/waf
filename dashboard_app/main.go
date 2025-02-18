package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/drive-deep/waf/dashboard_app/influxdb"
	dashboard_route "github.com/drive-deep/waf/dashboard_app/routes"
)

func main() {
	// Initialize InfluxDB connection
	INFLUXDB_URL := "http://influxdb:8086"
	INFLUXDB_ORG := "my-org"
	INFLUXDB_BUCKET := "api_metrics"
	INFLUXDB_TOKEN := "my-influxdb-token-12345abcdef67890"

	// Initialize InfluxDB connection without authentication
	influxdb.InitInfluxDB(INFLUXDB_URL, INFLUXDB_TOKEN, INFLUXDB_ORG, INFLUXDB_BUCKET)
	defer influxdb.GetInfluxDBClient().Close()

	mux := http.NewServeMux()
	dashboard_route.RegisterRoutes(mux)

	port := 8090
	fmt.Printf("🚀 Server running on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
