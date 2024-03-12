package model

type Pagination struct {
	Page         int `form:"page"`
	Limit        int `form:"limit"`
	Offset       int
	TotalElement int `json:"total_element"`
	CurrentPage  int `json:"current_page"`
	TotalPage    int `json:"total_page"`
}

func CreatePagination(param Pagination) *Pagination {
	pagination := param

	if param.Limit == 0 {
		pagination.Limit = 5
	}

	if param.Page == 0 {
		pagination.Page = 1
	}

	pagination.Offset = (param.Page - 1) * pagination.Limit

	return &pagination
}
