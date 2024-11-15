package main

import "fmt"

/*
* @TEST 1 空接口
1. interface {} 即表示空接口, 或者通用万能类型
2. int, string, float32, float64, struct ... 这些都实现了 interface{}
3. 于是就可以通过 interface 来引用任意的数据类型

* @TEST 2 interface 如何区分这个数据底层是什么数据类型, 又如何根据数据类型的不同来决定做不同的事
* 1. Golang 给 interface{} 提供了 “类型断言” 的机制
* 2. value, ok := arg.(string) 这种语法表示这个 变量 底层类型是不是 string 类型的
* 3. 如果是这个类型，ok 为 true
*/

func interfaceTest(arg interface{}) {
	fmt.Println("interfaceTest called")

	value, ok := arg.(string)
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Not a string")
	}
}

type Books struct {
	author string
}

func main() {
	b := Books{"me"}
	interfaceTest(b)
	interfaceTest("string")
}
