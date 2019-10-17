// 这个示例程序展示公开的结构类型中如何访问
// 未公开的内嵌类型的例子
package main

import (
	"fmt"

	"entities"
)

// main 是应用程序的入口
func main() {
	// 创建 entities 包中的 Admin 类型的值
	a := entities.Admin{
		Rights: 10,
	}

	// 设置未公开的内部类型的公开字段的值
	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("User: %v\n", a)
}

// G:\GitHub\learnGo\Go语言实战\ch05_Go语言的类型系统\listing74\src>go run listing74.go
// User: {{Bill bill@email.com} 10}
