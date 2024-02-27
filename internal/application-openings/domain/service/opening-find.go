package service

import (
	"fmt"
	"go-basic/internal/application-openings/infra/schema"
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Find One Opening
// @Description Update job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id path int	true "Opening id"
// @Success 200 {object} OpeningFindOutput
// @Failure 400 {object} ErrorOutput
// @Failure 500 {object} ErrorOutput
// @Router /opening [get]
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

type OpeningFindOutput struct {
	Message string `json:"message"`
	Data    schema.OpeningResponse
}
