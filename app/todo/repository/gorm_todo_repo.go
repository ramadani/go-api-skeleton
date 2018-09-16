package repository

import (
	"github.com/ramadani/go-api-skeleton/app/todo/model"
	"github.com/ramadani/go-api-skeleton/db"
)

type GormTodoRepo struct {
	db db.Orm
}

func (td *GormTodoRepo) All() []model.Todo {
	var todos []model.Todo
	td.db.Find(&todos)

	return todos
}

func (td *GormTodoRepo) Create(title, body string) model.Todo {
	todo := model.Todo{Title: title, Body: body}
	td.db.Create(&todo)

	return todo
}

func (td *GormTodoRepo) Find(id uint) (model.Todo, error) {
	var todo model.Todo
	err := td.db.First(&todo, id).Error

	return todo, err
}

func (td *GormTodoRepo) Update(title, body string, id uint) (model.Todo, error) {
	todo, err := td.Find(id)

	if err == nil {
		todo.Title = title
		todo.Body = body
		td.db.Save(&todo)
	}

	return todo, err
}

func (td *GormTodoRepo) Delete(id uint) error {
	todo, err := td.Find(id)

	if err != nil {
		return err
	}

	return td.db.Delete(&todo).Error
}

func NewGormRepo(db db.Orm) *GormTodoRepo {
	return &GormTodoRepo{db}
}
