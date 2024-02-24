package service

import (
	"go-basic/internal/application-openings/infra/schema"
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

func OpeningListHandler(ctx *gin.Context) {
	openings := []schema.Opening{}

	if err := databaseInstance.Find(&openings).Error; err != nil {
		outputErrorFormat(ctx, http.StatusInternalServerError, "error listening openings")
		return
	}

	outputSuccessFormat(ctx, "list-openings", openings)
}
