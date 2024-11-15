package main

/*
* @ TEST 1 channel And select
* 1. 单流程下的 go 只能监控一个 channel 的状态，select 可以监控多个 channel
 */

import (
	"fmt"
)

func fib(c1, c2 chan int) {
	x, y := 1, 1

	for {
		select {
		// 如果可写
		case c1 <- x:
			x, y = y, x+y
		// 如果可写
		case <-c2:
			fmt.Println("quit")
			return
		}
	}
}

func test2() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c1)
		}

		c2 <- 0
	}()

	fib(c1, c2)
}
