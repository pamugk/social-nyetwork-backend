package models

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	ErrorText string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrRender(err error, status int) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: status,

		ErrorText: err.Error(),
	}
}

type NewEntityResponse struct {
	Id int64 `json:"id"`
}

func (e *NewEntityResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusCreated)
	return nil
}

type SuccessResponse struct{}

func (e *SuccessResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// @Description Pagination response info
type Page struct {
	PageNumber int `json:"page_number"` // Number of returned page
	PageSize   int `json:"page_size"` // Max page size
	TotalItems int `json:"total_items"` // Total count of found items
}
