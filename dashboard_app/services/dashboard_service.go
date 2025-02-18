package services

import (
	"log"

	"github.com/drive-deep/waf/dashboard_app/influxdb"
	"github.com/drive-deep/waf/dashboard_app/models"
)

func GetDashboardData() ([]models.Dashboard, error) {
	endpoints := []string{"/", "/dashboard"} // Add more API endpoints if needed
	var dashboardData []models.Dashboard

	for _, endpoint := range endpoints {
		hitsToday, hitsWeek, err := influxdb.QueryHits(endpoint)
		if err != nil {
			log.Printf("Error fetching data for %s: %v", endpoint, err)
			continue
		}

		dashboardData = append(dashboardData, models.Dashboard{
			APIEndpoint: endpoint,
			HitsToday:   hitsToday,
			HitsWeek:    hitsWeek,
		})
	}

	return dashboardData, nil
}
