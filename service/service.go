package service

import (
	"todo/dto"
	"todo/pkg/errs"
)

type TodoService interface {
	CreateTodo(payload *dto.TodoRequest) (*dto.NewTodoResponse, errs.MessageErr)
	GetAllTodo() (*dto.GetAllTodoResponse, errs.MessageErr)
}
