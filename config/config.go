package config

import (
	"errors"
	"fmt"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
	"time"
)

var Set = wire.NewSet(NewConfig)

type AppConfig struct {
	Server    Server
	DB        DB
	Cache     Cache
	Adapter   Adapter
	LogConfig LogConfig
}

type Server struct {
	Name                string
	Version             string
	Port                string
	Mode                string
	RequestTimeout      time.Duration
	ReadTimeout         time.Duration
	WriteTimeout        time.Duration
	CtxDefaultTimeout   time.Duration
	Debug               bool
	MaxCountRequest     int
	ExpirationLimitTime time.Duration
}

type DB struct {
	AutoMigrate bool
	MySQL       Database
}

type Database struct {
	Host            string
	Port            int
	UserName        string
	Password        string
	Database        string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	Options         string
}

type Cache struct {
	Redis Redis
}

type Redis struct {
	Mode    string
	Address []string
}

type Adapter struct {
	CentralAuthService HttpAdapter
}

type HttpAdapter struct {
	URL    string
	APIKey string
}

type LogConfig struct {
	Encoding string
	Level    string
}

// Get config path for local or docker
func getDefaultConfig() string {
	return "/config/config-example"
}

// NewConfig Load config file from given path
func NewConfig() (*AppConfig, error) {
	config := AppConfig{}
	path := os.Getenv("cfgPath")
	if path == "" {
		path = getDefaultConfig()
	}
	fmt.Printf("config path:%s\n", path)

	v := viper.New()

	v.SetConfigName(path)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	err := v.Unmarshal(&config)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &config, nil
}
