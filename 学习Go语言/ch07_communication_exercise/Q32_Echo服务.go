/*
Q32. (1) Echo 服务
1. 编写一个简单的 echo 服务。使其监听于本地的 TCP 端口 8053 上。它应当可以
读取一行（以换行符结尾），将这行原样返回然后关闭连接。
2. 让这个服务可以并发，这样每个请求都可以在独立的 goroutine 中进行处理。
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}
