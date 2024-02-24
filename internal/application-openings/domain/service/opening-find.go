package service

import (
	"fmt"
	"go-basic/internal/application-openings/infra/schema"
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

func OpeningFindHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		outputErrorFormat(ctx, http.StatusBadRequest, outputErrorMessageFormat("id", "queryParameter").Error())
		return
	}

	opening := schema.Opening{}

	if err := databaseInstance.First(&opening, id).Error; err != nil {
		outputErrorFormat(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	outputSuccessFormat(ctx, "find-opening", opening)
}
