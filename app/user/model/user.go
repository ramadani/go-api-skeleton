package model

import "time"

// User is the model for user's table.
type User struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"type:varchar(100);unique_index;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
