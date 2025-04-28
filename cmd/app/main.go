package main

import (
	"github.com/gin-gonic/gin"
	"go-todo-app/internal/handler"
	"go-todo-app/internal/infra"
	"go-todo-app/internal/service"
)

func main() {
	db := infra.ConnectDB()
	infra.MigrateDB(db)

	todoService := service.NewTodoService(db)
	todoHandler := handler.NewTodoHandler(todoService)

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/todos", todoHandler.CreateTodo)
		v1.GET("/todos", todoHandler.GetTodos)
		v1.GET("/todos/:id", todoHandler.GetTodo)
		v1.PUT("/todos/:id", todoHandler.UpdateTodo)
		v1.DELETE("/todos/:id", todoHandler.DeleteTodo)
	}

	router.Run(":8080")
}
