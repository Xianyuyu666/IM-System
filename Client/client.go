package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
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

// 客户端处理服务器回复信息的方法
func (client *Client) DealResponse() {
	//一旦服务器发送信息直接打印到屏幕，永久阻塞等待
	io.Copy(os.Stdout, client.Conn)
}

// 客户端更新用户名的方法
func (client *Client) Rename() bool {
	var newName string
	fmt.Print(">>>请输入新用户名: \n")
	fmt.Scanln(&newName)

	msg := "rename|" + newName + "\n"
	_, err := client.Conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("Conn.Write error", err)
		return false
	}
	return true
}

func (client *Client) PublicChat() {
	for {
		//提示用户输入信息
		var msg string
		fmt.Print(">>>请输入广播信息(exit退出): \n")
		fmt.Scanln(&msg)
		if msg == "" {
			fmt.Println(">>>消息不可为空!")
			continue
		} else if msg == "exit" {
			fmt.Println(">>>已退出公聊模式")
			break
		}

		//发送到服务器
		_, err := client.Conn.Write([]byte(msg + "\n"))
		if err != nil {
			fmt.Println("Conn.Write error", err)
			return
		}
	}
}

func (client *Client) Run() {
	for client.flag != 4 {
		for client.menu() != true {
		}
		switch client.flag {
		case 1:
			client.PublicChat()
		case 2:
			fmt.Println(2)
		case 3:
			client.Rename()
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

	go client.DealResponse()

	fmt.Println(">>>链接服务器成功\n", "IP:", serverIp, "POST:", serverPort)
	client.Run()
}
