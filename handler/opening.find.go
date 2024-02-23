package handler

import (
	"application-openings/schema"
	"fmt"
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

func FindOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, formatOutputErrorMessage("id", "queryParameter").Error())
		return
	}

	opening := schema.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	sendSuccess(ctx, "find-opening", opening)
}
