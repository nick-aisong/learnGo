// 5-46 listing46.go
// 这个示例程序展示不是总能获取值的地址
package main

import "fmt"

// duration 是一个基于 int 类型的类型
type duration int

// 使用更可读的方式格式化 duration 值
func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}

// main 是应用程序的入口
func main() {
	duration(42).pretty()

	// ./listing46.go:17: 不能通过指针调用 duration(42)的方法
	// ./listing46.go:17: 不能获取 duration(42)的地址
}

// 如果传值，只用给值接收方法用
// 如果传指针，可以给值接收方法和指针接收方法用
// Values 		Methods Receivers
// -----------------------------------------------
//  T 			(t T)
// *T 			(t T) and (t *T)

// 值接收方法可以接收值和指针
// 指针接收方法只可以接收指针
//  Methods 	Receivers Values
// -----------------------------------------------
// (t T) 		T and *T
// (t *T) 		*T

// G:\GitHub\learnGo\Go语言实战\ch05_Go语言的类型系统>go run listing46.go
// # command-line-arguments
// .\listing46.go:17: cannot call pointer method on duration(42)
// .\listing46.go:17: cannot take the address of duration(42)
