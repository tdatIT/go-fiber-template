package taskServ

import (
	"context"
	"go-service-template/config"
	"go-service-template/internal/domain/dto"
	"go-service-template/internal/domain/entity"
	"go-service-template/internal/infrastructure/repository/taskRepo"
	"go-service-template/pkgs/gplog"
	"go-service-template/pkgs/utils/mapper"
	"go-service-template/pkgs/utils/pagable"
)

func NewTaskService(cfg *config.AppConfig, taskRepo taskRepo.TaskRepository) TaskService {
	return &taskService{
		cfg:      cfg,
		taskRepo: taskRepo,
	}
}

type taskService struct {
	cfg      *config.AppConfig
	taskRepo taskRepo.TaskRepository
}

func (t taskService) CreateTask(ctx context.Context, req *dto.CreateTaskReq) (*dto.CreateTaskResp, error) {
	timeReq := req.Deadline.GetTime()
	model := &entity.Task{
		Title:       req.Title,
		Status:      entity.DBModelActive,
		Description: req.Description,
		Deadline:    &timeReq,
	}

	savedModel, err := t.taskRepo.Save(ctx, model)
	if err != nil {
		gplog.Errorf("[CreateTask] failed to save task: %v", err)
		return nil, err
	}

	return &dto.CreateTaskResp{
		ID: savedModel.ID,
	}, nil
}

func (t taskService) GetTaskList(ctx context.Context, req *dto.GetTaskListReq) (*dto.GetTaskListResp, error) {
	entities, total, err := t.taskRepo.FindAndCountTaskList(ctx, req.Query)
	if err != nil {
		gplog.Errorf("[GetTaskList] failed to get task list: %v", err)
		return nil, err
	}

	var tasks []*dto.TaskDTO
	if err := mapper.BindingStruct(entities, &tasks); err != nil {
		gplog.Errorf("[GetTaskList] failed to binding task list: %v", err)
		return nil, err
	}

	resp := pagable.ListResponse{
		Items:   tasks,
		Total:   total,
		Page:    req.Query.GetPage(),
		Size:    req.Query.GetSize(),
		HasMore: req.Query.GetHasMore(int(total)),
	}

	return &dto.GetTaskListResp{
		ListResponse: resp,
	}, nil
}

func (t taskService) GetTaskByID(ctx context.Context, id int) (*dto.TaskDTO, error) {
	model, err := t.taskRepo.FindTaskByID(ctx, id)
	if err != nil {
		gplog.Errorf("[GetTaskByID] failed to get task by id: %v", err)
		return nil, err
	}

	var task dto.TaskDTO
	if err := mapper.BindingStruct(model, &task); err != nil {
		gplog.Errorf("[GetTaskByID] failed to binding task: %v", err)
		return nil, err
	}

	return &task, nil
}

func (t taskService) UpdateTask(ctx context.Context, req *dto.UpdateTaskReq) error {
	timeReq := req.Deadline.GetTime()
	model := &entity.Task{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Deadline:    &timeReq,
	}

	_, err := t.taskRepo.Update(ctx, model)
	if err != nil {
		gplog.Errorf("[UpdateTask] failed to update task: %v", err)
		return err
	}

	return nil
}

func (t taskService) DeleteTask(ctx context.Context, id int) error {
	err := t.taskRepo.Delete(ctx, id)
	if err != nil {
		gplog.Errorf("[DeleteTask] failed to delete task: %v", err)
		return err
	}

	return nil
}
