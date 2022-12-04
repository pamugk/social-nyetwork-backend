package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/social-nyetwork/backend/docs"
	"github.com/swaggo/http-swagger"

	"github.com/social-nyetwork/backend/internal/api"
	"github.com/social-nyetwork/backend/internal/service"
)

// @title Social Nyetwork Swagger
// @version 1.0
// @description This is a Social Nyetwork server.

// @tag.name users
// @tag.description User service

// @license.name WTFPL 2.0
// @license.url http://www.wtfpl.net/txt/copying/

// @securitydefinitions.oauth2.password	 TokenAuth
// tokenUrl /api/v1/users/login

// @BasePath /api/v1
func main() {
	defer service.CloseDbPool()

	r := chi.NewRouter()

	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		api.InitMessageRoutes(r)
		api.InitUserRoutes(r)
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	http.ListenAndServe(":8080", r)
}
