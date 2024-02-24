package service

import (
	"fmt"
	"go-basic/internal/application-openings/infra/schema"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OpeningDeleteService(ctx *gin.Context) {
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

	if err := databaseInstance.Delete(&opening).Error; err != nil {
		outputErrorFormat(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s ", id))
		return
	}

	outputSuccessFormat(ctx, "delete-opening", opening)
}
