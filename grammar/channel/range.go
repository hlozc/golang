package main

/*
* @ TEST 1 channel And range
* 1. range channel 会阻塞地等待 channel 的结果
 */

import (
	"fmt"
)

func test1() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Write Data ", i)
			c <- i
		}

		close(c)
	}()

	for data := range c {
		fmt.Println("Read Data ", data)
	}

	fmt.Println("Main Exit")

}
