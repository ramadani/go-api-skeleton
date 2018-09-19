package data

type UserPaginate struct {
	Data    []User `json:"data"`
	Total   uint   `json:"total"`
	PerPage uint   `json:"per_page"`
	Page    uint   `json:"page"`
	Pages   uint   `json:"pages"`
}
