package cacheDB

import (
	"context"
	"go-service-template/config"
	"time"
)

type CacheEngine interface {
	Get(ctx context.Context, key string) ([]byte, error)
	GetInt(ctx context.Context, key string) int
	Keys(ctx context.Context, pattern string) ([]string, error)
	Set(ctx context.Context, key string, val interface{}, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
	AddToSet(ctx context.Context, key string, val interface{}) error
	RemoveFromSet(ctx context.Context, key string, val interface{}) error
	IsMember(ctx context.Context, key string, val interface{}) (bool, error)
	Ping(ctx context.Context) error
	Close() error
}

func NewCacheEngine(cfg *config.AppConfig) (CacheEngine, error) {
	if cfg.Cache.Redis.Mode == "cluster" {
		return NewRedisCacheCluster(cfg)
	}
	return NewRedisCacheStandalone(cfg)
}
