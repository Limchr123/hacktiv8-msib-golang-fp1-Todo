package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"todo/database"
	_ "todo/docs"
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
	r.GET("/todos", todoHandler.GetAllTodo)
	r.GET("/todos/:id", todoHandler.GetTodoById)
	r.PUT("/todos/:id", todoHandler.UpdateTodoById)
	r.DELETE("/todos/:id", todoHandler.DeleteTodoById)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
