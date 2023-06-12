package dto

import (
	"todo/entity"
)

type TodoRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (t *TodoRequest) TodoRequestToEntity() *entity.Todo {
	return &entity.Todo{
		Title:     t.Title,
		Completed: t.Completed,
	}
}

type TodoResponse struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type NewTodoResponse struct {
	StatusCode int          `json:"statusCode"`
	Message    string       `json:"message"`
	Data       TodoResponse `json:"data"`
}
