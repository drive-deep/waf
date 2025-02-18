package routes

import (
	"net/http"

	"github.com/drive-deep/waf/dashboard_app/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.DashboardHandler)
}

