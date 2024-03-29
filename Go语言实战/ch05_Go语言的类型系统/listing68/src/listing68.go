// 这个示例程序展示如何访问另一个包的未公开的
// 标识符的值
package main

import (
	"fmt"

	"counters"
)

// main 是应用程序的入口
func main() {
	// 使用 counters 包公开的 New 函数来创建一个未公开的类型的变量
	counter := counters.New(10)

	fmt.Printf("Counter: %d\n", counter)
}

// G:\GitHub\learnGo\Go语言实战\ch05_Go语言的类型系统\listing68\src>go run listing68.go
// Counter: 10
