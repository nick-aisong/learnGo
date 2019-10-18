/*
Q9. (1) 变参
1. 编写函数接受整数类型变参，并且每行打印一个数字
*/
package main

import (
	"fmt"
)

func prtthem(numbers ...int) {
	for _, d := range numbers {
		fmt.Printf("%d\n", d)
	}
}

func main() {
	prtthem(1, 4, 5, 7, 4)
	prtthem(1, 2, 4)
}
