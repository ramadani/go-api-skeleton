package providers

import (
	"github.com/ramadani/go-api-skeleton/db"
	"github.com/ramadani/go-api-skeleton/models"
)

// DbMigration contains db instance.
type DbMigration struct {
	db *db.Database
}

// Boot the db migration.
func (p *DbMigration) Boot() {
	p.db.DB.AutoMigrate(
		&models.User{},
		&models.Todo{},
	)
}

// NewDbMigration returns db migration.
func NewDbMigration(db *db.Database) *DbMigration {
	return &DbMigration{db}
}
