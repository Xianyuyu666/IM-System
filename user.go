package main

import "net"

type User struct {
	Name   string
	Addr   string
	C      chan string
	Conn   net.Conn
	server *Server
}

// 创建User的方法
func NewUser(conn net.Conn, server *Server) *User {
	UserAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   UserAddr,
		Addr:   UserAddr,
		C:      make(chan string),
		Conn:   conn,
		server: server,
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

// 用户上线的方法
func (this *User) Online() {
	//用户上线，添加进用户列表中
	this.server.MapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.MapLock.Unlock()

	//广播上线信息
	this.server.BroadCast(this, "上线")
}

// 用户下线的方法
func (this *User) Offline() {
	//用户下线，从用户列表中删除
	this.server.MapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.MapLock.Unlock()

	//广播下线信息
	this.server.BroadCast(this, "下线")
}

// 用户处理信息的方法
func (this *User) DoMessage(msg string) {
	//广播这个用户发送的信息
	this.server.BroadCast(this, msg)
}
