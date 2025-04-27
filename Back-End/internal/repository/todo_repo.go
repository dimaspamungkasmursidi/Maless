package repository

import (
	"github.com/Bobby-P-dev/todo-listgo.git/internal/helpers"
	models "github.com/Bobby-P-dev/todo-listgo.git/internal/models/domains"
)

type TodoRepository interface {
	GetAllTodos() ([]*models.Todos, *helpers.CustomError)
	GetTodoByID(id int) (*models.Todos, *helpers.CustomError)
	CreateTodo(todo *models.Todos) *helpers.CustomError
	UpdateTodoByID(todo *models.Todos) *helpers.CustomError
	DeleteTodoByID(id int) *helpers.CustomError
	GetByTagTodos(tag string) ([]*models.Todos, *helpers.CustomError)
}
