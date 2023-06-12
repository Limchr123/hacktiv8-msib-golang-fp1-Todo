package service

import (
	"todo/dto"
	"todo/pkg/errs"
)

type TodoService interface {
	CreateTodo(payload *dto.TodoRequest) (*dto.NewTodoResponse, errs.MessageErr)
	GetAllTodo() (*dto.GetAllTodoResponse, errs.MessageErr)
	GetTodoById(id uint) (*dto.GetTodoByIdResponse, errs.MessageErr)
	UpdateTodoById(id uint, payload *dto.TodoRequest) (*dto.UpdateTodoByIdResponse, errs.MessageErr)
	DeleteTodoById(id uint) (*dto.DeleteTodoByIdResponse, errs.MessageErr)
}
