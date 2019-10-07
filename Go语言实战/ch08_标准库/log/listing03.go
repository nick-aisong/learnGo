// 8-3 listing03.go
// 这个示例程序展示如何使用最基本的 log 包
package main

import (
	"log"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// Println 写到标准日志记录器
	log.Println("message")

	// Fatalln 在调用 Println()之后会接着调用 os.Exit(1)
	log.Fatalln("fatal message")

	// Panicln 在调用 Println()之后会接着调用 panic()
	log.Panicln("panic message")
}

// golang.org/src/log/log.go
const (
	// 将下面的位使用或运算符连接在一起，可以控制要输出的信息。没有
	// 办法控制这些信息出现的顺序（下面会给出顺序）或者打印的格式
	// （格式在注释里描述）。这些项后面会有一个冒号：
	// 2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	// 日期: 2009/01/23
	Ldate = 1 << iota
	// 时间: 01:23:23
	Ltime
	// 毫秒级时间: 01:23:23.123123。该设置会覆盖 Ltime 标志
	Lmicroseconds
	// 完整路径的文件名和行号: /a/b/c/d.go:23
	Llongfile
	// 最终的文件名元素和行号: d.go:23
	// 覆盖 Llongfile
	Lshortfile
	// 标准日志记录器的初始值
	LstdFlags = Ldate | Ltime
)

// G:\GitHub\learnGo\Go语言实战\ch08_标准库\log>go run listing03.go
// TRACE: 2019/10/07 18:30:59.261599 G:/GitHub/learnGo/Go语言实战/ch08_标准库/log/listing03.go:16: message
// TRACE: 2019/10/07 18:30:59.289524 G:/GitHub/learnGo/Go语言实战/ch08_标准库/log/listing03.go:19: fatal message
// exit status 1
