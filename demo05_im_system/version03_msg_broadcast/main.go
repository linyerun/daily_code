//对中文信息好像还是存在问题

package main

import "version03_msg_broadcast/server"

func main() {
	server.NewServer("127.0.0.1", 8888).Start()
}
