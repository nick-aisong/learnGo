/*
Q6. (0) 整数顺序
1. 编写函数，返回其（两个）参数正确的（自然）数字顺序：
	f(7,2) → 2,7
	f(2,7) → 2,7
*/
package main

import (
	"fmt"
)

func f1(i, j int) (m, n int) {
	if i < j {
		m, n = i, j
	} else {
		m, n = j, i
	}
	return
}

func order(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func main() {
	x, y := f1(7, 2)
	fmt.Printf("%d, %d\n", x, y)
	fmt.Printf("%d, %d\n", x, y)
}
