package handler

import "fmt"

type CreateOpeningRequest struct {
	Role        string `json:"role"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Remote      *bool  `json:"remote"`
	Link        string `json:"link"`
	Salary      int64  `json:"salary"`
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Description == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Description == "" {
		return errParamIsRequired("description", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if r.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}
	return nil
}

type UpdateOpeningRequest struct {
	Role        string `json:"role"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Remote      *bool  `json:"remote"`
	Link        string `json:"link"`
	Salary      int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role != "" || r.Description != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Salary > 0 {
		return nil
	}

	return fmt.Errorf("at least one valid field must be provided")
}
