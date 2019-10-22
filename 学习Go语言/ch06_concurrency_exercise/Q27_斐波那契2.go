/*
Q27. (2) 斐波那契 II
1. 这是类似的练习，第一个在第 34 页的练习 10。完整的问题描述：
斐波那契数列以： 1; 1; 2; 3; 5; 8; 13; : : : 开头。或用数学形式：
x1 = 1; x2 = 1; xn = xn−1 + xn−2 8n > 2。
编写一个函数接收 int 值，并给出同样数量的斐波那契数列。
但是现在有额外条件：必须使用 channel。
*/
package main

import (
	"fmt"
)

func dup3(in <-chan int) (<-chan int, <-chan int, <-chan int) {
	a, b, c := make(chan int, 2), make(chan int, 2), make(chan int, 2)
	go func() {
		for {
			x := <-in
			a <- x
			b <- x
			c <- x
		}
	}()
	return a, b, c
}

func fib() <-chan int {
	x := make(chan int, 2)
	a, b, out := dup3(x)
	go func() {
		x <- 0
		x <- 1
		<-a
		for {
			x <- <-a + <-b
		}
	}()
	return out
}

func main() {
	x := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(<-x)
	}
}
