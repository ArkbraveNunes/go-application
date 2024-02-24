package service

import (
	"fmt"
	"go-basic/internal/application-openings/infra/schema"
	"net/http"

	gin "github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body OpeningCreateInput true "Request body"
// @Success 200 {object} OpeningCreateOutput
// @Failure 400 {object} ErrorOutput
// @Failure 500 {object} ErrorOutput
// @Router /opening [post]
func OpeningCreateService(ctx *gin.Context) {
	body := OpeningCreateInput{}
	ctx.BindJSON(&body)

	if err := body.Validate(); err != nil {
		loggerInstance.ErrorWithValues("validation error: %v", err.Error())
		outputErrorFormat(ctx, http.StatusBadRequest, err.Error())
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

	if err := databaseInstance.Create(&opening).Error; err != nil {
		loggerInstance.ErrorWithValues("Error creating new opening in the database: %v", err.Error())
		outputErrorFormat(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	outputSuccessFormat(ctx, "create-opening", opening)
}

type OpeningCreateInput struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

type OpeningCreateOutput struct {
	Message string `json:"message"`
	Data    schema.OpeningResponse
}

func (input *OpeningCreateInput) Validate() error {
	if input.Role == "" && input.Company == "" && input.Location == "" && input.Link == "" && input.Remote == nil && input.Salary <= 0 {
		return fmt.Errorf("Request body is invalid format")
	}
	if input.Role == "" {
		return outputErrorMessageFormat("role", "string")
	}
	if input.Company == "" {
		return outputErrorMessageFormat("company", "string")
	}
	if input.Location == "" {
		return outputErrorMessageFormat("location", "string")
	}
	if input.Link == "" {
		return outputErrorMessageFormat("link", "string")
	}
	if input.Remote == nil {
		return outputErrorMessageFormat("remote", "boolean")
	}
	if input.Salary <= 0 {
		return outputErrorMessageFormat("salary", "int64")
	}

	return nil
}
