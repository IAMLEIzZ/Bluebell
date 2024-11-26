package redis

import (
	"context"
	"fmt"

	"github.com/iamleizz/bluebell/setting"
	"github.com/redis/go-redis/v9"
)

// redisClient 用于存储redis连接
var rdb *redis.Client

// initRedis 初始化redis连接
func Init(cfg *setting.RedisConfig) (err error) {
	// 初始化redis连接
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", 
			cfg.Host,
			cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	_, err = rdb.Ping(context.Background()).Result()

	return 
}

func Close() {
	_ = rdb.Close()
}