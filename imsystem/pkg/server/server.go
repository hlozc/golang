package server

import (
	"fmt"
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

// 处理该连接的所有事件
func (s *Server) handle(conn net.Conn) {
	user := NewUser(conn)

	// New User Online, put the msg in channel
	s.mapLock.Lock()
	s.onlineMap[user.Name] = user
	s.mapLock.Unlock()

	s.broatcast(user, "online")

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
