package taskHandle

import "github.com/gofiber/fiber/v2"

type TaskHandle interface {
	GetAllTaskList(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	CreateTask(ctx *fiber.Ctx) error
	UpdateTask(ctx *fiber.Ctx) error
	DeleteTask(ctx *fiber.Ctx) error
}
