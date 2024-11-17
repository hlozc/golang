package server

// 每个连接（User）启动的时候，都有协程监听管道，管道里面的数据就是要发送给客户端的数据
// 服务器的管道用来广播给所有的 User
// 每个连接（User）启动的时候，也要启动一个协程监听该连接的所有读事件，有数据就读取出来并广播

import (
	"fmt"
	"io"
	"log/slog"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	// Record All User
	onlineMap map[string]*User
	mapLock   sync.RWMutex
	Messages  chan string
}

func NewServer(ip string, port int) *Server {
	svr := &Server{
		Ip:        ip,
		Port:      port,
		onlineMap: make(map[string]*User),
		Messages:  make(chan string),
	}
	return svr
}

// 需要广播的消息，放在 Messages 里面，后续交给协程处理
func (s *Server) broatcast(user *User, msg string) {
	msg = fmt.Sprintf("[user: %v] %v", user.Name, msg)
	// debug
	slog.Info(msg)

	s.Messages <- msg
}

// 监听 Messages，随时同步广播数据
func (s *Server) listenMessages() {
	for {
		msg := <-s.Messages

		s.mapLock.Lock()
		for _, user := range s.onlineMap {
			user.C <- msg
		}
		s.mapLock.Unlock()
	}
}

// 监听这个连接是否有发送数据过来
func (s *Server) listenConn(user *User, conn net.Conn) {
	buffer := make([]byte, 4096)
	for {
		n, err := conn.Read(buffer)
		// 说明对方关闭连接了
		if n == 0 {
			s.broatcast(user, "offline")
			return
		}

		// 如果有错误，但是并不是正常读取到 EOF
		if err != nil && err != io.EOF {
			fmt.Println("Connect Error")
			return
		}

		// 去掉最后的 '\n', 得到用户发送过来的数据
		msg := string(buffer[:n])
		if msg[n-1] == '\n' {
			msg = msg[:n-1]
		}
		if n > 2 && msg[n-1] == '\r' {
			msg = msg[:n-1]
		}

		s.broatcast(user, msg)
	}
}

// 处理该连接的所有事件
// 其实做的事情就是：新连接到来了，为他做准备工具，然后广播其他用户，现在有一个新用户上线了
// 随后进入 select 准备随时处理该连接后续的所有操作
func (s *Server) handle(conn net.Conn) {
	user := NewUser(conn)

	// New User Online, put the msg in channel
	s.mapLock.Lock()
	s.onlineMap[user.Name] = user
	s.mapLock.Unlock()

	// 通知其他用户现在有新用户上线了
	s.broatcast(user, "online")

	// 还要监听这个新连接是否有数据到来
	go s.listenConn(user, conn)

	select {}
}

func (s *Server) Run() {
	// Listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		slog.Info(fmt.Sprintf("Listen Error, Reason: %v", err.Error()))
		return
	}
	// Dont Forget To Close
	defer listener.Close()

	go s.listenMessages()

	// Accept Connection, Accept() will return a conn obj
	for {
		conn, err := listener.Accept()
		if err != nil {
			slog.Info(fmt.Sprintf("Accept Error, Reason: %v", err.Error()))
			continue
		}

		// Have a new connect, begin to handle with goroutine
		go s.handle(conn)
	}
}
