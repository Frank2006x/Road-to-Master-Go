package router

import (
	"github.com/Frank2006x/Fibre/src/controller"
	"github.com/Frank2006x/Fibre/src/middleware"
	"github.com/gofiber/fiber/v3"
)

func TodoRoutes(app *fiber.App) {
	todoGroup := app.Group("/todos")
	todoGroup.Use(middleware.AuthMiddleware)
	todoGroup.Post("/", controller.CreateTodo)
	todoGroup.Get("/", controller.GetTodos)
	todoGroup.Put("/:id", controller.UpdateTodo)
	todoGroup.Delete("/:id", controller.DeleteTodo)

}
