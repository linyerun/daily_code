package main

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:   ip,
		Port: port,
	}
}

func (s *Server) handler(conn net.Conn) {
	defer conn.Close()
	fmt.Println("连接成功")
}

func (s *Server) Start() {
	//监听ip:port
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		log.Printf("net.Listen err: %v\n", err)
		return
	}
	//关闭监听
	defer listener.Close()
	for {
		//获取连接
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept err: %v\n", err)
			continue
		}
		//连接处理
		go s.handler(conn)
	}
}
