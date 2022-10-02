package main

import (
	"github.com/go-redis/redis"
)

func main() {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.133:6379",
		Password: "redis",
		DB:       0,
	})
	//体现不出事务特性
	pipeline := redisCli.TxPipeline()
	defer pipeline.Close()
	pipeline.IncrBy("p1", -100)
	pipeline.IncrBy("p2", 100)
	pipeline.IncrBy("p3", -100)
	if _, err := pipeline.Exec(); err != nil {
		pipeline.Discard()
	}
}
