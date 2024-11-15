package server

import (
	"fmt"
	"log/slog"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

func NewServer(ip string, port int) *Server {
	svr := &Server{
		Ip:   ip,
		Port: port,
	}
	return svr
}

func (s *Server) handle(conn net.Conn) {
	slog.Info("Connect Ready")
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
