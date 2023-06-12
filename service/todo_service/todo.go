package todo_service

import (
	"net/http"
	"todo/dto"
	"todo/pkg/errs"
	"todo/repository/todo_repository"
	"todo/service"
)

type todoService struct {
	todoRepo todo_repository.TodoRepository
}

func NewTodoService(todoRepo todo_repository.TodoRepository) service.TodoService {
	return &todoService{
		todoRepo: todoRepo,
	}
}

func (t *todoService) CreateTodo(payload *dto.TodoRequest) (*dto.NewTodoResponse, errs.MessageErr) {
	todoPayload := payload.TodoRequestToEntity()

	todoCreate, err := t.todoRepo.CreateTodo(todoPayload)
	if err != nil {
		return nil, errs.NewBadRequest("Error occurred while trying to create request")
	}

	response := &dto.NewTodoResponse{
		StatusCode: http.StatusCreated,
		Message:    "Success",
		Data: dto.TodoResponse{
			Title:     todoCreate.Title,
			Completed: todoCreate.Completed,
		},
	}

	return response, nil
}

func (t *todoService) GetAllTodo() (*dto.GetAllTodoResponse, errs.MessageErr) {
	listTodo, err := t.todoRepo.GetAllTodo()
	if err != nil {
		return nil, errs.NewNotFound("Error occurred while trying to get data")
	}
	list := []dto.Todo{}

	for _, todo := range listTodo {
		list = append(list, dto.Todo{
			ID:        todo.ID,
			Title:     todo.Title,
			Completed: todo.Completed,
			CreatedAt: todo.CreatedAt,
			UpdatedAt: todo.UpdatedAt,
		})
	}

	response := &dto.GetAllTodoResponse{
		StatusCode: http.StatusOK,
		Message:    "Succes",
		Data:       list,
	}

	return response, nil
}

func (t *todoService) GetTodoById(id uint) (*dto.GetTodoByIdResponse, errs.MessageErr) {
	todo, err := t.todoRepo.GetTodoById(id)
	if err != nil {
		return nil, errs.NewBadRequest("Error occurred while trying to get item")
	}

	response := &dto.GetTodoByIdResponse{
		StatusCode: http.StatusOK,
		Message:    "Succes",
		Data: dto.Todo{
			ID:        todo.ID,
			Title:     todo.Title,
			Completed: todo.Completed,
			CreatedAt: todo.CreatedAt,
			UpdatedAt: todo.UpdatedAt,
		},
	}

	return response, nil
}

func (t *todoService) UpdateTodoById(id uint, payload *dto.TodoRequest) (*dto.UpdateTodoByIdResponse, errs.MessageErr) {
	todoPayload := payload.TodoRequestToEntity()
	todo, err := t.todoRepo.UpdateTodoById(id, todoPayload)
	if err != nil {
		return nil, errs.NewNotFound("Error occurred while trying to find data")
	}
	response := &dto.UpdateTodoByIdResponse{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data: dto.Todo{
			ID:        todo.ID,
			Title:     todo.Title,
			Completed: todo.Completed,
			CreatedAt: todo.CreatedAt,
			UpdatedAt: todo.UpdatedAt,
		},
	}

	return response, nil
}

func (t *todoService) DeleteTodoById(id uint) (*dto.DeleteTodoByIdResponse, errs.MessageErr) {
	if err := t.todoRepo.DeleteTodoById(id); err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to delete")
	}

	response := &dto.DeleteTodoByIdResponse{
		StatusCode: http.StatusOK,
		Message:    "Success",
	}

	return response, nil
}
