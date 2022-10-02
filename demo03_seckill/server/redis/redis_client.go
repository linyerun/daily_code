package redis

import redigo "github.com/gomodule/redigo/redis"

// NewRedisClient 单例模式创建一个redis-client
func NewRedisClient() redigo.Conn {
	return newRedisPool().Get()
}
