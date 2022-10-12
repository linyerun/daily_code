package server

import "net"

type User struct {
	Name string
	Addr string
	Ch   chan string
	Conn net.Conn
}

func NewUser(conn net.Conn) (user *User) {
	addr := conn.RemoteAddr().String()
	user = &User{
		Name: addr,
		Addr: addr,
		Ch:   make(chan string), //无缓冲通道
		Conn: conn,
	}
	go user.StartListen() //开启用户发送协程
	return
}

func (u *User) StartListen() {
	for {
		//获取信息
		msg := <-u.Ch
		//发送信息
		u.Conn.Write([]byte(msg))
	}
}
