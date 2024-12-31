package taskHandle

import (
	"github.com/gofiber/fiber/v2"
	"go-service-template/internal/biz/taskServ"
	"go-service-template/internal/domain/dto"
	"go-service-template/pkgs/gplog"
	responses "go-service-template/pkgs/utils/common/response"
	errors "go-service-template/pkgs/utils/common/servErr"
	"go-service-template/pkgs/utils/pagable"
)

func NewTaskHandle(taskServ taskServ.TaskService) TaskHandle {
	return &taskHandleImpl{
		taskServ: taskServ,
	}
}

type taskHandleImpl struct {
	taskServ taskServ.TaskService
}

func (t taskHandleImpl) GetAllTaskList(ctx *fiber.Ctx) error {
	query, err := pagable.GetQueryFromFiberCtx(ctx)
	if err != nil {
		gplog.Errorf("failed to get query from fiber ctx: %v", err)
		return errors.ErrBadRequest
	}

	req := new(dto.GetTaskListReq)
	req.Query = query

	resp, err := t.taskServ.GetTaskList(ctx.Context(), req)
	if err != nil {
		gplog.Errorf("failed to get task list: %v", err)
		return taskErrHandle(err)
	}

	response := responses.DefaultSuccess
	response.Data = resp
	return response.JSON(ctx)
}

func (t taskHandleImpl) GetById(ctx *fiber.Ctx) error {
	req := new(dto.GetTaskByIdReq)

	if err := ctx.ParamsParser(req); err != nil {
		gplog.Errorf("failed to parse request: %v", err)
		return errors.ErrBadRequest
	}

	resp, err := t.taskServ.GetTaskByID(ctx.Context(), req.ID)
	if err != nil {
		gplog.Errorf("failed to get task by id: %v", err)
		return taskErrHandle(err)
	}

	response := responses.DefaultSuccess
	response.Data = resp
	return response.JSON(ctx)
}

func (t taskHandleImpl) CreateTask(ctx *fiber.Ctx) error {
	req := new(dto.CreateTaskReq)

	if err := ctx.BodyParser(req); err != nil {
		gplog.Errorf("failed to parse request: %v", err)
		return errors.ErrBadRequest
	}

	resp, err := t.taskServ.CreateTask(ctx.Context(), req)
	if err != nil {
		gplog.Errorf("failed to create task: %v", err)
		return taskErrHandle(err)
	}

	response := responses.DefaultSuccess
	response.Data = resp
	return response.JSON(ctx)
}

func (t taskHandleImpl) UpdateTask(ctx *fiber.Ctx) error {
	req := new(dto.UpdateTaskReq)
	if err := ctx.BodyParser(req); err != nil {
		gplog.Errorf("failed to parse request: %v", err)
		return errors.ErrBadRequest
	}

	if err := ctx.ParamsParser(&req); err != nil {
		gplog.Errorf("failed to parse request: %v", err)
		return errors.ErrBadRequest
	}

	err := t.taskServ.UpdateTask(ctx.Context(), req)
	if err != nil {
		gplog.Errorf("failed to update task: %v", err)
		return taskErrHandle(err)
	}

	return responses.DefaultSuccess.JSON(ctx)
}

func (t taskHandleImpl) DeleteTask(ctx *fiber.Ctx) error {
	var id int
	if err := ctx.ParamsParser(&id); err != nil {
		gplog.Errorf("failed to parse request: %v", err)
		return errors.ErrBadRequest
	}

	err := t.taskServ.DeleteTask(ctx.Context(), id)
	if err != nil {
		gplog.Errorf("failed to delete task: %v", err)
		return taskErrHandle(err)
	}

	return responses.DefaultSuccess.JSON(ctx)
}
