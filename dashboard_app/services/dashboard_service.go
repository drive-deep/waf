package services

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/drive-deep/waf/dashboard_app/influxdb"
	"github.com/drive-deep/waf/dashboard_app/models"
)

// GetDashboardData fetches dashboard data for either "day" or "week"
func GetDashboardData(duration string) ([]models.Dashboard, error) {
	var apiHits []influxdb.APIHits
	var err error

	// Select the appropriate function based on the duration
	switch duration {
	case "day":
		apiHits, err = influxdb.QueryAllEndpointHitsDay()
	case "week":
		apiHits, err = influxdb.QueryAllEndpointHitsWeek()
	default:
		return nil, errors.New("invalid duration: must be 'day' or 'week'")
	}

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
