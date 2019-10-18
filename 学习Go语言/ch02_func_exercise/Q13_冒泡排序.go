/*
Q13. (1) 冒泡排序
1. 编写一个针对 int 类型的 slice 冒泡排序的函数
	它在一个列表上重复步骤来排序，比较每个相邻的元素，并且顺序错
	误的时候，交换它们。一遍一遍扫描列表，直到没有交换为止，这意
	味着列表排序完成。算法得名于更小的元素就像 "泡泡" 一样冒到列表的別端
*/
package main

import (
	"fmt"
)

func bubblesort(n []int) {
	var swapped bool = true
	for i := 0; i < len(n)-1 && swapped; i++ {
		swapped = false
		for j := i + 1; j < len(n); j++ {
			if n[j] < n[i] {
				n[i], n[j] = n[j], n[i]
				swapped = true
			}
		}
		//fmt.Printf("%v\n", n)
	}
}

func main() {
	var slice = []int{11, 433, 132, 54, 6, 21, 424, 3246, 7, 564}
	fmt.Printf("Before: %v\n", slice)
	bubblesort(slice)
	fmt.Printf("After:  %v\n", slice)
}
