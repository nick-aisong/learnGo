/*
Q10. (1) 斐波那契
1. 斐波那契数列以：1, 1, 2, 3, 5, 8, 13, . . . 开始
或者用数学形式表达：x1 = 1; x2 =1; xn = xn−1 + xn−2 ∀n > 2
编写一个接受 int 值的函数，并给出这个值得到的斐波那契数列
*/
package main

import (
	"fmt"
)

func fibonacci(value int) []int {
	x := make([]int, value)
	x[0], x[1] = 1, 1
	for n := 2; n < value; n++ {
		x[n] = x[n-1] + x[n-2]
	}
	return x
}

func main() {
	for _, term := range fibonacci(10) {
		fmt.Printf("%v ", term)
	}
}
