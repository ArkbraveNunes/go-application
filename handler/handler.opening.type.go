package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func formatOutputErrorMessage(name, dataType string) error {
	return fmt.Errorf("param : %s (type: %s) is required", name, dataType)
}

func sendError(ctx *gin.Context, statusCode int, message string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"message":    message,
	})
}

func sendSuccess(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successful", operation),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message    string `json:message`
	StatusCode string `json:statusCode`
}
