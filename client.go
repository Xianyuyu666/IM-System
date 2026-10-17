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
	flag      int //用户选择的模式
}

var serverIp string
var serverPort int

func (client *Client) menu() bool {
	var flag int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("4.退出")

	fmt.Scanln(&flag)

	if flag > 0 && flag <= 4 {
		client.flag = flag
		return true
	}
	fmt.Println(">>>请输入合法选择")
	return false
}

func NewClient(severip string, serverport int) *Client {
	client := &Client{
		SeverIP:   severip,
		SeverPort: serverport,
		flag:      999,
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

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {
		}
		switch client.flag {
		case 1:
			fmt.Println(1)
		case 2:
			fmt.Println(2)
		case 3:
			fmt.Println(3)
		case 4:
			fmt.Println(">>>退出成功")
			return
		}
	}
}

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址(默认127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口号(默认8888)")
}

func main() {
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>链接服务器失败")
		return
	}
	fmt.Println(">>>链接服务器成功\n", "IP:", serverIp, "POST:", serverPort)
	client.Run()
}
