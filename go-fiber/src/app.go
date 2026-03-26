package src

import (
	"log"

	"github.com/Frank2006x/Fibre/src/db"
	"github.com/Frank2006x/Fibre/src/router"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
)

func SetupApp() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	err:=godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env file")
	}
	db.ConnectDB()
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	router.AuthRoutes(app)
	router.TodoRoutes(app)

	return app
}