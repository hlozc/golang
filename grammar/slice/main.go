package main

import "fmt"

/*
* @TEST 1 len 用来获取集合的长度，时间复杂度是 O(1)

* @TEST 2 range 通过这种方式遍历，每次遍历可以获取两个值，i 和 value
for range ar {

}

* @TEST 3 静态数组
* [4]int 和 [10]int 是两种不同类型的变量，无法进行传递
* [4] int 这种类型是静态数组，并且这种类型的数据在传递的时候是值拷贝

* @TEST 4 动态数组
* 在声明数组的时候，比如下面这一行，不声明长度，那么这个数组就是动态数组
ar := []int{1, 2, 3}
* 并且传递的时候是引用的方式传递的

* TEST 5 动态数组的追加和截取
numbs := make([]int, 3)
fmt.Printf("len = %d, cap = %d, content = %v\n", len(numbs), cap(numbs), numbs)

numbs = append(numbs, 10)
fmt.Printf("len = %d, cap = %d, content = %v\n", len(numbs), cap(numbs), numbs)
添加元素的时候如果长度不够，会两倍扩容

* TEST 6 切片的截取
numbs2 := numbs[1:3]
fmt.Printf("len = %d, cap = %d, content = %v\n", len(numbs2), cap(numbs2), numbs2)
这里截取之后实际上 numbs2 和 numbs 的内存空间是一致的

要想实现不同内存区域，可以使用 copy，他会完成深拷贝
numbs3 := make([]int, 2)
copy(numbs3, numbs2)
fmt.Printf("len = %d, cap = %d, content = %v\n", len(numbs3), cap(numbs3), numbs3)

* TEST 7 append 就尽量在原地修改，就算地址变化，也会返回一个新的有效地址
删除元素
ori = append(ori[0:i], ori[i + 1:] ...)
... 表示展开，将切片里面的所有元素展开成为独立的参数
*/

func modifyArray(ar [10]int) {
	ar[1] = 100
}

func modifySlice(ar []int) {
	ar[0] = 100
}

func main() {
	var ar [10]float32
	ar1 := [10]int{1, 2, 3}
	modifyArray(ar1)

	for i := 0; i < len(ar); i++ {
		fmt.Println(i, ar[i])
	}

	for i, v := range ar1 {
		fmt.Println(i, v)
	}

	// 查看数组数据类型
	fmt.Printf("Array Type is %T\n", ar)
	fmt.Printf("Array Type is %T\n", ar1)

	ar2 := []int{1, 2, 3}
	for i, v := range ar2 {
		fmt.Println(i, v)
	}
	fmt.Printf("Array Type is %T\n", ar2)

	modifySlice(ar2)

	for i, v := range ar2 {
		fmt.Println(i, v)
	}

	// 切片部分
	fmt.Println(" ------------------------------- ")
	slice1 := []int{1, 2, 3} // 声明一个切片, 并初始化

	var slice2 []int // 声明了一个切片，但是没有初始化，没有空间
	slice2 = make([]int, 5)

	fmt.Printf("len = %d, content = %v\n", len(slice1), slice1)

	fmt.Printf("len = %d, content = %v\n", len(slice2), slice2)

	// 切片容量的追加和截取
	numbs := []int{9, 99, 999}
	fmt.Printf("len = %d, cap = %d, content = %v\n", len(numbs), cap(numbs), numbs)

	numbs = append(numbs, 10)
	fmt.Printf("len = %d, cap = %d, content = %v\n", len(numbs), cap(numbs), numbs)

	// 左闭右开，表示 [left, right) 下标的所有元素
	numbs2 := numbs[1:3]
	fmt.Printf("len = %d, cap = %d, content = %v\n", len(numbs2), cap(numbs2), numbs2)

	numbs3 := make([]int, 2)
	copy(numbs3, numbs2)
	fmt.Printf("len = %d, cap = %d, content = %v\n", len(numbs3), cap(numbs3), numbs3)
}
