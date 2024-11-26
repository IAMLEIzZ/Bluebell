package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

// redisClient 用于存储redis连接
var rdb *redis.Client

// initRedis 初始化redis连接
func Init() (err error) {
	// 初始化redis连接
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", 
			viper.GetString("redis.host"),
			viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
	})
	_, err = rdb.Ping(context.Background()).Result()

	return 
}

func Close() {
	_ = rdb.Close()
}