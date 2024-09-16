package redis

import (
	"fmt"
	"web_App/settings"

	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	// 使用 fmt.Sprintf 格式化 Redis 地址
	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	rdb = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	_, err = rdb.Ping().Result()
	return
}

func Close() {
	_ = rdb.Close()
}
