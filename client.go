package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	SeverIP   string
	SeverPort int
	Name      string
	Conn      net.Conn
}

var serverIp string
var serverPort int

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

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址(默认127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口号(默认8888)")
}

func main() {
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println("链接服务器失败")
		return
	}
	fmt.Println("链接服务器成功\n", "IP:", serverIp, "POST:", serverPort)
	select {}
}
