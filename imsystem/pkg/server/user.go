package server

import (
	"net"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

// User Constructor
func NewUser(conn net.Conn) *User {
	// RemoteAddr() can get remote addr
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}

	// Start the user goroutine, Ready to work
	go user.listen()

	return user
}

// 监听当前的管道，有数据就可以发送给对端, 也就是客户端那边收到的数据
func (u *User) listen() {
	for {
		msg := <-u.C
		u.conn.Write([]byte(msg + "\r\n"))
	}
}
