/*
Q32. (1) Echo 服务
1. 编写一个简单的 echo 服务。使其监听于本地的 TCP 端口 8053 上。它应当可以
读取一行（以换行符结尾），将这行原样返回然后关闭连接。
2. 让这个服务可以并发，这样每个请求都可以在独立的 goroutine 中进行处理。
*/
package main

import (
	"bufio"
	"fmt"
	"net"
)

func Echo(c net.Conn) {
	defer c.Close()
	line, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Printf("Failure to read: %s\n", err.Error())
		return
	}
	_, err = c.Write([]byte(line))
	if err != nil {
		fmt.Printf("Failure to write %s\n", err.Error())
		return
	}
}

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8053")
	if err != nil {
		fmt.Printf("Failure to listen: %s\n", err.Error())
	}
	for {
		if c, err := l.Accept(); err != nil {
			Echo(c)
		}
	}
}

/*
2. 为了使其能够并发处理链接，只需要修改一行代码，就是：
if c, err := l.Accept() ; err == nil { Echo(c) }
改为：
if c, err := l.Accept() ; err == nil { go Echo(c) }
*/
