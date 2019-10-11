// hellochannels
// 这个程序在书里没列出，GitHub源码库里有
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printer(ch chan int) {
	for i := range ch {
		fmt.Printf("Received %d ", i)
	}
	wg.Done()
}

// main is the entry point for the program.
func main() {
	c := make(chan int)
	go printer(c)
	wg.Add(1)

	// Send 10 integers on the channel.
	for i := 1; i <= 10; i++ {
		c <- i
	}

	close(c)
	wg.Wait()
}

// G:\GitHub\learnGo\Go语言实战\ch01_关于Go语言的介绍\channels>go run hellochannels.go
// Received 1 Received 2 Received 3 Received 4 Received 5 Received 6 Received 7 Received 8 Received 9 Received 10
