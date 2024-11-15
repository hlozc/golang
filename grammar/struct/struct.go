package main

import "fmt"

/*
* @TEST 1 声明新的数据类型
type newint int

* @TEST 2 定义一个结构体
type Book struct {
	id     int
	title  string
	author string
}
* 对象的传递仍然是值传递

*/

type newint int

type Book struct {
	id     int
	title  string
	author string
}

func visitBook(book *Book) {
	(*book).author = "new author"
}

func test1() {
	var i newint
	fmt.Println(i)

	var b1 Book
	b1.title = "golang"
	b1.author = "book author"
	fmt.Println(b1)

	visitBook(&b1)
	fmt.Println(b1)
}
