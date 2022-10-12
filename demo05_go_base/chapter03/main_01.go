//测试关闭连接，read的反应

package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	go func() {
		defer func() { fmt.Println("退出协程") }()
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("n == 0")
			return
		}
		if err != nil {
			fmt.Println("err != nil")
			return
		}
	}()
	//关闭试一下开启的那个协程是什么反映
	conn.Close()
	select {}
}
