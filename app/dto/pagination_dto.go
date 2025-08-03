package dto

type PaginationResponse struct {
	Items      any   `json:"items"`
	TotalItems int64 `json:"total_items"`
	Page       int   `json:"page"`
	Size       int   `json:"size"`
	TotalPages int   `json:"total_pages"`
}
