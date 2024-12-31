package server

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/bytedance/sonic"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hellofresh/health-go/v5"
	"go-service-template/config"
	"go-service-template/pkgs/gplog"
	"go-service-template/pkgs/hltchk"
	errors "go-service-template/pkgs/utils/common/servErr"
)

func InitFiberApp(cfg *config.AppConfig) *fiber.App {
	app := fiber.New(fiber.Config{
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		JSONDecoder:  sonic.Unmarshal,
		JSONEncoder:  sonic.Marshal,
		ErrorHandler: errors.CustomErrorHandler,
	})

	//init default gplog
	appLog := gplog.NewLogger(&gplog.LogConfig{
		ServiceName: cfg.Server.Name,
		Level:       cfg.LogConfig.Level,
		LogFormat:   gplog.JsonFormat,
		TimeFormat:  gplog.RFC3339NanoTimeEncoder,
	})

	// recover middleware
	app.Use(recover.New())

	//init health check
	h, _ := hltchk.InitHealthCheckService(cfg)
	app.Get("/health", adaptor.HTTPHandlerFunc(h.HandlerFunc))

	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/liveness",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			result := h.Measure(c.Context())
			return result.Status == health.StatusOK
		},
		ReadinessEndpoint: "/readiness",
	}))
	gplog.SetLogger(appLog)

	//swagger
	swgCfg := swagger.Config{
		BasePath: "/api/v1/",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
		Title:    "Swagger API Docs",
	}
	app.Use(swagger.New(swgCfg))

	//metrics
	prometheus := fiberprometheus.New(cfg.Server.Name)
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	return app
}
