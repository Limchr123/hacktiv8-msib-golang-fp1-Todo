package dto

import (
	"time"
	"todo/entity"
)

type TodoRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Todo struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
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

type GetAllTodoResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       []Todo `json:"data"`
}
