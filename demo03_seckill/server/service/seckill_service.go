package service

import (
	"fmt"
	"log"
	. "server/redis"
)

type SeckillService struct {
}

func (s SeckillService) AddProduct() (err error) {
	client := NewRedisClient()
	defer client.Close()
	reply, err := client.Do("set %s %d", "product_apples", 100)
	log.Println(reply)
	return
}

func (s SeckillService) SeckillProduct(userId int64) error {
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
