package repository

import (
	"github.com/ramadani/go-api-skeleton/app/user/model"
	"github.com/ramadani/go-api-skeleton/db"
)

type GormUserRepo struct {
	db *db.Gorm
}

func (rp *GormUserRepo) Create(name, email, password string) (model.User, error) {
	user := model.User{Name: name, Email: email, Password: password}
	err := rp.db.DB.Create(&user).Error

	return user, err
}

func (rp *GormUserRepo) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := rp.db.DB.Where(&model.User{Email: email}).First(&user).Error

	return user, err
}

func NewGormRepo(db *db.Gorm) *GormUserRepo {
	return &GormUserRepo{db}
}
