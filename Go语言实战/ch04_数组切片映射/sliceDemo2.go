// sliceDemo2
package main

import (
	"fmt"
)

// 4-44 在函数间传递切片
func main() {
	// 分配包含 100 万个整型值的切片
	slice := make([]int, 1e6)
	// 将 slice 传递到函数 foo
	slice = foo(slice)

}

// 函数 foo 接收一个整型切片，并返回这个切片（值传递）
func foo(slice []int) []int {
	//...
	return slice
}

//在 64 位架构的机器上，一个切片需要 24 字节的内存：
//指针字段需要 8 字节
//长度和容量字段分别需要 8 字节
//复制时只会复制切片本身，不会涉及底层数组
