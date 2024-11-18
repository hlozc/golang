package client

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/fatih/color" // 引入颜色包
)

const (
	PUBLIC_CHAT  = 1
	PRIVATE_CHAT = 2
	RENAME       = 3
	WHO_ONLINE   = 4
	EXIT         = 0
	NUM_OPT      = 5
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	opt        int
}

var RemoteIp string
var RemotePort int

func NewClient(ip string, port int) *Client {
	conn, err := net.Dial("tcp", fmt.Sprintf("%v:%d", ip, port))
	if err != nil {
		fmt.Println("Connection failed. Please try again.")
		return nil
	}

	client := &Client{
		ServerIp:   ip,
		ServerPort: port,
		conn:       conn,
		opt:        NUM_OPT,
	}

	go client.dealResponse()

	return client
}

// 永久阻塞，监听服务器给我发送了什么，然后显示到终端上
// io.Copy() 就是将后者的信息拷贝到前者 (当然前提是 io)
func (c *Client) dealResponse() {
	io.Copy(os.Stdout, c.conn)
}

// 美化后的菜单
func (c *Client) menu() bool {
	var opt int

	// 美化输出菜单
	fmt.Println()
	color.Set(color.FgGreen, color.Bold) // 设置绿色加粗
	fmt.Println("╔════════════════════════════════════════════╗")
	fmt.Println("║              Welcome to Client             ║")
	fmt.Println("╚════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Please choose an option:")

	// 分隔线
	fmt.Println("───────────────────────────────────────────────")

	color.Set(color.FgCyan) // 设置蓝色
	fmt.Println("1. Public Chat")
	fmt.Println("2. Private Chat")
	fmt.Println("3. Rename")
	fmt.Println("4. Who Online")
	fmt.Println("0. Exit")

	fmt.Println("───────────────────────────────────────────────")
	color.Unset() // 重置颜色

	// 输入选项
	fmt.Print("Enter option (0-4): ")
	fmt.Scanln(&opt)
	if opt >= 0 && opt < NUM_OPT {
		c.opt = opt
		return true
	} else {
		fmt.Println("Invalid option. Please try again.")
		return false
	}
}

func (c *Client) sendMsg(msg string) {
	_, err := c.conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("Client send msg fail")
		return
	}
}

func (c *Client) rename() {
	var name string

	fmt.Println("Please select your new name")
	fmt.Scanln(&name)
	c.sendMsg("rename|" + name + "\r\n")
}

func (c *Client) publicChat() {
	var msg string

	fmt.Println("Entry your message > ")
	fmt.Scanln(&msg)
	c.sendMsg(msg + "\r\n")
}

func (c *Client) privateChat() {
	var who, msg string

	fmt.Println("Who to talk > ")
	fmt.Scanln(&who)

	fmt.Println("And tell to him/her > ")
	fmt.Scanln(&msg)

	c.sendMsg("to|" + who + "|" + msg)
}

func (c *Client) whoOnline() {
	c.sendMsg("who")
}

func (c *Client) Run() {
	for c.opt != 0 && c.menu() {
		switch c.opt {
		case PUBLIC_CHAT:
			c.publicChat()
		case PRIVATE_CHAT:
			c.privateChat()
		case RENAME:
			c.rename()
		case WHO_ONLINE:
			c.whoOnline()
		case EXIT:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

func init() {
	// 后面三个参数分别表示：ip 表示 -ip，第二个表示默认值，第三个表示 -help 的时候该选项的说明
	flag.StringVar(&RemoteIp, "ip", "127.0.0.1", "default 127.0.0.1")
	flag.IntVar(&RemotePort, "port", 8080, "default 8080")
}
