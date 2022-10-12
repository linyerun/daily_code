package server

import "net"

type User struct {
	Name   string
	Addr   string
	Ch     chan string
	Conn   net.Conn
	server *Server
}

func NewUser(conn net.Conn, server *Server) (user *User) {
	addr := conn.RemoteAddr().String()
	user = &User{
		Name:   addr,
		Addr:   addr,
		Ch:     make(chan string), //无缓冲通道
		Conn:   conn,
		server: server,
	}
	go user.StartListen() //开启用户发送协程
	user.Online()         //广播用户上线消息
	return
}

// StartListen 获取通道信息发送到客户端
func (u *User) StartListen() {
	for {
		//获取信息
		msg := <-u.Ch
		//发送信息
		u.Conn.Write([]byte(msg))
	}
}

// Online 用户上线
func (u *User) Online() {
	s := u.server
	s.lock.Lock()
	s.onlineUsers[u.Name] = u
	s.lock.Unlock()
	//广播这条信息
	s.broadcast(u.Name+"，上线！", u)
}

// Offline 用户下线
func (u *User) Offline() {
	s := u.server
	s.broadcast(u.Name+"，下线！", u)
	s.lock.Lock()
	delete(s.onlineUsers, u.Name)
	s.lock.Unlock()
}
