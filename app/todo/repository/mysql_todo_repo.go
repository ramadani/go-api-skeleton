package repository

import (
	"github.com/ramadani/go-api-skeleton/app/todo/model"
	"github.com/ramadani/go-api-skeleton/db"
)

type MySQLTodoRepo struct {
	db *db.Database
}

func (td *MySQLTodoRepo) All() []model.Todo {
	var todos []model.Todo
	td.db.DB.Find(&todos)

	return todos
}

func (td *MySQLTodoRepo) Create(title, body string) model.Todo {
	todo := model.Todo{Title: title, Body: body}
	td.db.DB.Create(&todo)

	return todo
}

func (td *MySQLTodoRepo) Find(id uint) model.Todo {
	var todo model.Todo
	td.db.DB.First(&todo, id)

	return todo
}

func (td *MySQLTodoRepo) Update(title, body string, id uint) model.Todo {
	todo := td.Find(id)
	todo.Title = title
	todo.Body = body

	td.db.DB.Save(&todo)

	return todo
}

func (td *MySQLTodoRepo) Delete(id uint) bool {
	todo := td.Find(id)

	td.db.DB.Delete(&todo)

	return true
}

func NewMySQLRepo(db *db.Database) *MySQLTodoRepo {
	return &MySQLTodoRepo{db}
}
