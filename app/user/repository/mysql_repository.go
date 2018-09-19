package repository

import (
	"database/sql"

	"github.com/ramadani/go-api-skeleton/app/user/data"
)

const (
	PaginateQuery = "SELECT id, name, email FROM users LIMIT ? OFFSET ?"
	CountQuery    = "SELECT COUNT(id) AS total FROM users"
)

// MySQLRepository of user repo
type MySQLRepository struct {
	db *sql.DB
}

// Paginate user from mysql
func (repo *MySQLRepository) Paginate(offset, limit uint) ([]data.User, uint, error) {
	var users []data.User
	var total uint

	// get users
	userRows, pErr := repo.db.Query(PaginateQuery, limit, offset)
	if pErr != nil {
		return users, total, pErr
	}

	defer userRows.Close()
	for userRows.Next() {
		user := data.User{}
		if rErr := userRows.Scan(&user.ID, &user.Name, &user.Email); rErr != nil {
			return users, total, rErr
		}
		users = append(users, user)
	}

	totalRows, tErr := repo.db.Query(CountQuery)
	if tErr != nil {
		panic(tErr)
	}

	defer totalRows.Close()
	for totalRows.Next() {
		totalRows.Scan(&total)
	}

	return users, total, nil
}

// NewMySQLRepository new mysql user repo
func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db}
}
