package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/drive-deep/waf/web_app/routes"
)

func main() {
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	port := 8080
	fmt.Printf("🚀 Server running on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), mux))
}
