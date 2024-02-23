package handler

import (
	"application-openings/schema"
	"net/http"

	gin "github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequestInput true "Request body"
// @Success 200 {object} CreateOpeningRequestOutput
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	body := CreateOpeningRequestInput{}
	ctx.BindJSON(&body)

	if err := body.Validate(); err != nil {
		logger.ErrorWithValues("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schema.Opening{
		Role:     body.Role,
		Company:  body.Company,
		Location: body.Location,
		Remote:   *body.Remote,
		Link:     body.Link,
		Salary:   body.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.ErrorWithValues("Error creating new opening in the database: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
