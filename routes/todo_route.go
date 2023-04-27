package routes

import (
	"github.com/gin-gonic/gin"
	"todo-online-api/handlers"
)

func RouteTodo(r *gin.RouterGroup) {
	todo_route := r.Group("/todos")
	{
		todo_route.POST("/", handlers.CreateOneTodo)
		todo_route.GET("/", handlers.GetAllTodo)
		todo_route.DELETE("/", handlers.DeleteOneTodo)
		todo_route.PUT("/", handlers.ToggleCheckTodo)
	}
}