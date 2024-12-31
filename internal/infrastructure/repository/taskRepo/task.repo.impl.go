package taskRepo

import (
	"context"
	"go-service-template/internal/domain/entity"
	"go-service-template/pkgs/database/ormDB"
	"go-service-template/pkgs/utils/pagable"
	"gorm.io/gorm"
)

type taskRepoImpl struct {
	orm ormDB.Gorm
}

func NewTaskRepo(orm ormDB.Gorm) TaskRepository {
	return &taskRepoImpl{orm: orm}
}

func (t taskRepoImpl) FindTaskByID(ctx context.Context, id int) (*entity.Task, error) {
	task := new(entity.Task)
	err := t.orm.Exec(func(tx *gorm.DB) error {
		return tx.Find(task, "id = ? and deleted_at is not null", id).Error
	}, ctx)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t taskRepoImpl) FindAndCountTaskList(ctx context.Context, query *pagable.Query) ([]*entity.Task, int64, error) {
	var tasks []*entity.Task
	var count int64
	err := t.orm.Exec(func(tx *gorm.DB) error {
		return tx.Model(&entity.Task{}).Where("deleted_at is not null").Count(&count).Error
	}, ctx)
	if err != nil {
		return nil, 0, err
	}

	err = t.orm.Exec(func(tx *gorm.DB) error {
		return tx.Limit(query.GetLimit()).Offset(query.GetOffset()).Find(&tasks, "deleted_at is not null").Error
	}, ctx)
	if err != nil {
		return nil, 0, err
	}

	return tasks, count, nil
}

func (t taskRepoImpl) Save(ctx context.Context, req *entity.Task) (*entity.Task, error) {
	err := t.orm.Transaction(func(tx *gorm.DB) error {
		return tx.Create(req).Error
	})
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (t taskRepoImpl) Update(ctx context.Context, req *entity.Task) (*entity.Task, error) {
	err := t.orm.Transaction(func(tx *gorm.DB) error {
		return tx.Save(req).Error
	})
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (t taskRepoImpl) UpdateByMap(ctx context.Context, id int, data map[string]interface{}) error {
	err := t.orm.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&entity.Task{}).Where("id = ?", id).Updates(data).Error
	})
	if err != nil {
		return err
	}

	return nil
}

func (t taskRepoImpl) Delete(ctx context.Context, id int) error {
	err := t.orm.Transaction(func(tx *gorm.DB) error {
		return tx.Delete(&entity.Task{}, id).Error
	})
	if err != nil {
		return err
	}

	return nil
}
