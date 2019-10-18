/*
1. 下面的程序有什么错误？
package main
import "fmt"
func main() {
	for i := 0 ; i < 10 ; i++ {
		fmt.Printf("%v\n", i)
	}
	fmt.Printf("%v\n", i)
}
*/
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", i)
	}
	fmt.Printf("%v\n", i)
}

// 这个程序不能被编译，由于第 9 行的变量 i，未定义：i 仅在 for 循环中有效
