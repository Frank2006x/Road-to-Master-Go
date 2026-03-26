package router

import (
	"urlshorter/src/controller"

	"github.com/gofiber/fiber/v3"
)

func ShortenRoute(app *fiber.App) {
	shortenGrp:= app.Group("/shorten")
	shortenGrp.Post("/", controller.ShortenURL)
	shortenGrp.Get("/", controller.GetURL)
	app.Get("/:id", controller.RedirectURL)
}