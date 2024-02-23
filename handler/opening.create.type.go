package handler

import (
	"application-openings/schema"
	"fmt"
)

type CreateOpeningRequestInput struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

type CreateOpeningRequestOutput struct {
	Message string `json:"message"`
	Data    schema.OpeningResponse
}

func (input *CreateOpeningRequestInput) Validate() error {
	if input.Role == "" && input.Company == "" && input.Location == "" && input.Link == "" && input.Remote == nil && input.Salary <= 0 {
		return fmt.Errorf("Request body is invalid format")
	}
	if input.Role == "" {
		return formatOutputErrorMessage("role", "string")
	}
	if input.Company == "" {
		return formatOutputErrorMessage("company", "string")
	}
	if input.Location == "" {
		return formatOutputErrorMessage("location", "string")
	}
	if input.Link == "" {
		return formatOutputErrorMessage("link", "string")
	}
	if input.Remote == nil {
		return formatOutputErrorMessage("remote", "boolean")
	}
	if input.Salary <= 0 {
		return formatOutputErrorMessage("salary", "int64")
	}

	return nil
}
