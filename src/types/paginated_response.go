package types

type PaginatedResponse[T any] struct {
	Items []T `json:"items"`
	HasNextPage bool `json:"has_next_page"`
	CurrentPage int `json:"current_page"`
	IsFirstPage bool `json:"is_first_page"`
}