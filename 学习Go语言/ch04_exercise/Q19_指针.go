/*
Q19. (1) 指针
1. 假设定义了下面的结构：
type Person struct {
	name string
	age int
}
下面两行之间的区别是什么？
var p1 Person
p2 := new(Person)
2. 下面两个内存分配的区别是什么？
func Set(t *T) {
	x = t
}
和
func Set(t T) {
	x= &t
}
*/
/*
A19. (1) 指针
1. 第一行：var p1 Person 分配了 Person-值给 p1。p1 的类型是 Person。
第二行：p2 := new(Person) 分配了内存并且将指针赋值给 p2。p2 的类型是*Person。

2. 在第二个函数中，x 指向一个新的（堆上分配的）变量 t，其包含了实际参数值的副本。
在第一个函数中，x 指向了 t 指向的内容，也就是实际上的参数指向的内容。
因此在第二个函数，我们有了 “额外” 的变量存储了相关值的副本。
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}
