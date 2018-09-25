package repository

import (
	"database/sql"
)

// SQLRepository helpers
type SQLRepository struct{}

// DoTx for transaction
func (repo *SQLRepository) DoTx(tx *sql.Tx, err error) {
	switch err {
	case nil:
		err = tx.Commit()
	default:
		tx.Rollback()
	}
}
