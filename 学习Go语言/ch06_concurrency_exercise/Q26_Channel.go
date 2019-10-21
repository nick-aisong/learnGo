/*
Q26. (1) Channel
1. 修改在练习 Q1 中创建的程序，换句话说，主体中调用的函数现在是一个
goroutine 并且使用 channel 通讯。不用担心 goroutine 是如何停止的。

2. 在完成了问题 1 后，仍有一些待解决的问题。其中一个麻烦是 goroutine 在
main.main() 结束的时候，没有进行清理。更糟的是，由于 main.main() 和
main.shower() 的竞争关系，不是所有数字都被打印了。本应该打印到 9，但
是有时只打印到 8。添加第二个退出 channel，可以解决这两个问题。
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}
