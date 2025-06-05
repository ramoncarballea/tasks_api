package dto

type PagedResponse[T any] struct {
	PageNumber uint  `json:"page_number"`
	PageSize   uint  `json:"page_size"`
	TotalItems int64 `json:"total_items"`
	Data       []T   `json:"data"`
}
