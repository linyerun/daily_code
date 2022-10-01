package config

import (
	"fmt"
	"github.com/go-redis/redis"
)

var (
	RedisCli *redis.Client
)

func InitRedis() {
	initConfig() //初始化一下，加载需要的配置
	RedisCli = redis.NewClient(&redis.Options{
		Addr:     Addr + ":" + fmt.Sprintf("%d", Port),
		Password: Pwd,
		DB:       0,
	})
}
