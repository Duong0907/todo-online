package handlers

import (
	"log"
	"net/http"
	"todo-online-api/models"
	"todo-online-api/pkgs/auth"

	"github.com/gin-gonic/gin"
)


func GetAllTodo(c *gin.Context) {
	token := auth.GetTokenString(c)
	claims, err := auth.ParseTokenString(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, 
			models.NewErrorResponse("Invalid token"))
		return
	}

	var todo models.Todo
	todo.UserID = claims.UserID
	todos, err := todo.GetAllByUserID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, 
			models.NewErrorResponse("Failed to get all todos"))
		return
	}

	c.JSON(http.StatusOK,
		models.NewSuccessResponse("Succesful to get all todo", gin.H{
			"todos": todos,
		}))
}

func CreateOneTodo(c *gin.Context) {
	token := auth.GetTokenString(c)
	claims, err := auth.ParseTokenString(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, 
			models.NewErrorResponse("Invalid token"))
		return
	}

	var todo models.Todo
	todo.UserID = claims.UserID
	if err != nil {
		c.JSON(http.StatusBadRequest, 
			models.NewErrorResponse("Wrong user ID"))
		return
	}

	if err := c.BindJSON(&todo); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, 
			models.NewErrorResponse("Fail to read todo"))
		return
	}

	if err = todo.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, 
			models.NewErrorResponse("Fail to create todo"))
		return
	}
	
	c.JSON(http.StatusOK, 
		models.NewSuccessResponse("Success to create the todo", gin.H{
			"todo": todo,
		}))
}

func DeleteOneTodo(c *gin.Context) {
	token := auth.GetTokenString(c)
	claims, err := auth.ParseTokenString(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, 
			models.NewErrorResponse("Invalid token"))
		return
	}

	var todo models.Todo
	todo.UserID = claims.UserID

	if err != nil {
		c.JSON(http.StatusBadRequest, 
			models.NewErrorResponse("Wrong user ID"))
		return
	}
		
	id := c.Query("id")
	if err = todo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, 
			models.NewErrorResponse("Failed to delete the todo"))
		return
	}

	c.JSON(http.StatusOK, 
		models.NewSuccessResponse("Success to delete the todo", nil))
}


func ToggleCheckTodo(c *gin.Context) {
	token := auth.GetTokenString(c)
	claims, err := auth.ParseTokenString(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, 
			models.NewErrorResponse("Invalid token"))
		return
	}

	var todo models.Todo
	todo.UserID = claims.UserID

	if err != nil {
		c.JSON(http.StatusBadRequest, 
			models.NewErrorResponse("Wrong user ID"))
			return
		}
		
	id := c.Query("id")
	if err = todo.ToggleCheck(id); err != nil {
		c.JSON(http.StatusInternalServerError, 
			models.NewErrorResponse("Failed to toggle check the todo"))
		return
	}

	c.JSON(http.StatusOK, 
		models.NewSuccessResponse("Success to toggle check the todo", nil))
}
