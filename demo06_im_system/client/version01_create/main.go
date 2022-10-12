package main

import "version01_create/client"

func main() {
	cli := client.NewClient("127.0.0.1", 8888)
	if cli == nil {
		panic("远程连接失败")
	}
	println(cli.GetMsg())
	//阻塞一下看效果
	select {}
}
