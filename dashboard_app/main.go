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
	INFLUXDB_ORG := "myorg"
	INFLUXDB_BUCKET := "mybucket"
	INFLUXDB_TOKEN := "1un_2_x6qWT8k-g2vSD3UDzH2qENtlK7kUuydbJwJkiMzq5lI0I7-mFERZ7HKqeDqK1UObpr_mHqZV0Uchw2_Q=="

	// Initialize InfluxDB connection without authentication
	influxdb.InitInfluxDB(INFLUXDB_URL, INFLUXDB_ORG, INFLUXDB_BUCKET, INFLUXDB_TOKEN)

	mux := http.NewServeMux()
	dashboard_route.RegisterRoutes(mux)

	port := 8090
	fmt.Printf("🚀 Server running on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), mux))
}
