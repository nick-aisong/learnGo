/*
Q24. (1) 指针和反射
1. 在第 “自省和反射” 节，第 76 页的最后一段中，有这样的描述：
	右边的代码没有问题，并且设置了成员变量 Name 为 “Albert Einstein”。
	当然，这仅仅工作于调用 Set() 时传递一个指针参数。
为什么是这样的情况？
######################################################
Listing 5.8. 私有成员的反射
type Person struct {
	name string // 名称
	age int
}
func Set(i interface {}) {
	switch i.(type) {
	case *Person:
		r := reflect.ValueOf(i)
		r.Elem(0).Field(0).SetString("Albert Einstein")
	}
}

Listing 5.9. 公有成员的反射
type Person struct {
	Name string // Name
	age int
}
func Set(i interfac e { }) {
	switch i.(type) {
		case *Person:
		r := reflect.ValueOf(i)
		r.Elem().Field(0).SetString("Albert Einstein")
	}
}
左边的代码可以编译并运行，但是当运行的时候，将得到打印了栈的运行时错误：
panic: reflect.Value.SetString using value obtained using unexported field
右边的代码没有问题，并且设置了成员变量 Name 为 “Albert Einstein”。
当然，这仅仅工作于调用 Set() 时传递一个指针参数
######################################################
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}
