/*
1. 创建一个固定大小保存整数的栈。它无须超出限制的增长。定义 push 函数——
将数据放入栈，和 pop 函数——从栈中取得内容。栈应当是后进先出（LIFO）
的
2. 更进一步。编写一个 String 方法将栈转化为字符串形式的表达。可以这样的
方式打印整个栈：fmt.Printf("My stack %v\n", stack)
栈可以被输出成这样的形式：[0:m] [1:l] [2:k]
*/
package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(e int) {
	s.data[s.i] = e
	s.i++
}

func (s *stack) pop() (e int) {
	s.i--
	return s.data[s.i]
}

func (s stack) String() string {
	var str string
	for i := 0; i < s.i; i++ {
		str = str + "[" + strconv.Itoa(i) +
			":" + strconv.Itoa(s.data[i]) +
			"] "
	}
	return str
}

func main() {
	var s stack
	s.push(25)
	s.push(14)
	s.push(1)
	s.pop()
	fmt.Printf("stack %v\n", s)
}
