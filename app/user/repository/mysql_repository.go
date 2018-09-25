package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ramadani/go-api-skeleton/app/user/data"
	"github.com/ramadani/go-api-skeleton/helpers/format"
)

const (
	// PaginateQuery get user paginate
	PaginateQuery = `SELECT id, name, email FROM users WHERE deleted_at IS NULL OFFSET ? LIMIT ?`
	// CountQuery get total of user
	CountQuery = `SELECT COUNT(id) AS total FROM users WHERE deleted_at IS NULL`
	// CreateQuery to create a new user query
	CreateQuery = `INSERT INTO users (name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
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
	userRows, err := repo.db.Query(PaginateQuery, offset, limit)
	if err != nil {
		return users, total, err
	}

	defer userRows.Close()
	for userRows.Next() {
		user := data.User{}
		if err = userRows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return users, total, err
		}
		users = append(users, user)
	}

	// get total of users by given query
	totalRows, err := repo.db.Query(CountQuery)
	if err != nil {
		return users, total, err
	}

	defer totalRows.Close()
	for totalRows.Next() {
		totalRows.Scan(&total)
	}

	return users, total, nil
}

// Create a new user
func (repo *MySQLRepository) Create(name, email, password string) (data.User, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return data.User{}, err
	}

	now := time.Now().Format(format.DateTimeToString)
	fmt.Println(now)

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	_, err = tx.Exec(CreateQuery, name, email, password, now, now)

	if err != nil {
		return data.User{}, err
	}

	return data.User{ID: 1, Name: name, Email: email}, nil
}

// FindByID to find user by id
func (repo *MySQLRepository) FindByID(id uint) (data.User, error) {
	return data.User{}, nil
}

// Update an existing user
func (repo *MySQLRepository) Update(name string, id uint) (data.User, error) {
	return data.User{}, nil
}

// Delete an existing user
func (repo *MySQLRepository) Delete(id uint) error {
	return nil
}

// NewMySQLRepository new mysql user repo
func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db}
}
