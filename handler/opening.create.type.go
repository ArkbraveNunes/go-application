package handler

import (
	"fmt"
)

func validateError(name, dataType string) error {
	return fmt.Errorf("param : %s (type: %s) is required", name, dataType)
}

type CreateOpeningRequestInput struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (input *CreateOpeningRequestInput) Validate() error {
	if input.Role == "" && input.Company == "" && input.Location == "" && input.Link == "" && input.Remote == nil && input.Salary <= 0 {
		return fmt.Errorf("Request body is invalid format")
	}
	if input.Role == "" {
		return validateError("role", "string")
	}
	if input.Company == "" {
		return validateError("company", "string")
	}
	if input.Location == "" {
		return validateError("location", "string")
	}
	if input.Link == "" {
		return validateError("link", "string")
	}
	if input.Remote == nil {
		return validateError("remote", "boolean")
	}
	if input.Salary <= 0 {
		return validateError("salary", "int64")
	}

	return nil
}
