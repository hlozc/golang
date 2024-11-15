package main

import (
	"fmt"
	"time"
)

/*
* @ TEST 1 channel
* 用来让 goroutine 中的不同协程进行通信
* 1. channel <- value 表示发送 value 到 channel 中
* 2. <- channel 接收管道数据并丢弃
* 3. buf := <-channel 表示从 channel 中读取数据并拿出来
* 4. buf, ok := <-channel 功能和第三点差不多，但是会检查管道是否已经关闭或者是否为空

* @ TEST 2 无缓冲区的例子
* 1. 如果主线程先执行，那么尝试读取 channel 中的数据的时候就会读不到数据而阻塞，直到协程写入数据到管道
* 2. 但是如果 协程 先运行，那么由于是 【无缓冲区】的，所以这里的协程也会进行阻塞，等待对方读取
*    所以【无缓冲区】的 channel，必须要双方都阻塞到 channel，直到两边完成对接

* @ TEST 3 channel 有缓冲区和无缓冲区的问题
* 1. 无缓冲区：谁先到谁阻塞，直到“两边都把手伸进来”，完成对接，双方释放
* 2. 有缓冲区：生产者消费者队列, 和命名管道和匿名管道差不多，满的时候和空的时候都有同步机制

* @ TEST 4 channel 关闭

data, ok := <-c

* 1. ok 为 true, 说明 channel 没有关闭，ok 为 false 说明 channel 顺利关闭
* 2. close 可以主动关闭管道
*/

func test9() {
	// 定义一个无缓存的管道，里面用来存储一个 int 数据
	c := make(chan int)

	// 从 channel 中读取数据并接收
	nb := <-c
	fmt.Println("read channel data: ", nb)

	go func() {
		defer fmt.Println("goroutine exit")

		fmt.Println("goroutine running")

		// 将 666 写入到 channel 中
		c <- 666
	}()
}

func test10() {
	// 主线程写数据，routine 读取数据
	// 表示带有缓冲的 channel，大小为 3
	c := make(chan int, 3)
	fmt.Println("channel len: ", len(c), " and cap: ", cap(c))

	go func() {
		defer fmt.Println("routine exit")
		time.Sleep(10 * time.Second)
		sep := 0
		var nb int
		var isopen bool

		for {
			if nb, isopen = <-c; !isopen {
				break
			}
			fmt.Println("[", sep, "] routine read data: ", nb)
			sep++
		}
	}()

	nb := 100
	sep := 0
	for {
		fmt.Println("[", sep, "] main write data: ", nb)
		c <- nb
		nb++
		sep++

		time.Sleep(1 * time.Second)
		if nb == 110 {
			close(c)
		}
	}
}
