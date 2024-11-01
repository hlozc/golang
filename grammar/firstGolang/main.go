package main

/*
* @TITLE 初识 Golang
* 1. GOPATH 表示 Golang 项目目录
* 2. GOROOT 表示 Go 编辑器相关的可执行文件目录

* @TEST 1 Golang 的优势
* 1. 非常快捷方便，依赖的库会被静态打包到可执行文件中（并且是机器码），一般的机器都可以执行运行，不需要重新配置环境
* 2. 静态语言，编译的时候大部分错误就可以直接检查出来，减少跑的时候出现的问题
* 3. 天生支持并发
* 4. 标准库强大，runtime 系统调度机制，高效的 GC 回收
*/

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello world")
	time.Sleep(5 * time.Second);
}
