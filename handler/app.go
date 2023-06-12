package handler

import (
	"github.com/gin-gonic/gin"
	"todo/database"
	"todo/handler/http_handler/todo_handler"
	"todo/repository/todo_repository/todo_pg"
	"todo/service/todo_service"
)

func StartApp() {
	db := database.GetDataBaseInstance()

	todoRepo := todo_pg.NewTodoPG(db)
	todoService := todo_service.NewTodoService(todoRepo)
	todoHandler := todo_handler.NewTodoHandler(todoService)

	r := gin.Default()

	r.POST("/todos", todoHandler.CreateTodo)

	r.Run(":8080")
}
