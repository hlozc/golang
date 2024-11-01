/*
* @TEST 1 变量的声明, 总共有四种
* 1. var i int 声明这个变量类型，这时候会赋值默认值
* 2. var i int = 100 , 或者在 1 的基础上直接赋值
* 3. var i = 100, 可以省去变量的类型，通过初始化的值来推断变量的类型
* 4. 省略 var 关键字，通过 := 来初始化，表示初始化并赋值，只能在函数体中使用

* @TEST 2 全局变量的声明
* 全局变量的声明适用于上面三条规则，第四条规则不适用
* var gA int = 200

* @ TEST 3 Golang 支持同时声明多个变量
* 也支持多行多变量声明
* var (
*  ...
* )
 */

package main

import (
	"fmt"
)

func main() {
	// var a int
	// var b int = 100
	var c = 100
	var d = "string"
	e := 3.14

	var x, y = 20, "hello world"
	var (
		xx int    = 20
		yy string = "hello world"
	)

	fmt.Printf("c = %d And type is %T\n", c, c)
	fmt.Println("var d = ", d)
	fmt.Printf("var e = %v, type is %T\n", e, e)
	fmt.Printf("x = %d, y = %d\n", x, y)
	fmt.Printf("xx = %d, yy = %d\n", xx, yy)
}
