package ormDB

import (
	"fmt"
	"go-service-template/config"
	"go-service-template/pkgs/gplog"
)

func InitMySQLConn(configuration *config.AppConfig) Gorm {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configuration.DB.MySQL.UserName,
		configuration.DB.MySQL.Password,
		configuration.DB.MySQL.Host,
		configuration.DB.MySQL.Port,
		configuration.DB.MySQL.Database,
	)
	cfg := Config{
		DSN:             dataSourceName,
		MaxOpenConn:     configuration.DB.MySQL.MaxOpenConn,
		MaxIdleConn:     configuration.DB.MySQL.MaxIdleConn,
		ConnMaxLifetime: configuration.DB.MySQL.ConnMaxLifetime,
		ConnMaxIdleTime: configuration.DB.MySQL.ConnMaxIdleTime,
		Debug:           true,
	}
	conn, err := New(cfg)
	if err != nil {
		panic(err)
	}

	gplog.Info("Gorm has created database connection - MYSQL")
	return conn

}
