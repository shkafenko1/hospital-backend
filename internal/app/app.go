package app

import (
	"hospital-backend/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func App() {
	logger.InitLogger()

	app := fiber.New()
	api := app.Group("/api")
}
