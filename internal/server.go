//go:build wireinject
// +build wireinject

// this code to enable wire inject
package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"go-service-template/config"
	"go-service-template/internal/biz"
	"go-service-template/internal/handler"
	"go-service-template/internal/infrastructure/adapter"
	"go-service-template/internal/infrastructure/repository"
	"go-service-template/internal/middleware"
	"go-service-template/internal/router"
	"go-service-template/pkgs/database"
	"go-service-template/pkgs/database/cacheDB"
	"go-service-template/pkgs/database/ormDB"
	"go-service-template/pkgs/gplog"
)

type Server struct {
	cfg      *config.AppConfig
	app      *fiber.App
	_ormDB   ormDB.Gorm
	_cacheDB cacheDB.CacheEngine
}

func New() (*Server, error) {
	panic(wire.Build(wire.NewSet(
		NewServer,
		config.Set,
		database.Set,
		adapter.Set,
		repository.Set,
		biz.Set,
		handler.Set,
		middleware.Set,
		router.Set,
	)))
}

func NewServer(
	cfg *config.AppConfig,
	//insert all router here
	taskRouter router.TaskRoute,
	//insert all middleware here

	_ormDB ormDB.Gorm,
	_cacheDB cacheDB.CacheEngine,
) *Server {
	app := InitFiberApp(cfg)
	//init router
	v1 := app.Group("/v1")
	taskRouter.Init(&v1)

	return &Server{
		cfg:      cfg,
		app:      app,
		_ormDB:   _ormDB,
		_cacheDB: _cacheDB,
	}
}

func (serv Server) App() *fiber.App {
	return serv.app
}

func (serv Server) Config() *config.AppConfig {
	return serv.cfg
}

func (serv Server) Shutdown() {
	//close all connection
	_ = serv._ormDB.Close()
	_ = serv._cacheDB.Close()

	// Shutdown fiber app
	if err := serv.app.Shutdown(); err != nil {
		gplog.Errorf("Failed to shutdown fiber app cause:%v", err)
	}
}
