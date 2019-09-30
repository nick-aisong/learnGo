// thread
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}
func main() {
	lock := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go Count(lock)
	}

	for {
		// 需要加锁，不然可能最后还没打印10，程序就退出了
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
}

//G:\GitHub\learnGo\Go语言编程\ch04_并发编程>go run thread.go
//1
//2
//3
//4
//5
//6
//7
//8
//9
//10
