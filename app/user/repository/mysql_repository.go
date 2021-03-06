package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/ramadani/go-api-skeleton/app/user/data"
	"github.com/ramadani/go-api-skeleton/commons/format"
	"github.com/ramadani/go-api-skeleton/commons/sqlutils"
)

const (
	// PaginateQuery get user paginate
	PaginateQuery = `SELECT id, name, email FROM users WHERE deleted_at IS NULL LIMIT ? OFFSET ?`
	// CountQuery get total of user
	CountQuery = `SELECT COUNT(id) AS total FROM users WHERE deleted_at IS NULL`
	// CreateQuery to create a new user query
	CreateQuery = `INSERT INTO users (name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	// FindByIDQuery to get an existing user
	FindByIDQuery = `SELECT id, name, email FROM users WHERE id = ? AND deleted_at IS NULL LIMIT 1`
	// UpdateQuery to update an existing user
	UpdateQuery = `UPDATE users SET name = ?, updated_at = ? WHERE id = ? AND deleted_at IS NULL`
	// SoftDeleteQuery to delete an existing user
	SoftDeleteQuery = `UPDATE users SET deleted_at = ? WHERE id = ? AND deleted_at IS NULL`
)

// MySQLRepository of user repo
type MySQLRepository struct {
	db *sql.DB
}

// Paginate user from mysql
func (repo *MySQLRepository) Paginate(limit, offset uint) ([]data.User, uint, error) {
	var users []data.User
	var total uint

	// get users by given query
	userRows, err := repo.db.Query(PaginateQuery, limit, offset)
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
	err = repo.db.QueryRow(CountQuery).Scan(&total)
	if err != nil {
		return users, total, err
	}

	return users, total, nil
}

// Create a new user
func (repo *MySQLRepository) Create(name, email, password string) (uint, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return 0, err
	}

	now := time.Now().UTC().UTC().Format(format.DateTimeToString)
	res, err := tx.Exec(CreateQuery, name, email, password, now, now)

	defer sqlutils.DoTx(tx, err)
	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()

	return uint(id), nil
}

// FindByID to find user by id
func (repo *MySQLRepository) FindByID(id uint) (data.User, error) {
	var user data.User

	err := repo.db.QueryRow(FindByIDQuery, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return data.User{}, err
	}

	return user, nil
}

// Update an existing user
func (repo *MySQLRepository) Update(name string, id uint) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	now := time.Now().UTC().UTC().Format(format.DateTimeToString)
	_, err = tx.Exec(UpdateQuery, name, now, id)
	defer sqlutils.DoTx(tx, err)

	return err
}

// Delete an existing user
func (repo *MySQLRepository) Delete(id uint) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	now := time.Now().UTC().Format(format.DateTimeToString)
	res, err := tx.Exec(SoftDeleteQuery, now, id)
	defer sqlutils.DoTx(tx, err)

	if res != nil {
		if affected, _ := res.RowsAffected(); affected == 0 {
			return errors.New("no found")
		}
	}

	return err
}

// NewMySQLRepository new mysql user repo
func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db: db}
}
