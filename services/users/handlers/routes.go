package handlers

import (
	"synapsis-backend-challenge/services/users"

	"github.com/gofiber/fiber/v2"
)

func MapUser(routes fiber.Router, h users.Handler) {
	routes.Post("/login", h.Login)
	routes.Post("/register", h.Register)
}
