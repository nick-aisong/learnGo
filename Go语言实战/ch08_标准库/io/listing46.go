// 8-46 listing46.go
// 这个示例程序展示如何使用 io.Reader 和 io.Writer 接口
// 写一个简单版本的 curl
package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// main 是应用程序的入口
func main() {
	// 这里的 r 是一个响应， r.Body 是 io.Reader
	r, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	// 创建文件来保存响应内容
	file, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// 使用 MultiWriter，这样就可以同时向文件和标准输出设备
	// 进行写操作
	dest := io.MultiWriter(os.Stdout, file)

	// 读出响应的内容，并写到两个目的地
	io.Copy(dest, r.Body)
	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}
}

// G:\GitHub\learnGo\Go语言实战\ch08_标准库\io>go run listing46.go http://www.baidu.com listing46_output.txt
