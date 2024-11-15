/*
* TEST 1 哈希表的声明
1. var hashmap map[int]string

2. hashmap2 := make(map[int]string)

3. hashmap3 := map[int]string {
		1: "php",
		2: "public",
	}

* TEST 2 哈希表的传递
* 这里哈希表的传递和 slice 一样，都是传递的引用
*/

package main

import "fmt"

func main() {
	// 第一部分
	var hashmap map[int]string
	if hashmap == nil {
		fmt.Println("hashmap is nil")
	}

	// 使用 map 前需要先分配空间
	hashmap = make(map[int]string, 10)
	hashmap[1] = "c++"
	fmt.Println(hashmap)

	// 也可以这样声明, 不定义大小
	hashmap2 := make(map[int]string)
	hashmap2[1] = "python"
	fmt.Println(hashmap2)

	// 第三种方式
	hashmap3 := map[int]string{
		1: "php",
		2: "public",
	}
	fmt.Println(hashmap3)

	// 第二部分, 遍历
	cityMap := map[string]int{
		"beijing":  1,
		"shenzhen": 2,
		"hangzhou": 3,
	}

	for key, value := range cityMap {
		fmt.Println(key, value)
	}

	// 第三，删除
	delete(cityMap, "shenzhen")

	// 第四，修改
	cityMap["beijing"] = 0

	fmt.Println(" --------- SEP -----------")
	for key, value := range cityMap {
		fmt.Println(key, value)
	}
}
