package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"  // mysql dialect
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite dialect
)

// Gorm contains db library.
type Gorm struct {
	DB *gorm.DB
}

// Close the connection
func (db *Gorm) Close() error {
	return db.DB.Close()
}

// New the database connection and returns db.
func New(driver, connection string) *Gorm {
	db, err := gorm.Open(driver, connection)

	if err != nil {
		panic("failed to connect database")
	}

	return &Gorm{db}
}
