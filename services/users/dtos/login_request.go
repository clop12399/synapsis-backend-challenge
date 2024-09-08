package dtos

import "github.com/go-playground/validator/v10"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (request LoginRequest) Validate() error {
	return validator.New().Struct(request)
}
