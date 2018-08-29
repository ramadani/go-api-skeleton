package providers

import (
	"github.com/ramadani/go-api-skeleton/db"
	"github.com/ramadani/go-api-skeleton/models"
)

// DbMigrationProvider contains db instance.
type DbMigrationProvider struct {
	db *db.Database
}

// Boot the db migration.
func (p *DbMigrationProvider) Boot() {
	p.db.DB.AutoMigrate(
		&models.User{},
		&models.Todo{},
	)
}

// NewDbMigration returns db migration.
func NewDbMigration(db *db.Database) *DbMigrationProvider {
	return &DbMigrationProvider{db}
}
