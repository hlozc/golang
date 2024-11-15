package main

import "fmt"

/*
* TEST 1 继承
 */

type Human struct {
	name string
	age  int
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk")
}

type SuperMan struct {
	Human // 表示 SuperMan 继承了 Human 类的方法

	level int
}

func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly")
}

func test3() {
	human := Human{name: "zhangsan", age: 1}
	human.Walk()
	human.Eat()

	superman := SuperMan{Human: Human{name: "super", age: 21}, level: 20}
	superman.Eat()
	superman.Fly()
}
