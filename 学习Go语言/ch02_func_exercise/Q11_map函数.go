/*
Q11. (1) map 函数 map() 函数是一个接受一个函数和一个列表作为参数的函数
函数应用于列表中的每个元素，而一个新的包含有计算结果的列表被返回
因此: map(f(),(a1, a2, . . . , an−1, an)) = (f(a1), f(a2), . . . , f(an−1), f(an))

*/
package main

import (
	"fmt"
)

func Map(f func(int) int, l []int) []int {
	j := make([]int, len(l))
	for k, v := range l {
		j[k] = f(v)
	}
	return j
}

func main() {
	m := []int{1, 3, 4}
	f := func(i int) int {
		return i * i
	}
	fmt.Printf("%v", (Map(f, m)))
}
