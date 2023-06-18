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

// CreateTodo godoc
// @Summary Create Todo
// @Description Create Todo List
// @Tags		todos
// @Accept		json
// @Produce		json
// @Param		todo	body		dto.TodoRequest	true	"Create todo request body"
// @Success		201		{object}	dto.NewTodoResponse
// @Failure		422		{object}	errs.MessageErrData
// @Failure		400		{object}	errs.MessageErrData
// @Router		/todos  [post]
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

// GetAllTodo godoc
// @Summary 	List Todo
// @Description Get All Todo List
// @Tags		todos
// @Accept		json
// @Produce		json
// @Success		200		{object}	dto.GetAllTodoResponse
// @Failure		422		{object}	errs.MessageErrData
// @Failure		400		{object}	errs.MessageErrData
// @Router		/todos  [get]
func (t *todoHandler) GetAllTodo(ctx *gin.Context) {
	listTodo, err := t.todoService.GetAllTodo()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(listTodo.StatusCode, listTodo)
}

// GetOrderByID godoc
//
//	@Summary		Find todos
//	@Description	Get an todo list by id
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint	true	"Todo ID"
//	@Success		200	{object}	dto.GetTodoByIdResponse
//	@Failure		400	{object}	errs.MessageErrData
//	@Failure		404	{object}	errs.MessageErrData
//	@Router			/todos/{id} [get]
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

// UpdateOrderByID godoc
//
//	@Summary		Update Todo
//	@Description	Update todo by id todo
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			id		path		uint				true	"Todo ID"
//	@Param			todo	body		dto.TodoRequest		true	"Update order request body"
//	@Success		200		{object}	dto.UpdateTodoByIdResponse
//	@Failure		400		{object}	errs.MessageErrData
//	@Failure		404		{object}	errs.MessageErrData
//	@Failure		422		{object}	errs.MessageErrData
//	@Router			/todos/{id} [put]
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

// DeleteOrderByID godoc
//
//	@Summary		Delete todo
//	@Description	Delete an todo by id
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint	true	"Todo ID"
//	@Success		200	{object}	dto.DeleteTodoByIdResponse
//	@Failure		400	{object}	errs.MessageErrData
//	@Failure		404	{object}	errs.MessageErrData
//	@Failure		500	{object}	errs.MessageErrData
//	@Router			/todos/{id} [delete]
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
