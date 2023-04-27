package handlers

import (
	// "log"
	"log"
	"net/http"
	"todo-online-api/models"
	"todo-online-api/pkgs/auth"

	"github.com/gin-gonic/gin"
)

func CreateOneUser(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	email := c.Query("email")

	// Hash password
	hashedPassword, _ := auth.HashPassword(password)

	var user models.User 
	if user.Existed(username, email) {
		c.JSON(http.StatusBadRequest, 
			models.NewErrorResponse("Username or email existed"))
		return
	}
	
	if err := user.Create(username, email, hashedPassword); err != nil {
		c.JSON(http.StatusInternalServerError, 
			models.NewErrorResponse("Failed to create a new user"))
		return
	}

	c.JSON(http.StatusOK, 
		models.NewSuccessResponse("Successful to create a new user", nil))
}


func GetUser(c *gin.Context) {
	token := auth.GetTokenString(c)
	log.Println(token)
	claims, err := auth.ParseTokenString(token)
	log.Println(claims)
	if err != nil {
		c.JSON(http.StatusUnauthorized, 
			models.NewErrorResponse("Invalid token"))
		return
	}

	var user models.User 
	user.ID = claims.UserID
	if err := user.GetOne(); err != nil {
		c.JSON(http.StatusBadRequest, 
			models.NewErrorResponse("User not found"))
		return
	}

	// Hide password
	user.Password = "********"

	c.JSON(http.StatusOK, 
		models.NewSuccessResponse("Successful to get user", gin.H{
			"user": user,
		}))
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	var user models.User 
	err := user.GetByUsername(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, 
			models.NewErrorResponse("Wrong username"))
		return
	}

	if err := auth.ValidatePassword(user.Password, password); err != nil {
		c.JSON(http.StatusBadRequest, 
			models.NewErrorResponse("Wrong password"))
		return
	}

	claims := auth.NewClaims(user.ID)
	tokenString, err := auth.GenerateTokenStringFromClaims(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, 
			models.NewErrorResponse("Failed to generat token string"))
	}
	c.SetCookie("todo_api_jwt", tokenString, 3600, "/", "", false, false)

	c.JSON(http.StatusOK, 
		models.NewSuccessResponse("Login successfully", gin.H{
			"user": user,
		}))
}
