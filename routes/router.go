package routes

import (
	"todolist/pkg/todo"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	app := gin.Default()

	todoRepo := todo.NewTodoRepository()
	todoUsecase := todo.NewTodoUsecase(todoRepo)
	todoHandler := todo.NewTodoHandler(todoUsecase)

	v1 := app.Group("/v1")
	{
		v1.GET("/get", todoHandler.GetTodo)
		v1.GET("/search", todoHandler.SearchTodo)
		v1.POST("/create", todoHandler.CreateTodo)
		v1.PATCH("/update", todoHandler.UpdateTodo)
	}

	return app
}