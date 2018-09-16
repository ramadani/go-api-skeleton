package db

import (
	"github.com/jinzhu/gorm"
)

// Orm interface for orm concrete
type Orm interface {
	Create(value interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	Delete(value interface{}, where ...interface{}) *gorm.DB
}
