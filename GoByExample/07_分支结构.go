// 07_分支结构
package main

import (
	"fmt"
	"time"
)

func main() {

	// 一个基本的 switch
	i := 2
	fmt.Println("write", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// 在一个 case 语句中，你可以使用逗号来分隔多个表达式
	// 在这个例子中，我们很好的使用了可选的default分支
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	// 不带表达式的 switch 是实现 if/else 逻辑的另一种方式
	// 这里展示了 case 表达式是如何使用非常量的
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}
}
