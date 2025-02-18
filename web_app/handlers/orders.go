package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/drive-deep/waf/web_app/models"
)

func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	orders := []models.Order{
		{ID: 101, Item: "Laptop", Amount: 1200},
		{ID: 102, Item: "Phone", Amount: 800},
	}

	json.NewEncoder(w).Encode(orders)
}
