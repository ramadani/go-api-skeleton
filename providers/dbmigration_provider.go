package providers

import (
	"github.com/ramadani/go-api-skeleton/app/models"
	"github.com/ramadani/go-api-skeleton/db"
)

// DbMigration contains db instance.
type DbMigration struct {
	db *db.Database
}

// Boot the db migration.
func (dbm DbMigration) Boot() {
	dbm.db.DB.AutoMigrate(
		&models.User{},
		&models.Todo{},
	)
}

// NewDbMigration returns db migration.
func NewDbMigration(db *db.Database) *DbMigration {
	return &DbMigration{db}
}
