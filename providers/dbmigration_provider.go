package providers

import (
	todo "github.com/ramadani/go-api-skeleton/app/todo/model"
	user "github.com/ramadani/go-api-skeleton/app/user/model"
	"github.com/ramadani/go-api-skeleton/db"
)

// DbMigrationProvider contains db instance.
type DbMigrationProvider struct {
	db *db.Database
}

// Boot the db migration.
func (p *DbMigrationProvider) Boot() {
	p.db.DB.AutoMigrate(
		&user.User{},
		&todo.Todo{},
	)
}

// NewDbMigration returns db migration.
func NewDbMigration(db *db.Database) *DbMigrationProvider {
	return &DbMigrationProvider{db}
}
