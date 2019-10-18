/*
Q12. (0) 最小值和最大值
	1. 编写一个函数，找到 int slice ([]int) 中的最大值
	2. 编写一个函数，找到 int slice ([]int) 中的最小值
*/
package main

import (
	"fmt"
)

func max(l []int) (max int) {
	max = l[0]
	for _, v := range l {
		if v > max {
			max = v
		}
	}
	return
}

func min(l []int) (min int) {
	min = l[0]
	for _, v := range l {
		if v < min {
			min = v
		}
	}
	return
}
func main() {
	var slice = []int{1, 43, 46, -4, 64, -8, 100, 2, 11, 99}
	fmt.Printf("%v\n", slice)
	fmt.Printf("max: %d\n", max(slice))
	fmt.Printf("min: %d\n", min(slice))
}
