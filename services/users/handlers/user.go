package handlers

import (
	"synapsis-backend-challenge/services/users"
	"synapsis-backend-challenge/services/users/dtos"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	uc users.Usecase
}

func NewUserHandler(uc users.Usecase) users.Handler {
	return &UserHandler{
		uc: uc,
	}
}

// Login implements users.Handler.
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var request dtos.LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	if err := request.Validate(); err != nil {
		return err
	}

	if err := h.uc.Login(c.Context()); err != nil {
		return err
	}

	return c.JSON(request)
}

// Register implements users.Handler.
func (h *UserHandler) Register(c *fiber.Ctx) error {
	var request dtos.RegisterRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	if err := request.Validate(); err != nil {
		return err
	}

	if err := h.uc.Register(c.Context()); err != nil {
		return err
	}

	return c.JSON(request)
}
