package handler

import (
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Route post opening is run",
	})
}
