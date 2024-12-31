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

type redisCacheCluster struct {
	client *redis.ClusterClient
}

func NewRedisCacheCluster(cfg *config.AppConfig) (CacheEngine, error) {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: cfg.Cache.Redis.Address,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		gplog.Fatalf("Failed to connect to redis: %v", err)
	}

	cacheEngine := redisCacheCluster{
		client: client,
	}
	return &cacheEngine, nil
}

func (r redisCacheCluster) Get(ctx context.Context, key string) ([]byte, error) {
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

func (r redisCacheCluster) GetInt(ctx context.Context, key string) int {
	gplog.Debugf("Get key: %s", key)
	result := r.client.Get(ctx, key)
	val, err := result.Int()
	if err != nil {
		return 0
	}

	return val
}

func (r redisCacheCluster) Keys(ctx context.Context, pattern string) ([]string, error) {
	gplog.Debugf("Get keys with pattern: %s", pattern)
	result := r.client.Keys(ctx, pattern)
	return result.Val(), result.Err()
}

func (r redisCacheCluster) Set(ctx context.Context, key string, val interface{}, ttl time.Duration) error {
	gplog.Debugf("Set key: %s", key)
	data, err := sonic.Marshal(val)
	if err != nil {
		return err
	}

	result := r.client.Set(ctx, key, data, ttl*time.Second)
	return result.Err()
}

func (r redisCacheCluster) Delete(ctx context.Context, key string) error {
	gplog.Debugf("Deleting key: %s", key)
	result := r.client.Del(ctx, key)
	return result.Err()
}

func (r redisCacheCluster) AddToSet(ctx context.Context, key string, val interface{}) error {
	gplog.Debugf("Add to set key: %s", key)
	data, err := sonic.Marshal(val)
	if err != nil {
		return err
	}

	result := r.client.SAdd(ctx, key, data)
	return result.Err()
}

func (r redisCacheCluster) RemoveFromSet(ctx context.Context, key string, val interface{}) error {
	gplog.Debugf("Remove from set key: %s", key)
	data, err := sonic.Marshal(val)
	if err != nil {
		return err
	}

	result := r.client.SRem(ctx, key, data)
	return result.Err()
}

func (r redisCacheCluster) IsMember(ctx context.Context, key string, val interface{}) (bool, error) {
	gplog.Debugf("Check member key: %s", key)
	data, err := sonic.Marshal(val)
	if err != nil {
		return false, err
	}

	result := r.client.SIsMember(ctx, key, data)
	return result.Val(), result.Err()
}

func (r redisCacheCluster) Ping(ctx context.Context) error {
	gplog.Debug("Ping redis")
	result := r.client.Ping(ctx)
	return result.Err()
}

func (r redisCacheCluster) Close() error {
	gplog.Debugf("Closing redis connection host: %s", r.client.Options().Addrs)
	return r.client.Close()
}
