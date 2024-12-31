package database

import (
	"github.com/google/wire"
	"go-service-template/pkgs/database/cacheDB"
	"go-service-template/pkgs/database/ormDB"
)

var Set = wire.NewSet(
	ormDB.InitMySQLConn,
	cacheDB.NewCacheEngine,
)
