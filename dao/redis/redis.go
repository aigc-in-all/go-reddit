package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"goreddit/setting"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

func Init(cfg *setting.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return
}

func Close() {
	_ = client.Close()
}
