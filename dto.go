package mockapi

type PaginatedBooks struct {
	Data       []Book `json:"data"`
	Page       int    `json:"page"`
	PageSize   int    `json:"page_size"`
	TotalItems int64  `json:"total_items"`
	TotalPages int    `json:"total_pages"`
}

type APIError struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}