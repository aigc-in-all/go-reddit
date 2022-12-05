package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"goreddit/setting"
)

var rdb *redis.Client

func Init(cfg *setting.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return
}

func Close() {
	_ = rdb.Close()
}
