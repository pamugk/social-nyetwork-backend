package api

import(
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
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
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request   body      models.CreateMessageRequest  true  "New message data"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Failure      401  {object}  string
// @Failure      403  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /messages [post]
// @Security TokenAuth
func createMessage(w http.ResponseWriter, r *http.Request) {

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

}

// editMessage godoc
// @Summary      Edit message
// @Description  Update message text
// @Tags         messages
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Message ID"
// @Param        request   body      models.MessageData  true  "Updated password data"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Failure      401  {object}  string
// @Failure      403  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /messages/{id} [put]
// @Security TokenAuth
func editMessage(w http.ResponseWriter, r *http.Request) {

}

// getMessages godoc
// @Summary      Get messages
// @Description  Get list of messages by filters
// @Tags         messages
// @Accept       json
// @Produce      json
// @Param        page   query      int  false  "Page number"
// @Param        pageSize   query  int  false  "Page size"
// @Success      200  {object}  models.GetUsersResponse
// @Failure      400  {object}  string
// @Failure      401  {object}  string
// @Failure      403  {object}  string
// @Failure      500  {object}  string
// @Router       /messages [get]
// @Security TokenAuth
func getMessages(w http.ResponseWriter, r *http.Request) {

}
