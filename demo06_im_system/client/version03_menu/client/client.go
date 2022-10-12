//服务端反馈信息写得不好，这里无从下手处理信息，比如ChangeName

package client

import (
	"fmt"
	"log"
	"net"
)

type client struct {
	serverIP   string
	serverPort int
	name       string
	conn       net.Conn
}

func NewClient(serverIP string, serverPort int) (cli *client) {
	//创建客户端
	cli = new(client)
	cli.serverIP = serverIP
	cli.serverPort = serverPort
	cli.name = fmt.Sprintf("%s:%d", serverIP, serverPort)
	//建立连接
	conn, err := net.Dial("tcp", cli.name)
	if err != nil {
		log.Printf("net.Dial err: %v\n", err)
		return nil
	}
	cli.conn = conn
	go cli.getMsg() //开启输出服务端返回的信息
	return
}

// ChangeName 更改姓名
func (cli *client) ChangeName(name string) {
	cli.sendMsg("rename|" + name)
	cli.name = name //要是上面的操作失败了，这里更改就错误了
}

// PrivateChat 私聊信息
func (cli *client) PrivateChat(recipient string, msg string) {
	cli.sendMsg("to|" + recipient + "|" + msg)
}

// PublicChat 群聊信息
func (cli *client) PublicChat(msg string) {
	cli.sendMsg(msg)
}

// ExitConn 退出连接
func (cli *client) ExitConn() {
	cli.conn.Close()
}

// 做成一个单独的协程输出服务端返回的信息
func (cli *client) getMsg() {
	buf := make([]byte, 1024*4)
	for {
		n, _ := cli.conn.Read(buf)
		fmt.Print(string(buf[:n]))
	}
}

// 发送信息
func (cli *client) sendMsg(msg string) {
	cli.conn.Write([]byte(msg + "\n"))
}
