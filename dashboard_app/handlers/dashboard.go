package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/drive-deep/waf/dashboard_app/services"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	timeframe := r.URL.Query().Get("interval")

	if timeframe == "" {
		// Fetch both day (-1d) and week (-7d) data
		dayDashboardData, err := services.GetDashboardData("-1d")
		if err != nil {
			http.Error(w, `{"error": "Failed to fetch daily data from InfluxDB"}`, http.StatusInternalServerError)
			return
		}

		weekDashboardData, err := services.GetDashboardData("-7d")
		if err != nil {
			http.Error(w, `{"error": "Failed to fetch weekly data from InfluxDB"}`, http.StatusInternalServerError)
			return
		}

		// Return both datasets in a JSON response
		response := map[string]interface{}{
			"day":  dayDashboardData,
			"week": weekDashboardData,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		// Check if timeframe is a valid integer
		if _, err := strconv.Atoi(timeframe); err != nil {
			http.Error(w, `{"error": "Invalid interval value. Must be an integer."}`, http.StatusBadRequest)
			return
		}

		// Fetch data for the requested timeframe
		dashboardData, err := services.GetDashboardData(fmt.Sprintf("-%ss", timeframe))
		if err != nil {
			http.Error(w, `{"error": "Failed to fetch data from InfluxDB"}`, http.StatusInternalServerError)
			return
		}

		// Return single dataset
		response := map[string]interface{}{
			"timeframe": timeframe,
			"data":      dashboardData,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
