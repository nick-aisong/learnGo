/*
Q23. (1) 接口和编译
1. 在第 72 页的代码 5.3 编译正常——就像文中开始描述的那样。但是当运行的时
候，会得到运行时错误，因此有些东西有错误。为什么代码编译没有问题呢？

######################################################
type R struct {i int}
func (p *R) Get() int {return p.i}
func (p *R) Put(v int) {p.i = v}

func g(any interface {}) int {
	return any.(I).Get()
}

	i := 5 // 声明 i 是一个 int
fmt.Println(g(i))
这能编译，但是当运行的时候会得到:
panic: interface conversion: int is not main.I: missing method Get
这是绝对没问题，内建类型 int 没有 Get() 方法
######################################################
*/
/*
1. 代码能够编译是因为整数类型实现了空接口，这是在编译时检查的。
修复这个正确的途径是测试这个空接口可以被转换，如果可以，调用对应的方
法。5.2 列出的 Go 代码中定义了函数 g——这里重复一下：
func g(any interface{}) int {return any.(I).Get()}
应当修改为：
func g(any interface {}) int {
	if v, ok := any.(I) ; ok { // 检查是否可以转换
	return v.Get() // 如果可以，调用 Get()
}
	return -1 // 随便返回个什么
}
如果现在调用 g()，就不会有运行时错误了。在 Go 中这种用法被称作 “comma ok”
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}
