/*
Q34. (1) *Finger 守护进程
1. 编写一个 finger 守护进程，可以工作于 finger(1) 命令。
来自 Debian 的包描述：
	Fingerd 是一个基于 RFC 1196 [28] 的简单的守护进程，它为许多站点
	提供了 “finger” 程序的接口。这个程序支持返回一个友好的、面向用
	户的系统或用户当前状况的详细报告。
最基本的只需要支持用户名参数。如果用户有 .plan 文件，则显示该文件内容。
因此程序需要能够提供：
	• 用户存在吗？
	• 如果用户存在，显示 .plan 文件的内容。
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}
