package data

type UserPaginate struct {
	Data    []User
	Total   int64
	PerPage int
	Page    int
	Pages   int
}
