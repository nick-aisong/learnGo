// 6-7 listing07.go
// 这个示例程序展示如何创建 goroutine02 // 以及 goroutine 调度器的行为
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// main 是所有 Go 程序的入口
func main() {
	// 分配 2 个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(2)

	// wg 用来等待程序完成
	// 计数加 2，表示要等待两个 goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// 声明一个匿名函数，并创建一个 goroutine
	go func() {
		// 在函数退出时调用 Done 来通知 main 函数工作已经完成
		defer wg.Done()

		// 显示字母表 3 次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// 声明一个匿名函数，并创建一个 goroutine
	go func() {
		// 在函数退出时调用 Done 来通知 main 函数工作已经完成
		defer wg.Done()

		// 显示字母表 3 次
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// 等待 goroutine 结束
	fmt.Println("Waiting To Finish")

	wg.Wait()

	fmt.Println("\nTerminating Program")
}

// G:\GitHub\learnGo\Go语言实战\ch06_并发>go run listing07.go
// Start Goroutines
// Waiting To Finish
// a b c d e f g h i j k l m n o p q r s t u v w x y z a A B C D E b c d e f g h i j k l m n o p q r s t F G H I J K L M N O P Q R S T U V W X Y Z A B C D E F G H I J K L M N O P Q R S T u v w x y z a b c d e f g h i j k l m n o p q r s t u v w x y z U V W X Y Z A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
// Terminating Program
