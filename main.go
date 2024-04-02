package main

import (
	"fmt"
	"todo-online-api/pkgs/auth"
	"todo-online-api/pkgs/middleware"
	"todo-online-api/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)


func main() {
	auth.GenerateJWTKey()
	router := gin.Default()

	router.Use(middleware.CorsMiddleware)

	api := router.Group("/api")
	{
		api.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Ping successful",
			})
		})
	}

	website := router.Group("/website")
	{
		website.Static("/", "./website/")
	}


	routes.RouteTodo(api)
	routes.RouteUser(api)

	router.Run(":8080")
}