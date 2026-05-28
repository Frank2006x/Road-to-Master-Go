package handler

import (
	"fiber/internal/data"

	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{}
}

func (h *TaskHandler) GetTasks(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"tasks": data.Tasks,
	})
}