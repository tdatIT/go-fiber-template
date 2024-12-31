package biz

import (
	"github.com/google/wire"
	"go-service-template/internal/biz/taskServ"
)

var Set = wire.NewSet(
	taskServ.NewTaskService,
)
