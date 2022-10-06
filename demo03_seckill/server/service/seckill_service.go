package service

import (
	lua "github.com/yuin/gopher-lua"
	"log"
	. "server/redis"
	"sync"
)

var (
	s        *seckillService
	once     sync.Once
	prodKey  = "product_apples"
	usersKey = "usersInProductApples"
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
	reply, err := client.Do("set", prodKey, 100)
	log.Printf("result: %v\n", reply.(string))
	client.Do("del", usersKey)
	return
}

// SeckillProduct 使用Lua脚本(失败的，lua脚本连怎么连接我的远程redis都不知道，还调用个啥)
//不知道尚硅谷的redis6怎么调用到的
func (s seckillService) SeckillProduct(userId int64) (msg string, err error) {
	//创建一个lua虚拟机
	L := lua.NewState()
	defer L.Close()

	//设置要调用的lua文件
	if err = L.DoFile("static/redis.lua"); err != nil {
		msg = err.Error()
		log.Println(msg)
		return
	}

	//调用函数
	err = L.CallByParam(lua.P{
		Fn:      L.GetGlobal("seckill"),
		NRet:    1,
		Protect: true,
	}, lua.LString(userId))
	if err != nil {
		msg = err.Error()
		log.Println(msg)
		return
	}

	res := L.Get(1).(lua.LNumber)

	switch res {
	case 0:
		msg = "抢购结束"
	case 1:
		msg = "抢购成功"
	case 2:
		msg = "重复抢购"
	}
	return
}

//func (s seckillService) SeckillProduct(userId int64) error {
//
//	// 1. 连接redis客户端
//	client := NewRedisClient()
//	defer client.Close()
//
//	//2. 拼接秒杀用户key
//	user := fmt.Sprintf("userId_%d", userId)
//
//	//3. 获取库存，如果库存null，秒杀还没有开始，为0表示秒杀结束
//	reply, _ := client.Do("get", prodKey)
//	if reply == nil {
//		return fmt.Errorf("秒杀还没开始")
//	} else if reply.([]uint8)[0] == '0' {
//		return fmt.Errorf("秒杀结束")
//	}
//
//	//4. 判断用户是否重复秒杀操作
//	reply, _ = client.Do("sismember", usersKey, user)
//	if reply.(int64) == 1 {
//		return fmt.Errorf("秒杀用户重复")
//	}
//
//	//5. 执行秒杀操作
//	if _, err := client.Do("watch", prodKey, usersKey); err != nil {
//		// _是string类型
//		return fmt.Errorf("秒杀失败")
//	}
//	if _, err := client.Do("multi"); err != nil {
//		// _是string类型
//		return fmt.Errorf("秒杀失败")
//	}
//
//	if _, err := client.Do("decr", prodKey); err != nil {
//		client.Do("discard")
//		return fmt.Errorf("秒杀失败")
//	}
//
//	if _, err := client.Do("sadd", usersKey, user); err != nil {
//		client.Do("discard")
//		return fmt.Errorf("秒杀失败")
//	}
//
//	if _, err := client.Do("exec"); err != nil {
//		client.Do("discard")
//		return fmt.Errorf("秒杀失败")
//	}
//
//	return nil
//}
