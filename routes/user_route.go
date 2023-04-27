package routes

import (
	"github.com/gin-gonic/gin"
	"todo-online-api/handlers"
)

func RouteUser(r *gin.RouterGroup) {
	user_route := r.Group("/users")
	{
		user_route.POST("/", handlers.CreateOneUser)
		user_route.POST("/login", handlers.Login)
		user_route.GET("/", handlers.GetUser)
	}
}