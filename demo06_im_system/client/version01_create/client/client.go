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
	return
}

func (cli *client) ChangeName(name string) {

}

func (cli *client) GetMsg() string {
	buf := make([]byte, 1024)
	n, _ := cli.conn.Read(buf)
	return string(buf[:n-1])
}
