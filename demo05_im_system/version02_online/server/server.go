package server

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type Server struct {
	IP   string
	Port int
	//下面这三个还是不要被访问到为妙
	onlineUsers map[string]*User
	lock        sync.RWMutex
	ch          chan string
}

func NewServer(ip string, port int) (server *Server) {
	server = &Server{
		IP:          ip,
		Port:        port,
		onlineUsers: map[string]*User{},
		ch:          make(chan string),
	}
	go server.startChan() //开启广播通道协程
	return
}

func (s *Server) startChan() {
	for {
		msg := <-s.ch
		s.lock.RLock()
		for _, user := range s.onlineUsers {
			user.Ch <- msg
		}
		s.lock.RUnlock()
	}
}

func (s *Server) broadcast(msg string, user *User) {
	//拼接发送信息
	msg = "[" + user.Addr + "]" + user.Name + "：" + msg + "\n"
	s.ch <- msg
}

func (s *Server) handler(user *User) {

	//广播这条信息
	s.broadcast(user.Name+"，上线！", user)

	//阻塞协程
	select {}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		log.Printf("net.Listen err: %v \n", err)
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("listener.Accept err: %v \n", err)
			continue
		}
		user := NewUser(conn)
		s.lock.Lock()
		s.onlineUsers[user.Name] = user
		s.lock.Unlock()
		go s.handler(user)
	}
}
