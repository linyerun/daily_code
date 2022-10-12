package main

import (
	"flag"
	"version03_menu/client"
)

var (
	serverIP   string
	serverPort int
)

func init() {
	//可以通过 -h 查看这些提示
	//获取值：main.exe -port 9999 -ip 127.0.0.1
	flag.StringVar(&serverIP, "ip", "127.0.0.1", "请输入服务器ip地址(默认是127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "请输入服务器端口(默认是8888)")
}

func main() {
	//flag解析
	flag.Parse()

	//创建客户端
	cli := client.NewClient(serverIP, serverPort)
	if cli == nil {
		panic("远程连接失败")
	}
	//var choice int
	//var msg string
}
