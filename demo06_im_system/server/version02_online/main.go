package main

import "vsersion02_online/server"

func main() {
	server.NewServer("127.0.0.1", 8888).Start()
}
