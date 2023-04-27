package models

import (
	// "go.mongodb.org/mongo-driver/bson"
	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type Response struct {
	Message string `json:"message"`
	Error   bool   `json:"error"`
	Data    gin.H  `json:"data,omitempty"`
}

func NewSuccessResponse(message string, data gin.H) Response {
	return Response {
		message,
		false,
		data,
	}
}

func NewErrorResponse(message string) Response {
	return Response {
		message,
		true,
		nil,
	}
}
