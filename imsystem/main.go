package main

import "github.com/hlozc/imsystem/pkg/server"

// Compile command: go build -o server.exe main.go pkg/server/server.go

func main() {
	svr := server.NewServer("127.0.0.1", 8080)
	svr.Run()
}
