package src

import (
	"log"
	"urlshorter/src/db"
	"urlshorter/src/router"

	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3"
)

func SetupApp() *fiber.App{
	app := fiber.New()
	app.Use(logger.New())
	_,err:=db.ConnectDB();
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	router.ShortenRoute(app)	
	return app
}