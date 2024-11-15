package main

/*
* TEST 1 pair
* 1. 对于 Golang 的每一个变量，都存在两个值，type 和 value
* 2. type 包括 static type 和 concrete type
* 3. 比如下面这个表达式

str = "hello"
var alltype interface{}
alltype = string

* 那么这个变量对应的是：pair<statictype: string, value: "hello"> 这样的格式
* 不管这个 str 赋值给谁，哪怕是 interface{}，那么 pair<statictype: string, value: "hello"> 永远都会完整地转递
* 不管变量怎么传递，pair<> 对都是完整传递的

eg:
	// 那么这时候 tty 这个变量对应的 pair 就是：
	// pair<type: *os.File, value: "/dev/tty"文件描述符>
	// 类型就是 os.File，即文件类型，对应的 value 就是这个文件的文件描述符，控制体
	tty, err := io.OpenFile("/dev/tty", os.O_RDWR, 0)

	if err != nil {
		fmt.Println("open file erro")
		return
	}

	var reader io.Reader
	// reader 本身是 io.Reader 类型的，但是却接收了 tty
	// 所以 reader 本质上的类型是 pair<type: *os.File, value: "/dev/tty"文件描述符>
	reader = tty

	var writer io.Writer
	// 不管这里怎么传递，哪怕 reader 进行了类型转换
	// 最终 writer 的类型都是 tty 一开始的类型
	writer = reader.(io.Writer)

* 核心：
* 1. 每个变量都有自己的 pair
* 2. 并且这个 pair 传递的时候会保证准确性
 */

import (
	"fmt"
)

func test6() {
	var str string = "hello world"
	

	var alltype interface{}
	alltype = str

	if v, ok := alltype.(string); ok {
		fmt.Println("value is ", v)
	}
}

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {

}

func (this *Book) ReadBook() {
	fmt.Println("Read Book")
}

func (this *Book) WriteBook() {
	fmt.Println("Write Book")
}

func test7() {
	// 对应的 pair 是：pair<type: Book, value: book{} 对应的地址}
	b := &Book{}

	var r Reader
	r = b

	r.ReadBook()

	var w Writer

	// 这里为什么可以成功断言：因为 w 和 r 本质上的 pair 是一致的
	w = r.(Writer)
	w.WriteBook()
}
