package repository

import (
	"github.com/ramadani/go-api-skeleton/app/todo/model"
	"github.com/ramadani/go-api-skeleton/db"
)

type GormTodoRepo struct {
	db *db.Database
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

func (td *GormTodoRepo) Find(id uint) model.Todo {
	var todo model.Todo
	td.db.DB.First(&todo, id)

	return todo
}

func (td *GormTodoRepo) Update(title, body string, id uint) model.Todo {
	todo := td.Find(id)
	todo.Title = title
	todo.Body = body

	td.db.DB.Save(&todo)

	return todo
}

func (td *GormTodoRepo) Delete(id uint) bool {
	todo := td.Find(id)

	td.db.DB.Delete(&todo)

	return true
}

func NewGormRepo(db *db.Database) *GormTodoRepo {
	return &GormTodoRepo{db}
}
