package taskServ

import (
	"context"
	"go-service-template/internal/domain/dto"
)

type TaskService interface {
	CreateTask(ctx context.Context, req *dto.CreateTaskReq) (*dto.CreateTaskResp, error)
	GetTaskList(ctx context.Context, req *dto.GetTaskListReq) (*dto.GetTaskListResp, error)
	GetTaskByID(ctx context.Context, id int) (*dto.TaskDTO, error)
	UpdateTask(ctx context.Context, req *dto.UpdateTaskReq) error
	DeleteTask(ctx context.Context, id int) error
}
