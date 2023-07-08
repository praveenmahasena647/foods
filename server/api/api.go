package api

import (
	"net/http"

	"github.com/praveenmahasena647/foodApi/internels/handlers"
)

func RunServer() error {
	http.HandleFunc("/", handlers.ServeData)
	return http.ListenAndServe(":8080", nil)
}
