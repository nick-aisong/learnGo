// 4-1 add
package main

import "fmt"

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}
func main() {
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}
}

// 主函数启动了10个goroutine，然后返回，这时程序就退出了，而被启动的
// 执行Add(i, i)的goroutine没有来得及执行，所以程序没有任何输出
