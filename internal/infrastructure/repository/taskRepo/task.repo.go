package taskRepo

import (
	"context"
	"go-service-template/internal/domain/entity"
	"go-service-template/pkgs/utils/pagable"
)

type TaskRepository interface {
	FindTaskByID(ctx context.Context, id int) (*entity.Task, error)
	FindAndCountTaskList(ctx context.Context, query *pagable.Query) ([]*entity.Task, int64, error)
	Save(ctx context.Context, req *entity.Task) (*entity.Task, error)
	Update(ctx context.Context, req *entity.Task) (*entity.Task, error)
	UpdateByMap(ctx context.Context, id int, data map[string]interface{}) error
	Delete(ctx context.Context, id int) error
}
