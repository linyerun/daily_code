package main

import "version07_out/server"

func main() {
	server.NewServer("127.0.0.1", 8888).Start()
}
