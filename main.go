package main

import (
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"

    "github.com/swaggo/http-swagger"
    _ "github.com/social-nyetwork/backend/docs"
)

// @title Social Nyetwork Swagger
// @version 1.0
// @description This is a Social Nyetwork server.

// @license.name WTFPL 2.0
// @license.url http://www.wtfpl.net/txt/copying/

// @BasePath /api/v1
func main() {
    r := chi.NewRouter()

    r.Use(middleware.Heartbeat("/ping"))
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    r.Route("/messages", func(r chi.Router) {

    }

    r.Route("/users", func(r chi.Router) {

    }

    r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
	)

    http.ListenAndServe(":8080", r)
}
