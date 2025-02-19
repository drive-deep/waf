package services

import (
	"encoding/json"
	"log"

	"github.com/drive-deep/waf/dashboard_app/influxdb"
	"github.com/drive-deep/waf/dashboard_app/models"
)

// GetDashboardData fetches dashboard data for either "day" or "week"
func GetDashboardData(duration string) ([]models.Dashboard, error) {
	var apiHits []influxdb.APIHits
	var err error

	apiHits, err = influxdb.QueryAllEndpointHitsDuration(duration)
	if err != nil {
		log.Printf("Error fetching data from InfluxDB: %v", err)
		return nil, err
	}

	// Convert APIHits slice to JSON
	jsonData, err := json.Marshal(apiHits)
	if err != nil {
		log.Printf("Error marshaling JSON data: %v", err)
		return nil, err
	}

	var dashboardData []models.Dashboard
	err = json.Unmarshal(jsonData, &dashboardData)
	if err != nil {
		log.Printf("Error unmarshaling JSON data: %v", err)
		return nil, err
	}

	return dashboardData, nil
}
