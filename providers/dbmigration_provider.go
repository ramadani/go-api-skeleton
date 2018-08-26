package providers

import (
	"github.com/jinzhu/gorm"
	"github.com/ramadani/go-api-skeleton/app/models"
)

// DbMigration contains db instance.
type DbMigration struct {
	db *gorm.DB
}

// Boot the db migration.
func (dbm DbMigration) Boot() {
	dbm.db.AutoMigrate(&models.User{})
}

// NewDbMigration returns db migration.
func NewDbMigration(db *gorm.DB) *DbMigration {
	return &DbMigration{db}
}
