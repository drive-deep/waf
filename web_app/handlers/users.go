package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/drive-deep/waf/web_app/models"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	users := []models.User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	json.NewEncoder(w).Encode(users)
}
