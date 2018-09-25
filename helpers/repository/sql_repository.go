package repository

import (
	"database/sql"
)

// SQLRepository helpers
type SQLRepository struct{}

// CommitRollback for transaction
func (repo *SQLRepository) CommitRollback(tx *sql.Tx, err error) {
	switch err {
	case nil:
		err = tx.Commit()
	default:
		tx.Rollback()
	}
}
