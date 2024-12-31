package ormDB

import (
	"context"
	"database/sql"
	"go-service-template/config"
	"go-service-template/internal/domain/entity"
	"go-service-template/pkgs/gplog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// Gorm defines interface for access the database.
type Gorm interface {
	DB() *gorm.DB
	SqlDB() *sql.DB
	Exec(fc func(tx *gorm.DB) error, ctx context.Context) (err error)
	Transaction(fc func(tx *gorm.DB) error) (err error)
	Close() error
	DropTableIfExists(value interface{}) error
}

// Config GORM Config
type Config struct {
	Debug           bool
	DBType          string
	DSN             string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	TablePrefix     string
}

type _gorm struct {
	db    *gorm.DB
	sqlDB *sql.DB
	cfg   *config.AppConfig
}

func New(c Config) (Gorm, error) {
	dial := mysql.Open(c.DSN)
	gConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.TablePrefix,
			SingularTable: true,
		},
	}

	db, err := gorm.Open(dial, gConfig)
	if err != nil {
		return nil, err
	}

	if c.Debug {
		db = db.Debug()
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if c.MaxOpenConn != 0 {
		sqlDB.SetMaxOpenConns(c.MaxOpenConn)
	}
	if c.ConnMaxLifetime != 0 {
		sqlDB.SetConnMaxLifetime(c.ConnMaxLifetime)
	}
	if c.MaxIdleConn != 0 {
		sqlDB.SetMaxIdleConns(c.MaxIdleConn)
	}
	if c.ConnMaxIdleTime != 0 {
		sqlDB.SetConnMaxIdleTime(c.ConnMaxIdleTime)
	}

	// Auto migration
	err = db.AutoMigrate(&entity.Task{})
	if err != nil {
		gplog.Errorf("failed to auto migrate table: %v", err)
		return nil, err
	}

	return &_gorm{
		db:    db,
		sqlDB: sqlDB,
	}, nil
}
