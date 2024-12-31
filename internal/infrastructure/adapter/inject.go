package adapter

import (
	"github.com/google/wire"
	"go-service-template/internal/infrastructure/adapter/cas"
)

var Set = wire.NewSet(cas.NewCasAdapter)
