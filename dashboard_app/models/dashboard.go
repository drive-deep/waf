package models

type Dashboard struct {
	APIEndpoint string `json:"api_endpoint"`
	Hits   int    `json:"hits"`
}
