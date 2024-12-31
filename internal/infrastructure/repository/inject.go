package repository

import (
	"github.com/google/wire"
	"go-service-template/internal/infrastructure/repository/taskRepo"
)

var Set = wire.NewSet(
	taskRepo.NewTaskRepo,
)
