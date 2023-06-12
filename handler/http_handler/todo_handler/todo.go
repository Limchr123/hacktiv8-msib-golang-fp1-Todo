package todo_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo/dto"
	"todo/service"
)

type todoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(todoService service.TodoService) *todoHandler {
	return &todoHandler{todoService: todoService}
}

func (t *todoHandler) CreateTodo(ctx *gin.Context) {
	var request dto.TodoRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	todo, err := t.todoService.CreateTodo(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(todo.StatusCode, todo)
}
