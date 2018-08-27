package models

import (
	"time"
)

// Todo represent todo's structure table
type Todo struct {
	ID        uint   `gorm:"primary_key"`
	Title     string `gorm:"not null"`
	Body      string `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
