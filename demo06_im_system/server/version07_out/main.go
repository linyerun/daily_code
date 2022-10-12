//通过调用user.Conn.SetReadDeadline(time.Now().Add(10 * time.Minute))完成我想要的效果

package main

import "version07_out/server"

func main() {
	server.NewServer("127.0.0.1", 8888).Start()
}
