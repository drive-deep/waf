package handlers

import (
	"net/http"
)


func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	
	w.Write([]byte("Hello, welcome to the API!"))
}