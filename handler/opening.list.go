package handler

import (
	"application-openings/schema"
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schema.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listening openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
