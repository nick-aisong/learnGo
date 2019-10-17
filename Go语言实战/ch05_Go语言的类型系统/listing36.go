// 5-36 listing36.go
// 这个示例程序展示 Go 语言里如何使用接口
package main

import (
	"fmt"
)

// notifier 是一个定义了
// 通知类行为的接口
type notifier interface {
	notify()
}

// user 在程序里定义一个用户类型
type user struct {
	name  string
	email string
}

// notify 是使用指针接收者实现的方法
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// main 是应用程序的入口
func main() {
	// 创建一个 user 类型的值，并发送通知
	u := user{"Bill", "bill@email.com"}

	sendNotification(u) // 需要改成&u才能通过编译

	// ./listing36.go:32: 不能将 u（类型是 user）作为
	// sendNotification 的参数类型 notifier：
	// user 类型并没有实现 notifier
	// （notify 方法使用指针接收者声明）
}

// sendNotification 接受一个实现了 notifier 接口的值
// 并发送通知
func sendNotification(n notifier) {
	n.notify()
}

// G:\GitHub\learnGo\Go语言实战\ch05_Go语言的类型系统>go run listing36.go
// # command-line-arguments
// .\listing36.go:33: cannot use u (type user) as type notifier in argument to sendNotification:
//         user does not implement notifier (notify method has pointer receiver)
