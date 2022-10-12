package server

import (
	"fmt"
	"io"
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
	//监听信息的接收
	go func() {
		buf := make([]byte, 1024*4) //大小4KB
		for {
			n, err := user.Conn.Read(buf)
			//nc被退出时触发，下线
			if n == 0 {
				user.Offline()
				return
			}
			//异常处理(ctrl+c退出不属于io.EOF，所以丢在下面吧)
			if err != nil && err != io.EOF {
				log.Printf("Conn.Read err: %v\n", err)
				return
			}
			//广播信息
			s.broadcast(string(buf[:n-1]), user) //不发送最后的'\n'
		}
	}()
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
		user := NewUser(conn, s)
		go s.handler(user)
	}
}
