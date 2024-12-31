package handler

import (
	"github.com/google/wire"
	"go-service-template/internal/handler/taskHandle"
)

var Set = wire.NewSet(
	taskHandle.NewTaskHandle,
)
