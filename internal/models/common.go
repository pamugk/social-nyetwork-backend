package models

// @Description Pagination response info
type Page struct {
	PageNumber int // Number of returned page
	PageSize   int // Max page size
	TotalItems int // Total count of found items
}
