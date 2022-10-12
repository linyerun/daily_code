//添加功能：发送who返回所有在线用户(包括自己)

package main

import "version05_search/server"

func main() {
	server.NewServer("127.0.0.1", 8888).Start()
}
