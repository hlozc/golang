package server

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
	svr  *Server // 表示这个 User 是属于哪个 Server 管理的
}

// User Constructor
func NewUser(svr *Server, conn net.Conn) *User {
	// RemoteAddr() can get remote addr
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
		svr:  svr,
	}

	// Start the user goroutine, Ready to work
	go user.listenChan()

	return user
}

// 监听当前的管道，有数据就可以发送给对端, 也就是客户端那边收到的数据
func (u *User) listenChan() {
	for {
		msg := <-u.C
		u.conn.Write([]byte(msg + "\r\n"))
	}
}

// 给当前用户发送消息
func (u *User) sendMsg(msg string) {
	fmt.Println("Debug Msg: ", msg)

	_, err := u.conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("Send msg to user fail")
	}
}

// 上线功能
func (u *User) online() {
	// New User Online, put the msg in channel
	u.svr.addUser(u)

	// 通知其他用户现在有新用户上线了
	u.svr.broatcast(u, "online")
}

// 下线功能
func (u *User) offline() {
	// User Offline
	u.svr.delUser(u)

	u.svr.broatcast(u, "offline")
}

// 收到用户的消息之后，所做的事情
func (u *User) handleMessage(msg string) {
	if msg == "who" {
		msg := u.svr.showOnlines()
		u.sendMsg(msg)
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		name := strings.Split(msg, "|")[1]
		res := u.svr.updateName(u, name)
		u.sendMsg(res)
	} else {
		u.svr.broatcast(u, msg)
	}

}
