package main

import (
	"fiber/internal/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	router.SetupTasksRouter(app)

	app.Listen(":8080")
}