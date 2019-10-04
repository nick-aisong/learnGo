// arrayDemo3
package main

import (
	"fmt"
)

// 4-15 使用指针在函数间传递大数组
func main() {
	// 分配一个需要 8 MB 的数组
	var array [1e6]int
	// 将数组的地址传递给函数 foo
	foo(&array)
}

// 函数 foo 接受一个指向 100 万个整型值的数组的指针
func foo(array *[1e6]int) {
	//...
}

// //  4-14 使用值传递，在函数间传递大数组
// func main() {
// 	// 声明一个需要 8 MB 的数组
// 	var array [1e6]int
// 	// 将数组传递给函数 foo
// 	foo(array) // 函数 foo 接受一个 100 万个整型值的数组
// }

// func foo(array [1e6]int) {
// 	//...
// }
