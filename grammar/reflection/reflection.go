package main

/*
* @TEST 1 反射

func ValueOf(i interface{}) Value {...}
- ValueOf 用来获取输入参数接口中的数据的值，如果接口为空返回 0

func TypeOf(i interface{}) Type {...}
- TypeOf 用来动态获取输入参数接口中的值的类型，如果接口为空，返回 nil

*/

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
	Sex  string
}

func (this User) Show() {
	fmt.Printf("User Info : \n{  name = %v\n  age = %d\n  sex = %v\n}", this.Name, this.Age, this.Sex)
}

func reflectNum(i interface{}) {
	argType := reflect.TypeOf(i)
	fmt.Println("Type is ", argType.Name())

	argValue := reflect.ValueOf(i)
	fmt.Println("Value is ", argValue)

	// 获取里面的字段
	// 1. 通过 interface 的 reflect.Type, 通过 Type 获取到 NumField 进行遍历
	// 2. 然后就可以得到每一个 field, 即数据类型
	// 3. 通过 field 有一个 Interface() 方法得到对应的 value
	for i := 0; i < argType.NumField(); i++ {
		field := argType.Field(i)              // 获取第 i 个字段
		value := argValue.Field(i).Interface() // 通过类型来获取到对应的变量

		fmt.Printf("[Type: %v] Field(%s) = %v\n", field.Type, field.Name, value)
	}

	// 调用里面的函数
	// 1. 并且 NumMethod 这个函数会返回 [导出] 并且 [值接收的函数]
	// 2. 如果是 [导出] 并且 [指针接收 即 *User] 的函数，那么需要使用 PtrTo(argType) 把类型转成指针
	fmt.Println("User Method Num is ", argType.NumMethod())
	for i := 0; i < argType.NumMethod(); i++ {
		method := argType.Method(i)
		fmt.Printf("[Function: %s] %v", method.Name, method.Type)
	}
}

func test8() {
	user := User{"Liu", 20, "female"}
	reflectNum(user)
}
