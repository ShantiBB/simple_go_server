package main

import (
	"SimpleGoProject/handlers"
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthcheck/", app.HealthcheckHandler)
	mux.HandleFunc("GET /v1/books/", handlers.CreateBookHandler)
	mux.HandleFunc("GET /v1/books/{id}/", handlers.ShowBookHandler)

	return mux
}
