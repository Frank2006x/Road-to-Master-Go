package router

import (
	"fiber/internal/handler"

	"github.com/gofiber/fiber/v3"
)

func SetupTasksRouter(app *fiber.App) {
	handler := handler.NewTaskHandler()

	taskGroup := app.Group("/tasks")
	taskGroup.Get("/", handler.GetTasks)
	taskGroup.Post("/", handler.CreateTask)
	taskGroup.Delete("/:id", handler.DeleteTask)

}