package cacheDB

import (
	"context"
	"errors"
	"github.com/bytedance/sonic"
	"github.com/redis/go-redis/v9"
	"go-service-template/config"
	"go-service-template/pkgs/gplog"
	"time"
)

type redisCacheStandalone struct {
	client *redis.Client
}

func NewRedisCacheStandalone(cfg *config.AppConfig) (CacheEngine, error) {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Cache.Redis.Address[0],
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		gplog.Fatalf("Failed to connect to redis: %v", err)
	}

	cacheEngine := redisCacheStandalone{
		client: client,
	}
	return &cacheEngine, nil
}

func (r redisCacheStandalone) Get(ctx context.Context, key string) ([]byte, error) {
	gplog.Debugf("Get key: %s", key)
	result := r.client.Get(ctx, key)
	val, err := result.Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			gplog.Debugf("Key not found: %s", key)
			return nil, nil
		}
		return nil, err
	}

	return val, err
}

func (r redisCacheStandalone) GetInt(ctx context.Context, key string) int {
	gplog.Debugf("Get key: %s", key)
	result := r.client.Get(ctx, key)
	val, err := result.Int()
	if err != nil {
		return 0
	}

	return val
}
func (r redisCacheStandalone) Keys(ctx context.Context, pattern string) ([]string, error) {
	gplog.Debugf("Get keys with pattern: %s", pattern)
	result := r.client.Keys(ctx, pattern)
	return result.Val(), result.Err()
}

func (r redisCacheStandalone) Set(ctx context.Context, key string, val interface{}, ttl time.Duration) error {
	gplog.Debugf("Set key: %s", key)
	data, err := sonic.Marshal(val)
	if err != nil {
		return err
	}

	result := r.client.Set(ctx, key, data, ttl*time.Second)
	return result.Err()
}

func (r redisCacheStandalone) AddToSet(ctx context.Context, key string, val interface{}) error {
	gplog.Debugf("Add to set key: %s", key)
	data, err := sonic.Marshal(val)
	if err != nil {
		return err
	}

	result := r.client.SAdd(ctx, key, data)
	return result.Err()
}

func (r redisCacheStandalone) RemoveFromSet(ctx context.Context, key string, val interface{}) error {
	gplog.Debugf("Remove from set key: %s", key)
	data, err := sonic.Marshal(val)
	if err != nil {
		return err
	}

	result := r.client.SRem(ctx, key, data)
	return result.Err()
}

func (r redisCacheStandalone) IsMember(ctx context.Context, key string, val interface{}) (bool, error) {
	gplog.Debugf("Check member key: %s", key)
	data, err := sonic.Marshal(val)
	if err != nil {
		return false, err
	}

	result := r.client.SIsMember(ctx, key, data)
	return result.Val(), result.Err()
}

func (r redisCacheStandalone) Delete(ctx context.Context, key string) error {
	gplog.Debugf("Deleting key: %s", key)
	result := r.client.Del(ctx, key)
	return result.Err()

}

func (r redisCacheStandalone) Ping(ctx context.Context) error {
	gplog.Debugf("Ping redis")
	result := r.client.Ping(ctx)
	return result.Err()
}

func (r redisCacheStandalone) Close() error {
	gplog.Infof("Closing redis connection host: %s", r.client.Options().Addr)
	return r.client.Close()
}
