package main

import "version04_optimize/server"

func main() {
	server.NewServer("127.0.0.1", 8888).Start()
}
