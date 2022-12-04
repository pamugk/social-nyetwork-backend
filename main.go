package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"

	_ "github.com/social-nyetwork/backend/docs"
	"github.com/swaggo/http-swagger"

	"github.com/social-nyetwork/backend/internal/models"
	"github.com/social-nyetwork/backend/internal/service"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

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
		r.Route("/messages", func(r chi.Router) {
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator)
		})

		r.Route("/users", func(r chi.Router) {
			r.Post("/", createUser)
			r.Post("/login", login)

			r.Route("/", func(r chi.Router) {
				r.Use(jwtauth.Verifier(tokenAuth))
				r.Use(jwtauth.Authenticator)

				r.Get("/", getUsers)
				r.Post("/logout", logout)
				r.Route("/{id:^-?\\d+}", func(r chi.Router) {
					r.Get("/", getUser)
					r.Put("/", editUser)
					r.Delete("/", removeUser)

					r.Put("/password", changeUserPassword)
				})
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
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /users/{id}/password [put]
// @Security TokenAuth
func changeUserPassword(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err == nil {
		_, claims, _ := jwtauth.FromContext(r.Context())
		if claims["userId"] != id {
			http.Error(w, "Not authorized", http.StatusForbidden)
			return
		}

		data := &models.PasswordData{}
		if err = render.Bind(r, data); err == nil {
			err = service.ChangeUserPassword(id, data)
		}
	}

	if err == service.NotFoundError {
		http.Error(w, "User not found", http.StatusNotFound)
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// createUser godoc
// @Summary      Create user
// @Description  Create new user with set password
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request   body      models.CreateUserRequest  true  "New user data"
// @Success      201  {object}  models.NewEntityResponse
// @Failure      400  {object}  string
// @Failure      500  {object}  string
// @Router       /users [post]
func createUser(w http.ResponseWriter, r *http.Request) {
	data := &models.CreateUserRequest{}
	err := render.Bind(r, data)
	var id int64 = -1
	if err == nil {
		id, err = service.CreateUser(data)
	}

	var renderer render.Renderer = nil
	if err == nil {
		renderer = &models.NewEntityResponse{Id: id}
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	render.Render(w, r, renderer)
}

// editUser godoc
// @Summary      Edit user info
// @Description  Update user information
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Param        request   body      models.EditUserRequest  true  "Updated user data"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /users/{id} [put]
// @Security TokenAuth
func editUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err == nil {
		_, claims, _ := jwtauth.FromContext(r.Context())
		if claims["userId"] != id {
			http.Error(w, "Not authorized", http.StatusForbidden)
			return
		}

		data := &models.EditUserRequest{}
		if err = render.Bind(r, data); err == nil {
			err = service.EditUser(id, data)
		}
	}

	if err == service.NotFoundError {
		http.Error(w, "User not found", http.StatusNotFound)
	} else if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
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
// @Security TokenAuth
func getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	var response* models.GetUserResponse = nil
	if err == nil {
		response, err = service.GetUser(id)
	}

	if err == nil {
		render.Render(w, r, response)
	} else if err == service.NotFoundError {
		http.Error(w, "User not found", http.StatusNotFound)
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
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
// @Security TokenAuth
func getUsers(w http.ResponseWriter, r *http.Request) {
	response, err := service.GetUsers(0, 10)
	if err == nil {
		render.Render(w, r, response)
	} else if err == service.NotFoundError {
		http.Error(w, "User not found", http.StatusNotFound)
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// login godoc
// @Summary      Log in
// @Description  Logs user in
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}	string
// @Failure      400  {object}  string
// @Failure      401  {object}  string
// @Failure      403  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /users/login [post]
func login(w http.ResponseWriter, r *http.Request) {
	if login, password, ok := r.BasicAuth(); ok {
		var tokenString string
		id, err := service.StartSession(login, password)
		if err == nil {
			_, tokenString, err = tokenAuth.Encode(map[string]interface{}{ "userId": id })
		}

		if err == nil {
			json.NewEncoder(w).Encode(tokenString)
		} else if err == service.WrongCredentialsError {
			http.Error(w, err.Error(), http.StatusForbidden)
		} else if err == service.NotFoundError {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	} else {
		w.Header().Set("WWW-Authenticate", "")
		http.Error(w, "Authentication required", http.StatusUnauthorized)
	}
}

// logout godoc
// @Summary      Log out
// @Description  Logs user out
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}	string
// @Failure      400  {object}  string
// @Failure      500  {object}  string
// @Router       /users/logout [post]
// @Security TokenAuth
func logout(w http.ResponseWriter, r *http.Request) {

}

// removeUser godoc
// @Summary      Delete user
// @Description  Delete user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /users/{id} [delete]
// @Security TokenAuth
func removeUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err == nil {
		_, claims, _ := jwtauth.FromContext(r.Context())
		if claims["userId"] != id {
			http.Error(w, "Not authorized", http.StatusForbidden)
			return
		}

		err = service.RemoveUser(id)
	}

	if err == service.NotFoundError {
		http.Error(w, "User not found", http.StatusNotFound)
	}  else if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
