package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	Conn net.Conn
}

// 创建User的方法
func NewUser(conn net.Conn) *User {
	UserAddr := conn.RemoteAddr().String()

	user := &User{
		Name: UserAddr,
		Addr: UserAddr,
		C:    make(chan string),
		Conn: conn,
	}

	go user.ListenMassage()

	return user
}

// 监听服务器消息的方法，负责发送到客户端
func (this *User) ListenMassage() {
	for {
		msg := <-this.C

		this.Conn.Write([]byte(msg + "\n"))
	}
}
