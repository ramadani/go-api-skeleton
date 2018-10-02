package sqlutils

import "database/sql"

// DoTx for transaction
func DoTx(tx *sql.Tx, err error) {
	switch err {
	case nil:
		err = tx.Commit()
	default:
		tx.Rollback()
	}
}
