package providers

import (
	"github.com/jinzhu/gorm"
	"github.com/ramadani/go-api-skeleton/app/models"
)

type DbMigration struct {
	db *gorm.DB
}

func (dbm DbMigration) Boot() {
	dbm.db.AutoMigrate(&models.User{})
}

func NewDbMigration(db *gorm.DB) *DbMigration {
	return &DbMigration{db}
}
