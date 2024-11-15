package main

import (
	"fmt"
	"reflect"
)

/*
* @TEST 1 结构体标签
* 1. 语法就是 `` 的格式，并且需要是键值对的格式
* 2. 表示给这个结构体的该字段绑定一个 Key Value
* 3. 可以写多个, 用空格隔开
* 4. 有点相当于注释


* @TEST 2 Elem() 函数
* * 1. 作用：获取当前 *指针* 所指的结构体类型, 所以传递的时候需要传递指针
 */

type resume struct {
	Name string `info:"name" doc:"Liu"`
	Sex  string `info:"sex"`
}

func showTag(i interface{}) {
	// ELem 会返回当前类型的所有成员类型
	e := reflect.TypeOf(i).Elem()

	for i := 0; i < e.NumField(); i++ {
		// 获取每一个字段，并根据每一个字段中的关键词搜索出对应的 Value
		info := e.Field(i).Tag.Get("info")
		fmt.Println("info: ", info)

		doc := e.Field(i).Tag.Get("doc")
		if len(doc) != 0 {
			fmt.Println("doc: ", doc)
		}
	}
}

func test9() {
	var r resume = resume{"Liu", "Male"}
	// 这里需要的是地址
	showTag(&r)
}
