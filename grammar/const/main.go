package main

import "fmt"

/*
* TEST 1 常量的声明
* 1. 不允许使用 := 进行声明
* 2. 声明为常量之后不可修改

* TEST 2 常量可以用来定义枚举类型
* 类似于多变量声明

const (
	...
)
*

* TEST 3 const 定义枚举的时候，可以添加关键字 iota，第一行默认是 0，后面都会累加
* 并且可以套用公式, 比如 10 * iota, 那么后面就会是 10 * 1, 10 * 2, ... 10 * iota
* 反正最终 iota 从 0 开始递增

const (
	EMPTY = 10 * iota
	FULL
)

* TEST 4 const 甚至可以变换公式, 反正最终 iota 保持递增

const (
	a, b = iota + 1, iota + 2
	c, d

	e, f = iota + 2, iota + 3
)

* TEST 5 多返回值
* 如果一个函数没有返回值，那么：
func f() {
}

* 如果一个函数有多个返回值，那么：
func f() (r1 int, r2 string) {
}
并且这些返回值可以命名, 并且 r1, r2 一开始就是作为一个局部变量来进行初始化
*/

const (
	EMPTY = 10 * iota
	FULL
)

func TwoValue(a int, b int) (r1 int, r2 int) {
	r1 = a + b
	r2 = a - b
	return
}

// 如果函数返回值类型一致，那么可以省略
func ThreeValue(a int, b int) (r1, r2 int) {
	return a + b, a - b
}

func main() {
	const len int = 10
	fmt.Println("len = ", len)
	fmt.Printf("%d, %d", EMPTY, FULL)
}
