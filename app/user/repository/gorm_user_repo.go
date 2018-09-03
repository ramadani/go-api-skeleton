package repository

import (
	"github.com/ramadani/go-api-skeleton/app/user/model"
	"github.com/ramadani/go-api-skeleton/db"
)

type GormUserRepo struct {
	db *db.Database
}

func (rp *GormUserRepo) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := rp.db.DB.Where(&model.User{Email: email}).First(&user).Error

	return user, err
}
