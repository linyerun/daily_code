package redis

import (
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"sync"
	"time"
)

var (
	redisPool *redigo.Pool
	once      sync.Once
)

// redis pool
func poolInitRedis(server string, password string) *redigo.Pool {
	return &redigo.Pool{
		MaxIdle:     2, //空闲数
		IdleTimeout: 240 * time.Second,
		MaxActive:   3, //最大数
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

//单例模式创建redis pool
func newRedisPool() *redigo.Pool {
	once.Do(func() {
		redisPool = poolInitRedis(Addr+fmt.Sprintf(":%d", Port), Pwd)
	})
	return redisPool
}
