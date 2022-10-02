package service

import (
	"fmt"
	"log"
	. "server/redis"
	"sync"
)

var (
	s    *seckillService
	once sync.Once
)

type seckillService struct {
}

func NewSeckillService() *seckillService {
	once.Do(func() {
		s = &seckillService{}
	})
	return s
}

func (s seckillService) AddProduct() (err error) {
	client := NewRedisClient()
	defer client.Close()
	reply, err := client.Do("set", "product_apples", 100)
	log.Printf("result: %v\n", reply.(string))
	return
}

func (s seckillService) SeckillProduct(userId int64) error {
	// 1. 连接redis客户端
	client := NewRedisClient()
	defer client.Close()
	//2. 拼接秒杀用户key
	userKey := fmt.Sprintf("userId_%d", userId)

	//3. 判断用户是否重复秒杀操作
	reply, _ := client.Do("sismember %s %s", "usersInProductApples", userKey)
	if reply.(int) == 1 {
		return fmt.Errorf("秒杀用户重复")
	}

	//3. 监视库存
	reply, _ = client.Do("get product_apples")
	if reply.(int) == 0 {
		return fmt.Errorf("秒杀结束")
	}
	//4. 获取库存，如果库存null，秒杀还没有开始

	//5. 判断用户是否重复秒杀操作

	//6. 判断如果商品数量，库存数量小于1，秒杀结束

	return nil
}
