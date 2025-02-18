package models

type Dashboard struct {
	APIEndpoint string `json:"api_endpoint"`
	HitsToday   int    `json:"hits_today"`
	HitsWeek    int    `json:"hits_week"`
}
