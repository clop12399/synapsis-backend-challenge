package main

import (
	"synapsis-backend-challenge/services/users/handlers"
	"synapsis-backend-challenge/services/users/repositories"
	"synapsis-backend-challenge/services/users/usecases"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	userRepo := repositories.NewuserRepository()

	userUC := usecases.NewUserUsecase(userRepo)

	userHandler := handlers.NewUserHandler(userUC)

	handlers.MapUser(app, userHandler)

	app.Listen(":3000")
}
