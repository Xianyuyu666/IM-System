package main

import (
	"fmt"
	"net"
)

type Server struct {
	IP   string
	Port int
}

func newServer(ip string, port int) *Server {
	return &Server{IP: ip, Port: port}
}

func (this *Server) Handler(conn net.Conn) {
	// 业务逻辑
	fmt.Println("链接成功")
}

func (this *Server) Start() {
	//listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.IP, this.Port))
	if err != nil {
		fmt.Println("Listen error:", err)
		return
	}
	defer listener.Close()

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
