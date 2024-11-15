package main

/*
* @TEST 1 struct tag 在 json 中的运用
* 1. 就可以通过这里的注释来修改 json 中对应的字段名
* 2. 也就是：【自定义序列化之后的名称】
 */

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string 	`json:"title"`
	Time int 		`json:"time"`
	Price int 		`json:"100"`
	Actors []string 
}


func main()  {
	// 1. 将结构体转化成字符串
	m := Movie{"喜剧之王", 2000, 10, []string{"大帅逼", "周星驰"}}

	// 将这个结构体转成 json 格式
	movieJson, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Json Marshal Fail")
	}
	fmt.Printf("Json String is %s\n", movieJson)

	// 2. 将字符串重新转为结构体
	var m1 Movie
	err = json.Unmarshal(movieJson, &m1)
	if err != nil {
		fmt.Println("Unmarshal Error")
	}
	fmt.Println(m1)
}
