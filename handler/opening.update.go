package handler

import (
	"application-openings/schema"
	"fmt"
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, formatOutputErrorMessage("id", "queryParameter").Error())
		return
	}

	body := UpdateOpeningRequestInput{}
	ctx.BindJSON(&body)

	if err := body.Validate(); err != nil {
		logger.ErrorWithValues("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schema.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
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

	if err := db.Save(&opening).Error; err != nil {
		logger.ErrorWithValues("error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}

	sendSuccess(ctx, "update-opening", opening)
}
