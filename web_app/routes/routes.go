package routes

import (
	"net/http"
	"github.com/drive-deep/waf/web_app/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/users", handlers.UsersHandler)
	mux.HandleFunc("/orders", handlers.OrdersHandler)
}
