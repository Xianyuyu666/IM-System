package main

import (
	"fmt"
	"net"
)

type Client struct {
	SeverIP   string
	SeverPort int
	Name      string
	Conn      net.Conn
}

func NewClient(severip string, serverport int) *Client {
	client := &Client{
		SeverIP:   severip,
		SeverPort: serverport,
	}

	//链接
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", severip, serverport))
	if err != nil {
		fmt.Println("Dial error:", err)
		return nil
	}

	client.Conn = conn
	return client
}

func main() {
	client := NewClient("127.0.0.1", 8888)
	if client == nil {
		fmt.Println("链接服务器失败")
		return
	}
	fmt.Println("链接服务器成功")
	select {}
}
