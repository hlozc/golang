package main

/*
* TEST 1 init 函数和 import 导包
* 1. 进入 main 函数
* 2. 开始导包，如果被导入的包里面依然要导入，那么就持续导入，直到没有新包为止
* 3. 开始初始化各种变量，const, var
* 4. 执行包的 init 函数

* TEST 2 如果要导包
* 那么需要将这个包在 GOPATH 中的相对路径来导入

* TEST 3 大小写规则
* 1. 文件中如果函数名首字母为大写，那么表示对外开放
* 2. 文件中如果函数名首字母为小写，那么表示这个函数只能在包里面使用，不对外开放

* TEST 4 匿名导入
* 有时候可能不想去调用某个包中的函数，但是却需要去调用里面的初始化函数
* 那么就会进行匿名导包

import (
	_ "grammar/init/lib1"		// 通过下划线空格的方式
	p "grammar/init/lib2"		// 也可以起别名
	. "grammar/init/lib3" 		// 也可以通过 . 的方式，这种方式表示展开作用域
)

但是如果起了别名就一定要用
*/

import (
	"fmt"
	_ "grammar/init/lib1"
	lib "grammar/init/lib2"
)

func modify(v *int) {
	*v = 100
}

func main() {
	lib.Lib2Test()

	a := 10
	modify(&a)
	fmt.Println(a)
}
