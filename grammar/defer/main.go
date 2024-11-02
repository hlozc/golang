package main

/*
* TEST 1 defer 关键字
* 后面需要跟一个具体的表达式, 有点类似于 finally
* 一个函数或者一个流程结束之前要触发的行为
* 但是如果存在多个 defer，那么类似于栈，先 defer 的会后执行

* TEST 2 defer 和 return 谁先执行
* return 会比 defer 先调用
 */

import "fmt"

func main() {
	defer fmt.Println("public void main")
	defer fmt.Println("public void main2")

	fmt.Println("hello world")
}
