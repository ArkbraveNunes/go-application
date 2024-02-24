package service

import (
	"fmt"
	"go-basic/internal/application-openings/infra/schema"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Update opening
// @Description Update job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body OpeningUpdateInput true "Request body"
// @Success 200 {object} OpeningUpdateOutput
// @Failure 400 {object} ErrorOutput
// @Failure 500 {object} ErrorOutput
// @Router /opening [post]
func OpeningUpdateService(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		outputErrorFormat(ctx, http.StatusBadRequest, outputErrorMessageFormat("id", "queryParameter").Error())
		return
	}

	body := OpeningUpdateInput{}
	ctx.BindJSON(&body)

	if err := body.Validate(); err != nil {
		loggerInstance.ErrorWithValues("validation error: %v", err.Error())
		outputErrorFormat(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schema.Opening{}

	if err := databaseInstance.First(&opening, id).Error; err != nil {
		outputErrorFormat(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	if body.Role != "" {
		opening.Role = body.Role
	}
	if body.Company != "" {
		opening.Company = body.Company
	}
	if body.Location != "" {
		opening.Location = body.Location
	}
	if body.Link != "" {
		opening.Link = body.Link
	}
	if body.Remote != nil {
		opening.Remote = *body.Remote
	}
	if body.Salary > 0 {
		opening.Salary = body.Salary
	}

	if err := databaseInstance.Save(&opening).Error; err != nil {
		loggerInstance.ErrorWithValues("error updating opening: %v", err.Error())
		outputErrorFormat(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}

	outputSuccessFormat(ctx, "update-opening", opening)
}

type OpeningUpdateInput struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

type OpeningUpdateOutput struct {
	Message string `json:"message"`
	Data    schema.OpeningResponse
}

func (input *OpeningUpdateInput) Validate() error {
	if input.Role != "" || input.Company != "" || input.Location != "" || input.Link != "" || input.Remote != nil || input.Salary > 0 {
		return nil
	}

	return fmt.Errorf("At least one valid field must be provided")
}
