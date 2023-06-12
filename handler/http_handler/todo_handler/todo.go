package todo_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	}

	todo, err := t.todoService.CreateTodo(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(todo.StatusCode, todo)
}

func (t *todoHandler) GetAllTodo(ctx *gin.Context) {
	listTodo, err := t.todoService.GetAllTodo()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(listTodo.StatusCode, listTodo)
}

func (t *todoHandler) GetTodoById(ctx *gin.Context) {
	id := ctx.Param("id")
	todoId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	todo, err := t.todoService.GetTodoById(uint(todoId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
	}

	ctx.JSON(todo.StatusCode, todo)
}

func (t *todoHandler) UpdateTodoById(ctx *gin.Context) {
	var request dto.TodoRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	id := ctx.Param("id")
	todoId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
	}

	todo, err := t.todoService.UpdateTodoById(uint(todoId), &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(todo.StatusCode, todo)
}

func (t *todoHandler) DeleteTodoById(ctx *gin.Context) {
	id := ctx.Param("id")
	todoId, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
	}

	todo, err := t.todoService.DeleteTodoById(uint(todoId))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(todo.StatusCode, todo)
}
