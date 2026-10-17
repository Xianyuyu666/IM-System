package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	IP   string
	Port int

	//用户列表
	OnlineMap map[string]*User
	MapLock   sync.RWMutex

	//消息广播的channel
	Message chan string
}

// 创建一个Server的方法
func NewServer(ip string, port int) *Server {
	return &Server{
		IP:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		MapLock:   sync.RWMutex{},
		Message:   make(chan string),
	}
}

// 监听Message广播消息channel的goroutine,一旦有消息发送到所有在线的用户
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message

		this.MapLock.Lock()
		for _, u := range this.OnlineMap {
			u.C <- msg
		}
		this.MapLock.Unlock()
	}
}

// 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendmsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	this.Message <- sendmsg
}

// 执行业务逻辑方法
func (this *Server) Handler(conn net.Conn) {
	//用户上线，加入用户列表OnlineMap中
	user := NewUser(conn)

	this.MapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.MapLock.Unlock()
	//把上线信息广播到所有在线用户中
	this.BroadCast(user, "上线了")

	//handler阻塞
	select {}
}

// 启动服务的方法
func (this *Server) Start() {
	//listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.IP, this.Port))
	if err != nil {
		fmt.Println("Listen error:", err)
		return
	}
	defer listener.Close()
	//启动监听Message
	go this.ListenMessage()
	//accept
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}

		go this.Handler(conn)
	}
}
