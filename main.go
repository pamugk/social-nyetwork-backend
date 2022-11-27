package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	_ "github.com/social-nyetwork/backend/docs"
	"github.com/swaggo/http-swagger"

	"github.com/social-nyetwork/backend/internal/models"
)

// @title Social Nyetwork Swagger
// @version 1.0
// @description This is a Social Nyetwork server.

// @tag.name users
// @tag.description User service

// @license.name WTFPL 2.0
// @license.url http://www.wtfpl.net/txt/copying/

// @BasePath /api/v1
func main() {
	r := chi.NewRouter()

	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/messages", func(r chi.Router) {

		})

		r.Route("/users", func(r chi.Router) {
			r.Get("/", getUsers)
			r.Post("/", createUser)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", getUser)
				r.Put("/", editUser)
				r.Delete("/", removeUser)

				r.Put("/password", changeUserPassword)
			})
		})
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	http.ListenAndServe(":8080", r)
}

// changeUserPassword godoc
// @Summary      Change user password
// @Description  Change user password
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Param        request   body      models.PasswordData  true  "New password data"
// @Success      200  {object}  models.UserData
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /users/{id}/password [put]
func changeUserPassword(w http.ResponseWriter, r *http.Request) {
	data := &models.PasswordData{}
	if err := render.Bind(r, data); err != nil {
		render.Status(r, http.StatusBadRequest)
		return
	}
}

// createUser godoc
// @Summary      Create user
// @Description  Create new user with set password
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request   body      models.CreateUserRequest  true  "New user data"
// @Success      201  {object}  models.UserData
// @Failure      400  {object}  string
// @Failure      500  {object}  string
// @Router       /users [post]
func createUser(w http.ResponseWriter, r *http.Request) {
	data := &models.CreateUserRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Status(r, http.StatusBadRequest)
		return
	}
}

// editUser godoc
// @Summary      Edit user info
// @Description  Update user information
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Param        request   body      models.EditUserRequest  true  "Updated user data"
// @Success      200  {object}  models.UserData
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /users/{id} [put]
func editUser(w http.ResponseWriter, r *http.Request) {
	data := &models.EditUserRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Status(r, http.StatusBadRequest)
		return
	}
}

// getUser godoc
// @Summary      Get full user info
// @Description  Get user info by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  models.UserData
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /users/{id} [get]
func getUser(w http.ResponseWriter, r *http.Request) {

}

// getUsers godoc
// @Summary      Search users
// @Description  Search users by some filters
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        page   query      int  false  "Page number"
// @Param        pageSize   query  int  false  "Page size"
// @Success      200  {object}  models.GetUsersResponse
// @Failure      400  {object}  string
// @Failure      500  {object}  string
// @Router       /users [get]
func getUsers(w http.ResponseWriter, r *http.Request) {

}

// removeUser godoc
// @Summary      Delete user
// @Description  Delete user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      204  {object}  models.UserData
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /users/{id} [delete]
func removeUser(w http.ResponseWriter, r *http.Request) {

}
