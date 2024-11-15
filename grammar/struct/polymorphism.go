package main

import "fmt"

/*
* TEST 1 接口 interface
* 1. 如果想要继承一个接口，只要把接口的所有方法都实现，就可以了
* 2. 也就可以实现通过 AnimalIF 去接收对象
* 3. 如果这个类没有实现这个接口的所有方法，那么这个 【接口指针】 就无法接收这个具体的类

* TEST 2 实现多态
func Sleep(animal AnimalIF) {
	animal.Sleep()
}

func main() {
	Sleep(&Cat{"yellow"})

	Sleep(&Dog{"black"})
}

* TEST 3 总结
1. 首先定义接口（父类），并声明所有需要的函数
2. 有子类实现了父类的所有接口方法
3. 接口指针引用子类对象
*/

// 本质上是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string // 获取动物的颜色
	GetType() string  // 获取动物的类型
}

type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat.Sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

// -----------

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog.Sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Cat"
}

// animal 接口指针，用来接收不同的子类对象，完成多态
func Sleep(animal AnimalIF) {
	animal.Sleep()
}

func test4() {
	Sleep(&Cat{"yellow"})

	Sleep(&Dog{"black"})
}
