package service

import (
	"go-basic/internal/application-openings/infra/schema"
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary List Openings
// @Description Find List Openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} OpeningListOutput
// @Failure 400 {object} ErrorOutput
// @Failure 500 {object} ErrorOutput
// @Router /opening [get]
func OpeningListHandler(ctx *gin.Context) {
	openings := []schema.Opening{}

	if err := databaseInstance.Find(&openings).Error; err != nil {
		outputErrorFormat(ctx, http.StatusInternalServerError, "error listening openings")
		return
	}

	outputSuccessFormat(ctx, "list-openings", openings)
}

type OpeningListOutput struct {
	Message string `json:"message"`
	Data    []schema.OpeningResponse
}
