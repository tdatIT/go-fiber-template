package router

import (
	"github.com/gofiber/fiber/v2"
	"go-service-template/internal/handler/taskHandle"
	"go-service-template/internal/middleware"
)

type TaskRoute interface {
	Init(root *fiber.Router)
}

type taskRouter struct {
	taskHandle taskHandle.TaskHandle
	authMdw    *middleware.AuthMiddleware
}

func NewTaskRoute(
	authMdw *middleware.AuthMiddleware,
	taskHandle taskHandle.TaskHandle,
) TaskRoute {
	//return task router
	return &taskRouter{
		authMdw:    authMdw,
		taskHandle: taskHandle,
	}
}

func (s taskRouter) Init(root *fiber.Router) {
	senderRouter := (*root).Group("/tasks")
	{
		senderRouter.Post("", s.authMdw.RequiredAuthentication(), s.taskHandle.CreateTask)
		senderRouter.Get("", s.taskHandle.GetAllTaskList)
		senderRouter.Get("/:id", s.taskHandle.GetAllTaskList)
		senderRouter.Patch("/:id", s.taskHandle.UpdateTask)
		senderRouter.Delete("/:id", s.taskHandle.DeleteTask)
	}
}
