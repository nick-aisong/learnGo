//  5-11 listing11
// 这个示例程序展示如何声明
// 并使用方法
package main

import (
	"fmt"
)

// user 在程序里定义一个用户类型
type user struct {
	name  string
	email string
}

// notify 使用值接收者实现了一个方法
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// changeEmail 使用指针接收者实现了一个方法
func (u *user) changeEmail(email string) {
	u.email = email
}

// main 是应用程序的入口
func main() {
	// user 类型的值可以用来调用
	// 使用值接收者声明的方法
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	// 指向 user 类型值的指针也可以用来调用
	// 使用值接收者声明的方法
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	// user 类型的值可以用来调用
	// 使用指针接收者声明的方法
	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	// 指向 user 类型值的指针可以用来调用
	// 使用指针接收者声明的方法
	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}

// 值接收者使用值的副本来调用方法，而指针接受者使用实际值来调用方法

// G:\GitHub\learnGo\Go语言实战\ch05_Go语言的类型系统>go run listing11.go
// Sending User Email To Bill<bill@email.com>
// Sending User Email To Lisa<lisa@email.com>
// Sending User Email To Bill<bill@newdomain.com>
// Sending User Email To Lisa<lisa@newdomain.com>
