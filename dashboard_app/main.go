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
	influxdbURL := "http://influxdb:8086"
	influxdbBucket := "api_metrics"

	// Initialize InfluxDB connection without authentication
	influxdb.InitInfluxDB(influxdbURL, "", "", influxdbBucket)
	defer influxdb.GetInfluxDBClient().Close()

	mux := http.NewServeMux()
	dashboard_route.RegisterRoutes(mux)

	port := 8090
	fmt.Printf("🚀 Server running on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
