package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/drive-deep/waf/dashboard_app/services"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dashboardData, err := services.GetDashboardData()
	if err != nil {
		http.Error(w, `{"error": "Failed to fetch data from InfluxDB"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dashboardData)
}
