package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	go func() {
	Label:
		for {
			select {
			case <-ch:
				break Label
			default:
				//必须等待这个结束了才行，不然main里面的一直阻塞
				fmt.Println("睡眠中")
				time.Sleep(time.Second * 1000)
				fmt.Println("睡眠结束")
			}
		}
		fmt.Println("协程结束")
	}()
	fmt.Println("等待5秒")
	time.Sleep(time.Second * 5)
	ch <- struct{}{}
	fmt.Println("5秒结束")
}
