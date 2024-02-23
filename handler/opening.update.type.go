package handler

import "fmt"

type UpdateOpeningRequestInput struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (input *UpdateOpeningRequestInput) Validate() error {
	if input.Role != "" || input.Company != "" || input.Location != "" || input.Link != "" || input.Remote != nil || input.Salary > 0 {
		return nil
	}

	return fmt.Errorf("At least one valid field must be provided")
}
