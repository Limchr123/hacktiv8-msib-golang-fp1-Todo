package todo_repository

import (
	"todo/entity"
	"todo/pkg/errs"
)

type TodoRepository interface {
	CreateTodo(todoPayload *entity.Todo) (*entity.Todo, errs.MessageErr)
	GetAllTodo() ([]entity.Todo, errs.MessageErr)
	GetTodoById(id uint) (*entity.Todo, errs.MessageErr)
	UpdateTodoById(id uint, todoPayload *entity.Todo) (*entity.Todo, errs.MessageErr)
	DeleteTodoById(id uint) errs.MessageErr
}
