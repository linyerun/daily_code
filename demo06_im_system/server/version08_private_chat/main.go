package main

import "version08_private_chat/server"

func main() {
	server.NewServer("127.0.0.1", 8888).Start()
}
