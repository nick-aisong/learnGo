// 这个示例程序展示无法从另一个包里
// 访问未公开的标识符
package main

import (
	"fmt"

	"counters"
)

// main 是应用程序的入口
func main() {
	// 创建一个未公开的类型的变量
	// 并将其初始化为 10
	counter := counters.alertCounter(10)

	// ./listing64.go:15: 不能引用未公开的名字 counters.alertCounter
	// ./listing64.go:15: 未定义：counters.alertCounter

	fmt.Printf("Counter: %d\n", counter)
}

// G:\GitHub\learnGo\Go语言实战\ch05_Go语言的类型系统\listing64\src>go run listing64.go
// # command-line-arguments
// .\listing64.go:15: cannot refer to unexported name counters.alertCounter
// .\listing64.go:15: undefined: counters.alertCounter
