// 35_速率限制
// 速率限制(英) 是一个重要的控制服务资源利用和质量的途径
// Go 通过 Go 协程、通道和打点器优美的支持了速率限制
package main

import (
	"fmt"
	"time"
)

func main() {

	// 首先我们将看一下基本的速率限制
	// 假设我们想限制我们接收请求的处理，我们将这些请求发送给一个相同的通道
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// 这个 limiter 通道将每 200ms 接收一个值
	// 这个是速率限制任务中的管理器
	limiter := time.Tick(time.Microsecond * 200)

	// 通过在每次请求前阻塞 limiter 通道的一个接收，我们限制自己每 200ms 执行一次请求
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 有时候我们想临时进行速率限制，并且不影响整体的速率控制我们可以通过通道缓冲来实现
	// 这个 burstyLimiter 通道用来进行 3 次临时的脉冲型速率限制
	burstyLimiter := make(chan time.Time, 3)

	// 想将通道填充需要临时改变3次的值，做好准备
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 每 200 ms 我们将添加一个新的值到 burstyLimiter中，直到达到 3 个的限制
	go func() {
		for t := range time.Tick(time.Microsecond * 200) {
			burstyLimiter <- t
		}
	}()

	// 现在模拟超过 5 个的接入请求
	// 它们中刚开始的 3 个将由于受 burstyLimiter 的“脉冲”影响
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

// 运行程序，我们看到第一批请求意料之中的大约每 200ms 处理一次

// 第二批请求，我们直接连续处理了 3 次，这是由于这个“脉冲”速率控制，然后大约每 200ms 处理其余的 2 个
