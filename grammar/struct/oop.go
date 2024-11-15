package main

import "fmt"

/*
* @TEST 1 面向对象
* 给对象定义其函数, (this Hero) 表示这个函数是属于哪个结构体的
func (this Hero) GetName() {
	fmt.Println("Name = ", this.name)
}

* TEST 2 声明和初始化一个自定义对象
hero := Hero{name: "zhangsan", level: 20}

* TEST 3 对象的函数中的 this，也是值拷贝，所以要修改或者访问对象，那么就需要指针

* TEST 4 一个类
* 1. 如果类名首字母大写，那么其他类可以访问
* 2. 如果属性首字母大写，那么表示公开
* 3. 如果属性首字母小写，那么表示私有
*/

type Hero struct {
	name  string
	level int
}

func (this Hero) Show() {
	fmt.Printf("{\n  Hero: %s\n  Level: %d\n}\n", this.name, this.level)
}

// 这里的 (this Hero) 表示这个函数绑定到了 Hero 这个结构体
func (this Hero) GetName() string {
	return this.name
}

func (this *Hero) SetName(name string) {
	this.name = name
}

func test2() {
	hero := Hero{name: "zhangsan", level: 20}
	hero.Show()

	hero.SetName("ppp")
	hero.Show()
}
