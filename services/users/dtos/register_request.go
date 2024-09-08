package dtos

import "github.com/go-playground/validator/v10"

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (request RegisterRequest) Validate() error {
	return validator.New().Struct(request)
}
