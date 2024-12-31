package hltchk

import (
	"fmt"
	"github.com/hellofresh/health-go/v5"
	healthMysql "github.com/hellofresh/health-go/v5/checks/mysql"
	healthRedis "github.com/hellofresh/health-go/v5/checks/redis"
	"go-service-template/config"

	"time"
)

func InitHealthCheckService(config *config.AppConfig) (*health.Health, error) {
	h, err := health.New(health.WithComponent(health.Component{
		Name:    "template-service",
		Version: config.Server.Version,
	}))
	if err != nil {
		return nil, err
	}

	//mysql check
	err = h.Register(health.Config{
		Name:      "mysql",
		Timeout:   time.Second * 2,
		SkipOnErr: false,
		Check: healthMysql.New(healthMysql.Config{
			DSN: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
				config.DB.MySQL.UserName,
				config.DB.MySQL.Password,
				config.DB.MySQL.Host,
				config.DB.MySQL.Port,
				config.DB.MySQL.Database,
			),
		}),
	})
	if err != nil {
		return nil, err
	}

	//redis check
	err = h.Register(health.Config{
		Name:      "redis",
		Timeout:   time.Second * 2,
		SkipOnErr: false,
		Check: healthRedis.New(healthRedis.Config{
			DSN: fmt.Sprintf("redis://%s", config.Cache.Redis.Address[0]),
		}),
	})
	if err != nil {
		return nil, err
	}

	return h, nil
}
