package models

import (
	"net/http"

	"github.com/go-chi/render"
)

type NewEntityResponse struct {
	Id int64 `json:"id"`
}

func (e *NewEntityResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusCreated)
	return nil
}

// @Description Pagination response info
type Page struct {
	PageNumber int `json:"page_number"` // Number of returned page
	PageSize   int `json:"page_size"` // Max page size
	TotalItems int `json:"total_items"` // Total count of found items
}
