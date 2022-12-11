package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"

	"github.com/social-nyetwork/backend/internal/models"
	"github.com/social-nyetwork/backend/internal/service"
)

func InitMessageRoutes(r chi.Router) {
	r.Route("/messages", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/", getMessages)
		r.Post("/", createMessage)
		r.Route("/{id}", func(r chi.Router) {
			r.Put("/", editMessage)
			r.Delete("/", deleteMessage)
		})
	})
}

// createMessage godoc
// @Summary      Create message
// @Description  Create new message, sent from one user to another
// @Tags         messages
// @Accept       json
// @Produce      json
// @Param        request   body      models.CreateMessageRequest  true  "New message data"
// @Success      201  {object}  models.NewEntityResponse[string]
// @Failure      400  {object}  string
// @Failure      401  {object}  string
// @Failure      403  {object}  string
// @Failure      500  {object}  string
// @Router       /messages [post]
// @Security TokenAuth
func createMessage(w http.ResponseWriter, r *http.Request) {
	var id string
	var err error
	data := &models.CreateMessageRequest{}
	if err = render.Bind(r, data); err == nil {
		_, claims, _ := jwtauth.FromContext(r.Context())
		if claims["userId"] != data.Sender {
			http.Error(w, "Not authorized", http.StatusForbidden)
			return
		}

		_, err = service.CreateMessage(data)
	}

	if err == nil {
		render.Render(w, r, &models.NewEntityResponse[string]{Id: id})
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// deleteMessage godoc
// @Summary      Delete message
// @Description  Delete message
// @Tags         messages
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Message ID"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Failure      401  {object}  string
// @Failure      403  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /messages/{id} [delete]
// @Security TokenAuth
func deleteMessage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := service.DeleteMessage(id)

	if err == service.NotFoundError {
		http.Error(w, "Message not found", http.StatusNotFound)
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// editMessage godoc
// @Summary      Edit message
// @Description  Update message text
// @Tags         messages
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Message ID"
// @Param        request   body      models.MessageData  true  "Updated message data"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Failure      401  {object}  string
// @Failure      403  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /messages/{id} [put]
// @Security TokenAuth
func editMessage(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	userId := claims["userId"].(int64)

	var err error
	id := chi.URLParam(r, "id")
	data := &models.MessageData{}
	if err = render.Bind(r, data); err == nil {
		err = service.EditMessage(id, userId, data)
	}

	if err == service.NotFoundError {
		http.Error(w, "Message not found", http.StatusNotFound)
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// getMessages godoc
// @Summary      Get messages
// @Description  Get list of messages by filters
// @Tags         messages
// @Accept       json
// @Produce      json
// @Param        from   query      int  true  "Sending user id"
// @Param        to   query      int  true  "Receiving user id"
// @Param        page   query      int  false  "Page number"
// @Param        pageSize   query  int  false  "Page size"
// @Success      200  {object}  models.GetMessagesResponse
// @Failure      400  {object}  string
// @Failure      401  {object}  string
// @Failure      403  {object}  string
// @Failure      500  {object}  string
// @Router       /messages [get]
// @Security TokenAuth
func getMessages(w http.ResponseWriter, r *http.Request) {
	var from, to int64 = -1, -1
	 if r.URL.Query().Has("from") && r.URL.Query().Has("to") {
		 val, err := strconv.ParseInt(r.URL.Query().Get("from"), 10, 64)
		 if err == nil {
			 from = val
		}

		 val, err = strconv.ParseInt(r.URL.Query().Get("to"), 10, 64)
		 if err == nil {
			 to = val
		}
	}
	if from == -1 || to == -1 {
		http.Error(w, "Correspondents not defined", http.StatusBadRequest)
	}

	page := 0
	pageSize := 10

	 if r.URL.Query().Has("page") {
		 val, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 32)
		 if err == nil {
			 page = int(val)
		}
	}
	 if r.URL.Query().Has("pageSize") {
		 val, err := strconv.ParseInt(r.URL.Query().Get("pageSize"), 10, 32)
		 if err == nil {
			 page = int(val)
		}
	}

	response, err := service.GetMessages(from, to, page, pageSize)
	if err == nil {
		render.Render(w, r, response)
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
