/*
Q1. (0) For-loop
1. 创建一个基于 for 的简单的循环。使其循环 10 次，并且使用 fmt 包打印出计数器的值
2. 用 goto 改写 1 的循环。关键字 for 不可使用
3. 再次改写这个循环，使其遍历一个 array，并将这个 array 打印到屏幕上
*/

package main

import (
	"fmt"
)

func For_loop1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func For_loop2() {
	i := 0

Loop:
	println(i)
	if i < 10 {
		i++
		goto Loop
	}
}

func For_loop3() {
	array := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(array)
}

func main() {
	For_loop1()
	For_loop2()
	For_loop3()
}
