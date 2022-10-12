package main

import "version06_rename/server"

func main() {
	server.NewServer("127.0.0.1", 8888).Start()
}
