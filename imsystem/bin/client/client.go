package main

import (
	"flag"
	"fmt"

	"github.com/hlozc/imsystem/pkg/client"
)

func main() {
	// 解析命令行参数
	flag.Parse()

	cli := client.NewClient(client.RemoteIp, client.RemotePort)
	if cli == nil {
		return
	}

	fmt.Println("╔════════════════════════════════════════════╗")
	fmt.Println("║              Connect Success!              ║")
	fmt.Println("╚════════════════════════════════════════════╝")
	cli.Run()
}
