// pipeline
package main

import (
	"fmt"
)

type PipeData struct {
	value   int
	handler func(int) int
	next    chan int
}

func handle(queue chan *PipeData) {
	for data := range queue {
		data.next <- data.handler(data.value)
	}
}

func main() {

	var ch chan int
	var queue chan *PipeData
	var p PipeData

	p.value = 1
	p.handler = func(i int) int {
		return i + 1
	}

	p.next = ch

	queue <- &p

	handle(queue)

	i := <-ch

	fmt.Print("%d", i)
}

// 会死锁
