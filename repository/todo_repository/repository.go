package todo_repository

import (
	"todo/entity"
	"todo/pkg/errs"
)

type TodoRepository interface {
	CreateTodo(todoPayload *entity.Todo) (*entity.Todo, errs.MessageErr)
	GetAllTodo() ([]entity.Todo, errs.MessageErr)
}
