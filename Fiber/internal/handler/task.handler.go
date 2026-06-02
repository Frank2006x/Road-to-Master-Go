package handler

import (
	"fiber/internal/data"
	"fiber/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type TaskHandler struct {
}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{}
}

func (h *TaskHandler) GetTasks(c fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"tasks": data.Tasks,
	})
}

func (h *TaskHandler) CreateTask(c fiber.Ctx) error {
	var newTask models.Task

	if err:=c.Bind().Body(&newTask); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	data.Tasks = append(data.Tasks, newTask)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"task": newTask,
	})
}

func (h *TaskHandler) DeleteTask(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid task ID",
		})
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid task ID",
		})
	}
	for i, task := range data.Tasks {
		if task.ID == idInt {
			data.Tasks = append(data.Tasks[:i], data.Tasks[i+1:]...)
			return c.JSON(fiber.Map{
				"message": "Task deleted successfully",
			})
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Task not found",
	})
}

