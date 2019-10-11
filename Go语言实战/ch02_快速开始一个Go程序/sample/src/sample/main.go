// 直接从源码库copy过来的，修改了导入包的位置
package main

import (
	"log"
	"os"

	_ "sample/matchers"
	"sample/search"
)

// init is called prior to main.
// init 在 main 之前调用
func init() {
	// Change the device for logging to stdout.
	// 将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
// main 是整个程序的入口
func main() {
	// Perform the search for the specified term.
	// 使用特定的项做搜索
	search.Run("president")
}

// G:\GitHub\learnGo\Go语言实战\ch02_快速开始一个Go程序\sample\src\sample>set GOPATH=G:\GitHub\learnGo\Go语言实战\ch02_快速开始一个Go程序\sample
// G:\GitHub\learnGo\Go语言实战\ch02_快速开始一个Go程序\sample\src\sample>go run main.go > output.txt
// 2019/10/12 01:00:56 Register default matcher
// 2019/10/12 01:00:56 Register rss matcher
