package repository

import (
	"github.com/ramadani/go-api-skeleton/app/todo/model"
	"github.com/ramadani/go-api-skeleton/db"
)

type GormTodoRepo struct {
	db *db.Gorm
}

func (td *GormTodoRepo) All() []model.Todo {
	var todos []model.Todo
	td.db.DB.Find(&todos)

	return todos
}

func (td *GormTodoRepo) Create(title, body string) model.Todo {
	todo := model.Todo{Title: title, Body: body}
	td.db.DB.Create(&todo)

	return todo
}

func (td *GormTodoRepo) Find(id uint) (model.Todo, error) {
	var todo model.Todo
	err := td.db.DB.First(&todo, id).Error

	return todo, err
}

func (td *GormTodoRepo) Update(title, body string, id uint) (model.Todo, error) {
	todo, err := td.Find(id)

	if err == nil {
		todo.Title = title
		todo.Body = body
		td.db.DB.Save(&todo)
	}

	return todo, err
}

func (td *GormTodoRepo) Delete(id uint) error {
	todo, err := td.Find(id)

	if err != nil {
		return err
	}

	return td.db.DB.Delete(&todo).Error
}

func NewGormRepo(db *db.Gorm) *GormTodoRepo {
	return &GormTodoRepo{db}
}
