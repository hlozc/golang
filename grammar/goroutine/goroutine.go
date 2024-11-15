package main

/*
* @ TEST 1 Golang 早期的 Goroutine 调度器非常低效
* 1. 存在一个 全局的 go 协程队列，存放所有的协程
* 2. 然后存在多个不同的线程，这些线程，获取锁之后就会将获取到的 goroutine 进行执行，执行完成之后再还锁

* @ TEST 2 Golang 早期的 Goroutine 调度器优缺点
* 1. 每次创建、获取、销毁协程，都需要获取锁
* 2. 线程在转移协程的时候，执行的过程会造成低效的 CPU 开销
* 3. 比如一个线程在使用某个协程的时候，这个协程又创建了一个协程
* 4. 那么按理来说，既然创建了，那么大概率马上就会用到这个协程 

* @ TEST 3 GMP
* G：goroutine
* P: processor 处理器, 用来管理 goroutine 的
* M: thread
* 参考文章：https://juejin.cn/post/7095300413376725006

* @ TEST 4 调度器的涉及策略
* @ 复用线程
* 1. working stealing
* 	 全局队列主要用来存放一些空闲的 G，如果创建了某个 G，那么这个 G 优先放在本地的 G 队列中
*    如果满了，那么才会放在全局中。
*    
*    如果某个线程的 P 处于空闲状态，那么就会尝试其他队列中取协程

* 2. hand off 
*  	 如果现在 M1 在执行某个 G，但是这个 G 阻塞了，这时候就会创建或者唤醒一个 M2
*    然后让 M2 来接手 M1 原来的 P（也就是逻辑 CPU）  
*    也就是让 M2 接收 M1 的 P 和 G，然后 M1 继续进行阻塞
*	 接着 M1 可能就会进行睡眠或者销毁 

* @ 利用并行
* GOMAXPROCS 用来限定 P 的个数，这个 GOMAXPROCS <= CPU 核心

* @ 抢占
* 传统的 co-routine 除非协程主动释放，否则这个协程不会放弃 CPU 使用权
* 但是 goroutine 会限定一个 协程 的最大使用时间，超过就会被其他抢占，并且优先级一般相等

* @ 全局队列
* 线程的 P 会优先 working steal 其他线程的本地协程，偷完了就会去全局队列里面取

* @ TEST 5 协程使用
* 1. 通过 go 关键词来启动一个协程
* 2. go 中的协程是异步的方式来运行的，并且是 go 运行时管理的轻量级，并且这个函数就会在后台进行
* 3. 启动一个 goroutine 之后，然后这个 goroutine 会在后台执行，goroutine 之间通过 channel 通信

* @ TEST 6 协程退出
* 1. 如果想让协程退出，那么需要调用

runtime.Goexit()
*/

import (
	"fmt"
	"runtime"
	"time"
)

func task() {
	i := 0
	for {
		i ++
		fmt.Println("Goroutine Num: ", i)
		time.Sleep(1 * time.Second)
	}
}

func main()  {
	// 创建一个协程，去执行 task() 函数
	// go task() 

	go func() {
		defer fmt.Println("out defer")

		func() {
			defer fmt.Println("in defer")
			runtime.Goexit()
			fmt.Println("in")
		}()

		fmt.Println("out")
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}