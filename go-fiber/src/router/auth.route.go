package router

import (
	"github.com/Frank2006x/Fibre/src/controller"
	"github.com/Frank2006x/Fibre/src/middleware"
	"github.com/gofiber/fiber/v3"
)

func AuthRoutes(app *fiber.App) {
	authGroup := app.Group("/auth")
	authGroup.Post("/register", controller.RegisterUser)
	authGroup.Post("/login", controller.LoginUser)
	authGroup.Post("/logout", middleware.AuthMiddleware, controller.LogoutUser)
}