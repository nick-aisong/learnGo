/*
Q14. (1) 函数返回一个函数
1. 编写一个函数返回另一个函数，返回的函数的作用是对一个整数 +2
函数的名称叫做 plusTwo
然后可以像下面这样使用：
	p := plusTwo()
	fmt.Printf("%v\n", p(2))
应该打印 4
2. 使 1 中的函数更加通用化，创建一个 plusX(x) 函数，返回一个函数用于对整数加上 x。
*/
package main

import (
	"fmt"
)

func plusTwo() func(int) int {
	return func(x int) int {
		return x + 2
	}
}

// 这里用到了闭包
func plusX(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	p1 := plusTwo()
	p2 := plusX(9)

	fmt.Printf("%v\n", p1(9))
	fmt.Printf("%v\n", p2(9))
}
