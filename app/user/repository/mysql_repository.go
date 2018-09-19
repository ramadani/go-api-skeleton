package repository

import (
	"database/sql"

	"github.com/ramadani/go-api-skeleton/app/user/data"
)

const (
	// PaginateQuery get user paginate
	PaginateQuery = `SELECT id, name, email FROM users WHERE deleted_at IS NULL OFFSET ? LIMIT ?`
	// CountQuery get total of user
	CountQuery = `SELECT COUNT(id) AS total FROM users WHERE deleted_at IS NULL`
)

// MySQLRepository of user repo
type MySQLRepository struct {
	db *sql.DB
}

// Paginate user from mysql
func (repo *MySQLRepository) Paginate(offset, limit uint) ([]data.User, uint, error) {
	var users []data.User
	var total uint

	// get users by given query
	userRows, pErr := repo.db.Query(PaginateQuery, offset, limit)
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

	// get total of users by given query
	totalRows, tErr := repo.db.Query(CountQuery)
	if tErr != nil {
		return users, total, tErr
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
